package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

type TowarModel struct {
	Name string
	Price int
}

func handleFunc(w http.ResponseWriter, req *http.Request) {
	pathArgs, err := parsePathPrefix(req.URL.Path[1:])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("URL path is incorrect"))
		return
	}

	var jsonContent []TowarModel
	sum := 0
	response := "<table border=\"1\">"
	for _, each := range pathArgs {
		_, dir, _, _ := runtime.Caller(0)
		dir = filepath.Dir(dir)
		data, err := ioutil.ReadFile(dir + "/" + each + ".json")
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		err = json.Unmarshal(data, &jsonContent)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		response += towarModelToHtml(jsonContent)
		sum += sumTowarModelPrice(jsonContent)
	}
	response += "</table><div>" + strconv.Itoa(sum) + "</div>"

	fmt.Fprintf(w, "%s", response)
}

func parsePathPrefix(path string) ([]string, error) {
	if len(path) <= 0 {
		return nil, errors.New("")
	}

	args := strings.Split(path, "_")
	for _, each := range args {
		if len(each) > 1 {
			return nil, errors.New("")
		}
	}

	return args, nil
}

func towarModelToHtml(list []TowarModel) string {
	formatResult := ""
	for _, each := range list {
		formatResult +=  "<tr>" +
			"<td>" + each.Name + "</td>" +
			"<td>" + strconv.Itoa(each.Price) + "</td>" +
			"</tr>"
	}


	return formatResult
}

func sumTowarModelPrice(list []TowarModel) int {
	sum := 0
	for _, each := range list {
		sum += each.Price
	}
	return sum
}

func main() {
	http.HandleFunc("/", handleFunc)
	log.Panic(http.ListenAndServe(":8080", nil))
}