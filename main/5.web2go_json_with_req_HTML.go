package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func sendRequest(w http.ResponseWriter, r *http.Request) {
	// Добавление заголовка Access-Control-Allow-Origin
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Чтение данных из тела запроса
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Ошибка чтения тела запроса", http.StatusBadRequest)
		return
	}

	// Вывод полученных данных
	fmt.Println("Получен запрос:", string(body))

	// Проверка наличия поля "message" в JSON-сообщении
	var requestData map[string]interface{}
	err = json.Unmarshal(body, &requestData)
	if err != nil {
		http.Error(w, "Некорректное JSON-сообщение", http.StatusBadRequest)
		return
	}

	message, exists := requestData["message"]
	if !exists || message == nil {
		http.Error(w, "Некорректное JSON-сообщение: отсутствует или имеет некорректное значение поле 'message'", http.StatusBadRequest)
		return
	}

	// Проверка типа значения поля "message"
	_, ok := message.(string)
	if !ok {
		http.Error(w, "Некорректное JSON-сообщение: поле 'message' должно быть строкой", http.StatusBadRequest)
		return
	}

	// "Отправка" JSON-ответа (в реальности сервер должен отправлять ответ веб-сервису)
	response := map[string]interface{}{
		"status":  "success",
		"message": "Данные успешно получены",
	}

	// Преобразование данных в формат JSON
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Ошибка преобразования данных в JSON", http.StatusInternalServerError)
		return
	}

	// Запись JSON-ответа в тело HTTP-ответа
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func main() {
	// Зарегистрируйте обработчик запросов
	http.HandleFunc("/", sendRequest)

	// Запуск HTTP-сервера на порту 8080
	fmt.Println("Сервер запущен на порту 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Ошибка при запуске сервера:", err)
	}
}
