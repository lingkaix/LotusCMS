package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

//response data format
type responeseData struct {
	Status  int
	Massege string
	Data    interface{}
}

func (r *responeseData) toJSON() string {
	data, err := json.Marshal(r)
	if err != nil {
		fmt.Println("error:", err)
		res := responeseData{
			Status:  500,
			Massege: "json data error",
		}
		data, _ = json.Marshal(res)
	}
	return string(data)
}

//Router : used for add router handlers which are with url pattern matching parameters
type Router struct {
	httpInfo       string
	getHandlers    []RouterHandler
	postHandlers   []RouterHandler
	putHandlers    []RouterHandler
	deleteHandlers []RouterHandler
}

func (r *Router) open(port string) error {
	r.httpInfo = port

	return http.ListenAndServe(port, r)
}

func (r *Router) get(url string, handler func(http.ResponseWriter, *http.Request, map[string]string)) error {
	var rh RouterHandler
	err := rh.init(url, handler)
	if err != nil {
		return errors.New("error adding router -> " + err.Error())
	}
	r.getHandlers = append(r.getHandlers, rh)
	return nil
}

func (r *Router) post(url string, handler func(http.ResponseWriter, *http.Request, map[string]string)) error {
	var rh RouterHandler
	err := rh.init(url, handler)
	if err != nil {
		return errors.New("error adding router -> " + err.Error())
	}
	r.postHandlers = append(r.postHandlers, rh)
	return nil
}

func (r *Router) put(url string, handler func(http.ResponseWriter, *http.Request, map[string]string)) error {
	var rh RouterHandler
	err := rh.init(url, handler)
	if err != nil {
		return errors.New("error adding router -> " + err.Error())
	}
	r.putHandlers = append(r.putHandlers, rh)
	return nil
}

func (r *Router) delete(url string, handler func(http.ResponseWriter, *http.Request, map[string]string)) error {
	var rh RouterHandler
	err := rh.init(url, handler)
	if err != nil {
		return errors.New("error adding router -> " + err.Error())
	}
	r.deleteHandlers = append(r.deleteHandlers, rh)
	return nil
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	isMatched := false
	var handlers []RouterHandler
	switch req.Method {
	case "":
		fallthrough
	case "GET":
		handlers = r.getHandlers
	case "POST":
		handlers = r.postHandlers
	case "UPDATE":
		fallthrough
	case "PATCH":
		fallthrough
	case "PUT":
		handlers = r.putHandlers
	case "DELETE":
		handlers = r.getHandlers
	default:
		r.notFoundHandler(w, req)
	}
	for _, rh := range handlers {
		if a, b, _ := rh.matchURL(req.URL.Path); a {
			fmt.Println("matched:", rh.url, rh.pattern)
			isMatched = true
			rh.handleFunc(w, req, *b)
			break
		}
	}
	if isMatched == false {
		r.notFoundHandler(w, req)
	}
}

func (r *Router) notFoundHandler(w http.ResponseWriter, req *http.Request) {
	res := responeseData{
		Status:  404,
		Massege: "data not found",
	}
	fmt.Fprintf(w, res.toJSON())
}

//RouterHandler :
type RouterHandler struct {
	url        string
	pattern    *regexp.Regexp
	handleFunc func(http.ResponseWriter, *http.Request, map[string]string)
}

func (r *RouterHandler) init(url string, handler func(http.ResponseWriter, *http.Request, map[string]string)) error {
	if url[len(url)-1:] == "/" {
		url = url[0 : len(url)-1]
	}
	r.url = url
	r.handleFunc = handler
	regx, _ := regexp.Compile("{([a-zA-Z0-9]+)}")
	for regx.MatchString(url) {
		url = strings.Replace(url, regx.FindStringSubmatch(url)[0], "(?P<"+regx.FindStringSubmatch(url)[1]+">[a-z,A-Z,0-9]+)", -1)
	}
	if url[len(url)-4:] == "/***" {
		url = url[0:len(url)-4] + "(?P<_trail_>/.+)"
	}
	regx2, err := regexp.Compile(url)
	if err != nil {
		fmt.Println(err)
		return errors.New("url pattern complile error -> " + err.Error())
	}
	r.pattern = regx2
	return nil
}

func (r *RouterHandler) matchURL(url string) (bool, *map[string]string, error) {
	if url[len(url)-1:] == "/" {
		url = url[0 : len(url)-1]
	}
	if r.pattern.MatchString(url) && r.pattern.FindStringSubmatch(url)[0] == url {
		attrs := make(map[string]string)
		for i, attr := range r.pattern.FindStringSubmatch(url) {
			if i == 0 {
				continue
			}
			attrs[r.pattern.SubexpNames()[i]] = attr
		}
		return true, &attrs, nil
	}
	return false, nil, nil
}
