package colorify

import (
	"fmt"
	"github.com/mattn/go-isatty"
	"os"
)

type Colorify struct {
	color Color
	param Param
	noColor bool
}

type Param string

type Color string

const Base = "\033["

var (
	Reset Color = "0m"
	Red Color = "31m"
	Green Color = "32m"
	Yellow Color = "33m"
	Blue Color = "34m"
	Purple Color = "35m"
	Cyan Color = "36m"
	Gray Color = "37m"
	White Color= "97m"
	Bold Param = "1;"
	Underline Param = "4;"
	Regular Param = "0;"
)

func (c *Colorify) New() string {
	c.noColor = isatty.IsTerminal(os.Stdout.Fd()) && isatty.IsCygwinTerminal(os.Stdout.Fd())
	if c.param == "" {
		c.param = Regular
	}
	return fmt.Sprintf("%s%s%s", Base, c.param, c.color)
}

//func (c *Colorify) Reset() *Colorify {
//	return Colorify{color: Reset}
//}

func Println(color, msg string) (n int, err error) {
	return fmt.Println(color, msg)
}

