package controllers

import (
	"NetAppealIntelligence/models"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	_ "log"
	"net/http"
	"os"
	"path"
	"runtime/debug"
	_ "strings"
)

var templates = make(map[string]*template.Template)
var Intelligencedatas = models.Init()

const (
	UPLOAD_DIR   = "./uploads"
	TEMPLATE_DIR = "./views"
)

func init() {
	fileInfoArr, err := ioutil.ReadDir(TEMPLATE_DIR)
	if err != nil {
		panic(err)
		return
	}
	var templateName, templatePath string
	for _, fileInfo := range fileInfoArr {
		templateName = fileInfo.Name()
		if ext := path.Ext(templateName); ext != ".html" {
			continue
		}
		templatePath = TEMPLATE_DIR + "/" + templateName
		log.Println("Loading template:", templatePath)
		t := template.Must(template.ParseFiles(templatePath))
		templates[templateName] = t
	}
}

func DisplayIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("views/index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		t.Execute(w, nil)
		return
	}
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("views/upload.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		t.Execute(w, nil)
		return
	}

	if r.Method == "POST" {
		f, h, err := r.FormFile("image")
		if err != nil {
			http.Error(w, err.Error(),
				http.StatusInternalServerError)
			return
		}
		filename := h.Filename
		defer f.Close()
		t, err := os.Create(UPLOAD_DIR + "/" + filename)
		if err != nil {
			http.Error(w, err.Error(),
				http.StatusInternalServerError)
			return
		}
		defer t.Close()
		if _, err := io.Copy(t, f); err != nil {
			http.Error(w, err.Error(),
				http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/view?id="+filename,
			http.StatusFound)
	}
}

func ListfolderHandler(w http.ResponseWriter, r *http.Request) {
	imageId := r.FormValue("id")
	imagePath := UPLOAD_DIR + "/" + imageId
	fmt.Println(imageId)
	fmt.Println(imagePath)
	fmt.Println(r.URL.Path)
	if exists := isExists(imagePath); !exists {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "image")
	http.ServeFile(w, r, imagePath)

}

func SearchListHandler(w http.ResponseWriter, r *http.Request) {
	domain := r.FormValue("id")

	locals_search := make(map[string]interface{})

	Intelligencedata := models.Search(domain)

	locals_search["Data"] = Intelligencedata

	t_search, err := template.ParseFiles("views/search.html")
	if err != nil {
		http.Error(w, err.Error(),
			http.StatusInternalServerError)
		return
	}
	t_search.Execute(w, locals_search)
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("views/search.html")
	if err != nil {
		http.Error(w, err.Error(),
			http.StatusInternalServerError)
		return
	}
	t.Execute(w, "")
}

func ViewHandler(w http.ResponseWriter, r *http.Request) {
	imageId := r.FormValue("id")
	imagePath := UPLOAD_DIR + "/" + imageId
	fmt.Println(imageId)
	fmt.Println(imagePath)
	fmt.Println(r.URL.Path)
	if exists := isExists(imagePath); !exists {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "image")
	http.ServeFile(w, r, imagePath)
}

func isExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return os.IsExist(err)
}

func ListHandler(w http.ResponseWriter, r *http.Request) {

	locals := make(map[string]interface{})

	locals["Data"] = Intelligencedatas

	t, err := template.ParseFiles("views/list.html")
	if err != nil {
		http.Error(w, err.Error(),
			http.StatusInternalServerError)
		return
	}
	t.Execute(w, locals)
}

func renderHtml(w http.ResponseWriter, tmpl string, locals map[string]interface{}) {
	tmpl += ".html"
	err := templates[tmpl].Execute(w, locals)
	check(err)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func SafeHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if e, ok := recover().(error); ok {
				http.Error(w, e.Error(), http.StatusInternalServerError)

				// 或者输出自定义的 50x 错误页面
				// w.WriteHeader(http.StatusInternalServerError)
				// renderHtml(w, "error", e.Error())

				// logging
				log.Println("WARN: panic fired in %v.panic - %v", fn, e)
				log.Println(string(debug.Stack()))
			}
		}()
		fn(w, r)
	}
}
