// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "https://github.com/niklaus-code/nicloud",
        "contact": {
            "name": "NIKLAUS",
            "url": "https://github.com/niklaus-code/nicloud",
            "email": "1309584951@qq.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/user/login": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "用户登录接口1",
                "parameters": [
                    {
                        "type": "string",
                        "description": "name",
                        "name": "usernmae",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "name",
                        "name": "passwd",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/users.Vms_users"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "users.Vms_users": {
            "type": "object",
            "required": [
                "Passwd",
                "Username"
            ],
            "properties": {
                "Mobile": {
                    "type": "string"
                },
                "Passwd": {
                    "type": "string"
                },
                "Role": {
                    "type": "integer",
                    "enum": [
                        1,
                        2
                    ]
                },
                "Username": {
                    "type": "string"
                },
                "create_time": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "NILCOUD",
	Description:      "PRIVATE CLOUD PLATFORM",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
