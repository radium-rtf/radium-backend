// Package wave Code generated by swaggo/swag. DO NOT EDIT
package wave

import "github.com/swaggo/swag"

const docTemplatewave = `{
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
        "/v1/dialogue": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "dialogue"
                ],
                "parameters": [
                    {
                        "description": "Данные о реципиенте",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/create.DialogueCreate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "already exists",
                        "schema": {
                            "$ref": "#/definitions/model.Dialogue"
                        }
                    },
                    "201": {
                        "description": "created",
                        "schema": {
                            "$ref": "#/definitions/model.Dialogue"
                        }
                    }
                }
            }
        },
        "/v1/message": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "message"
                ],
                "parameters": [
                    {
                        "description": "Сообщение и направление",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/send.MessageSend"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "sent",
                        "schema": {
                            "$ref": "#/definitions/model.Message"
                        }
                    }
                }
            }
        },
        "/v1/messages/{chatId}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "message"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID группы/диалога",
                        "name": "chatId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": " ",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Message"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "create.DialogueCreate": {
            "type": "object",
            "required": [
                "userId"
            ],
            "properties": {
                "userId": {
                    "type": "string"
                }
            }
        },
        "model.Content": {
            "type": "object",
            "properties": {
                "text": {
                    "type": "string"
                }
            }
        },
        "model.Dialogue": {
            "type": "object",
            "properties": {
                "firstUserId": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "secondUserId": {
                    "type": "string"
                },
                "settings": {
                    "$ref": "#/definitions/model.DialogueSettings"
                }
            }
        },
        "model.DialogueSettings": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                }
            }
        },
        "model.Message": {
            "type": "object",
            "properties": {
                "content": {
                    "$ref": "#/definitions/model.Content"
                },
                "id": {
                    "type": "string"
                },
                "parentMessageId": {
                    "type": "string"
                },
                "senderId": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "send.MessageSend": {
            "type": "object",
            "required": [
                "chatId",
                "content"
            ],
            "properties": {
                "chatId": {
                    "type": "string"
                },
                "content": {
                    "$ref": "#/definitions/model.Content"
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

// SwaggerInfowave holds exported Swagger Info so clients can modify it
var SwaggerInfowave = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "wave",
	Description:      "",
	InfoInstanceName: "wave",
	SwaggerTemplate:  docTemplatewave,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfowave.InstanceName(), SwaggerInfowave)
}
