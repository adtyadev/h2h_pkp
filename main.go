package main

import (
	"h2h_pkp/controllers"
	"h2h_pkp/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.Connection()
	models.Migrate()
	// start := time.Now()
	// for i := 0; i <= 50; i++ {
	// 	models.Seeder()
	// }
	// elapsed := time.Since(start)
	// fmt.Printf("page took %s", elapsed)
	r.POST("/api/add", controllers.CreateRequest)
	r.POST("/api/agunan", controllers.AddAgunan)
	r.POST("/api/test/throttle", controllers.TestThrottle)
	//r.Run()
	s := &http.Server{
		Addr:           ":8181",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
	//jika tidak ada selama n waktu yang ditulis atau dibaca melalui server, maka server akan di terminated
}
