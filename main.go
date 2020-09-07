package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {

	dirPtr := flag.String("dir", ".", "root directoy")
	portPtr := flag.String("port", "8081", "server port")
	flag.Parse()

	port := ":"
	port = port + *portPtr

	router := gin.Default()
	//跨域问题
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowWildcard = true
	config.AllowCredentials = true
	config.AddAllowHeaders("Authorization")
	router.Use(cors.New(config))
	//静态文件服务
	router.StaticFS("/", http.Dir(*dirPtr))
	router.Use(static.Serve("/", static.LocalFile(*dirPtr, true)))
	//启动服务器
	router.Run(port)
	log.Printf("Listening and serving HTTP on %s\n", port)
}
