# gRPC for Go example (ToDo)
Simple gRPC server and client. 
This example is intended to fill the space between the example (by grpc.io) of Hello World with no field and examples by Aleksandr Sokolovskii and by Kyle Felter that demonstrates several features at once.
# Summary
gRPC - Remome Procedure Call
In memory database.
# Running the code
1. Start the server
```
cd %GOPATH%\src\todolist\cmd>
go run server\main.go
```
1. Start the client
```
cd %GOPATH%\src\todolist\cmd
go run client\main.go
```
# Resources
| Title | URL |
|-------|-----|
| gRPC site | https://grpc.io/ |
| Quick start | https://grpc.io/docs/languages/go/quickstart/ |
| Protocol Buffers | https://developers.google.com/protocol-buffers/ |
| proto | https://pkg.go.dev/google.golang.org/protobuf/proto |
| How to develop Go gRPC microservice with HTTP/REST endpoint, middleware, Kubernetes deployment, etc. | https://github.com/amsokol/go-grpc-http-rest-microservice-tutorial/tree/part1 |
| grpc-example using golang | https://github.com/kfelter/grpc-example |
# Proto file
The task is represented by a message with the following fields: 
id (int32) - identifier, 
title (string) - title, 
specification (string) - description, 
done (bool) - mark of completion.
ID is required in order not to require the uniqueness of the titles.
The list of tasks is represented by a message with repeated field.
service is a Remote Procedure Calls (rpc) definition.
The AddTask procedure adds a task to the database. The argument is the structure of the task, the result is an error message (empty if there is no error). 
The EditTask procedure modifies a task in the database. The argument is the structure of the task, the result is an error message (empty if there is no error).
The DoneTask procedure changes the status of the task in the database to Completed. The argument is the task ID, the result is an error message (empty if there is no error).
The Delete Task procedure deletes the task from the database. The argument is the task ID, the result is an error message (empty if there is no error).
The ShowTask procedure provides the contents of the task from the database. The argument is the task ID, the result is the task structure.
The Show Task List procedure provides a list of the contents of tasks from the database. There is no argument, the result is a list of tasks.
The server itself fills in the ID field when adding a task.
Sorting of tasks in the database is not required, since their order does not change.
So, the messages are: Task, TaskId, TaskList, Empty.
The empty message from Google is not used to reduce dependencies and reserve the ability to add any field to it.
# Test
For simplicity, the client does not include any user interface and plays the role of a test:
1) add a task to an empty database,
2) add another task,
3) edit one of the tasks (the second one),
4) show one of the tasks (the second one),
5) show the task list,
6) mark one of the tasks (the second one) as completed,
7) delete one of the tasks (the first one),
8) show the task list again,
9) show the remote task (to check the error message),
10) edit the remote task (similarly),
11) mark the deleted task as completed (similarly).
# File structure
todolist
```
L .git
L cmd
  L client
    L main.go
  L server
    L main.go
L third_party
  L google
L todolist
  L todo_list.proto
  L todo_list.pb.go
  L todo_list_rpc.pb.go
L go.mod
L go.sum
L README.md
```