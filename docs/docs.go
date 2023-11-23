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
        "/v1/auth/login/refresh-token": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "授权"
                ],
                "summary": "刷新令牌(延长时间)",
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":true,\"msg\":\"success\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/auth/login/using-password": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "授权"
                ],
                "summary": "使用密码登录",
                "parameters": [
                    {
                        "type": "string",
                        "description": "填手机号/email/用户名",
                        "name": "login_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "密码",
                        "name": "password",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "图形验证码id",
                        "name": "captcha_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "图形验证码",
                        "name": "captcha_answer",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":true,\"msg\":\"success\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/auth/login/using-phone": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "授权"
                ],
                "summary": "使用手机验证码进行登录",
                "parameters": [
                    {
                        "type": "string",
                        "description": "手机号",
                        "name": "phone",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "短信验证码",
                        "name": "verify_code",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":true,\"msg\":\"success\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/auth/password-reset/using-phone": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "授权"
                ],
                "summary": "找回密码(通过手机号+短信验证码)",
                "parameters": [
                    {
                        "type": "string",
                        "description": "手机号",
                        "name": "phone",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "短信验证码",
                        "name": "verify_code",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":true,\"msg\":\"success\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/auth/signup/email/exist": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "授权"
                ],
                "summary": "校验邮箱是否已注册",
                "parameters": [
                    {
                        "type": "string",
                        "description": "邮箱",
                        "name": "email",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":true,\"msg\":\"success\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/auth/signup/phone/exist": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "授权"
                ],
                "summary": "校验手机号是否已注册",
                "parameters": [
                    {
                        "type": "string",
                        "description": "手机号码",
                        "name": "phone",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"exist\":true,\"msg\":\"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/auth/signup/using-phone": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "授权"
                ],
                "summary": "使用手机验证码进行注册",
                "parameters": [
                    {
                        "type": "string",
                        "description": "手机号",
                        "name": "phone",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "name",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "密码",
                        "name": "password",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "确认密码",
                        "name": "password_confirm",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "短信验证码",
                        "name": "verify_code",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":true,\"msg\":\"success\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/categories": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "分类"
                ],
                "summary": "分类列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "排序(id/created_at/updated_at,默认id)",
                        "name": "sort",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "排序规则(仅支持 asc（正序）,desc（倒序）)",
                        "name": "order",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "每页条数(介于 2~100 之间)",
                        "name": "per_page",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "当前页",
                        "name": "page",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":true,\"msg\":\"success\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "分类"
                ],
                "summary": "创建分类",
                "parameters": [
                    {
                        "type": "string",
                        "description": "分类名",
                        "name": "name",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "描述",
                        "name": "description",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":true,\"msg\":\"success\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/categories/:id": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "分类"
                ],
                "summary": "分类详情",
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":true,\"msg\":\"success\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "分类"
                ],
                "summary": "修改分类",
                "parameters": [
                    {
                        "type": "string",
                        "description": "分类名",
                        "name": "name",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "描述",
                        "name": "description",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":true,\"msg\":\"success\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "分类"
                ],
                "summary": "删除分类",
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":true,\"msg\":\"success\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/topics": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "帖子"
                ],
                "summary": "帖子列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "排序(id/created_at/updated_at,默认id)",
                        "name": "sort",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "排序规则(仅支持 asc（正序）,desc（倒序）)",
                        "name": "order",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "每页条数(介于 2~100 之间)",
                        "name": "per_page",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "当前页",
                        "name": "page",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":true,\"msg\":\"success\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "帖子"
                ],
                "summary": "创建帖子",
                "parameters": [
                    {
                        "type": "string",
                        "description": "标题",
                        "name": "title",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "内容",
                        "name": "body",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "分类id",
                        "name": "category_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":true,\"msg\":\"success\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/topics/:id": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "帖子"
                ],
                "summary": "帖子详情",
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":true,\"msg\":\"success\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "帖子"
                ],
                "summary": "修改帖子",
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":true,\"msg\":\"success\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "帖子"
                ],
                "summary": "删除帖子",
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":true,\"msg\":\"success\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/user": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "获取当前用户",
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":true,\"msg\":\"success\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/users": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "获取所有用户",
                "parameters": [
                    {
                        "type": "string",
                        "description": "排序(id/created_at/updated_at,默认id)",
                        "name": "sort",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "排序规则(仅支持 asc（正序）,desc（倒序）)",
                        "name": "order",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "每页条数(介于 2~100 之间)",
                        "name": "per_page",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "当前页",
                        "name": "page",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":true,\"msg\":\"success\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/verify-codes/captcha": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "授权"
                ],
                "summary": "获取图片验证码",
                "responses": {
                    "200": {
                        "description": "{\"captcha_id\":1,\"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/verify-codes/phone": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "授权"
                ],
                "summary": "发送短信验证码(前提先获取图片验证码)",
                "parameters": [
                    {
                        "type": "string",
                        "description": "手机号码",
                        "name": "phone",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "图片验证码ID",
                        "name": "captcha_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "图片验证码答案",
                        "name": "captcha_answer",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
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
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
