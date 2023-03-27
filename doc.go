/*
Package colorify is an ANSI color package to output colorized or SGR defined
output to the standard output.
Quickstart on ways to use this:

		// Init colorify with color palette in struct.
		import "github.com/mainak90/colorify"
		red := colorify.Colorify{Color: colorify.Red, Attr: colorify.Underline}
		yellow := colorify.Colorify{Color: colorify.Yellow, Attr: colorify.Regular}

		// Start the init function
		redPrint := red.New()
		yellowPrint := yellow.New()

		// Use various fmt package interfaces to output in colour
		redPrint.Println("This", "is", "Sparta")
		yellowPrint.Println("This is yellow regular!!!")
		fmt.Println(yellowPrint.Sprintf("Sprintf message"))
		io.WriteString(os.Stdout, yellowPrint.Sprintln("This", "is", "Sparta"))
		redPrint.Printf("%s", "This", "is", "a", "word")

		// Colors supported
		Red, Green, Yellow, Blue, Purple, Cyan, White

	// Formatting attributes supported
	Bold, Italics, Underline

Help furnish this package with more add-on features.
*/
package colorify
