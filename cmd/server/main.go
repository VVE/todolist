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

// Package main implements a server for TodoList service.
package main

import (
	"context"
	//"errors"
	"fmt"
	"log"
	"net"
	"sort"

	//"google/protobuf/empty.proto"
	//grpc "google.golang.org/grpc"
	"google.golang.org/grpc"
	//emptypb "google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/VVE/todolist/todolist"
	//pb "todolist/todolist"
)

const (
	port = ":50051"
)

//var id int32 = 1

/* type tasks struct {
	title string  			// Наименование задачи.
    specification string  	// Содержание задачи.
	done bool		    	// Признак завершения.
} */

// server is used to implement todolist.TodoListServer.
type server struct {
	pb.UnimplementedTodoListServer
	//taskList []*pb.Task
	taskList pb.TaskList
}

func (s *server) AddTask(ctx context.Context, in *pb.Task) (*pb.Empty, error) {
	log.Printf("Received: AddTask %v %v %v", in.GetTitle(), in.GetSpecification(), in.GetDone())
	log.Printf("taskList before:\n")
	for i, t := range s.taskList.TaskList {
		log.Printf("%v %v %v %v %v\n", i, t.Id, t.Title, t.Specification, t.Done)
	}
	in.Id = int32(len(s.taskList.TaskList) + 1)
	log.Printf("Added: %v %v %v %v", in.Title, in.Specification, in.Done, in.Id)
	s.taskList.TaskList = append(s.taskList.TaskList, in)
	log.Printf("taskList after:\n")
	for i, t := range s.taskList.TaskList {
		log.Printf("%v %v %v %v %v\n", i, t.Id, t.Title, t.Specification, t.Done)
	}
	log.Printf("\n")
	return &pb.Empty{}, nil
}

func (s *server) EditTask(ctx context.Context, in *pb.Task) (*pb.Empty, error) {
	log.Printf("Received: EditTask %v %v %v %v", in.GetId(), in.GetTitle(), in.GetSpecification(), in.GetDone())
	log.Printf("taskList before:\n")
	for i, t := range s.taskList.TaskList {
		log.Printf("%v %v %v %v %v\n", i, t.Id, t.Title, t.Specification, t.Done)
	}
	idx := sort.Search(len(s.taskList.TaskList), func(i int) bool { return in.Id <= s.taskList.TaskList[i].Id })
	log.Printf("idx: %v\n", idx)
	if idx >= len(s.taskList.TaskList) || (in.Id != s.taskList.TaskList[idx].Id) {
		log.Printf("task is absent")
		log.Printf("\n")
		return &pb.Empty{}, fmt.Errorf("task is absent")
	}
	s.taskList.TaskList[idx].Title = in.Title
	s.taskList.TaskList[idx].Specification = in.Specification
	s.taskList.TaskList[idx].Done = in.Done
	log.Printf("task after: %v\n", s.taskList.TaskList[idx])
	log.Printf("\n")
	return &pb.Empty{}, nil
}

func (s *server) DoneTask(ctx context.Context, in *pb.TaskId) (*pb.Empty, error) {
	log.Printf("Received: DoneTask %v", in.GetId())
	log.Printf("taskList before:\n")
	for i, t := range s.taskList.TaskList {
		log.Printf("%v %v %v %v %v\n", i, t.Id, t.Title, t.Specification, t.Done)
	}
	idx := sort.Search(len(s.taskList.TaskList), func(i int) bool { return in.Id <= s.taskList.TaskList[i].Id })
	log.Printf("idx: %v\n", idx)
	if idx >= len(s.taskList.TaskList) || (in.Id != s.taskList.TaskList[idx].Id) {
		log.Printf("task is absent\n")
		return nil, fmt.Errorf("task is absent")
	}
	s.taskList.TaskList[idx].Done = true
	log.Printf("taskList after:\n")
	for i, t := range s.taskList.TaskList {
		log.Printf("%v %v %v %v %v\n", i, t.Id, t.Title, t.Specification, t.Done)
	}
	log.Printf("\n")
	return &pb.Empty{}, nil
}

func (s *server) DeleteTask(ctx context.Context, in *pb.TaskId) (*pb.Empty, error) {
	log.Printf("Received: DeleteTask %v", in.GetId())
	log.Printf("taskList before:\n")
	for i, t := range s.taskList.TaskList {
		log.Printf("%v %v %v %v %v\n", i, t.Id, t.Title, t.Specification, t.Done)
	}
	idx := sort.Search(len(s.taskList.TaskList), func(i int) bool { return in.Id <= s.taskList.TaskList[i].Id })
	log.Printf("idx: %v\n", idx)
	if idx >= len(s.taskList.TaskList) || (in.Id != s.taskList.TaskList[idx].Id) {
		log.Printf("task is absent\n")
		return nil, fmt.Errorf("task is absent")
	}
	s.taskList.TaskList = append(s.taskList.TaskList[:idx], s.taskList.TaskList[idx+1:]...)
	log.Printf("taskList after:\n")
	for i, t := range s.taskList.TaskList {
		log.Printf("%v %v %v %v %v\n", i, t.Id, t.Title, t.Specification, t.Done)
	}
	log.Printf("\n")
	return &pb.Empty{}, nil
}

func (s *server) ShowTask(ctx context.Context, in *pb.TaskId) (*pb.Task, error) {
	log.Printf("Received: ShowTask %v", in.GetId())
	log.Printf("taskList before:\n")
	for i, t := range s.taskList.TaskList {
		log.Printf("%v %v %v %v %v\n", i, t.Id, t.Title, t.Specification, t.Done)
	}
	idx := sort.Search(len(s.taskList.TaskList), func(i int) bool { return in.Id <= s.taskList.TaskList[i].Id })
	log.Printf("idx: %v\n", idx)
	if idx >= len(s.taskList.TaskList) || (in.Id != s.taskList.TaskList[idx].Id) {
		log.Printf("task is absent")
		log.Printf("\n")
		return nil, fmt.Errorf("task is absent")
	}
	var t pb.Task
	t.Id = s.taskList.TaskList[idx].Id
	t.Title = s.taskList.TaskList[idx].Title
	t.Specification = s.taskList.TaskList[idx].Specification
	t.Done = s.taskList.TaskList[idx].Done
	log.Printf("Reply: %v %v %v %v\n", t.Id, t.Title, t.Specification, t.Done)
	log.Printf("\n")
	return &t, nil
}

func (s *server) ShowTaskList(ctx context.Context, in *pb.Empty) (*pb.TaskList, error) {
	log.Printf("Received: ShowTaskList\n")
	log.Printf("TaskList:\n")
	for i, t := range s.taskList.TaskList {
		log.Printf("%v %v %v %v %v\n", i, t.Id, t.Title, t.Specification, t.Done)
	}
	log.Printf("\n")
	return &s.taskList, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("server is starting")
	s := grpc.NewServer()
	log.Printf("server is ready")
	pb.RegisterTodoListServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	log.Printf("server is closing")
}
