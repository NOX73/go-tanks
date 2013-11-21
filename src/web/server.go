package web

import (
  "net/http"
  "log"
  "html/template"
  "go/build"
  "path"
  "code.google.com/p/go.net/websocket"
)

type Server struct {
  templates     map[string]*template.Template
}

func NewServer () *Server {
  return &Server{ templates: make(map[string]*template.Template)}
}

func webPath () string {
  return path.Join( build.Default.GOPATH, "src/web" )
}

func viewPath ( filename string ) string {
  return  path.Join( webPath(), "app/views", filename )
}

func publicDir () http.Dir {
  return http.Dir(path.Join( webPath(), "public" ))
  //return http.Dir("/tmp")
}

func ( s *Server ) Run () {
  s.parseTemplates()

  // Root path
  http.HandleFunc("/", s.handler)

  // Static files
  http.Handle( "/public/", http.StripPrefix("/public", http.FileServer(publicDir())) )

  // WebSocket
  http.Handle("/ws", websocket.Handler(s.websocket))

  http.ListenAndServe(":9000", nil)
}

func ( s *Server ) handler ( w http.ResponseWriter, r *http.Request ) {
  err := s.templates["layout"].ExecuteTemplate(w, "index.html", s)
  if err != nil { log.Fatal(err) }
}

func ( s *Server ) websocket ( ws *websocket.Conn) {
  defer ws.Close()

  ws.Write([]byte("Hello"))
  buff := make([]byte, 255)

  for {
    n, err := ws.Read(buff)
    if err != nil { break }
    ws.Write(buff[0:n])
  }

}

func ( s *Server ) parseTemplates () {

  var t *template.Template

  t = template.New("layout")
  t.Funcs(template.FuncMap{"ng": func(s string)(string){return "{{" + s +"}}"}})

  _, err := t.ParseGlob(viewPath("*.html"))
  if err != nil { log.Fatal(err) }

  s.templates["layout"] = t

}
