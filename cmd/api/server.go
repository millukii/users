package main

import (
	"api-users/internal/users"

	"github.com/gin-gonic/gin"
	actuator "github.com/sinhashubham95/go-actuator"
	"github.com/spf13/viper"
)


func main() {
	
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
	
	userSvc := users.NewUserService()
	userHandler :=	NewUserHandler(userSvc)
	engine =	setupRouter(engine,userHandler)
  
	engine.Run(port)
}