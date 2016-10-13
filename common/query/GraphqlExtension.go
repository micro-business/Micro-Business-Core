package query

import (
	"github.com/graphql-go/graphql"
	ast "github.com/graphql-go/graphql/language/ast"
)

type StringKeyValuePair struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

var StringKeyValuePairType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "StringKeyValuePair",
		Fields: graphql.Fields{
			"key":   &graphql.Field{Type: graphql.String},
			"value": &graphql.Field{Type: graphql.String},
		},
	},
)

func GetSelectedFields(selectionPath []string,
	resolveParams graphql.ResolveParams) []string {
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
