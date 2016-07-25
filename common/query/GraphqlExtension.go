package query

import "github.com/graphql-go/graphql"

type StringKeyValuePair struct {
	Key   string `json:"id"`
	Value string `json:"name"`
}

var StringKeyValuePairType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "StringKeyValue",
		Fields: graphql.Fields{
			"key":   &graphql.Field{Type: graphql.String},
			"value": &graphql.Field{Type: graphql.String},
		},
	},
)
