// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "termsOfService": "https://github.com/go-programming-tour-book",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/articles/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "文章ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功",
                        "schema": {
                            "$ref": "#/definitions/model.ArticleSwagger"
                        }
                    },
                    "400": {
                        "description": "請求錯誤",
                        "schema": {
                            "$ref": "#/definitions/errcode.Error"
                        }
                    },
                    "500": {
                        "description": "內部錯誤",
                        "schema": {
                            "$ref": "#/definitions/errcode.Error"
                        }
                    }
                }
            }
        },
        "/api/v1/tags": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "取得多個標籤",
                "parameters": [
                    {
                        "maxLength": 100,
                        "type": "string",
                        "description": "標籤名稱",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "enum": [
                            0,
                            1
                        ],
                        "type": "integer",
                        "default": 1,
                        "description": "狀態",
                        "name": "state",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "頁碼",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "每頁數量",
                        "name": "page_size",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功",
                        "schema": {
                            "$ref": "#/definitions/model.TagSwagger"
                        }
                    },
                    "400": {
                        "description": "請求錯誤",
                        "schema": {
                            "$ref": "#/definitions/errcode.Error"
                        }
                    },
                    "500": {
                        "description": "內部錯誤",
                        "schema": {
                            "$ref": "#/definitions/errcode.Error"
                        }
                    }
                }
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "新增標籤",
                "parameters": [
                    {
                        "maxLength": 100,
                        "minLength": 3,
                        "description": "標籤名稱",
                        "name": "name",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "enum": [
                            0,
                            1
                        ],
                        "default": 1,
                        "description": "狀態",
                        "name": "state",
                        "in": "body",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    {
                        "maxLength": 100,
                        "minLength": 3,
                        "description": "建立者",
                        "name": "create_by",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功",
                        "schema": {
                            "$ref": "#/definitions/model.TagSwagger"
                        }
                    },
                    "400": {
                        "description": "請求錯誤",
                        "schema": {
                            "$ref": "#/definitions/errcode.Error"
                        }
                    },
                    "500": {
                        "description": "內部錯誤",
                        "schema": {
                            "$ref": "#/definitions/errcode.Error"
                        }
                    }
                }
            }
        },
        "/api/v1/tags/{id}": {
            "put": {
                "produces": [
                    "application/json"
                ],
                "summary": "更新標籤",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "標籤ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "maxLength": 100,
                        "minLength": 3,
                        "description": "標籤名稱",
                        "name": "name",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "enum": [
                            0,
                            1
                        ],
                        "default": 1,
                        "description": "狀態",
                        "name": "state",
                        "in": "body",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    {
                        "maxLength": 100,
                        "minLength": 3,
                        "description": "修改者",
                        "name": "modified_by",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功",
                        "schema": {
                            "$ref": "#/definitions/model.TagSwagger"
                        }
                    },
                    "400": {
                        "description": "請求錯誤",
                        "schema": {
                            "$ref": "#/definitions/errcode.Error"
                        }
                    },
                    "500": {
                        "description": "內部錯誤",
                        "schema": {
                            "$ref": "#/definitions/errcode.Error"
                        }
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "summary": "刪除標籤",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "標籤ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功",
                        "schema": {
                            "$ref": "#/definitions/model.TagSwagger"
                        }
                    },
                    "400": {
                        "description": "請求錯誤",
                        "schema": {
                            "$ref": "#/definitions/errcode.Error"
                        }
                    },
                    "500": {
                        "description": "內部錯誤",
                        "schema": {
                            "$ref": "#/definitions/errcode.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "app.Pager": {
            "type": "object",
            "properties": {
                "page": {
                    "type": "integer"
                },
                "pageSize": {
                    "type": "integer"
                },
                "total_rows": {
                    "type": "integer"
                }
            }
        },
        "errcode.Error": {
            "type": "object"
        },
        "model.Article": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "cover_image_url": {
                    "type": "string"
                },
                "created_by": {
                    "type": "string"
                },
                "created_on": {
                    "type": "integer"
                },
                "deleted_on": {
                    "type": "integer"
                },
                "desc": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "is_del": {
                    "type": "integer"
                },
                "modified_by": {
                    "type": "string"
                },
                "modified_on": {
                    "type": "integer"
                },
                "state": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "model.ArticleSwagger": {
            "type": "object",
            "properties": {
                "list": {
                    "$ref": "#/definitions/model.Article"
                },
                "pager": {
                    "$ref": "#/definitions/app.Pager"
                }
            }
        },
        "model.Tag": {
            "type": "object",
            "properties": {
                "created_by": {
                    "type": "string"
                },
                "created_on": {
                    "type": "integer"
                },
                "deleted_on": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "is_del": {
                    "type": "integer"
                },
                "modified_by": {
                    "type": "string"
                },
                "modified_on": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "state": {
                    "type": "integer"
                }
            }
        },
        "model.TagSwagger": {
            "type": "object",
            "properties": {
                "list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Tag"
                    }
                },
                "pager": {
                    "$ref": "#/definitions/app.Pager"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "部落格系統",
	Description: "GO語言程式設計之旅:一起用Go做專案",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
