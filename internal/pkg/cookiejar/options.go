package cookiejar

import (
	"net/http/cookiejar"

	"github.com/spf13/afero"
)

type Configurator interface {
	StoreToDisk() *Options
	StoreToMemory() *Options
	StoreToFilesystem(afero.Fs) *Options
	AsFile(string) *Options
	WithOptions(cookiejar.Options) *Options

	withDefaults() *Options
}

type Options struct {
	origin     *cookiejar.Options
	filename   string
	filesystem afero.Fs
}

func (opt *Options) StoreToDisk() *Options {
	opt.filesystem = afero.NewOsFs()
	return opt
}

func (opt *Options) StoreToMemory() *Options {
	opt.filesystem = afero.NewMemMapFs()
	return opt
}

func (opt *Options) StoreToFilesystem(fs afero.Fs) *Options {
	opt.filesystem = fs
	return opt
}

func (opt *Options) AsFile(name string) *Options {
	opt.filename = name
	return opt
}

func (opt *Options) WithOptions(options cookiejar.Options) *Options {
	opt.origin = &options
	return opt
}

func (opt *Options) withDefaults() *Options {
	if opt.filename == "" {
		opt.filename = "cookies.json"
	}
	if opt.filesystem == nil {
		opt.filesystem = afero.NewMemMapFs()
	}
	return opt
}
