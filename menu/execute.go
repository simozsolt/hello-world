package menu

import (
	"fmt"
	"github.com/simozsolt/hello-world/stringutil"
	"github.com/simozsolt/hello-world/numberutil"
)

func Execute(selected string) {
	fmt.Println()

	switch selected {
	case "1":
		sel1()
	case "2":
		sel2()
	}
}

func sel1() {
	print("Text: ")
	var inputString string
	fmt.Scanln(&inputString)

	println("Reversed string: ", stringutil.Reverse(inputString))
}

func sel2() {
	var nr int
	print("Number: ")
	fmt.Scanln(&nr)

	var divider int
	print("Divider: ")
	fmt.Scanln(&divider)

	one, two := numberutil.NrDividerRatioRest(nr, divider)
	fmt.Printf("%d / %d; Ratio: %d, rest: %d\n", nr, divider, one, two)
}
