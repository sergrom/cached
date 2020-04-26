package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/sergrom/cached/config"
	pb "github.com/sergrom/cached/pb"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type Cached struct {
	db *sql.DB
	config *config.Config
	users map[int32]pb.User
	userMux sync.Mutex
}

var cachedServer *Cached

// server is used to implement helloworld.GreeterServer.
type cachedGrpcServer struct {
	pb.UnimplementedCachedServer
}

func (s *cachedGrpcServer) GetUserList(ctx context.Context, req *pb.GetUserListReq) (*pb.GetUserListResp, error) {

	resp := &pb.GetUserListResp{}
	err := cachedServer.GetUserList(req, resp)

	//resp.UserList[1].Name = "rrrrrrrr"
	//fmt.Println(resp.UserList)

	return resp, err
}

func (c *Cached) Start(cfg *config.Config) {

	c.config = cfg

	// Connect to database
	log.Println("Trying to connect to database...")
	err := c.connectToDb()
	if err != nil {
		log.Fatalf("Cannot connect to database: %s", err.Error())
	}
	log.Println("Database connection was successfully established")

	// Handle interrupt signal, ping to database and invalidate cache ==================================================
	intSignal := make(chan os.Signal, 1)
	signal.Notify(intSignal, os.Interrupt)
	invChannel := time.Tick(time.Duration(c.config.Server.InvalidatePeriod) * time.Second)
	pingChannel := time.Tick(60 * time.Second)

	c.invalidateCache()

	go func() {
		for {
			select {
			case <-invChannel:
				log.Println("Invalidate cache")
				c.invalidateCache()
			case <- pingChannel:
				log.Println("Ping to database")
				pingErr := c.db.Ping()
				if pingErr != nil {
					log.Fatalf("Ping error: %s", pingErr.Error())
				}
			case <-intSignal:
				e := c.disconnectDb()
				if e != nil {
					log.Fatalf("Cannot close database connection: %s", e.Error())
				}
				log.Println("Bye")
				os.Exit(0)
			}
		}
	}()

	// serve grpc ===================================================================================================
	go func() {
		lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 50051))
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		s := grpc.NewServer()
		pb.RegisterCachedServer(s, &cachedGrpcServer{})
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// 4. serve jsonRpc =============================================================================================
	server := rpc.NewServer()
	regErr := server.Register(c)
	if regErr != nil {
		log.Fatalf("Cannot register rpc server: %s", regErr.Error())
	}

	unlinkErr := syscall.Unlink(c.config.Server.JsonRpcSocket)
	if unlinkErr != nil {
		log.Fatalf("Cannot unlink old socket \"%s\": %s", c.config.Server.JsonRpcSocket, unlinkErr.Error())
	}

	l, e := net.Listen("unix", c.config.Server.JsonRpcSocket)

	if e != nil {
		log.Fatal("Listen error: ", e)
	}

	chmodErr := os.Chmod(c.config.Server.JsonRpcSocket, 0666)
	if chmodErr != nil {
		log.Fatalf("Cannot change mod for socket \"%s\": %s", c.config.Server.JsonRpcSocket, unlinkErr.Error())
	}

	for {
		conn, err := l.Accept() // this blocks until connection or error
		if err != nil {
			// handle error
			continue
		}

		go func() {
			server.ServeCodec(jsonrpc.NewServerCodec(conn))
		}()
	}
}

func (c *Cached) connectToDb() error {
	dsn := fmt.Sprintf("%s:%s@%s(%s)/%s?interpolateParams=true",
		c.config.Db.User, c.config.Db.Password, c.config.Db.Proto, c.config.Db.Address, c.config.Db.Database)

	var connErr error
	c.db, connErr = sql.Open("mysql", dsn)
	if connErr != nil {
		return errors.New(fmt.Sprintf("Error on initializing database connection: %s", connErr.Error()))
	}

	pingErr := c.db.Ping() // This DOES open a connection if necessary. This makes sure the database is accessible
	if pingErr != nil {
		return errors.New(fmt.Sprintf("Ping to database error: %s", pingErr.Error()))
	}

	return nil
}

func (c *Cached) disconnectDb() error {
	return c.db.Close()
}

func (c *Cached) invalidateCache()  {
	users, err := c.getUsers()
	if err != nil {
		panic("cannot get users")
	}

	c.userMux.Lock()
	c.users = users
	c.userMux.Unlock()
}

// { "method": "Cached.GetUserById", "params": [{"id":"1"}]}
func (c *Cached) GetUserById(args *pb.GetUserByIdReq, response *pb.GetUserByIdResp) (err error) {
	fmt.Println("call GetUserById func")

	if user, ok := c.users[args.Id]; ok {
		response.User = &user
	}

	return nil
}

// { "method": "Cached.GetUserList", "params": []}
func (c *Cached) GetUserList(args *pb.GetUserListReq, response *pb.GetUserListResp) error {
	users := make([]*pb.User, 1)
	for id, item := range c.users {
		user := pb.User{Id: id, Name: item.Name}
		users = append(users, &user)
	}
	response.UserList = users

	return nil
}