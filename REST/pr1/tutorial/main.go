package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	contro "tutorial.com/api/tutorial/controllers"
	"tutorial.com/api/tutorial/services"
)

var (
	server         *gin.Engine
	userservice    services.UserService
	usercontroller contro.UserController
	ctx            context.Context
	usercollection *mongo.Collection
	mongoclient    *mongo.Client
	err            error
)

func init() {
	ctx = context.TODO()

	mongoconn := options.Client().ApplyURI("mongodb://localhost:27017")
	mongoclient, err = mongo.Connect(ctx, mongoconn)
	if err != nil {
		log.Fatal(err)
	}
	err = mongoclient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("mongo connected")

	usercollection = mongoclient.Database("userdb").Collection("users")
	userservice = services.NuserUserService(usercollection, ctx)

	usercontroller = contro.New(userservice)

	server = gin.Default()

}

//v1/user/creaete

func main() {
	defer mongoclient.Disconnect(ctx)
	basepath := server.Group("v1")
	usercontroller.RedgisterUserRoutes(basepath)

	log.Fatal(server.Run(":9090"))
}
