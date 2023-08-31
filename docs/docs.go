// Code generated by swaggo/swag. DO NOT EDIT.

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
        "/add_user_to_segment": {
            "post": {
                "description": "Добавление сегментов пользователю.",
                "consumes": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "description": "Пользовательские сегменты",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.UserSegments"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/create_segment": {
            "post": {
                "description": "Создание сегмента.",
                "consumes": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "description": "Название сегмента",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.Segment"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Segment"
                        }
                    }
                }
            }
        },
        "/create_user": {
            "post": {
                "description": "Создание пользователя.",
                "consumes": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "description": "Пользователь",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/delete_segment": {
            "post": {
                "description": "Удаление сегмента.",
                "consumes": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "description": "Название сегмента",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.Segment"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/delete_user_from_segment": {
            "post": {
                "description": "Удаление сегментов у пользователя.",
                "consumes": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "description": "Пользовательские сегменты",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.UserSegments"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/show_segments/{user_id}": {
            "get": {
                "description": "Вывод списка активных сегментов у пользователя.",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        }
    },
    "definitions": {
        "controllers.UserSegments": {
            "type": "object",
            "properties": {
                "segments": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "domain.Segment": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        },
        "domain.User": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "127.0.0.1:3000",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Dynamic service",
	Description:      "Документация для  сервиса, хранящего пользователя и сегменты, в которых он состоит.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
