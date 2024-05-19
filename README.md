# go web admin框架

zero-zone-admin 是一个基于 vue 和 go-zero 开发的全栈前后端分离的开发基础平台，集成jwt鉴权，动态路由，动态菜单，RBAC鉴权功能，把更多时间专注在业务开发上基础业务框架。

## server项目 

go-zero + mysql + redis

## web项目 

vue3 + vite + pinia + element-plus

## todo list 

- [x] 用户管理
- [x] 角色管理
- [x] 菜单管理
- [x] RBAC权限管理
- [x] 字典管理
- [ ] 自动化前端代码



## test api

```
cd api
goctl api go -api core.api --style goZero -dir . --home=../tpl/
```

## test model

```
cd core
goctl model mysql datasource -url="root:fucking@tcp(127.0.0.1:13306)/zero_zone" -table="td_firm"  -dir="./model" -cache=true --style=goZero --home=./cmd/tpl
```

