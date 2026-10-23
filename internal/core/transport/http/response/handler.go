package core_http_response

import (
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"log/slog"
	"net/http"

	core_logger "github.com/jabrail059/weather-dashboard/internal/core/logger"
	core_errors "github.com/jabrail059/weather-dashboard/internal/core/transport/http/errors"
	"go.uber.org/zap"
)

type HTTPResponseHandler struct {
	log *core_logger.Logger
	w   http.ResponseWriter
}

func NewHTTPResponseHandler(
	log *core_logger.Logger,
	w http.ResponseWriter,
) *HTTPResponseHandler {
	return &HTTPResponseHandler{
		log: log,
		w:   w,
	}
}

func (h *HTTPResponseHandler) JSONResponse(
	responseBody any,
	statusCode int,
) {
	h.w.WriteHeader(statusCode)

	if err := json.NewEncoder(h.w).Encode(responseBody); err != nil {
		h.log.Error("write HTTP response", zap.Error(err))
	}
}

func (h *HTTPResponseHandler) NoContentResponse() {
	h.w.WriteHeader(http.StatusNoContent)
}

func (h *HTTPResponseHandler) PanicResponse(p any, msg string) {
	statusCode := http.StatusInternalServerError
	err := fmt.Errorf("unexpected panic: %v", p)

	h.log.Error(msg, zap.Error(err))

	h.errorResponse(statusCode, err, msg)
}

func (h *HTTPResponseHandler) ErrorResponse(err error, msg string) {
	var (
		statusCode int
		logFunc    func(string, ...zap.Field)
	)

	switch {
	case errors.Is(err, core_errors.ErrNotFound):
		statusCode = http.StatusNotFound
		logFunc = h.log.Debug
	case errors.Is(err, core_errors.ErrInvalidArgument):
		statusCode = http.StatusBadRequest
		logFunc = h.log.Warn
	case errors.Is(err, core_errors.ErrConflict):
		statusCode = http.StatusConflict
		logFunc = h.log.Warn
	default:
		statusCode = http.StatusInternalServerError
		logFunc = h.log.Error
	}

	logFunc(msg, zap.Error(err))

	h.errorResponse(statusCode, err, msg)
}

func (h *HTTPResponseHandler) errorResponse(
	statusCode int,
	err error,
	msg string,
) {
	response := ErrorResponse{
		Error:   err.Error(),
		Message: msg,
	}

	h.JSONResponse(response, statusCode)
}

func (h *HTTPResponseHandler) HTMXMethodErrorResponse() {
	statusCode := http.StatusBadRequest

	response := ErrorResponse{
		Error:   "no HX-Request in request header",
		Message: "Invalid HTMX request header value",
	}

	h.JSONResponse(response, statusCode)
}

func (h *HTTPResponseHandler) RenderPartHTML(data any, file string) error {
	path := fmt.Sprintf("public/templates/%s", file)
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		slog.Warn("ошибка парсинга страницы: " + err.Error())
		return fmt.Errorf("возникла ошибка парсинга во время шаблонизации страницы")
	}

	err = tmpl.ExecuteTemplate(h.w, file, data)
	if err != nil {
		slog.Warn(fmt.Sprintf("ошибка шаблонизации страницы %v", err.Error()))
		return fmt.Errorf("возникла ошибка шаблонизации страницы")
	}
	return nil
}

func (h *HTTPResponseHandler) RenderHTML(data any, file []string) error {
	if len(file) == 0 {
		return fmt.Errorf("no HTML files to render")
	}

	paths := make([]string, len(file))
	for i, path := range file {
		path = fmt.Sprintf("public/templates/%s", path)
		paths[i] = path
	}

	tmpl, err := template.ParseFiles(paths...)
	if err != nil {
		return fmt.Errorf("template parse files: %w", err)
	}

	if err := tmpl.Execute(h.w, data); err != nil {
		return fmt.Errorf("template execute: %w", err)
	}

	return nil
}
