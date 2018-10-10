package main

import (
	"net/http"
	"sync"
	"html/template"
	"path/filepath"
	"log"
	"flag"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/google"
	"github.com/stretchr/objx"
)

var avatars Avatar = UseFileSystemAvatar

type templateHandler struct {
	once sync.Once
	filename string
	templ	 *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	t.once.Do(func(){
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates",t.filename)))
	})
	data := map[string] interface{}{
		"Host": r.Host,
	}
	if authCookie, err := r.Cookie("auth"); err == nil {
		data["UserData"] = objx.MustFromBase64(authCookie.Value)
	}
	t.templ.Execute(w, data)
}

func main() {

	var addr = flag.String("addr", ":8080","The address of the app")
	flag.Parse()

	var clientID = "213972874290-b8olmgl0mbvmvce6om43nudjrjdpjc4r.apps.googleusercontent.com"
	var clientSecret = "Xfd2LlZ5HkQUBIC1OSGSRxFt"
	gomniauth.SetSecurityKey("Xfd2LlZ5HkQUBIC1OSGSRxFt")
	gomniauth.WithProviders(google.New(clientID,clientSecret,"http://localhost:8080/auth/callback/google"))


	r := newRoom(avatars)
	http.Handle("/chat", MustAuth(&templateHandler{filename:"chat.html"}))
	http.Handle("/login",&templateHandler{filename:"login.html"})
	http.HandleFunc("/auth/",logindHandler)
	http.Handle("/room", r)
	http.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		http.SetCookie(w, &http.Cookie{
			Name: "auth",
			Value: "",
			Path: "/",
			MaxAge: -1,
		})
		w.Header().Set("Location","/chat")
		w.WriteHeader(http.StatusTemporaryRedirect)
	})
	http.Handle("/upload", &templateHandler{filename:"upload.html"})
	http.HandleFunc("/uploader", uploaderHandler)
	http.Handle("/avatars/",
		http.StripPrefix("/avatars/",
			http.FileServer(http.Dir("./avatars"))))
	go r.run()

	println("Starting web server on ",*addr)
	if err := http.ListenAndServe(*addr,nil); err != nil{
		log.Fatal(err)
	}
}
