// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2019-05-16 10:45:30.7840849 +0800 CST m=+0.115740401

package docs

import (
	"github.com/swaggo/swag"
)

var doc = `{
    "swagger": "2.0",
    "info": {
        "description": "go-admin",
        "title": "go-admin",
        "termsOfService": "https://github.com/hequan2017/go-admin",
        "contact": {
            "name": "hequan",
            "url": "https://github.com/hequan2017",
            "email": "hequan2011@sina.com"
        },
        "license": {
            "name": "MIT",
            "url": "https://github.com/hequan2017/go-admin/blob/master/LICENSE"
        },
        "version": "1.1.2"
    },
    "basePath": "/api/v1",
    "paths": {
        "/api/v1/menus": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "menu"
                ],
                "summary": "获取所有菜单",
                "responses": {
                    "200": {
                        "description": "{ \"code\": 200, \"data\": {}, \"msg\": \"ok\" }",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "menu"
                ],
                "summary": "增加菜单",
                "parameters": [
                    {
                        "type": "string",
                        "description": "name",
                        "name": "name",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "path",
                        "name": "path",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{ \"code\": 200, \"data\": {}, \"msg\": \"ok\" }",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/menus/": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "menu"
                ],
                "summary": "获取单个菜单",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{ \"code\": 200, \"data\": {}, \"msg\": \"ok\" }",
                        "schema": {
                            "type": "string"
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
                    "menu"
                ],
                "summary": "更新菜单",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "name",
                        "name": "name",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "path",
                        "name": "path",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "method",
                        "name": "method",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{ \"code\": 200, \"data\": {}, \"msg\": \"ok\" }",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "menu"
                ],
                "summary": "删除菜单",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{ \"code\": 200, \"data\": {}, \"msg\": \"ok\" }",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/roles": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "role"
                ],
                "summary": "获取所有角色",
                "responses": {
                    "200": {
                        "description": "{ \"code\": 200, \"data\": {}, \"msg\": \"ok\" }",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "role"
                ],
                "summary": "增加角色",
                "parameters": [
                    {
                        "type": "string",
                        "description": "name",
                        "name": "name",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "menu_id",
                        "name": "menu_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{ \"code\": 200, \"data\": {}, \"msg\": \"ok\" }",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/roles/": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "role"
                ],
                "summary": "获取单个角色",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{ \"code\": 200, \"data\": {}, \"msg\": \"ok\" }",
                        "schema": {
                            "type": "string"
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
                    "role"
                ],
                "summary": "更新角色",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "name",
                        "name": "name",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "menu_id",
                        "name": "menu_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{ \"code\": 200, \"data\": {}, \"msg\": \"ok\" }",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "role"
                ],
                "summary": "删除角色",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{ \"code\": 200, \"data\": {}, \"msg\": \"ok\" }",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/users": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "获取所有用户",
                "parameters": [
                    {}
                ],
                "responses": {
                    "200": {
                        "description": "{ \"code\": 200, \"data\": {}, \"msg\": \"ok\" }",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "增加用户",
                "parameters": [
                    {
                        "description": "username",
                        "name": "username",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object"
                        }
                    },
                    {
                        "description": "password",
                        "name": "password",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object"
                        }
                    },
                    {
                        "type": "integer",
                        "description": "role_id",
                        "name": "role_id",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{ \"code\": 200, \"data\": {}, \"msg\": \"ok\" }",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/users/": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "获取单个用户信息",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{ \"code\": 200, \"data\": {}, \"msg\": \"ok\" }",
                        "schema": {
                            "type": "string"
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
                    "users"
                ],
                "summary": "更新用户",
                "parameters": [
                    {},
                    {},
                    {
                        "type": "integer",
                        "description": "role_id",
                        "name": "role_id",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{ \"code\": 200, \"data\": {}, \"msg\": \"ok\" }",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "删除用户",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{ \"code\": 200, \"data\": {}, \"msg\": \"ok\" }",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/auth": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "获取登录token 信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "username",
                        "name": "username",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "password",
                        "name": "password",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{ \"code\": 200, \"data\": { \"token\": \"xxx\" }, \"msg\": \"ok\" }",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "{\"code\":400,  \"data\":null,\"msg\":\"请求参数错误\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "{ \"code\": 404, \"data\":null,\"msg\":\"请求参数错误\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    }
}`

type s struct{}

func (s *s) ReadDoc() string {
	return doc
}
func init() {
	swag.Register(swag.Name, &s{})
}
