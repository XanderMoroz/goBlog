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
        "/api/v1/current_user": {
            "get": {
                "description": "Get token from users cookee",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "get current user",
                "operationId": "get-current-user",
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
            }
        },
        "/api/v1/login": {
            "post": {
                "description": "Authenticate User in app with given request body",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "user authentication",
                "parameters": [
                    {
                        "description": "Введите данные для авторизации",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.LoginRequest"
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
        "/api/v1/logout": {
            "get": {
                "description": "Clear JWT token by setting an empty value and expired time in the cookie",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "logout current user",
                "operationId": "logout-current-user",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/register": {
            "post": {
                "description": "Register User in app with given request body",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "user registration",
                "parameters": [
                    {
                        "description": "Введите данные для регистрации",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.SignUpUserRequest"
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
        "/articles": {
            "get": {
                "description": "Get all articles from db",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Articles"
                ],
                "summary": "get all articles",
                "operationId": "get-all-articles",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.ArticleResponse"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Creating Article in DB with given request body",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Articles"
                ],
                "summary": "create new article",
                "parameters": [
                    {
                        "description": "Введите данные статьи",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreateArticleRequest"
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
        "/articles/{id}": {
            "get": {
                "description": "Get an article by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Articles"
                ],
                "summary": "get an article by ID",
                "operationId": "get-article-by-id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Article ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.ArticleResponse"
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
            "put": {
                "description": "Update article by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Articles"
                ],
                "summary": "update article by ID",
                "operationId": "delete-article-by-id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "articleID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Введите новые данные статьи",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UpdateArticleBody"
                        }
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
            },
            "delete": {
                "description": "Delete a article by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Articles"
                ],
                "summary": "delete a article by ID",
                "operationId": "delete-article-by-id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "articleID",
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
        },
        "/categories": {
            "get": {
                "description": "Get all categories from db",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Categories"
                ],
                "summary": "get all categories",
                "operationId": "get-all-categories",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.CategoryResponse"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Creating Category in DB with given request body",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Categories"
                ],
                "summary": "create new category",
                "parameters": [
                    {
                        "description": "Введите название категории статьи",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreateCategoryBody"
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
        }
    },
    "definitions": {
        "models.ArticleResponse": {
            "description": "Тело ответа после cоздания статьи",
            "type": "object",
            "required": [
                "content",
                "createdAt",
                "title",
                "updatedAt"
            ],
            "properties": {
                "content": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "models.CategoryResponse": {
            "description": "Тело ответа после извлечения статьи",
            "type": "object",
            "required": [
                "createdAt",
                "title",
                "updatedAt"
            ],
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "models.CreateArticleRequest": {
            "description": "Тело запроса для создания статьи",
            "type": "object",
            "required": [
                "content",
                "title"
            ],
            "properties": {
                "content": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "models.CreateCategoryBody": {
            "description": "Тело запроса для создания статьи",
            "type": "object",
            "required": [
                "title"
            ],
            "properties": {
                "title": {
                    "type": "string"
                }
            }
        },
        "models.LoginRequest": {
            "description": "Тело запроса для аутентификации пользователя",
            "type": "object",
            "required": [
                "email",
                "password"
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
        "models.SignUpUserRequest": {
            "description": "Тело запроса для регистрации пользователя",
            "type": "object",
            "required": [
                "email",
                "name",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "models.UpdateArticleBody": {
            "description": "Тело запроса для обновления статьи",
            "type": "object",
            "required": [
                "content",
                "title"
            ],
            "properties": {
                "content": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "models.UserResponse": {
            "description": "Тело ответа после cоздания пользователя",
            "type": "object",
            "required": [
                "createdAt",
                "email",
                "name",
                "updatedAt"
            ],
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
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
	Host:             "127.0.0.1:8080",
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
