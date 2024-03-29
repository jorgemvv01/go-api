// Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Jorge Mario Villarreal V.",
            "url": "https://www.linkedin.com/in/jorgemariovillarreal/",
            "email": "jorgemvv01@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/genres": {
            "get": {
                "description": "Get all Genres.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Movie Genre"
                ],
                "summary": "Get all Genres",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        },
        "/genres/create": {
            "post": {
                "description": "Create a new genre.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Movie Genre"
                ],
                "summary": "Create Genre",
                "parameters": [
                    {
                        "description": "Create genre",
                        "name": "tags",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.GenreRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        },
        "/genres/delete/{ID}": {
            "delete": {
                "description": "Delete Genre by ID.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Movie Genre"
                ],
                "summary": "Delete Genre",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Delete genre by ID",
                        "name": "ID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        },
        "/genres/update/{ID}": {
            "put": {
                "description": "Update Genre by ID.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Movie Genre"
                ],
                "summary": "Update Genre",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Update genre by ID",
                        "name": "ID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Update genre",
                        "name": "tags",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.GenreRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        },
        "/genres/{ID}": {
            "get": {
                "description": "Get a genre by ID.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Movie Genre"
                ],
                "summary": "Get Genre by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Get genre by ID",
                        "name": "ID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        },
        "/movies": {
            "get": {
                "description": "Get all Movies.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Movies"
                ],
                "summary": "Get all Movies",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        },
        "/movies/create": {
            "post": {
                "description": "Create a new movie.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Movies"
                ],
                "summary": "Create Movie",
                "parameters": [
                    {
                        "description": "Create movie",
                        "name": "tags",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.MovieRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        },
        "/movies/delete/{ID}": {
            "delete": {
                "description": "Delete Movie by ID.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Movies"
                ],
                "summary": "Delete Movie",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Delete Movie by ID",
                        "name": "ID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        },
        "/movies/update/{ID}": {
            "put": {
                "description": "Update Movie by ID.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Movies"
                ],
                "summary": "Update Movie",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Update movie by ID",
                        "name": "ID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Update movie",
                        "name": "tags",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.MovieRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        },
        "/movies/{ID}": {
            "get": {
                "description": "Get a movie by ID.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Movies"
                ],
                "summary": "Get Movie by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Get movie by ID",
                        "name": "ID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        },
        "/rent/create": {
            "post": {
                "description": "Create a new rent.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Rent"
                ],
                "summary": "Create rent",
                "parameters": [
                    {
                        "description": "Create rent",
                        "name": "tags",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.RentRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        },
        "/types/": {
            "get": {
                "description": "Get all Types",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Movie Type"
                ],
                "summary": "Get all Types",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        },
        "/types/{ID}": {
            "get": {
                "description": "Get Type by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Movie Type"
                ],
                "summary": "Get Type by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Get Type by ID",
                        "name": "ID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        },
        "/users": {
            "get": {
                "description": "Get all Users.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Get all Users",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        },
        "/users/create": {
            "post": {
                "description": "Create a new user.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Create User",
                "parameters": [
                    {
                        "description": "Create user",
                        "name": "tags",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        },
        "/users/delete/{ID}": {
            "delete": {
                "description": "Delete User by ID.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Delete User",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Delete user by ID",
                        "name": "ID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        },
        "/users/update/{ID}": {
            "put": {
                "description": "Update User by ID.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Update User",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Update user by ID",
                        "name": "ID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Update user",
                        "name": "tags",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        },
        "/users/{ID}": {
            "get": {
                "description": "Get a user by ID.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Get User by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Get user by ID",
                        "name": "ID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.GenreRequest": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        },
        "models.MovieRequest": {
            "type": "object",
            "properties": {
                "genre_id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "overview": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                },
                "release_date": {
                    "type": "string"
                },
                "type_id": {
                    "type": "integer"
                }
            }
        },
        "models.RentRequest": {
            "type": "object",
            "properties": {
                "end_date": {
                    "type": "string"
                },
                "movie_ids": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "start_date": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "models.Response": {
            "type": "object",
            "properties": {
                "data": {},
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "models.UserRequest": {
            "type": "object",
            "properties": {
                "lastname": {
                    "type": "string"
                },
                "surname": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "VideoClub / Go-REST-API",
	Description:      "A simple Go-REST-API to manage movie rentals \\n GitHub Repository: https://github.com/jorgemvv01/go-api",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
