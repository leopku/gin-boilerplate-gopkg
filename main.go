/**
* File: /main.go
* File Created: Wednesday, 21st February 2018 9:59:13 pm
* Author: liusong (liusong@daojia.com)
* -----
* Last Modified: Wednesday, 21st February 2018 9:59:13 pm
* Modified By: liusong (liusong@daojia.com>)
* -----
* Copyright 2018 - 2018 daojia.com, daojia
*/
package main

import (
	"os"

  "gopkg.in/appleboy/gin-jwt.v2"
  "gopkg.in/gin-gonic/gin.v1"
  "gopkg.in/joho/godotenv.v1"
  "gopkg.in/sirupsen/logrus.v1"
)

func helloHandler(c *gin.Context) {
  claims := jwt.ExtractClaims(c)
  c.JSON(200, gin.H{
    "userID": claims["id"],
    "text": "Hello World.",
  })
}

func main() {
  // err := godotenv.Load()
  if err := godotenv.Load(); err != nil {
    logrus.Fatal("Error loading .env file")
  }

  logrus.Info("[asdf]" + os.Getenv("DEMO"))

  authMiddleware := &jwt.GinJWTMiddleware{
    Realm: "test zone",
    Key: []byte("secret key"),
    // Timeout: time.Hour,
    // MaxRefresh: time.Hour,
    Authenticator: func(serId string, password string, c *gin.Context) (string, bool) {
      if (userId == "admin" && password == "admin") || (userId == "test" && password ="test") {
        return userId, true
      }
      return userId, false
    },
    Authorizator: func(userId string, c *gin.Context) bool {
      if userId == "admin" {
        return true
      }
      return false
    },
    Unauthorized: func(c *gin.Context, code int, message string) {
      c.JSON(code, gin.H{
        "code": code,
        "message": message,
      })
    },
    // TokenLookup: "header:Authorization",
    // TokenHeadName: "Bearer"
  }

  r := gin.Default()
  r.GET("/get", func(c *gin.Context) {
    c.JSON(200, gin.H{
      "headers": c.Request.Header,
    })
  })

  r.POST("/login", authMiddleware.LoginHandler)

  api := r.Group("/api")
  api.Use(authMiddleware.MiddlewareFunc())
  {
    api.GET("/hello", helloHandler)
    api.GET("/refresh_token", authMiddleware.RefreshHandler)
  }

  r.Run()
}
