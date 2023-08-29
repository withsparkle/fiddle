package presenter

type Option func(option) option

func Glamoured() Option {
	return func(opt option) option {
		opt.glamoured = true
		return opt
	}
}

type option struct {
	glamoured bool
}
