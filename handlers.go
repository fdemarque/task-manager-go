package main

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
)

// createTaskHandler lida com requisições POST /tasks
func createTaskHandler(w http.ResponseWriter, r *http.Request) {
	// 1. Decodificar o JSON que vem do Front/Curl
	var req CreateTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// 2. Validar dados básicos (Regra de Negócio)
	if req.Title == "" {
		http.Error(w, "Title is required", http.StatusBadRequest)
		return
	}
	if req.Effort < 1 || req.Effort > 5 {
		http.Error(w, "Effort must be between 1 and 5", http.StatusBadRequest)
		return
	}

	// 3. Inserir no Banco de Dados (SQL Puro)
	// Usamos $1, $2, etc para evitar SQL Injection
	sql := `
		INSERT INTO tasks (title, description, due_date, effort, difficulty, requirements, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id
	`

	var newID int
	// db é nossa variável global definida no main.go
	err := db.QueryRow(context.Background(), sql,
		req.Title,
		req.Description,
		req.DueDate,
		req.Effort,
		req.Difficulty,
		req.Requirements,
		time.Now(),
	).Scan(&newID)

	if err != nil {
		http.Error(w, "Database error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// 4. Retornar o ID criado
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]int{"id": newID})
}
