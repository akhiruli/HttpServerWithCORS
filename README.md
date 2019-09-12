# HttpServerWithCORS
A HTTP server which handles CORS.

# Dependency
go get github.com/gorilla/handlers
go get github.com/gorilla/mux

# Just run the below command to run
go run main.go

# some CURL commands
curl -X POST http://127.0.0.1:12345/ -d "enjoy"
curl http://127.0.0.1:12345/
