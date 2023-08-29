package cookiejar

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"

	"github.com/spf13/afero"
)

func Must(opt *Options) *Jar {
	jar, err := New(opt)
	if err != nil {
		panic(err)
	}
	return jar
}

func New(opt *Options) (*Jar, error) {
	opt.withDefaults()

	origin, err := cookiejar.New(opt.origin)
	if err != nil {
		return nil, err
	}
	file, err := opt.filesystem.OpenFile(opt.filename, os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		return nil, err
	}

	jar := &Jar{origin, file, make(map[string]url.URL)}
	return jar, jar.Init()
}

type Jar struct {
	origin  *cookiejar.Jar
	storage afero.File
	queries map[string]url.URL
}

func (jar *Jar) Init() error {
	var data map[string][]*http.Cookie
	if err := json.NewDecoder(jar.storage).Decode(&data); err != nil && !errors.Is(err, io.EOF) {
		return err
	}
	for raw, cookies := range data {
		u, err := url.Parse(raw)
		if err != nil {
			return err
		}
		jar.SetCookies(u, cookies)
	}
	return nil
}

func (jar *Jar) Dump() error {
	data := make(map[string][]*http.Cookie, len(jar.queries))
	for k, u := range jar.queries {
		data[k] = jar.Cookies(&u)
	}
	if err := jar.storage.Truncate(0); err != nil {
		return err
	}
	if _, err := jar.storage.Seek(0, 0); err != nil {
		return err
	}
	enc := json.NewEncoder(jar.storage)
	enc.SetIndent("", "  ")
	enc.SetEscapeHTML(false)
	return enc.Encode(data)
}

func (jar *Jar) SetCookies(u *url.URL, cookies []*http.Cookie) {
	key := jar.key(u)
	jar.queries[key.String()] = key
	jar.origin.SetCookies(u, cookies)
}

func (jar *Jar) Cookies(u *url.URL) []*http.Cookie {
	return jar.origin.Cookies(u)
}

func (jar *Jar) key(u *url.URL) url.URL {
	return url.URL{Scheme: u.Scheme, Host: u.Host, Path: "/"}
}
