Entity–Boundary–Interactor
==========================

EBI is a modern application architecture. Its goal is to separate business logic from presentation
logic. With an EBI architecture, by creating different delivery mechanisms, the same business logic
can be used on many platforms.

This repo is just the source for the website. 

.. class:: bold

Go to the official documentation_
---------------------------------

# Project structure 
```lua
/your-project-name
|--/cmd
|  |--/app
|  |  |--main.go
|--/internal
|  |--/entity
|  |  |--entity.go
|  |--/boundary
|  |  |--boundary.go
|  |--/interactor
|  |  |--interactor.go
|--/pkg
|  |--/database
|  |  |--database.go
|  |--/utils
|  |  |--utils.go
|--/config
|  |--config.go
|--/vendor
|--go.mod
|--go.sum
```
