package main

import (
	"dk-harat-backend-go/handlers"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Разрешаем CORS (чтобы Angular мог подключаться)
func enableCORS(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func main() {
	// Роут для API
	http.HandleFunc("/api/hello", func(w http.ResponseWriter, r *http.Request) {
		enableCORS(&w) // Включаем CORS

		// Отправляем JSON-ответ
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Hello from Go!",
		})
	})

	// Запускаем сервер на порту 8080
	http.ListenAndServe(":8080", nil)
	r := gin.Default()

	r.POST("/api/register", handlers.Register)
	r.POST("/api/login", handlers.Login)

	r.Run(":8080")
}
