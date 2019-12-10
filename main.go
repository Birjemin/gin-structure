package main

import (
	"context"
	"flag"
	conf2 "github.com/birjemin/gin-structure/conf"
	"github.com/birjemin/gin-structure/datasource"
	"github.com/birjemin/gin-structure/web/middleware"
	"github.com/birjemin/gin-structure/web/routers"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	var port = flag.String("port", "8081", "port flag")
	flag.Parse()

	router := gin.Default()
	router.Use(middleware.Cors())

	router.Static("/public", "/public")
	if conf2.RunMode == "pro" {
		gin.SetMode(gin.ReleaseMode)
	}
	routers.SetRouters(router)

	srv := &http.Server{
		Addr:           ":" + *port,
		Handler:        router,
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// gracefully shutdown
	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// Close database, redis, truncate message queues, etc.
	err := datasource.CloseDb()
	log.Println("[main]DB Pool Exited...", err)
	err = datasource.CloseRedis()
	log.Println("[main]Redis Pool Exited...", err)
	log.Println("[main]Iris shutdown...")
	log.Println("[main]Defer, Pool Redis and Db Stats", datasource.StatsRedis(), datasource.StatsDB())

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown: ", err)
	}

	log.Println("Server exiting")
}
