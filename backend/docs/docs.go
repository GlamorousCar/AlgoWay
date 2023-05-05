// Code generated by swaggo/swag. DO NOT EDIT.

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
        "/": {
            "get": {
                "description": "Страница с проверкой подключения к бд (тестовая)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "main"
                ],
                "summary": "Главная страница",
                "responses": {
                    "200": {
                        "description": "Успешное подключение -  версия базы данных",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/auth/login": {
            "post": {
                "description": "после успешной авторизации возвращается токен",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "авторизация пользователя",
                "parameters": [
                    {
                        "description": "Account Info",
                        "name": "message",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.LoginUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Token"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/auth/register": {
            "post": {
                "description": "после успешной регистрации возвращается сообщение об успешном выполнении",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "регистрация пользователя",
                "parameters": [
                    {
                        "description": "User Info",
                        "name": "message",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.RawUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/check_task": {
            "post": {
                "description": "Проверяет пользовательское решение какой-либо задачи",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "check_system"
                ],
                "summary": "Проверка решения",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Токен авторизации",
                        "name": "user_token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Solution info",
                        "name": "message",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CheckTaskRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Verdict"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/task": {
            "get": {
                "description": "Получение заданий к определенному алгоритму. Пока есть ошибка связанная с возвращаемым полем is_solved, будет исправлено скоро",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "main"
                ],
                "summary": "Список задач по алгоритму",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Получение задач по id алгоритма",
                        "name": "algo_id",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Возвращается список задач",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Task"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/themes/menu": {
            "get": {
                "description": "Получение всех тем и подтем меню сайта",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "main"
                ],
                "summary": "Меню сайта",
                "responses": {
                    "200": {
                        "description": "Возвращаются темы и подтемы",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.ThemeMenu"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/theory": {
            "get": {
                "description": "Получение теории к алгоритму по его id",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "main"
                ],
                "summary": "Теория к алгоритму",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Получение задачи по id алгоритма",
                        "name": "algo_id",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.AlgorithmTheory"
                        }
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Algorithm": {
            "type": "object",
            "properties": {
                "algorithm_id": {
                    "type": "integer"
                },
                "description": {
                    "type": "string"
                },
                "position": {
                    "type": "integer"
                },
                "theme_id": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "models.AlgorithmTheory": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string",
                    "example": "#Ligula. Vivamus nec ultrices. Himenaeos. \n Semper lorem volutpat urna at. Tincidunt purus vulputate neque velit senectus. Convallis hendrerit."
                },
                "id": {
                    "type": "integer",
                    "example": 4
                },
                "title": {
                    "type": "string",
                    "example": "civil"
                }
            }
        },
        "models.CheckTaskRequest": {
            "type": "object",
            "properties": {
                "lang": {
                    "type": "string"
                },
                "source_code": {
                    "type": "string"
                },
                "task_id": {
                    "type": "integer"
                }
            }
        },
        "models.LoginUser": {
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
        "models.RawUser": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "login": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "models.Task": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "is_solved": {
                    "type": "boolean"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "models.ThemeMenu": {
            "type": "object",
            "properties": {
                "algorithms": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Algorithm"
                    }
                },
                "position": {
                    "type": "integer"
                },
                "theme_id": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "models.Token": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "models.Verdict": {
            "type": "object",
            "properties": {
                "abbr": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "127.0.0.1:4000",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Swagger Algoway API",
	Description:      "Документация для нашего студенческого проекта - сервиса для изучения алгоритмов",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
