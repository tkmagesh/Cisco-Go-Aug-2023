package main

import (
	"fmt"

	// _ "github.com/tkmagesh/Cisco-Go-Aug-2023/09-modularity/calc" // to execute the init functions of the package

	"github.com/fatih/color"
	"github.com/tkmagesh/Cisco-Go-Aug-2023/09-modularity/calc"
	calcUtils "github.com/tkmagesh/Cisco-Go-Aug-2023/09-modularity/calc/utils"
)

func main() {
	color.Red("app executed")

	fmt.Println(calc.Add(100, 200))
	fmt.Println(calc.Subtract(100, 200))
	fmt.Println(calc.OpStats())

	// fmt.Println(utils.IsEven(21))
	fmt.Println(calcUtils.IsEven(21))
}
