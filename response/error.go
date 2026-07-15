package response

import (
	"encoding/json"
	"net/http"

	"github.com/fandrien/book-cabin/model"
)

func WriteError(
	w http.ResponseWriter,
	httpStatus int,
	code string,
	message string,
) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)

	_ = json.NewEncoder(w).Encode(model.ErrorResponse{
		Code:    code,
		Message: message,
	})
}
