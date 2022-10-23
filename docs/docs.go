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
        "/er": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "emergency-rooms"
                ],
                "summary": "get all emergency rooms",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "0-indexed page number, 0 is assumed when omitted",
                        "name": "page",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/routes.GetMultipleERsResponse"
                        }
                    },
                    "501": {
                        "description": "Not Implemented",
                        "schema": {
                            "$ref": "#/definitions/routes.HTTPErrorResponse"
                        }
                    }
                }
            },
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "emergency-rooms"
                ],
                "summary": "create a new emergency room",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer: \u003cTOKEN\u003e",
                        "name": "authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "ER to add",
                        "name": "emergency-room",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/routes.PutERRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/routes.GetSingleERResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/routes.HTTPErrorResponse"
                        }
                    },
                    "501": {
                        "description": "Not Implemented",
                        "schema": {
                            "$ref": "#/definitions/routes.HTTPErrorResponse"
                        }
                    }
                }
            }
        },
        "/er/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "emergency-rooms"
                ],
                "summary": "get an emergency room by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Emergency Room's ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/routes.GetSingleERResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/routes.HTTPErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "emergency-rooms"
                ],
                "summary": "delete an emergency room by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer: \u003cTOKEN\u003e",
                        "name": "authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Emergency Room's ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/routes.StatusResponse"
                        }
                    },
                    "501": {
                        "description": "Not Implemented",
                        "schema": {
                            "$ref": "#/definitions/routes.HTTPErrorResponse"
                        }
                    }
                }
            },
            "patch": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "emergency-rooms"
                ],
                "summary": "update an emergency room by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer: \u003cTOKEN\u003e",
                        "name": "authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Emergency Room's ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/routes.GetSingleERResponse"
                        }
                    },
                    "501": {
                        "description": "Not Implemented",
                        "schema": {
                            "$ref": "#/definitions/routes.HTTPErrorResponse"
                        }
                    }
                }
            }
        },
        "/healthz": {
            "get": {
                "description": "can be used for health checks",
                "summary": "health check route",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/routes.HealthzResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/routes.HealthzResponse"
                        }
                    }
                }
            }
        },
        "/version": {
            "get": {
                "description": "return the version",
                "summary": "version route",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.DepartmentBase": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "models.Point": {
            "type": "object",
            "required": [
                "lat",
                "long"
            ],
            "properties": {
                "lat": {
                    "type": "number"
                },
                "long": {
                    "type": "number"
                }
            }
        },
        "routes.GetMultipleERsResponse": {
            "type": "object",
            "properties": {
                "emergencyRooms": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/routes.GetSingleERResponse"
                    }
                },
                "lastPage": {
                    "type": "boolean"
                },
                "numPages": {
                    "type": "integer"
                },
                "page": {
                    "type": "integer"
                },
                "pageSize": {
                    "type": "integer"
                },
                "totalSite": {
                    "type": "integer"
                }
            }
        },
        "routes.GetSingleERResponse": {
            "type": "object",
            "required": [
                "displayableAddress",
                "location",
                "name"
            ],
            "properties": {
                "departments": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.DepartmentBase"
                    }
                },
                "displayableAddress": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "location": {
                    "$ref": "#/definitions/models.Point"
                },
                "name": {
                    "type": "string"
                },
                "open": {
                    "type": "boolean"
                },
                "utilization": {
                    "type": "integer"
                }
            }
        },
        "routes.HTTPError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "routes.HTTPErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "$ref": "#/definitions/routes.HTTPError"
                }
            }
        },
        "routes.HealthzResponse": {
            "type": "object",
            "properties": {
                "server": {
                    "type": "string"
                }
            }
        },
        "routes.PutERRequest": {
            "type": "object",
            "required": [
                "displayableAddress",
                "location",
                "name"
            ],
            "properties": {
                "departments": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "displayableAddress": {
                    "type": "string"
                },
                "location": {
                    "$ref": "#/definitions/models.Point"
                },
                "name": {
                    "type": "string"
                },
                "open": {
                    "type": "boolean"
                },
                "utilization": {
                    "type": "integer"
                }
            }
        },
        "routes.StatusResponse": {
            "type": "object",
            "properties": {
                "ok": {
                    "type": "boolean"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
