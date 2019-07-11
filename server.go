package main

import (
	"log"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/vimeo/go-magic/magic"
)

func (p *program) run() {
	loadConfig()
	magic.AddMagicDir("~/magicfiles")

	// GIN server
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.NoRoute(static.Serve("/", static.LocalFile("./www", false)))
	r.MaxMultipartMemory = 8 << 20 // Set a lower memory limit for multipart forms (default is 32 MiB)

	r.Use(func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("X-Frame-Options", "deny")
		/*
			if c.Request.TLS == nil {
				c.Header("Content-Security-Policy", "default-src 'self' 'unsafe-inline'; connect-src 'self'; frame-src 'none'; object-src 'none';")
			} else {
				c.Header("Content-Security-Policy", "default-src https: 'self' 'unsafe-inline'; frame-src 'none'; object-src 'none';")
			}
		*/

		c.Next()
	})

	api := r.Group("/api/v1")
	{
		api.POST("/auto-convert", apiAutoConvert)
	}

	// GIN spusteni serveru
	if config.Socket.Enabled {
		go r.RunUnix(config.Socket.Path)
		log.Println("Spusten HTTP server na socketu:", config.Socket.Path)
	}
	if config.HTTP.Enabled {
		go r.Run(config.HTTP.Address + ":" + config.HTTP.Port)
		log.Println("Spusten HTTP server na adrese:", config.HTTP.Address+":"+config.HTTP.Port)
	}
	if config.HTTPS.Enabled {
		go r.RunTLS(config.HTTPS.Address+":"+config.HTTPS.Port, config.HTTPS.Cert, config.HTTPS.Key)
		log.Println("Spusten HTTPS server na adrese:", config.HTTPS.Address+":"+config.HTTPS.Port)
	}

	<-p.exit
}
