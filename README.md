# JSONParse

Parser and schema validator for json schema draft v4

```
go get github.com/gdm85/JSONParse
```
GoDoc: https://godoc.org/github.com/gdm85/JSONParse

In addition to the library, there are three cmd files that perform the basic operations
* **jp-pretty** - output input file as formatted json
* **jp-parse** - parse a json file with error checking
* **jp-schema** - confirm json document conforms to schema

```
cd cmd
go build jp-pretty.go
go build jp-parse.go
go build jp-schema.go
```

The implementation is a recursive descent parser that tokenizes a json document and confirms the structure of the document against json specification grammar.  The result is loaded into a tree structure.  If validating against a schema, the schema is represented in a tree structure and the trees are walked in parallel.  The implementation does not use json marshal or unmarshal and does not use reflection.

The implementation will read from a file or from a http source that emits json.

The test suite used is https://github.com/json-schema/JSON-Schema-Test-Suite/tree/develop/tests/draft4.  All of the main test cases are used except for refRemote.json.  In the optional folder, the format.json test cases are also pulled in.  When executing the tests, the test files are pulled down from github, broken up into test cases, written to the file system in the tests/ and tests/optional/ folders, and then read in for processing.  In the source folder, run the following:

```
go test -test.v
```

An example of schema validation.  The first one will just emit errors if present.  The second one will show the trace of the validator.

```
jp-schema -file http://petstore.swagger.io/v2/swagger.json -schema http://swagger.io/v2/schema.json
jp-schema -file http://petstore.swagger.io/v2/swagger.json -schema http://swagger.io/v2/schema.json -level trace
```
The informational messages at the end of hte second are issues encountered while processing "oneOf" logic which do not affect the validity of the document.  Still working on error output.

Known issues - the (*JSONNode).GetJson() can cause a recursive loop if the json has a $ref back to itself.  The implementation is still in development, so there are other issues in there as well. 
