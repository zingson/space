package graph

import (
	"log"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

// Http Handler
var GraphqlHttpHandler = handler.New(&handler.Config{
	Schema:   GraphqlSchema(),
	Pretty:   true,
	GraphiQL: true,
})

// Graphql Schema
var GraphqlSchema = func() *graphql.Schema {
	newSchema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:        QueryType,
		Mutation:     MutationType,
		Subscription: nil,
		Types:        nil,
		Directives:   nil,
		Extensions:   nil,
	})
	if err != nil {
		// 异常退出
		log.Fatal(err)
	}
	log.Printf("GraphqlSchema Load Success...")
	return &newSchema
}

// Graphql Query Type
var QueryType = graphql.NewObject(graphql.ObjectConfig{
	Name:       "Query",
	Interfaces: nil,
	Fields: FieldsMerge(
		ConfigQueryFields,
	),
	IsTypeOf:    nil,
	Description: "查询操作",
})

// Graphql Mutation Type
var MutationType = graphql.NewObject(graphql.ObjectConfig{
	Name:       "Mutation",
	Interfaces: nil,
	Fields: FieldsMerge(
		ConfigMutationFields,
	),
	IsTypeOf:    nil,
	Description: "更新操作",
})

// Merge Graphql Fields
func FieldsMerge(args ...graphql.Fields) graphql.Fields {
	fields := make(graphql.Fields)
	for _, arg := range args {
		for k, v := range arg {
			fields[k] = v
		}
	}
	return fields
}
