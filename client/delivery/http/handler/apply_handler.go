package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/LibenHailu/fjobs/client/service"
)

type ApplyHandler struct {
	templ *template.Template
}

func NewApplyHandler(tmlp *template.Template) *ApplyHandler {
	return &ApplyHandler{templ: tmlp}
}
func (ah *ApplyHandler) Apply(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		jo := r.URL.Query().Get("id")
		joint, _ := strconv.ParseInt(jo, 10, 0)
		joid := uint(joint)
		data, err := service.GetJobTo(joid)
		fmt.Println(data)	
		if err != nil {
			panic(err)
		}
		ah.templ.ExecuteTemplate(w, "apply.layout", data)
		// http.Redirect(w, r, "http://localhost:8282/indexmain", http.StatusSeeOther)

	}
}
