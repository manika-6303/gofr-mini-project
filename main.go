package main

import "gofr.dev/pkg/gofr"


func initializeRouter(){

	app:=gofr.New()

	app.GET("/employees", GetEmp)
	app.GET("/employees/{id}", GetEmp)
	app.POST("/employees", PostEmp)
	app.PUT("employees/{id}", PutEmp)
	app.DELETE("/employees/{id}", DeleteEmp)

	log.fatal(http.ListenAndServe(":9000",app))
}

func main(){
	initializeRouter()
}
