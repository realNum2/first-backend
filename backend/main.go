package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"
)

type Message struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}

var (
	messages []Message
	nextID   = 1
	mu       sync.Mutex
)

func main() {
	http.HandleFunc("/messages", cors(messagesHandlers))
	http.HandleFunc("/", cors(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Message bord API is Running!"))
	}))

	log.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func messagesHandlers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		mu.Lock()
		defer mu.Unlock()
		json.NewEncoder(w).Encode(messages)

	case http.MethodPost:
		var input struct {
			Name string `json:"name"`
			Text string `json:"text"`
		}

		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}
		if input.Name == "" || input.Text == "" {
			http.Error(w, "Name and text are required", http.StatusBadRequest)
			return
		}

		mu.Lock()
		msg := Message{
			ID:        nextID,
			Name:      input.Name,
			Text:      input.Text,
			CreatedAt: time.Now().UTC(),
		}
		nextID++
		messages = append(messages, msg)
		mu.Unlock()

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(msg)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func cors(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next(w, r)
	}
}
