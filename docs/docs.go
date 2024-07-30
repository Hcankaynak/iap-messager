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
        "/automatic-message-sender": {
            "post": {
                "description": "Start/Stop Automatic Message Sender If you send a request with start=true, the automatic message sender",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Start/Stop Automatic Message Sender"
                ],
                "summary": "Start/Stop Automatic Message Sender",
                "parameters": [
                    {
                        "description": "Automatic Message Sender Payload",
                        "name": "messageSender",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.AutomaticMessageSender"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Response of start/stop automatic message sender",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/sent-messages": {
            "get": {
                "description": "Get Sent Messages",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Get Sent Messages"
                ],
                "summary": "Get Sent Messages",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/messages.MessageDTO"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handlers.AutomaticMessageSender": {
            "type": "object",
            "properties": {
                "start": {
                    "type": "boolean"
                }
            }
        },
        "messages.MessageDTO": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "phone_number": {
                    "type": "string"
                },
                "sending_status": {
                    "type": "boolean"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/v2",
	Schemes:          []string{},
	Title:            "IAP Messager API",
	Description:      "This is a sample server",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}