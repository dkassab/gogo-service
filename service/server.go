package service

import (
	//"fmt"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
	"net/http"
)

//NewServer() configures and returns a Server.

func NewServer() *negroni.Negroni {
	n := negroni.New()
	mx := mux.NewRouter()
	initRoutes(mx)
	n.Use(negroni.NewLogger())
	n.UseHandler(mx)

	return n
}

func initRoutes(mx *mux.Router) {
	formatter := render.New(render.Options{IndentJSON: true})
	mx.HandleFunc("/test", testHandler(formatter)).Methods("GET")
	mx.HandleFunc("/practice", practiceHandler(formatter)).Methods("GET")
}

func testHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {

		formatter.JSON(w, http.StatusOK,
			struct{ Test string }{"This is a test"})
	}
}

func practiceHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {

		var a [2]string
		a[0] = "Hello"
		a[1] = "World"
		formatter.JSON(w, http.StatusOK, a)
		//formatter.JSON(w, http.StatusOK, test1.y)
	}
}
