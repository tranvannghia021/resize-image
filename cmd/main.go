package main

import (
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/cors"
	"log"
	"os"
	appHandler "resize-image/cmd/app"
	"resize-image/configs"
	"time"
)

func init() {
	if configs.GetAppUrl() == "" || configs.GetPort() == 0 {
		log.Fatal("env is required")
	}

	// log init

	f, err := os.Create("log.txt")
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(f)
}

func main() {
	// init server
	domain := fmt.Sprintf(":%d", configs.GetPort())
	hertzS := server.Default(server.WithHostPorts(domain),
		server.WithHandleMethodNotAllowed(true),
		server.WithMaxRequestBodySize(5*1024*1024), // 5mb
	)
	// cors
	hertzS.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Authorization", "Origin", "X-Requested-With", "Content-Type", "Accept", "X-Permission-Checksum"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour, //3600
	}))
	log.Println("server is running " + domain)
	// router
	routes(hertzS)

	// start
	hertzS.Spin()
}

func routes(h *server.Hertz) {

	h.StaticFS("assets/", &app.FS{Root: "./", GenerateIndexPages: false})

	api := h.Group("api")
	{
		api.POST("image/re-size", appHandler.ResizeHandler)
	}

}
