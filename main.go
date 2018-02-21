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
  _ "gopkg.in/gin-gonic/gin.v1"
  log "gopkg.in/sirupsen/logrus.v1"
  "gopkg.in/joho/godotenv.v1"
)

func main() {
  // err := godotenv.Load()
  godotenv.Load()
  log.Info("[asdf]" + os.Getenv("DEMO"))
}
