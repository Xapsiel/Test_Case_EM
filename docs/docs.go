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
        "/info": {
            "get": {
                "description": "Получение списка песен из базы данных с фильтрацией по параметрам",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "songs"
                ],
                "summary": "Получение списка песен",
                "parameters": [
                    {
                        "type": "string",
                        "default": "Supermassive Black Hole",
                        "description": "Название песни",
                        "name": "song",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "default": "Muse",
                        "description": "Группа",
                        "name": "group",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "ID песни",
                        "name": "id",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Номер страницы",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "2006-07-19",
                        "description": "Дата с",
                        "name": "since",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "2006-07-19",
                        "description": "Дата по",
                        "name": "to",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Song"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    }
                }
            }
        },
        "/info/verse": {
            "get": {
                "description": "Получение текста конкретного куплета песни(Обязательные параметры - song,group или id)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "songs"
                ],
                "summary": "Получение текста куплета песни",
                "parameters": [
                    {
                        "type": "string",
                        "default": "Supermassive Black Hole",
                        "description": "Название песни",
                        "name": "song",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "default": "Muse",
                        "description": "Группа",
                        "name": "group",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 1,
                        "description": "ID песни",
                        "name": "id",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "example": 1,
                        "description": "Номер куплета",
                        "name": "verse",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.resultResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    }
                }
            }
        },
        "/songs": {
            "put": {
                "description": "Обновление информации о песне по предоставленным данным(Обязательные параметры- song,group или id)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "songs"
                ],
                "summary": "Обновление информации о песне",
                "parameters": [
                    {
                        "description": "Данные песни для обновления",
                        "name": "song",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Song"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.resultResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Добавление новой песни в базу данных(Обязательные параметры - song,group)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "songs"
                ],
                "summary": "Добавление новой песни",
                "parameters": [
                    {
                        "description": "Данные песни",
                        "name": "song",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Song"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.resultResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Удаление песни по предоставленным данным(Обязательные параметры- song,group или id)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "songs"
                ],
                "summary": "Удаление песни",
                "parameters": [
                    {
                        "description": "Данные песни для удаления",
                        "name": "song",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Song"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.resultResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handler.errorResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "error description"
                },
                "status": {
                    "type": "string",
                    "example": "fail"
                }
            }
        },
        "handler.resultResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "status": {
                    "type": "string",
                    "example": "success"
                },
                "text": {
                    "type": "string",
                    "example": "description"
                }
            }
        },
        "models.Song": {
            "type": "object",
            "properties": {
                "group": {
                    "type": "string",
                    "example": "Muse"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "link": {
                    "type": "string",
                    "example": "https://www.youtube.com/watch?v=Xsp3_a-PMTw"
                },
                "release_date": {
                    "type": "string",
                    "example": "2006-07-19"
                },
                "song_name": {
                    "type": "string",
                    "example": "Supermassive Black Hole"
                },
                "text": {
                    "type": "string",
                    "example": "Ooh baby, don't you know I suffer?\nOoh baby, can you hear me moan?"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Songs API",
	Description:      "This is an API for managing songs.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
