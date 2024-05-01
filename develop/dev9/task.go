package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	url := "https://automaticvertical.ru/"
	fileName := getFileName(url)

	// Скачиваем страницу
	err := downloadPage(url, fileName)
	if err != nil {
		fmt.Println("Error downloading page:", err)
		return
	}

	fmt.Println("Website downloaded successfully to:", fileName)
}

// Функция для определения имени файла на основе URL
func getFileName(url string) string {
	fileName := filepath.Base(url)
	// Добавляем расширение .html, если его нет
	if !strings.Contains(fileName, ".html") {
		fileName += ".html"
	}
	return fileName
}

// Функция для загрузки страницы
func downloadPage(url, fileName string) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return nil
}
