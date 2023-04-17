// Package Colorify is a simple but straightforward attempt to create a stdout colorification/beautification
// module. Its based on a very simple premise of creating a simple set of colored and formatted output, only
// the io.Writer/os.stdout intrerfaces.
// Implementation is inspired by fatih/color package which is by all means a more feature rich and mature project.
// This is just a very simple re-implementation of the same.
package colorify

import (
	"fmt"
	tb "github.com/mainak90/trickB"
	"github.com/mattn/go-isatty"
	"io"
	"os"
)

// Main struct for interfacing with all golang fmt functions.  Elems -->  1. Color --> The colored o/p choice.
// 2. Attr --> Added string attributes like bold/underline etc  3. NoColour --> Automatically set attribute which disables coloration if not o/p to tty
// 4. ColorScheme --> The string __repr__ of the generated color scheme.
type Colorify struct {
	Color       Color
	Attr        Attr
	NoColor     string
	ColorScheme string
}

// Custom type Attr stores the string attributes for the colored output. Typically Bold and Underlines are
// typical attributes.
type Attr string

// Color type represents the colors to be used for the output.
type Color string

// Base const for the base str __repr__ for colors.
const Base = "\033["

// Color str __repr__ vars for each supported colors.
var (
	Reset     		Color = "0m"
	Red       		Color = "31m"
	Green     		Color = "32m"
	Yellow    		Color = "33m"
	Blue      		Color = "34m"
	Purple    		Color = "35m"
	Cyan      		Color = "36m"
	Gray      		Color = "37m"
	White     		Color = "97m"
	Bold      		Attr  = "1;"
	Transparent 	Attr  = "2;"
	Italics   		Attr  = "3;"
	Underline 		Attr  = "4;"
	Reverse			Attr  = "7;"
	Strikethrough 	Attr  =  "9;"
	Regular   		Attr  = "0;"
)

func doColor(c *Colorify) bool {
	switch tb.TrickBFromString(c.NoColor) {
	case tb.UnSet:
		return isatty.IsTerminal(os.Stdout.Fd()) || isatty.IsCygwinTerminal(os.Stdout.Fd())
	case tb.True:
		return false
	case tb.False:
		return true
	default:
		return true
	}
}

func stringifyColumnar(c *Colorify, inf ...interface{}) string {
	var infString string
	defaultColor := makeColorString(c, c.Color)
	infString += defaultColor
	for _, val := range inf {
		switch val.(type) {
		case Color:
			if doColor(c) {
				colScheme := makeColorString(c, val.(Color))
				infString += fmt.Sprintf("%s", colScheme)
			} else {
				continue
			}
		case string:
			infString += fmt.Sprintf("%s%s", " ", val)
		default:
			infString += fmt.Sprintf("%s%s", " ", val)
		}
	}
	return infString
}


func makeColorString(c *Colorify, col Color) string {
	if col == "" {
		col = Color(Reset)
	}
	c.Color = col
	c.ColorScheme = fmt.Sprintf("%s%s%s", Base, c.Attr, c.Color)
	return c.ColorScheme
}

// Done: Resets the coloration and formatting, mostly used right after statment execution.
func Done() string {
	return fmt.Sprintf("%s%s%s", Base, Regular, Reset)
}

func (c *Colorify) Println(msg ...interface{}) (n int, err error) {
	return fmt.Println(stringifyColumnar(c, msg...), Done())
}

// Printf: Implements the fmt.Printf interface just with the colorification addon.
func (c *Colorify) Printf(format string, msg ...interface{}) (n int, err error) {
	return fmt.Printf(format, stringifyColumnar(c, msg...), Done())
}

// Fprintln: Implements the fmt.Fprintln interface just with the colorification addon.
func (c *Colorify) Fprintln(w io.Writer, msg ...interface{}) (n int, err error) {
	return fmt.Fprintln(w, stringifyColumnar(c, msg...), Done())
}

// Fprintf: Implements the fmt.Fprintf interface just with the colorification addon.
func (c *Colorify) Fprintf(w io.Writer, format string, msg ...interface{}) (n int, err error) {
	return fmt.Fprintf(w, format, stringifyColumnar(c, msg...), Done())
}

//Fprint: Implements the fmt.Fprint interface just with the colorification addon.
func (c *Colorify) Fprint(w io.Writer, msg ...interface{}) (n int, err error) {
	return fmt.Fprint(w, stringifyColumnar(c, msg...), Done())
}

// Sprintf: Implements the fmt.Sprintf interface just with the colorification addon.
func (c *Colorify) Sprintf(msg ...interface{}) string {
	return fmt.Sprintf("%s%s", stringifyColumnar(c, msg...), Done())
}

// Sprintln: Implements the fmt.Sprintln interface just with the colorification addon.
func (c *Colorify) Sprintln(msg ...interface{}) string {
	return fmt.Sprintln(stringifyColumnar(c, msg...), Done())
}
