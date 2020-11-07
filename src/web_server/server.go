package main

type Server struct {
	port string
}

func NewServer(port string) *Server{
	return &Server {
		port: port,
	}
}

func (servidor *Server) Listen() error {
	// puerto, handler
	err := http.ListenAndServe(servidor.port, nil)
	if err != nil {
		return err
	}
	return nil
}