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
        "/health": {
            "get": {
                "description": "Returns the health status, environment, and version of the application.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Health"
                ],
                "summary": "Health Check",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/main.jsonResponseEnvelope"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/main.healthCheckResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/main.jsonErrorResponseEnvelope"
                        }
                    }
                }
            }
        },
        "/todos": {
            "get": {
                "description": "Returns a list of todos based on query parameters such as limit, offset, order, and search.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Todos"
                ],
                "summary": "Get Todos",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 10,
                        "description": "Number of todos to return (min: 1, max: 100)",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 0,
                        "description": "Number of todos to skip (min: 0)",
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "default": "desc",
                        "description": "Order of the todos (asc or desc)",
                        "name": "order",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Search term to filter todos by title or description (max length: 100)",
                        "name": "search",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/main.jsonResponseEnvelope"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/store.Todo"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/main.jsonErrorResponseEnvelope"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/main.jsonErrorResponseEnvelope"
                        }
                    }
                }
            },
            "post": {
                "description": "Creates a new todo item.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Todos"
                ],
                "summary": "Create Todo",
                "parameters": [
                    {
                        "description": "Todo item to be created",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.todoCreatePayload"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/main.jsonResponseEnvelope"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/store.Todo"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/main.jsonErrorResponseEnvelope"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/main.jsonErrorResponseEnvelope"
                        }
                    }
                }
            }
        },
        "/todos/{id}": {
            "get": {
                "description": "Returns a todo item by its ID.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Todos"
                ],
                "summary": "Get Todo",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of the todo to retrieve",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/main.jsonResponseEnvelope"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/store.Todo"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/main.jsonErrorResponseEnvelope"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/main.jsonErrorResponseEnvelope"
                        }
                    }
                }
            },
            "put": {
                "description": "Updates an existing todo item by its ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Todos"
                ],
                "summary": "Update Todo",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of the todo to update",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Todo update payload",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.todoUpdatePayload"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/main.jsonResponseEnvelope"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/store.Todo"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/main.jsonErrorResponseEnvelope"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/main.jsonErrorResponseEnvelope"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/main.jsonErrorResponseEnvelope"
                        }
                    }
                }
            },
            "delete": {
                "description": "Deletes a todo item by its ID.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Todos"
                ],
                "summary": "Delete Todo",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of the todo to delete",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/main.jsonResponseEnvelope"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/store.Todo"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/main.jsonErrorResponseEnvelope"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/main.jsonErrorResponseEnvelope"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "main.healthCheckResponse": {
            "type": "object",
            "properties": {
                "env": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "version": {
                    "type": "string"
                }
            }
        },
        "main.jsonErrorResponseEnvelope": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "main.jsonResponseEnvelope": {
            "type": "object",
            "properties": {
                "data": {}
            }
        },
        "main.todoCreatePayload": {
            "type": "object",
            "required": [
                "item"
            ],
            "properties": {
                "completed": {
                    "type": "boolean"
                },
                "item": {
                    "type": "string",
                    "maxLength": 100,
                    "minLength": 2
                }
            }
        },
        "main.todoUpdatePayload": {
            "type": "object",
            "properties": {
                "completed": {
                    "type": "boolean"
                },
                "item": {
                    "type": "string",
                    "maxLength": 100,
                    "minLength": 2
                }
            }
        },
        "store.Todo": {
            "type": "object",
            "properties": {
                "completed": {
                    "type": "boolean"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "item": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
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

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "/v1",
	Schemes:          []string{},
	Title:            "Go Todo List CRUD Application",
	Description:      "This is a sample api for todo list application",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
