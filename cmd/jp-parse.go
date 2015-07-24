package main

import (
	"flag"
	"fmt"
	"github.com/gdm85/JSONParse"
	//	"strings"
)

func main() {
	filePtr := flag.String("file", "security.json", "file to parse")
	errorPtr := flag.Int("maxerror", 5, "maximum number of errors before abort")
	levelPtr := flag.String("level", "error", "level of logging (trace, info, warning, error)")

	flag.Parse()

	file := *filePtr
	maxError := *errorPtr
	level := *levelPtr

	fmt.Println("file", file)
	var parser *JSONParse.JSONParser
	if len(file) > 0 {
		parser = JSONParse.NewJSONParser(file, maxError, level)
		valDoc, errs := parser.Parse()
		if !valDoc {
			for i := range errs {
				parser.OutputError(errs[i])
			}
		}
	}
}
