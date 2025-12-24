package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Variável global para o Pool de conexões (em projetos maiores, injetaríamos isso)
var db *pgxpool.Pool

func main() {
	// 1. Configuração da Conexão (String de conexão do Postgres)
	// Formato: postgres://usuario:senha@host:porta/banco
	connStr := "postgres://kvervandi:strongpassword@localhost:5432/taskmanager"

	var err error
	// Criamos um Pool de conexões (muito mais eficiente que abrir/fechar toda hora)
	db, err = pgxpool.New(context.Background(), connStr)
	if err != nil {
		log.Fatal("Unable to connect to database:", err)
	}
	defer db.Close()

	// Testa se o banco está respondendo
	err = db.Ping(context.Background())
	if err != nil {
		log.Fatal("Could not ping database:", err)
	}
	fmt.Println("Connected to database!")

	// 2. Router Setup (Chi)
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type"},
	}))

	// 3. Rotas
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("System is running. Database is connected."))
	})

	r.Post("/tasks", createTaskHandler)
	r.Get("/tasks", listTasksHandler)
	r.Put("/tasks/{id}", updateTaskHandler)
	r.Delete("/tasks/{id}", deleteTaskHandler)

	// 4. Start Server
	port := "8080"
	fmt.Printf("Server is running at %s\n", port)
	http.ListenAndServe(":"+port, r)
}
