package main

import (
	"github.com/BleethNie/gin-wol/config"
	routes "github.com/BleethNie/gin-wol/routers"
	"github.com/BleethNie/gin-wol/utils"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"time"
)

func BackendServer() *http.Server {
	backPort := config.Cfg.Server.BackPort
	log.Printf("后台服务启动于 http://%s 端口", backPort)
	return &http.Server{
		Addr:         backPort,
		Handler:      routes.BackRouter(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
}

var g errgroup.Group

func main() {

	utils.InitViper()
	utils.InitDB()

	g.Go(func() error {
		return BackendServer().ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
