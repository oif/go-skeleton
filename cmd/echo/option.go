package echo

import "github.com/oif/go-skeleton/pkg/echo"

type Option struct {
}

func NewDefaultOption() *Option {
	return new(Option)
}

func (o *Option) Run() error {
	echo.Run()
	return nil
}
