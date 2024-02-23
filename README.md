# Todo Backend App with Golang

Todo application backend using Golang Gin and Gorm. There are two main endpoint groups here, users (for authentication and authorization) and todo (for related to todo models). I restrict the user, so that they only can read, update, and delete their own todos.

## How to run

1. Clone the repository
2. Install the needed packages, run with `go mod download` command
3. Run the application with `air` command

## Docker Image

you can find the docker image in the following link
[here](https://hub.docker.com/repository/docker/iqbalpa/todo-go/)
