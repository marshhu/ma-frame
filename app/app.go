package app

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type App struct{
	Engine *gin.Engine
	HostPort string
	HostPath string
}

var Instance *App

//构建App
func NewApp(hostPort,hostPath string) *App{
	app := &App{
		HostPort: hostPort,
		HostPath: hostPath,
	}
	Instance = app
  return Instance
}

//注册路由
func(a *App)RegisterRouter(router func(eng *gin.Engine) error) *App{
     router(a.Engine)
     return a
}

//启动服务
func(a *App)Run(){
	host := fmt.Sprintf(":%s", a.HostPort)
	s := &http.Server{
		Addr:           host,
		Handler:        a.Engine,
		ReadTimeout:    time.Duration(60) * time.Second,
		WriteTimeout:   time.Duration(60) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	go func() {
		log.Println("Server Listen at:" + host)
		if err := s.ListenAndServe(); err != nil {
			log.Printf("Listen:%s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("Shutdown Server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	s.RegisterOnShutdown(func() {
		log.Println("Server exited")
	})
	log.Println("Server exiting")
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
}