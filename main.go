package main

import (
	"fmt"
	"net/http"
)

/* Global Variables */
var mainRouter Router

func init() {
	//initDatabase()
	initRouter()

}

func main() {
}

// func initDatabase() {
// 	//read init.sql and execute
// }
func initRouter() {
	mainRouter.post("/login-form", f)
	mainRouter.open(":8080")
	fmt.Println("loading...")
}

func f(w http.ResponseWriter, r *http.Request, p map[string]string) {
	res := responeseData{
		Status:  200,
		Massege: "ok",
	}
	fmt.Fprintf(w, res.toJSON())
}
