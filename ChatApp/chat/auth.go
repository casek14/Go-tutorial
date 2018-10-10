package main

import (
	"net/http"
	"strings"
	"fmt"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/objx"
	"crypto/md5"
	"io"
	"log"
)

type authHandler struct {
	next http.Handler
}

func (h *authHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	_, err := r.Cookie("auth")
	if err == http.ErrNoCookie{
		//not authenticated, redirect to login
		w.Header().Set("Location","/login")
		w.WriteHeader(http.StatusTemporaryRedirect)
		return
	}

	if err != nil{
		//different error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//success call the next handler
	h.next.ServeHTTP(w,r)

	}

	func MustAuth(handler http.Handler) http.Handler{
		return &authHandler{next: handler}
	}


	func logindHandler(w http.ResponseWriter, r *http.Request){
		seqs := strings.Split(r.URL.Path,"/")
		action := seqs[2]
		provider := seqs[3]
		switch action {
		case "login":
			provider, err := gomniauth.Provider(provider)
			if err != nil{
				http.Error(w, fmt.Sprintf("ERROR when trying to get provider %s:%s", provider, err),http.StatusBadRequest)
				return
			}
			loginUrl, err := provider.GetBeginAuthURL(nil,nil)
			if err != nil {
				http.Error(w, fmt.Sprintf("Error when trying to GetBeginAuthURL" +
					"for %s:%s", provider, err), http. StatusInternalServerError)
				return
			}
			w.Header().Set("Location", loginUrl)
			w.WriteHeader(http.StatusTemporaryRedirect)
		case "callback":
			provider, err := gomniauth.Provider(provider)
			if err != nil {
				http.Error(w, fmt.Sprintf("ERROR when trying to get provider " +
					"%s:%s",provider,err),http.StatusBadRequest)
			}
			creds, err := provider.CompleteAuth(objx.MustFromURLQuery(r.URL.RawQuery))
			if err != nil{
				http.Error(w, fmt.Sprintf("Error whent trying to complete auth for " +
					"%s:%s",provider,err),http.StatusInternalServerError)
				return
			}
			user, err := provider.GetUser(creds)
			if err != nil {
				http.Error(w,fmt.Sprintf("Error trying to get user from " +
					"%s:%s", provider,err),http.StatusInternalServerError)
			}
				chatUser := &chatUser{User: user}
				m := md5.New()
				io.WriteString(m, strings.ToLower(user.Email()))
				chatUser.uniqueID = fmt.Sprintf("%x",m.Sum(nil))
				avatarURL, err := avatars.GetAvatarURL(chatUser)
				if err != nil {
					log.Fatalln("Error when trying to GetAvatarURL", "-", err)
				}
				return



			m := md5.New()
			io.WriteString(m, strings.ToLower(user.Email()))
			userID := fmt.Sprintf("%x", m.Sum(nil))
			authCookieValue := objx.New(map[string]interface{}{
				"userid": userID,
				"name": user.Name(),
				"avatar_url": avatarURL,
			}).MustBase64()


			http.SetCookie(w, &http.Cookie{
				Name: "auth",
				Value: authCookieValue,
				Path: "/",
			})
			w.Header().Set("Location", "/chat")
			w.WriteHeader(http.StatusTemporaryRedirect)

		default:
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "Auth action %s not supported", action)
		}

	}