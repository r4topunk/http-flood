package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// Logger é um middleware que registra as requisições HTTP
func Logger(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		inner.ServeHTTP(w, r)
		log.Printf(
			"%s %s %s %s",
			r.Method,
			r.RequestURI,
			r.RemoteAddr,
			time.Since(start),
		)
	})
}

// handler padrão que responde com "Hello, World!"
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func main() {
	// Configura o logger para registrar em um arquivo
	file, err := os.OpenFile("requests.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Erro ao abrir o arquivo de log: %v", err)
	}
	log.SetOutput(file)

	// Configura o servidor HTTP
	mux := http.NewServeMux()
	mux.HandleFunc("/", helloHandler)

	loggedMux := Logger(mux)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      loggedMux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	log.Println("Servidor iniciado na porta 8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
