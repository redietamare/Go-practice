package main

import (
	"task_manager/router"
)

func main() {
	r := router.SetupRouter()
	r.Run() // Listen and serve on 0.0.0.0:8080
}