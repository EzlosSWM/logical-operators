package server

import (
	"net/http"
	"strconv"
	"time"
)

const ON = "1"

func (s *Server) RegisterRoutes() http.Handler {

	mux := http.NewServeMux()

	// Static file server
	fs := http.FileServer(http.Dir("static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	mux.HandleFunc("/", s.handler)
	mux.HandleFunc("/and-operator", s.andOperator)
	mux.HandleFunc("/or-operator", s.orOperator)
	mux.HandleFunc("/not-operator", s.notOperator)

	return mux
}

func (s *Server) handler(w http.ResponseWriter, r *http.Request) {
	year := strconv.Itoa(time.Now().Year())
	s.templ.ExecuteTemplate(w, "index.html", map[string]string{"Year": year})
}

func (s *Server) andOperator(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		a := r.FormValue("a")
		b := r.FormValue("b")

		if (a == ON) && (b == ON) {
			s.templ.ExecuteTemplate(w, "bulb", map[string]bool{"SwapOOB": true, "IsOn": true})
		}

		if (a != ON) || (b != ON) {
			s.templ.ExecuteTemplate(w, "bulb", map[string]bool{"SwapOOB": true, "IsOn": false})
		}
	}

	s.templ.ExecuteTemplate(w, "bulb", nil)
	s.templ.ExecuteTemplate(w, "and-operator", nil)
}

func (s *Server) orOperator(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		a := r.FormValue("a")
		b := r.FormValue("b")

		if (a == ON) || (b == ON) {
			s.templ.ExecuteTemplate(w, "bulb", map[string]bool{"SwapOOB": true, "IsOn": true})
		}

		if (a != ON) && (b != ON) {
			s.templ.ExecuteTemplate(w, "bulb", map[string]bool{"SwapOOB": true, "IsOn": false})
		}
	}

	s.templ.ExecuteTemplate(w, "bulb", nil)
	s.templ.ExecuteTemplate(w, "or-operator", nil)
}

func (s *Server) notOperator(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		a := r.FormValue("a")

		if (a != ON) {
			s.templ.ExecuteTemplate(w, "bulb", map[string]bool{"SwapOOB": true})
		} else {
			s.templ.ExecuteTemplate(w, "bulb", map[string]bool{"SwapOOB": true, "IsOn": true})
		}
	}

	s.templ.ExecuteTemplate(w, "bulb", map[string]bool{"IsOn": true})
	s.templ.ExecuteTemplate(w, "not-operator", nil)
}
