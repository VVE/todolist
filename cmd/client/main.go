/*
 *
 * Copyright 2021 Vladimir Vanin.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a client for TodoList service.
package main

import (
	"context"
	"log"
	"time"

	pb "github.com/VVE/todolist/todolist"

	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewTodoListClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err = c.AddTask(ctx, &pb.Task{Title: "task2", Specification: "specification2.1", Done: false})
	if err != nil {
		log.Fatalf("could not add task: %v", err)
	}
	log.Printf("Task is added")

	_, err = c.AddTask(ctx, &pb.Task{Title: "task1", Specification: "specification1.1", Done: false})
	if err != nil {
		log.Fatalf("could not add task: %v", err)
	}
	log.Printf("Task is added")

	_, err = c.EditTask(ctx, &pb.Task{Id: 1, Title: "task2", Specification: "specification2.2", Done: false})
	if err != nil {
		log.Printf("could not edit task: %v", err)
	} else {
		log.Printf("Task is modified")
	}

	r2, err := c.ShowTask(ctx, &pb.TaskId{Id: 2})
	if err != nil {
		log.Printf("could not show task: %v", err)
	} else {
		log.Printf("Task show: %v %s %s %v", r2.GetId(), r2.GetTitle(), r2.GetSpecification(), r2.GetDone())
	}

	r3, err := c.ShowTaskList(ctx, &pb.Empty{})
	if err != nil {
		log.Fatalf("could not show task list: %v", err)
	}
	log.Printf("Task list: %s", r3.GetTaskList())

	_, err = c.DoneTask(ctx, &pb.TaskId{Id: 2})
	if err != nil {
		log.Printf("could not mark task as done: %v", err)
	} else {
		log.Printf("Task is marked done")
	}

	_, err = c.DeleteTask(ctx, &pb.TaskId{Id: 1})
	if err != nil {
		log.Printf("could not delete task: %v", err)
	} else {
		log.Printf("Task is deleted")
	}

	r3, err = c.ShowTaskList(ctx, &pb.Empty{})
	if err != nil {
		log.Fatalf("could not show task list: %v", err)
	}
	log.Printf("Task list: %s", r3.GetTaskList())

	r2, err = c.ShowTask(ctx, &pb.TaskId{Id: 1}) // тест показать удаленную задачу
	if err != nil {
		log.Printf("could not show task: %v", err)
	} else {
		log.Printf("ShowTask: %s %s %v", r2.GetTitle(), r2.GetSpecification(), r2.GetDone())
	}

	_, err = c.EditTask(ctx, &pb.Task{Id: 1, Title: "task2", Specification: "specification2.2", Done: false}) // тест изменить удаленную задачу
	if err != nil {
		log.Printf("could not edit task: %v", err)
	} else {
		log.Printf("Task is modified")
	}

	_, err = c.DoneTask(ctx, &pb.TaskId{Id: 1}) // тест пометить выполненной удаленную ветку
	if err != nil {
		log.Printf("could not mark task as done: %v", err)
	} else {
		log.Printf("Task is marked done")
	}
}
