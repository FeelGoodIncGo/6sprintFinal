package handlers

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

// Handler для корневого эндпоинта /
func HomeHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed"+r.Method, http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "index.html")

}

// Handler для эндпоинта /upload
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Парсим форму, устанавливаем максимальный размер загружаемых данных в 10 мегабайт.
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Error parsing form: "+err.Error(), http.StatusBadRequest)
		return
	}
	// Получаем файл из мультипартформы-запроса
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Читаем данные из файла
	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Error reading file", http.StatusInternalServerError)
		return
	}

	// Передаем данные в функцию автоопределения из пакета service
	convertedString, err := (&service.Converter{}).AutoConvert(string(data))
	if err != nil {
		http.Error(w, "Error converting data", http.StatusInternalServerError)
		return
	}

	// Создаем локальный файл и записываем в него результат конвертации
	newFilePath := time.Now().UTC().Format("2006-01-02_15-04-05") + filepath.Ext(handler.Filename)
	err = os.WriteFile(newFilePath, []byte(convertedString), 0644)
	if err != nil {
		http.Error(w, "Error writing to newfile: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	// Метод для отправки данных обратно клиенту через веб-ответ
	_, err = w.Write([]byte(convertedString))
	if err != nil {
		http.Error(w, "Error write response", http.StatusInternalServerError)
		return
	}
}
