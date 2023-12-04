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
        "/api/containers": {
            "get": {
                "description": "Возвращает все доступные контейнеры с опциональной фильтрацией по типу",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Контейнеры"
                ],
                "summary": "Получить все контейнеры",
                "parameters": [
                    {
                        "type": "string",
                        "description": "тип контейнера для фильтрации",
                        "name": "type",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schemes.GetAllContainersResponse"
                        }
                    }
                }
            }
        },
        "/api/containers/": {
            "post": {
                "description": "Добавить новый контейнер",
                "consumes": [
                    "multipart/form-data"
                ],
                "tags": [
                    "Контейнеры"
                ],
                "summary": "Добавить контейнер",
                "parameters": [
                    {
                        "type": "file",
                        "description": "Изображение контейнера",
                        "name": "image",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "Маркировка",
                        "name": "marking",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Тип",
                        "name": "type",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Длина",
                        "name": "length",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Высота",
                        "name": "height",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Ширина",
                        "name": "width",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Груз",
                        "name": "cargo",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Вес",
                        "name": "weight",
                        "in": "formData",
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
        "/api/containers/{container_id}": {
            "get": {
                "description": "Возвращает более подробную информацию об одном контейнере",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Контейнеры"
                ],
                "summary": "Получить один контейнер",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id контейнера",
                        "name": "container_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ds.Container"
                        }
                    }
                }
            },
            "put": {
                "description": "Изменить данные полей о контейнере",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Контейнеры"
                ],
                "summary": "Изменить котейнер",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Идентификатор контейнера",
                        "name": "container_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Маркировка",
                        "name": "marking",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "Тип",
                        "name": "type",
                        "in": "formData"
                    },
                    {
                        "type": "integer",
                        "description": "Длина",
                        "name": "length",
                        "in": "formData"
                    },
                    {
                        "type": "integer",
                        "description": "Высота",
                        "name": "height",
                        "in": "formData"
                    },
                    {
                        "type": "integer",
                        "description": "Ширина",
                        "name": "width",
                        "in": "formData"
                    },
                    {
                        "type": "file",
                        "description": "Изображение контейнера",
                        "name": "image",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "Груз",
                        "name": "cargo",
                        "in": "formData"
                    },
                    {
                        "type": "integer",
                        "description": "Вес",
                        "name": "weight",
                        "in": "formData"
                    }
                ],
                "responses": {}
            },
            "delete": {
                "description": "Удаляет контейнер по id",
                "tags": [
                    "Контейнеры"
                ],
                "summary": "Удалить контейнер",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id контейнера",
                        "name": "container_id",
                        "in": "path",
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
        "/api/containers/{container_id}/add_to_transportation": {
            "post": {
                "description": "Добавить выбранный контейнер в черновик перевозки",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Контейнеры"
                ],
                "summary": "Добавить в перевозку",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id контейнера",
                        "name": "container_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schemes.AllContainersResponse"
                        }
                    }
                }
            }
        },
        "/api/transportations": {
            "get": {
                "description": "Возвращает все перевозки с фильтрацией по статусу и дате формирования",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Перевозки"
                ],
                "summary": "Получить все перевозки",
                "parameters": [
                    {
                        "type": "string",
                        "description": "статус перевозки",
                        "name": "status",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "начальная дата формирования",
                        "name": "formation_date_start",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "конечная дата формирвания",
                        "name": "formation_date_end",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schemes.AllTransportationsResponse"
                        }
                    }
                }
            }
        },
        "/api/transportations/user_confirm": {
            "put": {
                "description": "Сформировать или удалить перевозку перевозку пользователем",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Перевозки"
                ],
                "summary": "Сформировать перевозку",
                "parameters": [
                    {
                        "description": "подтвердить",
                        "name": "confirm",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "boolean"
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
        "/api/transportations/{transportation_id}": {
            "get": {
                "description": "Возвращает подробную информацию о перевозке и её составе",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Перевозки"
                ],
                "summary": "Получить одну перевозку",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id перевозки",
                        "name": "transportation_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schemes.TransportationResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "Позволяет изменить транспорт перевозки и возвращает обновлённые данные",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Перевозки"
                ],
                "summary": "Указать транспорт перевозки",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id перевозки",
                        "name": "transportation_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Транспорт",
                        "name": "transport",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/app.SwaggerUpdateTransportationRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schemes.UpdateTransportationResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Удаляет первозку по id",
                "tags": [
                    "Перевозки"
                ],
                "summary": "Удалить перевозку",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id перевозки",
                        "name": "transportation_id",
                        "in": "path",
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
        "/api/transportations/{transportation_id}/delete_container/{container_id}": {
            "delete": {
                "description": "Удалить контейнер из перевозки",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Перевозки"
                ],
                "summary": "Удалить контейнер из перевозки",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id перевозки",
                        "name": "transportation_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "id контейнера",
                        "name": "container_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schemes.AllContainersResponse"
                        }
                    }
                }
            }
        },
        "/api/transportations/{transportation_id}/moderator_confirm": {
            "put": {
                "description": "Подтвердить или отменить перевозку модератором",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Перевозки"
                ],
                "summary": "Подтвердить перевозку",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id перевозки",
                        "name": "transportation_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "подтвердить",
                        "name": "confirm",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "boolean"
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
        "/auth/login": {
            "post": {
                "description": "Авторизует пользователя по логиню, паролю и отдаёт jwt токен для дальнейших запросов",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Авторизация"
                ],
                "summary": "Авторизация",
                "parameters": [
                    {
                        "description": "login and password",
                        "name": "user_credentials",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schemes.LoginReq"
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
        "/auth/loguot": {
            "post": {
                "description": "Выход из аккаунта",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Авторизация"
                ],
                "summary": "Выйти из аккаунта",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/auth/sign_up": {
            "post": {
                "description": "Регистрация нового пользователя",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Авторизация"
                ],
                "summary": "Регистрация",
                "parameters": [
                    {
                        "description": "login and password",
                        "name": "user_credentials",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schemes.RegisterReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schemes.RegisterResp"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "app.SwaggerUpdateTransportationRequest": {
            "type": "object",
            "properties": {
                "transport": {
                    "type": "string"
                }
            }
        },
        "ds.Container": {
            "type": "object",
            "required": [
                "cargo",
                "height",
                "length",
                "marking",
                "type",
                "weight",
                "width"
            ],
            "properties": {
                "cargo": {
                    "type": "string",
                    "maxLength": 50
                },
                "height": {
                    "type": "integer"
                },
                "image_url": {
                    "type": "string"
                },
                "length": {
                    "type": "integer"
                },
                "marking": {
                    "type": "string",
                    "maxLength": 11
                },
                "type": {
                    "type": "string",
                    "maxLength": 50
                },
                "uuid": {
                    "type": "string"
                },
                "weight": {
                    "type": "integer"
                },
                "width": {
                    "type": "integer"
                }
            }
        },
        "schemes.AllContainersResponse": {
            "type": "object",
            "properties": {
                "containers": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ds.Container"
                    }
                }
            }
        },
        "schemes.AllTransportationsResponse": {
            "type": "object",
            "properties": {
                "transportations": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/schemes.TransportationOutput"
                    }
                }
            }
        },
        "schemes.GetAllContainersResponse": {
            "type": "object",
            "properties": {
                "containers": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ds.Container"
                    }
                },
                "draft_transportation": {
                    "$ref": "#/definitions/schemes.TransportationShort"
                }
            }
        },
        "schemes.LoginReq": {
            "type": "object",
            "required": [
                "login",
                "password"
            ],
            "properties": {
                "login": {
                    "type": "string",
                    "maxLength": 30
                },
                "password": {
                    "type": "string",
                    "maxLength": 30
                }
            }
        },
        "schemes.RegisterReq": {
            "type": "object",
            "required": [
                "login",
                "password"
            ],
            "properties": {
                "login": {
                    "type": "string",
                    "maxLength": 30
                },
                "password": {
                    "type": "string",
                    "maxLength": 30
                }
            }
        },
        "schemes.RegisterResp": {
            "type": "object",
            "properties": {
                "ok": {
                    "type": "boolean"
                }
            }
        },
        "schemes.TransportationOutput": {
            "type": "object",
            "properties": {
                "completion_date": {
                    "type": "string"
                },
                "creation_date": {
                    "type": "string"
                },
                "customer": {
                    "type": "string"
                },
                "formation_date": {
                    "type": "string"
                },
                "moderator": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "transport": {
                    "type": "string"
                },
                "uuid": {
                    "type": "string"
                }
            }
        },
        "schemes.TransportationResponse": {
            "type": "object",
            "properties": {
                "containers": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ds.Container"
                    }
                },
                "transportation": {
                    "$ref": "#/definitions/schemes.TransportationOutput"
                }
            }
        },
        "schemes.TransportationShort": {
            "type": "object",
            "properties": {
                "container_count": {
                    "type": "integer"
                },
                "uuid": {
                    "type": "string"
                }
            }
        },
        "schemes.UpdateTransportationResponse": {
            "type": "object",
            "properties": {
                "transportation": {
                    "$ref": "#/definitions/schemes.TransportationOutput"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "127.0.0.1:8080",
	BasePath:         "/",
	Schemes:          []string{"http"},
	Title:            "Container loginstics",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
