package main

/**
 *
 * @author Sika Kay
 * @date 28/10/17
 *
 */
import (
	"log"
	"net/http"

	common "github.com/bitwiseTek/gmn-admin-dev/gmn-admin/api/common"
	"github.com/bitwiseTek/gmn-admin-dev/gmn-admin/api/routers"
	"github.com/codegangsta/negroni"
)

func main() {
	router := routers.InitRoutes()

	n := negroni.Classic()
	n.UseHandler(router)

	server := &http.Server{
		Addr:    common.AppConfig.Server,
		Handler: n,
	}
	log.Println("Listening on :80")
	server.ListenAndServe()
}
