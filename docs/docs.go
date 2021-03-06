// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Ashadi",
            "url": "https://ashadi-ch.xyz"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/login": {
            "post": {
                "description": "retrieve jwt token by given username and password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "login"
                ],
                "summary": "Login endpoint",
                "parameters": [
                    {
                        "description": "user payload",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UserLogin"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.CommonMessage"
                        }
                    }
                }
            }
        },
        "/user-billing": {
            "get": {
                "description": "billing user information",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "user billing endpoint",
                "parameters": [
                    {
                        "type": "string",
                        "description": "jwt token with Bearer prefix",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.CommonMessage"
                        }
                    }
                }
            }
        },
        "/user-info": {
            "get": {
                "description": "retrieve user information",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "userinfo endpoint",
                "parameters": [
                    {
                        "type": "string",
                        "description": "jwt token with Bearer prefix",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.CommonMessage"
                        }
                    }
                }
            }
        },
        "/user-update": {
            "post": {
                "description": "update user information by given payload",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "user update endpoint",
                "parameters": [
                    {
                        "type": "string",
                        "description": "jwt token with Bearer prefix",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "user payload",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UserInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.CommonMessage"
                        }
                    }
                }
            }
        },
        "/user-upload": {
            "post": {
                "description": "upload image by user",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "user upload image endpoint",
                "parameters": [
                    {
                        "type": "string",
                        "description": "jwt token with Bearer prefix",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "file to be upload in bytes",
                        "name": "file",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "primitive"
                        }
                    },
                    {
                        "description": "description of image",
                        "name": "description",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "primitive"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.CommonMessage"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.CommonMessage": {
            "type": "object",
            "properties": {
                "data": {},
                "errors": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "message": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "model.UserInput": {
            "type": "object",
            "required": [
                "email"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "model.UserLogin": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
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
	Host:        "gmr.ashadi-ch.xyz",
	BasePath:    "/v1/api",
	Schemes:     []string{},
	Title:       "GMR API",
	Description: "GMR API endpoint documentation",
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
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
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
