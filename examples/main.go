package main

import (
	"fmt"
	"github.com/mainak90/colorify"
	"io"
	"os"
)

func main() {
	// Int the struct first, now this struct can be reused for all further interface invocations.
	// In general as string attributes are seldom changed, you can initialize them directly at the struct level.
	col := colorify.Colorify{Attr: colorify.Bold, NoColor: "false"}
	colEmpty := colorify.Colorify{Attr: colorify.Regular}
	// The color schemes are columnar, which now means you can use one color for out print output
	//or as many colors as you like. Colors can be set directly on the interface function level making it easier.
	col.Println(colorify.Red, "[Error]", colorify.Yellow, "This is a new error")
	// The print functions implement the exact spec as its fmt counterpart, which means each message interface can be
	//  a collection of words or an individual string.
	col.Printf("%s%s", colorify.Yellow, "This", "is", "a", "word", colorify.Blue, "This is another string!")
	col.Fprintf(os.Stdout, "\n%s%s", colorify.Green, "This is green!", colorify.Red, "This", "is", "Red!")
	io.WriteString(os.Stdout, col.Sprintln("\n", colorify.Blue, "This", "is", "Sparta"))
	fmt.Println(col.Sprintf(colorify.Red, "This", "is", "Red Sprintf output!!!", colorify.Green, "While", "this", "is", "green Sprintf!"))
	io.WriteString(os.Stdout, colEmpty.Sprintln("This is regular", colorify.Green, "This is green!"))
	colWithCol := colorify.Colorify{Color: colorify.Green, Attr: colorify.Bold}
	colWithCol.Println("Should be green!")
}


