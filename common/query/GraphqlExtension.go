package query

import "github.com/graphql-go/graphql"

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
