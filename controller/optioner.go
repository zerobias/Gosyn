package controller

type OptionType string

func NewOptions(o, c, i bool) Options {
	opt := make(map[OptionType]bool, 3)
	opt[Optional] = o
	opt[Choises] = c
	opt[Iterative] = i
	return Options(opt)
}

func (ot OptionType) String() string { return string(ot) }

const (
	Optional  OptionType = "optional"
	Choises              = "choises"
	Iterative            = "iterative"
)

type Options map[OptionType]bool
type Optioner interface {
	HasNonDefaultOpt() bool
	Options() Options
}
