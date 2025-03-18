package apiserver

import (
	"fmt"
	"net"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

// webFilePath: ex) "./dist"
func Run(bindPort int, bindUrl string, useWebServer bool, webFilePath string, useHttps bool, keyFile string, certFile string, routeCB func(*gin.Engine)) error {
	// alloc
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowMethods: []string{"*"},
		AllowHeaders: []string{"*"},
		AllowOrigins: []string{"*"},
		//AllowCredentials: true,
		//AllowHeaders: []string{"Origin", "Accept", "Content-Type", "X-Requested-With"},
		//AllowOrigins: []string{"http://localhost:8080", "http://127.0.0.1:8080"},
	}))

	// webserver
	if useWebServer {
		r.Use(static.Serve("/", static.LocalFile(webFilePath, false)))
		r.NoRoute(func(c *gin.Context) {
			c.File("./dist/index.html")
		})
	}

	// route
	if routeCB != nil {
		routeCB(r)
	}

	// bind
	server := &http.Server{Handler: r}
	listner, err := net.Listen("tcp4", fmt.Sprintf("0.0.0.0:%d", bindPort))
	if err != nil {
		return err
	}

	// run
	if useHttps {
		err = server.ServeTLS(listner, certFile, keyFile)
	} else {
		err = server.Serve(listner)
	}

	return err
}
