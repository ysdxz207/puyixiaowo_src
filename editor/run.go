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
	"encoding/base64"
	"io/ioutil"
	"runtime"
)

type Result struct {
	Status   bool    `json:"status"`
	Msg      string   `json:"msg"`
	Title    string   `json:"title"`
	Filename string   `json:"filename"`
	Ext      string   `json:"ext"`
}

var server_config = map[string]string{
	"address": "http://localhost",
	"port": "1314",
}

func main() {
	//baseDir, _ := os.Getwd()
	go startPage()
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/article_images/", http.StripPrefix("/article_images/", http.FileServer(http.Dir("../content/article_images/"))))

	http.HandleFunc("/", index)
	http.HandleFunc("/create", create)
	http.HandleFunc("/published", published)
	http.HandleFunc("/upload", upload)

	http.ListenAndServe(":" + server_config["port"], nil)

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

func published(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/published" {
		t, err := template.ParseFiles("published.html")
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
		err := r.ParseForm()
		if err != nil {
			panic(err)
		}
		isPublish := r.PostFormValue("isPublish")

		generateMd(w, r)

		if (isPublish == "true") {
			publish(w, r)
		}
		http.Redirect(w, r, "/published", http.StatusFound)
	} else {
		NotFoundHandler(w, r);
	}

}

func upload(w http.ResponseWriter, r *http.Request) {
	var jsonRe Result
	if r.URL.Path == "/upload" {
		err := r.ParseForm()
		if err != nil {
			panic(err)
		}
		image := r.PostFormValue("image")
		filename := r.PostFormValue("filename")
		title := r.PostFormValue("title")
		ext := r.PostFormValue("ext")

		ddd, _ := base64.StdEncoding.DecodeString(image) //成图片文件并把文件写入到buffer
		path := "..\\content\\article_images\\" + title + "\\"
		os.MkdirAll(path, os.ModePerm)
		ioutil.WriteFile(path + filename + ext, ddd, 0666)   //buffer输出到jpg文件中（不做处理，直接写到文件）

		jsonRe.Msg = "成功"
		jsonRe.Status = true
		jsonRe.Title = title
		jsonRe.Filename = filename
		jsonRe.Ext = ext

		jr, err := json.Marshal(jsonRe)
		io.WriteString(w, string(jr))
	} else {
		NotFoundHandler(w, r);
	}
}

func publish(w http.ResponseWriter, r *http.Request) {
	cmdargs := [] string{"..\\deploy.bat"}
	execCommand("start", cmdargs)
}

func generateMd(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}
	title := r.PostFormValue("title")
	tags := r.PostFormValue("tags")
	categories := r.PostFormValue("categories")
	mdcontent := r.PostFormValue("mdcontent")

	pubDate := time.Now().Format("2006-01-02T15:04:05+08:00")

	filename := "..\\content\\post\\" + title + ".md"

	sub := "+++" + "\n"
	sub = sub + "date = \"" + pubDate + "\"\n"
	sub = sub + "title = \"" + title + "\"\n"
	jsonTags, _ := json.Marshal(strings.Split(tags, ","))
	sub = sub + "tags = " + string(jsonTags) + "\n"
	jsonCategories, _ := json.Marshal(strings.Split(categories, ","))
	sub = sub + "categories = " + string(jsonCategories) + "\n"
	sub = sub + "+++" + "\n\n"

	mdcontent = sub + mdcontent

	createfile(filename, mdcontent)
}

func createfile(filename string, str_content string) {

	fd, _ := os.OpenFile(filename, os.O_RDWR | os.O_CREATE, 0644)
	buf := []byte(str_content)
	fd.Write(buf)
	fd.Close()
}

func execCommand(commandName string, params []string) (bool, error) {
	cmd := exec.Command(commandName, params...)

	//显示运行的命令
	fmt.Println(cmd.Args)

	stdout, err := cmd.StdoutPipe()

	if err != nil {
		fmt.Println(err)
		return false, err
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
	return true, nil
}

func startPage() {
	addr := server_config["address"] + ":" + server_config["port"]
	for {
		time.Sleep(time.Second)

		log.Println("Checking if started...")
		resp, err := http.Get(addr)
		if err != nil {
			//log.Println("Failed:", err)
			continue
		}
		resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			//log.Println("Not OK:", resp.StatusCode)
			continue
		}

		// Reached this point: server is up and running!
		break
	}
	log.Printf("server has started")
	log.Printf("Press Ctrl + C to stop server.")
	open(addr)
}

func open(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}
	args = append(args, url)
	return exec.Command(cmd, args...).Start()
}