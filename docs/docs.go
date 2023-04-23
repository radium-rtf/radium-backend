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
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/account": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "account"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.UserDto"
                        }
                    }
                }
            }
        },
        "/account/course": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "account"
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "$ref": "#/definitions/entity.Course"
                        }
                    }
                }
            }
        },
        "/account/name": {
            "patch": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "account"
                ],
                "parameters": [
                    {
                        "description": "Новое имя",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.UserName"
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
        "/account/password": {
            "patch": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "account"
                ],
                "parameters": [
                    {
                        "description": "PasswordUpdate",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.PasswordUpdate"
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
        "/auth/refresh": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "parameters": [
                    {
                        "description": "RefreshToken",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.RefreshToken"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.Tokens"
                        }
                    }
                }
            }
        },
        "/auth/signIn": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "parameters": [
                    {
                        "description": "SignIn",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.SignIn"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.Tokens"
                        }
                    }
                }
            }
        },
        "/auth/signUp": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "parameters": [
                    {
                        "description": "SignUp",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.SignUp"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/entity.Tokens"
                        }
                    }
                }
            }
        },
        "/collaborator": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "course"
                ],
                "parameters": [
                    {
                        "description": "collaborator",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.Collaborator"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "ok",
                        "schema": {
                            "$ref": "#/definitions/entity.Collaborator"
                        }
                    }
                }
            }
        },
        "/course": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "course"
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "$ref": "#/definitions/entity.Course"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "multipart/form-data"
                ],
                "tags": [
                    "course"
                ],
                "parameters": [
                    {
                        "type": "file",
                        "description": "logo",
                        "name": "logo",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "name",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "description",
                        "name": "description",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "kotlin, math и тд",
                        "name": "type",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "created",
                        "schema": {
                            "$ref": "#/definitions/entity.Course"
                        }
                    }
                }
            }
        },
        "/course/{courseId}": {
            "get": {
                "tags": [
                    "course"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "course id",
                        "name": "courseId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "$ref": "#/definitions/entity.CourseTitle"
                        }
                    }
                }
            }
        },
        "/group": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "group"
                ],
                "parameters": [
                    {
                        "description": "GroupName",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.GroupName"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "created",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/join/course/{courseId}": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "course"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "course id",
                        "name": "courseId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "created",
                        "schema": {
                            "$ref": "#/definitions/entity.Course"
                        }
                    }
                }
            }
        },
        "/join/group/{groupId}": {
            "patch": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "group"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "group id",
                        "name": "groupId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "created",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/link/course": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "course"
                ],
                "parameters": [
                    {
                        "description": "link",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.LinkRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "ok",
                        "schema": {
                            "$ref": "#/definitions/entity.LinkDto"
                        }
                    }
                }
            }
        },
        "/module": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "module"
                ],
                "parameters": [
                    {
                        "description": "moduleRequest",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.ModuleRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "created",
                        "schema": {
                            "$ref": "#/definitions/entity.ModuleDto"
                        }
                    }
                }
            }
        },
        "/module/{courseId}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "module"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "course id",
                        "name": "courseId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "$ref": "#/definitions/entity.CourseModules"
                        }
                    }
                }
            }
        },
        "/slide": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "slide"
                ],
                "parameters": [
                    {
                        "description": "создание слайда",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.SlideRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "ok",
                        "schema": {
                            "$ref": "#/definitions/entity.Slide"
                        }
                    }
                }
            }
        },
        "/slide/{courseId}/{moduleNameEng}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "slide"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "course id",
                        "name": "courseId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "moduleName",
                        "name": "moduleNameEng",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "ok",
                        "schema": {
                            "$ref": "#/definitions/entity.ModuleSlides"
                        }
                    }
                }
            }
        },
        "/teacher/{userId}/group/{groupId}": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "group"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "user id",
                        "name": "userId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "group id",
                        "name": "groupId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "created",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "entity.Collaborator": {
            "type": "object",
            "properties": {
                "course_id": {
                    "type": "integer"
                },
                "user_email": {
                    "type": "string"
                }
            }
        },
        "entity.Course": {
            "type": "object",
            "properties": {
                "author_id": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "logo": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "entity.CourseModules": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "logo": {
                    "type": "string"
                },
                "modules": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.ModuleDto"
                    }
                },
                "name": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "entity.CourseTitle": {
            "type": "object",
            "properties": {
                "author": {
                    "$ref": "#/definitions/entity.UserDto"
                },
                "collaborators": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.UserDto"
                    }
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "links": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.LinkDto"
                    }
                },
                "logo": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "entity.GroupName": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        },
        "entity.LinkDto": {
            "type": "object",
            "properties": {
                "link": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "entity.LinkRequest": {
            "type": "object",
            "properties": {
                "course_id": {
                    "type": "integer"
                },
                "link": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "entity.ModuleDto": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "name_eng": {
                    "type": "string"
                }
            }
        },
        "entity.ModuleRequest": {
            "type": "object",
            "properties": {
                "course_id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "entity.ModuleSlides": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "name_eng": {
                    "type": "string"
                },
                "slides": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.SlideDto"
                    }
                }
            }
        },
        "entity.PasswordUpdate": {
            "type": "object",
            "properties": {
                "current": {
                    "type": "string"
                },
                "new": {
                    "type": "string"
                }
            }
        },
        "entity.RefreshToken": {
            "type": "object",
            "properties": {
                "refresh_token": {
                    "type": "string"
                }
            }
        },
        "entity.SignIn": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "entity.SignUp": {
            "type": "object",
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
        "entity.Slide": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "module": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "nameEng": {
                    "type": "string"
                }
            }
        },
        "entity.SlideDto": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "name_eng": {
                    "type": "string"
                }
            }
        },
        "entity.SlideRequest": {
            "type": "object",
            "properties": {
                "course_id": {
                    "type": "integer"
                },
                "module_name_eng": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "entity.Tokens": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "expires_in": {
                    "type": "string"
                },
                "refresh_token": {
                    "type": "string"
                }
            }
        },
        "entity.UserDto": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "entity.UserName": {
            "type": "object",
            "properties": {
                "name": {
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

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/v1",
	Schemes:          []string{},
	Title:            "без юлерна",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
