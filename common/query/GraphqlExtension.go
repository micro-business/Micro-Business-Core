package query

import (
	"github.com/graphql-go/graphql"
	ast "github.com/graphql-go/graphql/language/ast"
)

// GetSelectedFields returns the list of request fields listed under provider selection path in the Graphql query.
func GetSelectedFields(selectionPath []string, resolveParams graphql.ResolveParams) []string {
	fields := resolveParams.Info.FieldASTs

	for _, propName := range selectionPath {
		found := false

		for _, field := range fields {
			if field.Name.Value == propName {
				selections := field.SelectionSet.Selections
				fields = make([]*ast.Field, 0)

				for _, selection := range selections {
					fields = append(fields, selection.(*ast.Field))
				}

				found = true

				break
			}
		}

		if !found {
			return []string{}
		}
	}

	var collect []string

	for _, field := range fields {
		collect = append(collect, field.Name.Value)
	}

	return collect
}
