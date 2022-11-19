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
        "/departments": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "departments"
                ],
                "summary": "get all departments",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "0-indexed page number, 0 is assumed when omitted",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "page size, 100 is assumed when omitted",
                        "name": "page_size",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/routes.GetDepartmentsResponse"
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
            "put": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "departments"
                ],
                "summary": "create new department",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer: \u003cTOKEN\u003e",
                        "name": "authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Dep. to add",
                        "name": "department",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/routes.CreateDepartmentRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/routes.SingleDepartmentResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/routes.HTTPErrorResponse"
                        }
                    }
                }
            }
        },
        "/departments/{id}": {
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "departments"
                ],
                "summary": "delete a department",
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
                        "description": "department id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request",
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
                    "departments"
                ],
                "summary": "update an department by id",
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
                        "description": "Department's ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "ER to update",
                        "name": "department",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/routes.UpdateDepartmentRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/routes.HTTPErrorResponse"
                        }
                    }
                }
            }
        },
        "/emergency-room": {
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
                    },
                    {
                        "type": "integer",
                        "description": "page size, 100 is assumed when omitted",
                        "name": "page_size",
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
                    "400": {
                        "description": "Bad Request",
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
                    }
                }
            }
        },
        "/emergency-room/{id}": {
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
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request",
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
                    },
                    {
                        "description": "ER to update",
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
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request",
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
        "/organizations": {
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "organizations"
                ],
                "summary": "create a new organization",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer: \u003cTOKEN\u003e",
                        "name": "authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Org to add",
                        "name": "organization",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/routes.CreateOrgRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/routes.GetSingleOrgResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/routes.HTTPErrorResponse"
                        }
                    }
                }
            }
        },
        "/users": {
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "create a new user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer: \u003cTOKEN\u003e",
                        "name": "authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "user to add",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/routes.CreateUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/routes.CreateUserResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/routes.HTTPErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/routes.HTTPErrorResponse"
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
                    "type": "string",
                    "example": "pediatric surgery"
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
                    "type": "number",
                    "example": 51.9893434
                },
                "long": {
                    "type": "number",
                    "example": 7.62613583
                }
            }
        },
        "routes.CreateDepartmentRequest": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "name": {
                    "type": "string",
                    "example": "pediatric surgery"
                }
            }
        },
        "routes.CreateOrgRequest": {
            "type": "object",
            "required": [
                "contactEmail",
                "longName"
            ],
            "properties": {
                "contactEmail": {
                    "type": "string",
                    "example": "example@helpwave.de"
                },
                "longName": {
                    "type": "string",
                    "example": "Uniklinikum Münster"
                },
                "shortName": {
                    "type": "string",
                    "example": "UKM"
                }
            }
        },
        "routes.CreateUserRequest": {
            "type": "object",
            "required": [
                "email",
                "fullName",
                "password"
            ],
            "properties": {
                "admin": {
                    "type": "boolean"
                },
                "email": {
                    "type": "string",
                    "example": "example@helpwave.de"
                },
                "fullName": {
                    "type": "string",
                    "example": "Some Name"
                },
                "organization": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "maxLength": 100,
                    "minLength": 6,
                    "example": "hunter2"
                }
            }
        },
        "routes.CreateUserResponse": {
            "type": "object",
            "properties": {
                "userID": {
                    "type": "string"
                }
            }
        },
        "routes.GetDepartmentsResponse": {
            "type": "object",
            "properties": {
                "departments": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.DepartmentBase"
                    }
                },
                "lastPage": {
                    "type": "boolean"
                },
                "page": {
                    "type": "integer"
                },
                "pageSize": {
                    "type": "integer"
                },
                "totalSize": {
                    "type": "integer"
                }
            }
        },
        "routes.GetMultipleERsResponse": {
            "type": "object",
            "properties": {
                "emergencyRooms": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "lastPage": {
                    "type": "boolean"
                },
                "page": {
                    "type": "integer"
                },
                "pageSize": {
                    "type": "integer"
                },
                "totalSize": {
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
                    "type": "string",
                    "example": "Kardinal-von-Galen-Ring 10, 48149 Münster, Germany"
                },
                "id": {
                    "type": "string"
                },
                "location": {
                    "$ref": "#/definitions/models.Point"
                },
                "name": {
                    "type": "string",
                    "example": "Uniklinikum Münster"
                },
                "open": {
                    "type": "boolean"
                },
                "utilization": {
                    "type": "integer",
                    "example": 4
                }
            }
        },
        "routes.GetSingleOrgResponse": {
            "type": "object",
            "properties": {
                "avatarUrl": {
                    "type": "string"
                },
                "contactEmail": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "longName": {
                    "type": "string"
                },
                "shortName": {
                    "type": "string"
                }
            }
        },
        "routes.HTTPError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 500
                },
                "message": {
                    "type": "string",
                    "example": "Some complicated error message here"
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
                    "type": "string",
                    "example": "Kardinal-von-Galen-Ring 10, 48149 Münster, Germany"
                },
                "location": {
                    "$ref": "#/definitions/models.Point"
                },
                "name": {
                    "type": "string",
                    "example": "Uniklinikum Münster"
                },
                "open": {
                    "type": "boolean"
                },
                "utilization": {
                    "type": "integer",
                    "example": 4
                }
            }
        },
        "routes.SingleDepartmentResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string",
                    "example": "pediatric surgery"
                }
            }
        },
        "routes.UpdateDepartmentRequest": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string",
                    "example": "pediatric surgery"
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
