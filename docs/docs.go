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
        "/accounts": {
            "get": {
                "description": "Retrieve a list of all bank accounts",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "accounts"
                ],
                "summary": "List all accounts",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/github_com_anuraj2023_bank-account-management-be_internal_models.Account"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new bank account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "accounts"
                ],
                "summary": "Create a new account",
                "parameters": [
                    {
                        "description": "Account details",
                        "name": "account",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_anuraj2023_bank-account-management-be_internal_models.Account"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/github_com_anuraj2023_bank-account-management-be_internal_models.Account"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            }
        },
        "/health": {
            "get": {
                "description": "check if the web service is healthy",
                "produces": [
                    "application/json"
                ],
                "summary": "Check Health",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/internal_api_handlers.HealthResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "echo.HTTPError": {
            "type": "object",
            "properties": {
                "message": {}
            }
        },
        "github_com_anuraj2023_bank-account-management-be_internal_models.Account": {
            "type": "object",
            "properties": {
                "account_name": {
                    "type": "string",
                    "example": "Tom Cruise"
                },
                "account_number": {
                    "type": "string",
                    "example": "1234567890"
                },
                "address": {
                    "type": "string",
                    "example": "123 Becker Str, Berlin, DE 12345"
                },
                "amount": {
                    "type": "number",
                    "example": 1000.5
                },
                "iban": {
                    "type": "string",
                    "example": "DE89370400440532013000"
                },
                "type": {
                    "enum": [
                        "sending",
                        "receiving"
                    ],
                    "allOf": [
                        {
                            "$ref": "#/definitions/github_com_anuraj2023_bank-account-management-be_internal_models.AccountType"
                        }
                    ],
                    "example": "sending"
                }
            }
        },
        "github_com_anuraj2023_bank-account-management-be_internal_models.AccountType": {
            "type": "string",
            "enum": [
                "sending",
                "receiving"
            ],
            "x-enum-varnames": [
                "AccountTypeSending",
                "AccountTypeReceiving"
            ]
        },
        "internal_api_handlers.HealthResponse": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "string",
                    "example": "healthy"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "bank-account-management-be.onrender.com",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Swagger - Bank Account Management APIs",
	Description:      "This projects deals with creating and fetching bank accounts",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
