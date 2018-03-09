package main

import (
	"fmt"
	"net/http"
)

/* Global Variables */
var mainRouter Router

func init() {
	initDatabase()
}

func main() {
	initRouter()
}

func initDatabase() {
	//read init.sql and execute
}

func initRouter() {
	mainRouter.get("/login", f)    //get login form
	mainRouter.post("/login", f)   //login, return token if success
	mainRouter.get("/file/***", f) //get file
	mainRouter.post("/file", f)    //upload file
	//get list
	mainRouter.get("/users", f)
	mainRouter.get("/objects", f)
	mainRouter.get("/terms", f)
	mainRouter.get("/options", f)
	mainRouter.get("/comments", f)
	//get one
	mainRouter.get("/users/{id}", f)
	mainRouter.get("/objects/{id}", f)
	mainRouter.get("/terms/{id}", f)
	mainRouter.get("/options/{id}", f)
	mainRouter.get("/comments/{id}", f)
	//create one
	mainRouter.post("/users", f)
	mainRouter.post("/objects", f)
	mainRouter.post("/terms", f)
	mainRouter.post("/options", f)
	mainRouter.post("/comments", f)
	//update one
	mainRouter.put("/users/{id}", f)
	mainRouter.put("/objects/{id}", f)
	mainRouter.put("/terms/{id}", f)
	mainRouter.put("/options/{id}", f)
	mainRouter.put("/comments/{id}", f)
	//delete one
	mainRouter.delete("/users/{id}", f)
	mainRouter.delete("/objects/{id}", f)
	mainRouter.delete("/terms/{id}", f)
	mainRouter.delete("/options/{id}", f)
	mainRouter.delete("/comments/{id}", f)

	mainRouter.open(":8080")
}

//for developing
func f(w http.ResponseWriter, r *http.Request, p map[string]string) {
	initRouter()
	data := make(map[string]string)
	data["sub"] = "1234567890"
	data["name"] = "John Doe"
	data["iat"] = "1516239022"
	op, _ := token.encode(data)
	fmt.Println(op)
	res := responeseData{
		Status:  200,
		Massege: "ok",
		Data:    op,
	}
	fmt.Fprintf(w, res.toJSON())
}
