// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/health": {
            "get": {
                "description": "returns health status of the serivce",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "health"
                ],
                "summary": "HealthCheck",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/health.StatusResponse"
                        }
                    },
                    "503": {
                        "description": "Service Unavailable",
                        "schema": {
                            "$ref": "#/definitions/health.StatusResponse"
                        }
                    }
                }
            }
        },
        "/metrics": {
            "get": {
                "description": "returns all metrics descriptions",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "metrics"
                ],
                "summary": "GetMetrics",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/metrics.MetricDescriptor"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "adds new metric based on provided descriptor",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "metrics"
                ],
                "summary": "AddMetric",
                "parameters": [
                    {
                        "description": "Metric descriptor",
                        "name": "message",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/metrics.MetricDescriptor"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "400": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/metrics/:type/:name": {
            "delete": {
                "description": "removes metric with all its data",
                "tags": [
                    "metrics"
                ],
                "summary": "RemoveMetric",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Metric type",
                        "name": "type",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Metric name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "400": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/metrics/:type/:name/data": {
            "get": {
                "description": "returns data points for metric in given time range",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "metrics"
                ],
                "summary": "GetData",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Metric type",
                        "name": "type",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Metric name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Begin timestamp in RFC3339 format",
                        "name": "from",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "End timestamp in RFC3339 format",
                        "name": "to",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/metrics.MetricDataPoint"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad parameters provided by user",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Data retrieval field",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "adds new data point to the metric",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "metrics"
                ],
                "summary": "AddData",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Metric type",
                        "name": "type",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Metric name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Metric descriptor with value",
                        "name": "message",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/metrics.Value"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "400": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "health.StatusResponse": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "string"
                }
            }
        },
        "metrics.MetricDataPoint": {
            "type": "object",
            "properties": {
                "time": {
                    "type": "string"
                },
                "value": {
                    "type": "number"
                }
            }
        },
        "metrics.MetricDescriptor": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "metrics.Value": {
            "type": "object",
            "properties": {
                "value": {
                    "type": "number"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "localhost:8080",
	BasePath:    "/api/v1/",
	Schemes:     []string{},
	Title:       "Metrics API",
	Description: "metrics api supports storage and retrieval of various time-series metrics",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
