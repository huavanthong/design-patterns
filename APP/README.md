# Project Structure
This is a project structure for overview our architecture
```lua
├── common/
│   ├── common.go
├── entity/
│   ├── rectangle.go
│   └── circle.go
│
├── boundary/
│   ├── rectangle_input.go
│   ├── rectangle_output.go
│   ├── circle_input.go
│   ├── circle_output.go
│   ├── rectangle_boundary.go
│   └── circle_boundary.go
│
├── interactor/
│   ├── rectangle_interactor.go
│   └── circle_interactor.go
│
├── builder/
│   ├── shape_builder.go
│   ├── rectangle_builder.go
│   └── circle_builder.go
│
├── repository/
│   ├── rectangle_repository.go
│   └── circle_repository.go
│
├── go.mod
└── main.go
```

# Application Purpose
