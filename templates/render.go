package templates


import (
	"html/template"
	"net/http"
)

func LoginHandleTmpl (w http.ResponseWriter, r *http.Request) {
		
		tmpl, err := template.ParseFiles("templates/login.html" )
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	func LoginSuccessHandleTmpl (w http.ResponseWriter, r *http.Request) {
		
		tmpl, err := template.ParseFiles("templates/login-successful.templ" )
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	func LoginFailedHandleTmpl (w http.ResponseWriter, r *http.Request) {
		
		tmpl, err := template.ParseFiles("templates/login-failed.templ" )
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}


