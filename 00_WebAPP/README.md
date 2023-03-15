# Overview
After we've done for implementing the shape application running your local machine, we start migrating the application from local to web app.  
This is a project to demo your application can run through HTTP.
# Project Structure
There are some project structure templates for overview our architecture.

## Example 1
```lua
├── cmd/
│   └── main.go
├── internal/
│   ├── app/
│   │   ├── usecase/
│   │   │   ├── shape.go
│   │   │   ├── rectangle.go
│   │   │   └── circle.go
│   │   ├── entity/
│   │   │   ├── shape.go
│   │   │   ├── rectangle.go
│   │   │   └── circle.go
│   │   ├── repository/
│   │   │   ├── shape_repository.go
│   │   │   ├── rectangle_repository.go
│   │   │   └── circle_repository.go
│   │   ├── service/
│   │   │   ├── shape_service.go
│   │   │   ├── rectangle_service.go
│   │   │   └── circle_service.go
│   │   ├── handler/
│   │   │   ├── shape_handler.go
│   │   │   ├── rectangle_handler.go
│   │   │   └── circle_handler.go
│   │   ├── builder/
│   │   │   ├── shape_builder.go
│   │   │   ├── rectangle_builder.go
│   │   │   └── circle_builder.go
│   │   └── config/
│   │       ├── config.go
│   │       ├── database.go
│   │       ├── logger.go
│   │       └── router.go
│   └── pkg/
│       └── util/
│           ├── util.go
│           ├── string_util.go
│           ├── time_util.go
│           └── ...
├── vendor/
├── go.mod
├── go.sum
└── Dockerfile

```
## Example 2
```lua
├── cmd/
│   └── myapp/
│       └── main.go
├── internal/
│   ├── app/
│   │   ├── http/
│   │   │   ├── handler/
│   │   │   │   └── user_handler.go
│   │   │   └── server.go
│   │   ├── service/
│   │   │   ├── user_service.go
│   │   │   └── auth_service.go
│   │   └── repository/
│   │       ├── user_repository.go
│   │       └── auth_repository.go
│   ├── config/
│   │   ├── app.yaml
│   │   └── db.yaml
│   ├── db/
│   │   ├── mysql/
│   │   │   └── mysql.go
│   │   └── redis/
│   │       └── redis.go
│   ├── logger/
│   │   └── logger.go
│   ├── utils/
│   │   └── utils.go
│   └── errors/
│       ├── errors.go
│       └── codes.go
├── configs/
│   ├── app.yaml
│   └── db.yaml
├── build/
│   ├── myapp-linux
│   ├── myapp-mac
│   └── myapp-windows.exe
├── scripts/
│   ├── run.sh
│   ├── build.sh
│   └── deploy.sh
├── go.mod
└── README.md

```