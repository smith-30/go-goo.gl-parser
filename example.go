package main

import (
	"fmt"

	"github.com/smith-30/go-goo.gl-parser/parser"
)

func main() {
	p := parser.NewParser("your api token")
	d, err := p.DecodeURL("https://goo.gl/kMxDaw")
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	fmt.Println(d) // https://github.com/smith-30/go-goo.gl-parser
}
