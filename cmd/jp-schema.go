package main

import (
	"flag"
	"fmt"
	"github.com/gdm85/JSONParse"
	"os"
	//	"strings"
)

func main() {
	filePtr := flag.String("file", "security.json", "file to parse")
	errorPtr := flag.Int("maxerror", 8, "maximum number of errors before abort")
	schemaPtr := flag.String("schema", "", "schema to use for validation")
	levelPtr := flag.String("level", "error", "level of logging (trace, info, warning, error)")

	flag.Parse()

	file := *filePtr
	maxError := *errorPtr
	schemaFile := *schemaPtr
	level := *levelPtr

	fmt.Println("parse file", file)
	var parser *JSONParse.JSONParser
	if len(file) > 0 {
		parser = JSONParse.NewJSONParser(file, maxError, level)
		valDoc, errs := parser.Parse()
		if !valDoc {
			fmt.Println(errs)
			os.Exit(1)
		}
	}
	fmt.Println("  -valid json file")

	if len(schemaFile) > 0 {
		fmt.Println("load schema")
		schema := JSONParse.NewJSONSchema(schemaFile, level)
		fmt.Println("  -validate file against schema", schemaFile)
		valid, errs := schema.ValidateDocument(file)
		if !valid {
			errs.Output()
			fmt.Println("  -document is not valid against schema")
			os.Exit(1)
		}
	}

	fmt.Println("  -documente is valid against schema")
}
