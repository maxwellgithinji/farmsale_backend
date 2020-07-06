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
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "email": "email@email.com"
        },
        "license": {
            "name": "MIT",
            "url": "https://github.com/maxwellgithinji/farmsale_backend/blob/develop/LICENSE"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/auth/products": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get details of all products",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "summary": "Get details of all products",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/productsmodel.Product"
                            }
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "LOgin an existing user with their credentials",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "login"
                ],
                "summary": "Logs in an existing User",
                "parameters": [
                    {
                        "description": "login user",
                        "name": "login",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/usersmodel.LoginUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/usersmodel.User"
                            }
                        }
                    }
                }
            }
        },
        "/signup": {
            "post": {
                "description": "Signs Up a user with new credentials",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "signup"
                ],
                "summary": "Signs up a new user",
                "parameters": [
                    {
                        "description": "signup user",
                        "name": "signup",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/usersmodel.SignupUser"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/usersmodel.User"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "productsmodel.Product": {
            "type": "object",
            "properties": {
                "currentstock": {
                    "description": "` + "`" + `json:\"currentstock\" bson:\"currentstock\"` + "`" + `",
                    "type": "number"
                },
                "id": {
                    "type": "string"
                },
                "item": {
                    "description": "` + "`" + `json:\"item\" bson:\"item\"` + "`" + `",
                    "type": "string"
                },
                "standard": {
                    "description": "` + "`" + `json:\"standard\" bson:\"standard\"` + "`" + `",
                    "type": "string"
                },
                "unitprice": {
                    "description": "` + "`" + `json:\"unitprice\" bson:\"unitprice\"` + "`" + `",
                    "type": "number"
                }
            }
        },
        "usersmodel.LoginUser": {
            "type": "object",
            "properties": {
                "email": {
                    "description": "` + "`" + `json:\"email\" bson:\"email\"` + "`" + `",
                    "type": "string"
                },
                "password": {
                    "description": "` + "`" + `json:\"password\" bson:\"password\"` + "`" + `",
                    "type": "string"
                }
            }
        },
        "usersmodel.SignupUser": {
            "type": "object",
            "properties": {
                "email": {
                    "description": "` + "`" + `json:\"email\" bson:\"email\"` + "`" + `",
                    "type": "string"
                },
                "idnumber": {
                    "description": "` + "`" + `json:\"idnumber\" bson:\"idnumber\"` + "`" + `",
                    "type": "integer"
                },
                "password": {
                    "description": "` + "`" + `json:\"password\" bson:\"password\"` + "`" + `",
                    "type": "string"
                },
                "phonenumber": {
                    "description": "` + "`" + `json:\"phonenumber\" bson:\"phonenumber\"` + "`" + `",
                    "type": "string"
                },
                "username": {
                    "description": "` + "`" + `json:\"username\" bson:\"username\"` + "`" + `",
                    "type": "string"
                }
            }
        },
        "usersmodel.User": {
            "type": "object",
            "properties": {
                "email": {
                    "description": "` + "`" + `json:\"email\" bson:\"email\"` + "`" + `",
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "idnumber": {
                    "description": "` + "`" + `json:\"idnumber\" bson:\"idnumber\"` + "`" + `",
                    "type": "integer"
                },
                "isactive": {
                    "description": "` + "`" + `json:\"isactive\" bson:\"isactive\"` + "`" + `",
                    "type": "boolean"
                },
                "isadmin": {
                    "description": "` + "`" + `json:\"isadmin\" bson:\"isadmin\"` + "`" + `",
                    "type": "boolean"
                },
                "isblacklisted": {
                    "description": "` + "`" + `json:\"isblacklisted\" bson:\"isblacklisted\"` + "`" + `",
                    "type": "boolean"
                },
                "isvalid": {
                    "description": "` + "`" + `json:\"isvalid\" bson:\"isvalid\"` + "`" + `",
                    "type": "boolean"
                },
                "password": {
                    "description": "` + "`" + `json:\"password\" bson:\"password\"` + "`" + `",
                    "type": "string"
                },
                "phonenumber": {
                    "description": "` + "`" + `json:\"phonenumber\" bson:\"phonenumber\"` + "`" + `",
                    "type": "string"
                },
                "userclass": {
                    "description": "` + "`" + `json:\"userclass\" bson:\"userclass\"` + "`" + `",
                    "type": "string"
                },
                "username": {
                    "description": "` + "`" + `json:\"username\" bson:\"username\"` + "`" + `",
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
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
	BasePath:    "/api/v1",
	Schemes:     []string{},
	Title:       "Farmsale API",
	Description: "This is a service for connecting farmers to customers",
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
