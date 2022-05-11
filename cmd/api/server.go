package main

import (
	"api-users/internal/users"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	actuator "github.com/sinhashubham95/go-actuator"
	"github.com/spf13/viper"
)


func main() {

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	
	// create the gin engine
	engine := gin.Default()

		config := &actuator.Config{
		Endpoints: []int{
			actuator.Env,
			actuator.Info,
			actuator.Metrics,
			actuator.Ping,
			actuator.Shutdown,
			actuator.ThreadDump,
			},
			Env: "dev",
			Name: "API Users",
			Port: 8080,
			Version: "0.1.0",
	}
	// get the handler for actuator
	actuatorHandler := actuator.GetActuatorHandler(config)
	ginActuatorHandler := func(ctx *gin.Context) {
		actuatorHandler(ctx.Writer, ctx.Request)
	}

	engine.GET("/actuator/*endpoint", ginActuatorHandler)

  viper.SetConfigFile("./pkg/config/.env")
  viper.ReadInConfig()

  port := viper.Get("PORT").(string)

	engine.TrustedPlatform = gin.PlatformGoogleAppEngine
	
	httpClient := &http.Client{
		
	}
	client := users.NewClient(httpClient )

	initRepo := []users.User{}

	initRepo = append(initRepo, users.User{
		ID:   "1",
		Name: "John",
		Email: "",
	})

		initRepo = append(initRepo,  users.User{
		ID:   "2",
		Name: "Rosa",
		Email: "",
	})
	userRepo := users.NewUserRepository(initRepo)

	userSvc := users.NewUserService(client, userRepo)

	userHandler :=	users.NewUserHandler(userSvc)
	
	engine =	setupRouter(engine,userHandler)

	go func() {
			if err :=	engine.Run(port); err != nil {
				log.Println("failed to start server",err)
				os.Exit(1)
			}
	}()

	log.Println("ready to serve requests on " + port)
	<-c
	log.Println("gracefully shutting down")
	os.Exit(0)
}
// middleware setup
func setupRouter(r *gin.Engine, handler users.UserHandler) *gin.Engine {

	r.GET("/users", func(c *gin.Context) {
		handler.GetAll(c)
	})

	r.GET("/users/:id", func(c *gin.Context) {
		handler.Get(c)
	})

	r.POST("/users", func(c *gin.Context) {
		handler.Create(c)
	})

	return r
}
