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
	// Чтение файла, указываю абсолютный путь(иначе не находит)
	file, err := os.ReadFile("C:/Users/Alex/Dev/6sprintFinal/index.html")
	if err != nil {
		http.Error(w, "Error reading file: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	_, err = w.Write(file)
	if err != nil {
		http.Error(w, "Error write file", http.StatusInternalServerError)
		return
	}
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
	newFilePath := "C:/Users/Alex/Dev/6sprintFinal/" + time.Now().UTC().Format("2006-01-02_15-04-05") + filepath.Ext(handler.Filename)
	newFile, err := os.OpenFile(newFilePath, os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		http.Error(w, "Error open newfile: "+err.Error(), http.StatusInternalServerError)
		return
	}

	defer newFile.Close()

	// Метод используется для записи данных в файл на сервере
	_, err = newFile.Write([]byte(convertedString))
	if err != nil {
		http.Error(w, "Error accessing the server", http.StatusInternalServerError)
		return
	}
	// Метод для отправки данных обратно клиенту через веб-ответ
	_, err = w.Write([]byte(convertedString))
	if err != nil {
		http.Error(w, "Error write response", http.StatusInternalServerError)
		return
	}
}
