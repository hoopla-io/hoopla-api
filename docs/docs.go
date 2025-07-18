// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "email": "davronbekov.o@itv.uz"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v1/auth/confirm-sms": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "parameters": [
                    {
                        "description": "Confirm Sms Request",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth_request.ConfirmSmsRequest"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/v1/auth/login": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "parameters": [
                    {
                        "description": "Login Request",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth_request.LoginRequest"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/v1/auth/resend-sms": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "parameters": [
                    {
                        "description": "Resend Sms Request",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth_request.ResendSmsRequest"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/v1/shops/near-shops": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Shops"
                ],
                "parameters": [
                    {
                        "type": "number",
                        "name": "lat",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "number",
                        "name": "long",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "name",
                        "in": "query"
                    }
                ],
                "responses": {}
            }
        },
        "/v1/shops/shop": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Shops"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "name": "shopId",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/v1/subscriptions": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Subscriptions"
                ],
                "responses": {}
            }
        },
        "/v1/subscriptions/buy": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Subscriptions"
                ],
                "parameters": [
                    {
                        "description": "Buy Subscription",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/subscriptions_request.BuySubscriptionRequest"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/v1/user/deactivate": {
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "responses": {}
            }
        },
        "/v1/user/generate-qr-code": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "QR"
                ],
                "responses": {}
            }
        },
        "/v1/user/get-me": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "responses": {}
            }
        },
        "/v1/user/logout": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "parameters": [
                    {
                        "description": "Logout from an account",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user_request.LogoutRequest"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/v1/user/orders/create": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User/Orders"
                ],
                "parameters": [
                    {
                        "description": "New order",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user_orders_request.CreateRequest"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/v1/user/orders/drinks-stat": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "responses": {}
            }
        },
        "/v1/user/orders/orders-list": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User/Orders"
                ],
                "responses": {}
            }
        },
        "/v1/user/pay/services": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User/Pay"
                ],
                "responses": {}
            }
        },
        "/v1/user/pay/top-up": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User/Pay"
                ],
                "parameters": [
                    {
                        "type": "number",
                        "name": "amount",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/v1/user/refresh-token": {
            "patch": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "name": "refreshToken",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/vendors/iiko/webhook": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "IIKO"
                ],
                "parameters": [
                    {
                        "description": "Webhook for iiko",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/vendors_poster_request.WebhookRequest"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/vendors/poster/oauth": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Poster"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "name": "account",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "code",
                        "in": "query"
                    }
                ],
                "responses": {}
            }
        },
        "/vendors/poster/webhook": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Poster"
                ],
                "parameters": [
                    {
                        "description": "Webhook for poster",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/vendors_poster_request.WebhookRequest"
                        }
                    }
                ],
                "responses": {}
            }
        }
    },
    "definitions": {
        "auth_request.ConfirmSmsRequest": {
            "type": "object",
            "required": [
                "code",
                "sessionId"
            ],
            "properties": {
                "code": {
                    "type": "integer"
                },
                "sessionId": {
                    "type": "string"
                }
            }
        },
        "auth_request.LoginRequest": {
            "type": "object",
            "required": [
                "phoneNumber"
            ],
            "properties": {
                "phoneNumber": {
                    "type": "string"
                }
            }
        },
        "auth_request.ResendSmsRequest": {
            "type": "object",
            "required": [
                "sessionId"
            ],
            "properties": {
                "sessionId": {
                    "type": "string"
                }
            }
        },
        "subscriptions_request.BuySubscriptionRequest": {
            "type": "object",
            "properties": {
                "subscriptionId": {
                    "type": "integer"
                }
            }
        },
        "user_orders_request.CreateRequest": {
            "type": "object",
            "properties": {
                "drink_id": {
                    "type": "integer"
                },
                "shop_id": {
                    "type": "integer"
                }
            }
        },
        "user_request.LogoutRequest": {
            "type": "object"
        },
        "vendors_poster_request.WebhookRequest": {
            "type": "object",
            "properties": {
                "account": {
                    "type": "string"
                },
                "account_number": {
                    "type": "string"
                },
                "action": {
                    "type": "string"
                },
                "object": {
                    "type": "string"
                },
                "object_id": {
                    "type": "integer"
                },
                "time": {
                    "type": "string"
                },
                "verify": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "api.hoopla.uz",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Hoopla | Api",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
