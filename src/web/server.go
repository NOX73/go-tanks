package web

import (
  "net/http"
  "log"
  "html/template"
  "go/build"
  "path"
  "github.com/gorilla/websocket"
)

type Server struct {
  templates   map[string]*template.Template
  wsHandler   func(*websocket.Conn)
}

func NewServer ( wsHandler func(*websocket.Conn) ) *Server {
  return &Server{ templates: make(map[string]*template.Template), wsHandler: wsHandler }
}

func webPath () string {
  return path.Join( build.Default.GOPATH, "src/web" )
}

func viewPath ( filename string ) string {
  return  path.Join( webPath(), "app/views", filename )
}

func publicDir () http.Dir {
  return http.Dir(path.Join( webPath(), "public" ))
}

func ( s *Server ) Run () {
  s.parseTemplates()

  // Root path
  http.HandleFunc("/", s.handler)

  // Static files
  http.Handle( "/public/", http.StripPrefix("/public", http.FileServer(publicDir())) )

  // WebSocket
  http.HandleFunc("/ws", s.websocket)

  http.ListenAndServe(":9000", nil)
}

func ( s *Server ) handler ( w http.ResponseWriter, r *http.Request ) {
  err := s.templates["layout"].ExecuteTemplate(w, "index.html", s)
  if err != nil { log.Fatal(err) }
}

func ( s *Server ) websocket ( w http.ResponseWriter, r *http.Request ) {
  ws, _ := websocket.Upgrade(w, r, nil, 1024, 1024)

  s.wsHandler(ws)
}

func ( s *Server ) parseTemplates () {

  var t *template.Template

  t = template.New("layout")
  t.Funcs(template.FuncMap{"ng": func(s string)(string){return "{{" + s +"}}"}})

  _, err := t.ParseGlob(viewPath("*.html"))
  if err != nil { log.Fatal(err) }

  s.templates["layout"] = t

}
