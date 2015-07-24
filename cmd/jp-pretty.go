package main

import (
	"fmt"
	"github.com/gdm85/JSONParse"
	"os"
)

func main() {
	file := os.Args[1]

	var parser *JSONParse.JSONParser
	if len(file) > 0 {
		parser = JSONParse.NewJSONParser(file, 10, "error")
		valDoc, errs := parser.Parse()
		if !valDoc {
			fmt.Println(errs)
			os.Exit(1)
		}
		fmt.Println(parser.Pretty())
	} else {
		fmt.Println("must provide file or http service endpoint")
	}
}
