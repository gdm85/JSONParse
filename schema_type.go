package JSONParse

import (
)

// 5.5.2.  type
// 
// 5.5.2.1.  Valid values
// 
// The value of this keyword MUST be either a string or an array. If it is an array, elements of the array MUST be strings and MUST be unique.
// 
// String values MUST be one of the seven primitive types defined by the core specification.
// 
// 5.5.2.2.  Conditions for successful validation
// 
// An instance matches successfully if its primitive type is one of the types defined by keyword. Recall: "number" includes "integer".
//
// === validates that the type specified in the document is the same as specified in the schema
//  todo: if number, need to check if integer
//
func validType(mem *JSONNode, schema *JSONNode, parent *JSONNode) bool {
	schemaValue := schema.GetValue().(string)
	value := valueTypeToString(mem.GetValueType())

	if value == schemaValue {
		Trace.Println("  validType() - match on ", value)
		return true
	} else {
		Trace.Println("  validType() - invalid match on ", schemaValue, " -- was ", value)
		OutputError(mem, "invalid type: expecting - " + schemaValue + " - found - " + value + " instead")
		return false
	}
}

func valueTypeToString(valType ValueType) string {
	if valType == V_OBJECT {
		return "object"
	} else if valType == V_STRING {
		return "string"
	} else if valType == V_NUMBER {
		return "number"
	} else if valType == V_BOOLEAN {
		return "boolean"
	} else if valType == V_ARRAY {
		return "array"
	} else if valType == V_NULL {
		return "null"
	} else {
		return "unknown"
	}
}
