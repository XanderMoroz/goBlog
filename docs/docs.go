// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
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
        "/article": {
            "get": {
                "description": "Get all items in the article list",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Articles"
                ],
                "summary": "get all items",
                "operationId": "get-all-articles",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.article"
                        }
                    }
                }
            },
            "post": {
                "description": "Add a new item to the article list",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Articles"
                ],
                "summary": "add a new item",
                "operationId": "create-article",
                "parameters": [
                    {
                        "description": "article data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.article"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.article"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/main.message"
                        }
                    }
                }
            }
        },
        "/article/{id}": {
            "get": {
                "description": "Get a article item by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Articles"
                ],
                "summary": "get a article item by ID",
                "operationId": "get-article-by-id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "article ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.article"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/main.message"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a article item by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Articles"
                ],
                "summary": "delete a article item by ID",
                "operationId": "delete-article-by-id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "article ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.article"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/main.message"
                        }
                    }
                }
            }
        },
        "/nice": {
            "get": {
                "description": "get the status of server.",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "root"
                ],
                "summary": "Show the status of server.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/users": {
            "get": {
                "description": "Get all users from db",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "get all users",
                "operationId": "get-all-users",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.UserResponse"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Creating User in DB with given request body",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "create new user",
                "parameters": [
                    {
                        "description": "Request of Creating User Object",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreateUserRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/users/{id}": {
            "get": {
                "description": "Get a user item by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "get a user item by ID",
                "operationId": "get-user-by-id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "userUUID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.UserResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a user item by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "delete a user item by ID",
                "operationId": "delete-user-by-id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "userUUID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "main.article": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "task": {
                    "type": "string"
                }
            }
        },
        "main.message": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "models.CreateUserRequest": {
            "description": "Тело запроса для создания пользователя",
            "type": "object",
            "required": [
                "LastName",
                "age",
                "country",
                "email",
                "firstName",
                "role"
            ],
            "properties": {
                "LastName": {
                    "type": "string"
                },
                "age": {
                    "type": "integer"
                },
                "country": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "firstName": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                }
            }
        },
        "models.UserResponse": {
            "description": "Тело запроса для создания пользователя",
            "type": "object",
            "required": [
                "age",
                "country",
                "createdAt",
                "email",
                "firstName",
                "lastName",
                "role",
                "updatedAt"
            ],
            "properties": {
                "age": {
                    "type": "integer"
                },
                "country": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "firstName": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "lastName": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                },
                "updatedAt": {
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
	Title:            "Good News on Go",
	Description:      "Сервис с новостными статьями и блогами.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
