package main

import (
	"net/http"

	router "jrnc-db/server/adapter"
	"jrnc-db/server/adapter/controllers"
	"jrnc-db/server/driver/mysql"
)

type MyJob struct {}



// init initialize server
func main() {
	mysql.InitDev()

	//jobrunner.Start()
	//jobrunner.Schedule("@every 5s", MyJob{})

	r := router.Router(true)
	http.Handle("/", r)
	r.Run(":8080")
}

// Batch Job
func (e MyJob) Run() {
	ctrl := controllers.Controllers{}
	ctrl.JobSchedullerController()
}
