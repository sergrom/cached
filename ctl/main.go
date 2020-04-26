package main

import (
	"context"
	"fmt"
	pb "github.com/sergrom/cached/pb"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main()  {
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func() {
		err := conn.Close()
		if err != nil {
			log.Fatalf("Couldn't close connection: %s", err.Error())
		}
	}()
	c := pb.NewCachedClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	userListResp, err := c.GetUserList(ctx, &pb.GetUserListReq{})
	if err != nil {
		log.Fatalf("Could not get user list: %v", err)
	}

	fmt.Println(userListResp)

	for _, user := range userListResp.UserList {
		fmt.Println(user.Name)
	}
}