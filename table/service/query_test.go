package serviceTable_test

import (
	structureSpec "github.com/taubyte/go-specs/structure"
	serviceTable "github.com/taubyte/tau-cli/table/service"
)

func ExampleQuery() {
	service := &structureSpec.Service{
		Id:          "QmbAA8hRosp5BaXFXikADCtpkQCgQCPdRVhnxjiSHfXdWH",
		Name:        "someProject",
		Description: "this is a service of some type",
		Tags:        []string{"apple", "orange", "banana"},
		Protocol:    "/test/v1",
	}

	serviceTable.Query(service)

	// Output:
	// ┌─────────────┬────────────────────────────────────────────────┐
	// │ ID          │ QmbAA8hRosp5BaXFXikADCtpkQCgQCPdRVhnxjiSHfXdWH │
	// ├─────────────┼────────────────────────────────────────────────┤
	// │ Name        │ someProject                                    │
	// ├─────────────┼────────────────────────────────────────────────┤
	// │ Description │ this is a service of some type                 │
	// ├─────────────┼────────────────────────────────────────────────┤
	// │ Tags        │ apple, orange, banana                          │
	// ├─────────────┼────────────────────────────────────────────────┤
	// │ Protocol    │ /test/v1                                       │
	// └─────────────┴────────────────────────────────────────────────┘
}
