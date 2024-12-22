package main

import (
	"LMC/calc"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

// Запрос от пользователя
type Request struct {
	Expression string `json:"expression"`
}

// Ответ с результатом
type Response struct {
	Result string `json:"result,omitempty"`
	Error  string `json:"error,omitempty"`
}

// Хэндлер для обработки запросов
func calculateHandler(w http.ResponseWriter, r *http.Request) {
	// Устанавливаем Content-Type
	w.Header().Set("Content-Type", "application/json")

	// Проверяем метод запроса
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(Response{Error: "Method not allowed"})
		return
	}

	// Читаем тело запроса
	var req Request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || req.Expression == "" {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(Response{Error: "Expression is not valid"})
		return
	}

	// Вычисляем результат
	result, calcErr := calc.Calc(req.Expression)
	if calcErr != nil {
		if calcErr.Error() == "деление на ноль" || calcErr.Error() == "неизвестный символ" {
			w.WriteHeader(http.StatusUnprocessableEntity)
			json.NewEncoder(w).Encode(Response{Error: "Expression is not valid"})
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(Response{Error: "Internal server error"})
		}
		return
	}

	// Отправляем успешный результат
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response{Result: strconv.FormatFloat(result, 'f', -1, 64)})
}

func main() {
	// Регистрируем обработчик
	http.HandleFunc("/api/v1/calculate", calculateHandler)

	// Запускаем сервер
	log.Println("Starting server on :8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
