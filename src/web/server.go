package web

import (
  "net/http"
  "log"
  "html/template"
)

type Server struct {
  templates     map[string]*template.Template
}

func NewServer () *Server {
  return &Server{ templates: make(map[string]*template.Template)}
}

func viewPath ( filename string ) string {
  return  "/Users/nox73/Dropbox/proj/go-tanks/src/web/app/views/" + filename
}

func publicDir () http.Dir {
  return http.Dir("/Users/nox73/Dropbox/proj/go-tanks/src/web/app/public/")
}

func ( s *Server ) Run () {
  s.parseTemplates()
  http.HandleFunc("/", s.handler)
  static := http.FileServer(publicDir())
  http.HandleFunc("/public", static.ServeHTTP)
  http.ListenAndServe(":9000", nil)
}

func ( s *Server ) handler ( w http.ResponseWriter, r *http.Request ) {
  err := s.templates["layout"].ExecuteTemplate(w, "index.html", s)
  if err != nil { log.Fatal(err) }
}

func ( s *Server ) parseTemplates () {

  var t *template.Template

  t = template.New("layout")
  t.Funcs(template.FuncMap{"ng": func(s string)(string){return "{{" + s +"}}"}})

  _, err := t.ParseGlob(viewPath("*.html"))
  if err != nil { log.Fatal(err) }

  s.templates["layout"] = t

}
