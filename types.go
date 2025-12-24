package main

import (
    "time"
)

// Task representa nossa tarefa no sistema
type Task struct {
    ID           int       `json:"id"`
    Title        string    `json:"title"`
    Description  string    `json:"description"`
    DueDate      time.Time `json:"due_date"`
    Effort       int       `json:"effort"`
    Difficulty   string    `json:"difficulty"`
    Requirements []string  `json:"requirements"`
    CreatedAt    time.Time `json:"created_at"`
}

// CreateTaskRequest define o que esperamos receber do Front-end
// (O ID e CreatedAt são gerados pelo banco, então não pedimos aqui)
type CreateTaskRequest struct {
    Title        string    `json:"title"`
    Description  string    `json:"description"`
    DueDate      time.Time `json:"due_date"`
    Effort       int       `json:"effort"`
    Difficulty   string    `json:"difficulty"`
    Requirements []string  `json:"requirements"`
}
