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

type App struct {
	Config AppConfig
	Engine *gin.Engine
}

type AppConfig struct {
	HostPort     int
	ReadTimeout  int
	WriteTimeout int
}

var Instance *App

//构建App
func NewApp(config AppConfig) *App {
	app := &App{
		Config: config,
	}
	Instance = app
	return Instance
}

//注册路由
func (a *App) RegisterRouter(router func(eng *gin.Engine) error) *App {
	a.Engine = gin.New()
	router(a.Engine)
	return a
}

//启动服务
func (a *App) Run() {
	host := fmt.Sprintf(":%d", a.Config.HostPort)
	s := &http.Server{
		Addr:           host,
		Handler:        a.Engine,
		ReadTimeout:    time.Duration(a.Config.ReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(a.Config.WriteTimeout) * time.Second,
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
