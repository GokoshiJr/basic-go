package main

func main() {
	server := NewServer(":3000")
	server.Handle("/", "GET", HandleRoot)
	server.Handle("/api", "POST", server.AddMiddleware(HandleHome, CheckAuth(), Logging()))
	server.Handle("/create", "POST", PostRequest)
	server.Listen()
}
