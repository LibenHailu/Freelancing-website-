package main

import (
	"html/template"
	"net/http"

	"github.com/LibenHailu/fjobs/client/delivery/http/handler"
	"github.com/gorilla/mux"
)

var templ = template.Must(template.ParseGlob("../../ui/templates/*"))

func main() {
	userHandler := handler.NewUserHandler(templ)
	applyHandler := handler.NewApplyHandler(templ)
	router := mux.NewRouter()
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("../../ui/assets"))))
	router.HandleFunc("/", userHandler.Index)
	router.HandleFunc("/indexmain", userHandler.IndexAuth)
	router.HandleFunc("/login", userHandler.Login)
	router.HandleFunc("/signup", userHandler.SignUp)
	router.HandleFunc("/post", userHandler.Post)
	// this is for search
	router.HandleFunc("/search", userHandler.Search)
	router.HandleFunc("/searchauth", userHandler.SearchAuth)
	// router.HandleFunc("/home", userHandler.Home)
	// router.HandleFunc("/healthcenters", userHandler.Healthcenters)
	// router.HandleFunc("/logout", userHandler.Logout)
	// this for apply
	router.HandleFunc("/apply", applyHandler.Apply)
	http.ListenAndServe(":8282", router)
}

// import (
// 	"net/http"
// 	"text/template"
// )

// var tmpl = template.Must(template.ParseGlob("../../ui/templates/*.html"))

// func main() {

// 	mux := http.NewServeMux()
// 	fs := http.FileServer(http.Dir("ui/assets"))
// 	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

// 	mux.HandleFunc("/", index)
// 	mux.HandleFunc("/login", loginUser)
// 	mux.HandleFunc("/signup", signupUser)

// 	http.ListenAndServe(":8282", mux)
// }

// func index(w http.ResponseWriter, r *http.Request) {
// 	tmpl.ExecuteTemplate(w, "index.layout", nil)
// }
// func signupUser(w http.ResponseWriter, r *http.Request) {
// 	tmpl.ExecuteTemplate(w, "signup.layout", nil)
// }
// func loginUser(w http.ResponseWriter, r *http.Request) {
// 	// idraw := r.FormValue("id")
// 	// id, err := strconv.Atoi(idraw)

// 	// if err != nil {
// 	// 	w.WriteHeader(http.StatusNoContent)
// 	// 	tmpl.ExecuteTemplate(w, "error.layout", nil)
// 	// }

// 	// user, err := data.FetchUser(id)

// 	// if err != nil {
// 	// 	w.WriteHeader(http.StatusNoContent)
// 	// 	tmpl.ExecuteTemplate(w, "error.layout", nil)
// 	// }
// 	// tmpl.ExecuteTemplate(w, "user.layout", user)
// 	tmpl.ExecuteTemplate(w, "login.layout", nil)
// }

// // func allUsers(w http.ResponseWriter, r *http.Request) {
// // 	pageraw := r.FormValue("page")
// // 	page, err := strconv.Atoi(pageraw)

// // 	if err != nil {
// // 		w.WriteHeader(http.StatusNoContent)
// // 		tmpl.ExecuteTemplate(w, "error.layout", nil)
// // 	}

// // 	users, err := data.FetchUsers(page)

// // 	if err != nil {
// // 		w.WriteHeader(http.StatusNoContent)
// // 		tmpl.ExecuteTemplate(w, "error.layout", nil)
// // 	}

// // 	tmpl.ExecuteTemplate(w, "users.layout", users)
// // }
