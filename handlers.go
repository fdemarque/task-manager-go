package main

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
)

// createTaskHandler lida com requisições POST /tasks
func createTaskHandler(w http.ResponseWriter, r *http.Request) {
	var req CreateTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.Title == "" {
		http.Error(w, "Title is required", http.StatusBadRequest)
		return
	}
	if req.Effort < 1 || req.Effort > 5 {
		http.Error(w, "Effort must be between 1 and 5", http.StatusBadRequest)
		return
	}

	sql := `
		INSERT INTO tasks (title, description, due_date, effort, difficulty, requirements, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id
	`

	var newID int
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

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]int{"id": newID})
}

// listTasksHandler retorna todas as tarefas
func listTasksHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query(context.Background(), "SELECT id, title, description, due_date, effort, difficulty, requirements, created_at FROM tasks")
	if err != nil {
		http.Error(w, "Database error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	tasks := []Task{}
	for rows.Next() {
		var t Task
		if err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.DueDate, &t.Effort, &t.Difficulty, &t.Requirements, &t.CreatedAt); err != nil {
			http.Error(w, "Error scanning row: "+err.Error(), http.StatusInternalServerError)
			return
		}
		tasks = append(tasks, t)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

// updateTaskHandler atualiza uma tarefa existente
func updateTaskHandler(w http.ResponseWriter, r *http.Request) {
	// 1. Pegar o ID da URL
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// 2. Ler o JSON com os novos dados
	var req CreateTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// 3. Executar o UPDATE
	sql := `
		UPDATE tasks 
		SET title=$1, description=$2, due_date=$3, effort=$4, difficulty=$5, requirements=$6 
		WHERE id=$7
	`
	
	// ExecCommand retorna o resultado da operação
	res, err := db.Exec(context.Background(), sql,
		req.Title, req.Description, req.DueDate, req.Effort, req.Difficulty, req.Requirements, id,
	)
	if err != nil {
		http.Error(w, "Database error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Verifica se alguma linha foi afetada (se o ID existia)
	rowsAffected := res.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Task updated"}`))
}

// deleteTaskHandler remove uma tarefa
func deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	sql := "DELETE FROM tasks WHERE id=$1"
	
	res, err := db.Exec(context.Background(), sql, id)
	if err != nil {
		http.Error(w, "Database error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if res.RowsAffected() == 0 {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Task deleted"}`))
}