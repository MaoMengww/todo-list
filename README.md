
```
todo-list
├─ .hz
├─ app
│  ├─ gateway
│  │  ├─ handler
│  │  │  └─ api
│  │  │     └─ user
│  │  ├─ mw
│  │  │  └─ sentinel.go
│  │  ├─ router
│  │  └─ rpc
│  │     ├─ Init.go
│  │     └─ user.go
│  ├─ temp
│  │  └─ mw
│  │     └─ main.go
│  ├─ todo
│  │  ├─ controllers
│  │  │  └─ rpc
│  │  │     ├─ handler.go
│  │  │     └─ pack
│  │  ├─ domain
│  │  ├─ infrastructure
│  │  └─ usercase
│  └─ user
│     ├─ controllers
│     │  └─ rpc
│     │     ├─ handler.go
│     │     └─ pack
│     │        └─ pack.go
│     ├─ domain
│     │  ├─ repository.go
│     │  └─ user.go
│     ├─ infrastructure
│     │  └─ mysql.go
│     ├─ inject.go
│     └─ usercase
│        └─ usercase.go
├─ build.sh
├─ cmd
│  ├─ gateway
│  ├─ todo
│  │  └─ main.go
│  └─ user
│     └─ main.go
├─ config
│  ├─ config.go
│  └─ config.yaml
├─ go.mod
├─ go.sum
├─ idl
│  ├─ api
│  │  ├─ todo.thrift
│  │  └─ user.thrift
│  ├─ model.thrift
│  ├─ todo.thrift
│  └─ user.thrift
├─ kitex_gen
│  ├─ model
│  │  ├─ k-consts.go
│  │  ├─ k-model.go
│  │  └─ model.go
│  ├─ todo
│  │  ├─ k-consts.go
│  │  ├─ k-todo.go
│  │  ├─ todo.go
│  │  └─ todoservice
│  │     ├─ client.go
│  │     ├─ server.go
│  │     └─ todoservice.go
│  └─ user
│     ├─ k-consts.go
│     ├─ k-user.go
│     ├─ user.go
│     └─ userservice
│        ├─ client.go
│        ├─ server.go
│        └─ userservice.go
├─ kitex_info.yaml
├─ pkg
│  ├─ common
│  │  └─ tracer.go
│  ├─ middleware
│  │  └─ log.go
│  └─ utils
│     └─ jwt.go
└─ script
   └─ bootstrap.sh

```