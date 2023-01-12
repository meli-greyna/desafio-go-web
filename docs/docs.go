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
        "termsOfService": "https://fake.url/",
        "contact": {
            "name": "API Support",
            "url": "https://real.url"
        },
        "license": {
            "name": "MIT",
            "url": "https://opensource.org/licenses/MIT"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/ticket/getAverage/:dest": {
            "get": {
                "description": "Gets the percentage of total tickets bought over the last day with destination of given country",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tickets"
                ],
                "summary": "Destination percentage",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/ticket/getByCountry/:dest": {
            "get": {
                "description": "Gets the number of tickets bought for a trip to a given country passed by URL parameter",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tickets"
                ],
                "summary": "Get tickets by country",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost/8080",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "MELI Bootcamp API",
	Description:      "This API handles airline tickets.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}