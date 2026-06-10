package view

import (
	"fmt"
	"html/template"
	"log/slog"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, fileName string, data interface{}) error {
	path := fmt.Sprintf("templates/%s", fileName)
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		slog.Warn("Ошибка парсинга страницы: " + err.Error())
		return fmt.Errorf("Возникла ошибка парсинга во время шаболнизации страницы")
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		slog.Warn(fmt.Sprintf("Ошибка шаблонизации страницы %v", err.Error()))
		return fmt.Errorf("Возникла ошибка шаблонизации страницы")
	}
	return nil
}

func RenderPartOfTemplate(w http.ResponseWriter, fileName string, data interface{}) error {
	path := fmt.Sprintf("templates/%s", fileName)

	tmpl, err := template.ParseFiles(path)
	if err != nil {
		slog.Warn("Ошибка парсинга страницы: " + err.Error())
		return fmt.Errorf("Возникла ошибка парсинга во время шаболнизации страницы")
	}

	err = tmpl.ExecuteTemplate(w, fileName, data)
	if err != nil {
		slog.Warn(fmt.Sprintf("Ошибка шаблонизации страницы %v", err.Error()))
		return fmt.Errorf("Возникла ошибка шаблонизации страницы")
	}
	return nil
}

func RenderSeveralTemplates(w http.ResponseWriter, firstFileName string, secondFileName string, data interface{}) error {
	firstPath := fmt.Sprintf("templates/%s", firstFileName)
	secondPath := fmt.Sprintf("templates/%s", secondFileName)
	tmpl, err := template.ParseFiles(firstPath, secondPath)
	if err != nil {
		slog.Warn("Ошибка парсинга страницы: " + err.Error())
		return fmt.Errorf("Возникла ошибка парсинга во время шаболнизации страницы")
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		slog.Warn(fmt.Sprintf("Ошибка шаблонизации страницы %v", err.Error()))
		return fmt.Errorf("Возникла ошибка шаблонизации страницы")
	}
	return nil
}
