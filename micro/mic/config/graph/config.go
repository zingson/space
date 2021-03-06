package graph

import (
	"encoding/json"
	"errors"
	"log"

	ql "github.com/graphql-go/graphql"

	"config/mgdb"
)

func init() {
	MergeMutationFields(configMutationFields)
	MergeQueryFields(configQueryFields)
}

type config struct {
	Total int64
	List  interface{}
	Info  interface{}
	Value string
}

// Graphql Query
var configQueryFields = ql.Fields{
	"config": &ql.Field{
		Type: ql.NewObject(ql.ObjectConfig{
			Name: "ConfigQueryType",
			Fields: ql.Fields{
				"total": &ql.Field{
					Type: ql.Int,
					Args: ql.FieldConfigArgument{
						"name": &ql.ArgumentConfig{Type: ql.String, Description: "服务名"},
					},
					Resolve: func(p ql.ResolveParams) (i interface{}, e error) {
						return 0, nil
					},
					Description: "总记录数查询",
				},
				"list": &ql.Field{
					Type: ql.NewList(ql.NewObject(ql.ObjectConfig{
						Name: "ConfigListType",
						Fields: ql.Fields{
							"name": &ql.Field{
								Type:        ql.String,
								Description: "服务名",
							},
							"value": &ql.Field{
								Type:        ql.String,
								Description: "配置JSON字符串",
							},
						},
						Description: "集合对象",
					})),
					Args: ql.FieldConfigArgument{
						"limit": &ql.ArgumentConfig{
							Type:         ql.Int,
							DefaultValue: 100,
							Description:  "返回结果集数量",
						},
						"skip": &ql.ArgumentConfig{
							Type:         ql.Int,
							DefaultValue: 0,
							Description:  "跳过记录数",
						},
					},
					Resolve: func(p ql.ResolveParams) (i interface{}, e error) {
						return i, nil
					},
					Description: "数据集合",
				},
				"info": &ql.Field{
					Type: ql.NewObject(ql.ObjectConfig{
						Name: "ConfigInfoType",
						Fields: ql.Fields{
							"name":      &ql.Field{Type: ql.String, Description: "服务名"},
							"value":     &ql.Field{Type: ql.String, Description: "配置JSON字符串"},
							"remark":    &ql.Field{Type: ql.String, Description: "配置备注描述信息"},
							"updatedAt": &ql.Field{Type: ql.String, Description: "最后更新时间"},
						},
						Description: "配置信息",
					}),
					Args: ql.FieldConfigArgument{
						"name": &ql.ArgumentConfig{Type: ql.NewNonNull(ql.String), Description: "服务名"},
					},
					Resolve: func(p ql.ResolveParams) (i interface{}, e error) {
						log.Printf("执行 query config info")
						name := p.Args["name"].(string)
						msConfig, e := mgdb.NewConfig().FindOne(name)
						return msConfig, e
					},
					Description: "配置信息详情",
				},
				"value": &ql.Field{
					Name: "",
					Type: ql.String,
					Args: ql.FieldConfigArgument{
						"name": &ql.ArgumentConfig{Type: ql.NewNonNull(ql.String), Description: "服务名"},
					},
					Resolve: func(p ql.ResolveParams) (i interface{}, e error) {
						log.Printf("执行 query config value")
						name := p.Args["name"].(string)
						msConfig, e := mgdb.NewConfig().FindOne(name)
						if e != nil {
							e = errors.New("查询结果错误")
							return nil, e
						}
						v, e := json.Marshal(msConfig.Value)
						if e != nil {
							e = errors.New("查询结果转JSON错误")
							return nil, e
						}
						return string(v), e
					},
					DeprecationReason: "",
					Description:       "配置内容JSON字符串",
				},
			},
			Description: "配置信息查询类型",
		}),
		Resolve: func(p ql.ResolveParams) (i interface{}, e error) {
			log.Printf("执行 query config")
			rs := &config{}
			return rs, e
		},
		Description: "配置查询",
	},
}

// Graphql Mutation
var configMutationFields = ql.Fields{
	"configSave": &ql.Field{
		Type: updateResultType,
		Args: ql.FieldConfigArgument{
			"name": &ql.ArgumentConfig{
				Type:         ql.NewNonNull(ql.String),
				DefaultValue: nil,
				Description:  "服务名",
			},
			"value": &ql.ArgumentConfig{
				Type:         ql.NewNonNull(ql.String),
				DefaultValue: "{}",
				Description:  "JSON格式配置信息",
			},
		},
		Resolve: func(p ql.ResolveParams) (i interface{}, e error) {
			log.Printf("执行QL方法：%s", "config.configSave")
			value := p.Args["value"].(string)
			jsonMap := make(map[string]interface{})
			e = json.Unmarshal([]byte(value), &jsonMap)
			if e != nil {
				return nil, e
			}
			// updateResult := nil // (&(collection.MsConfig{Name: p.Args["name"].(string), Value: jsonMap})).Save()

			return nil, nil
		},
		Description: "新增/更新配置，响应成功更新条数",
	},
}

// Mongodb 更新结果,对应结构体：mongo.UpdateResult
var updateResultType = ql.NewObject(ql.ObjectConfig{
	Name: "UpdateResultType",
	Fields: ql.Fields{
		"matchedCount":  &ql.Field{Type: ql.Int, Description: "更新条件匹配文档数量"},
		"modifiedCount": &ql.Field{Type: ql.Int, Description: "更新文档数量"},
		"upsertedCount": &ql.Field{Type: ql.Int, Description: "upsert时插入文档数量"},
		"upsertedID":    &ql.Field{Type: ql.String, Description: "发生upsert时插入文档的标识符"},
	},
	Description: "Mongodb更新结果",
})
