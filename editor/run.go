package main

import (
	"net/http"
	"os"
	"log"
	"html/template"
	"time"
	"strings"
	"encoding/json"
	"io"
	"os/exec"
	"fmt"
	"bufio"
)

func main() {
	//baseDir, _ := os.Getwd()

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))


	http.HandleFunc("/", index)
	http.HandleFunc("/create", create)
	http.ListenAndServe(":1314", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		t, err := template.ParseFiles("index.html")
		if (err != nil) {
			log.Println(err)
		}
		t.Execute(w, nil)
	} else {
		NotFoundHandler(w, r);
	}

}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "404.html", http.StatusNotFound)
}

func create(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/create" {
		generateMd(w, r)
		io.WriteString(w, "<h1 style=\"text-align: center;margin-top: 300px;\">发布中..<h1>")
		publish(w, r)
	} else {
		NotFoundHandler(w, r);
	}


}
func publish(w http.ResponseWriter, r *http.Request) {
	cmdargs := [] string {"..\\develop.bat"}
	if (execCommand("call", cmdargs)) {
		http.Redirect(w, r , "success.html", http.StatusFound)
	}
}

func generateMd(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil{
		panic(err)
	}
	title := r.PostFormValue("title")
	tags := r.PostFormValue("tags")
	categories := r.PostFormValue("categories")
	mdcontent := r.PostFormValue("mdcontent")

	pubDate := time.Now().Format("2006-01-02T15:04:05+08:00")

	filename := title + ".md"

	sub := "+++" + "\n"
	sub = sub + "date = " + pubDate + "\n"
	sub = sub + "title = " + title + "\n"
	jsonTags, _ := json.Marshal(strings.Split(tags, ","))
	sub = sub + "tags = " + string(jsonTags) + "\n"
	jsonCategories, _ := json.Marshal(strings.Split(categories, ","))
	sub = sub + "categories = " + string(jsonCategories) + "\n"
	sub = sub + "+++" + "\n\n"

	mdcontent = sub + mdcontent


	createfile(filename, mdcontent)
}

func createfile(filename string, str_content string)  {
	fd,_:=os.OpenFile(filename,os.O_RDWR|os.O_CREATE,0644)
	buf:=[]byte(str_content)
	fd.Write(buf)
	fd.Close()
}

func execCommand(commandName string, params []string) bool {
	cmd := exec.Command(commandName, params...)

	//显示运行的命令
	fmt.Println(cmd.Args)

	stdout, err := cmd.StdoutPipe()

	if err != nil {
		fmt.Println(err)
		return false
	}

	cmd.Start()

	reader := bufio.NewReader(stdout)

	//实时循环读取输出流中的一行内容
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		fmt.Println(line)
	}

	cmd.Wait()
	return true
}