// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/execute": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Execute instructions",
                "parameters": [
                    {
                        "description": "Instruction list",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/executor.ExecuteRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/executor.Result"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "executor.ExecuteRequest": {
            "type": "object",
            "properties": {
                "instructions": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/executor.Instruction"
                    }
                }
            }
        },
        "executor.Instruction": {
            "type": "object",
            "properties": {
                "left": {},
                "op": {
                    "type": "string"
                },
                "right": {},
                "type": {
                    "$ref": "#/definitions/executor.InstructionType"
                },
                "var": {
                    "type": "string"
                }
            }
        },
        "executor.InstructionType": {
            "type": "string",
            "enum": [
                "calc",
                "print"
            ],
            "x-enum-varnames": [
                "CalcType",
                "PrintType"
            ]
        },
        "executor.Result": {
            "type": "object",
            "properties": {
                "value": {
                    "type": "integer"
                },
                "var": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
