package main

import (
	"net/http"
	"os"
	"time"
	"strings"
	"log"
	"html/template"
)

func main() {
	http.Handle("/css/", http.FileServer(http.Dir("/")))
	http.Handle("/js/", http.FileServer(http.Dir("/")))


	http.HandleFunc("/", index)
	http.HandleFunc("/create/", create)
	http.ListenAndServe(":1314", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("index.html")
	if (err != nil) {
		log.Println(err)
	}
	t.Execute(w, nil)
}

func create(w http.ResponseWriter, r *http.Request) {
	tracefile("hahaha")
}

func tracefile(str_content string)  {
	fd,_:=os.OpenFile("a.txt",os.O_RDWR|os.O_CREATE|os.O_APPEND,0644)
	fd_time:=time.Now().Format("2006-01-02 15:04:05");
	fd_content:=strings.Join([]string{"======",fd_time,"=====",str_content,"\n"},"")
	buf:=[]byte(fd_content)
	fd.Write(buf)
	fd.Close()
}
