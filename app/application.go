package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

//StartApp :
func StartApp() {
	router = gin.Default()
	srv := http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	mapUrls()
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Println(err)
			return
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGABRT)
	log.Println("Shutdown signal:", <-quit)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Println("Shutdown error", err)
	}
	select {
	case <-ctx.Done():
		fmt.Println("Timeout of 2 Seconds")
	}
	log.Println("Shutdown server gracefully")
}
