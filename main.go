package main

import (
	"crud/controllers"
	"net/http"
)

func main()  {
	controllers.IndexPage()
	controllers.Create()
	controllers.Update()
	controllers.Delete()
	http.ListenAndServe(":8080", nil)
}