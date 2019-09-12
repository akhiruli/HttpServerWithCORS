package main

import (
   "net/http"
   "github.com/gorilla/handlers"
   "github.com/gorilla/mux"
   "io/ioutil"
   "fmt"
)

func HttpHandler(client http.ResponseWriter, req *http.Request) {
   fmt.Println(req.Method) // print the HTTP method like GET, POST, PUT

   if req.Method == "POST" && req.Body != nil {
      b, err := ioutil.ReadAll(req.Body)
      if err == nil{
         fmt.Println(string(b)) //printing the body in case of POST
         http.Error(client, string(b), 200) // sending the same body as response
      }
      return
   }

   http.Error(client, "Something", 200) // sending response with something if body is nil
}

func startServer(){
   fmt.Println("Starting student tracking server.............!")

   router := mux.NewRouter()
   headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
   methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
   origins := handlers.AllowedOrigins([]string{"*"})

   router.HandleFunc("/", HttpHandler)

   listenSocket := "127.0.0.1" + ":" + "12345" //IP and PORT
   //listenSocket := ":" + "12345" //It will listen on *:12345

   err := http.ListenAndServe(listenSocket, handlers.CORS(headers, headers, methods, origins)(router))
   if err != nil {
      fmt.Println("Failed to start HTTP server....!")
   }
}

func main(){
   startServer()
}
