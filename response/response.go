package response

import (
	"encoding/json"
	"net/http"

	"github.com/fandrien/book-cabin/constant"
	"github.com/fandrien/book-cabin/model"
)

func WriteSuccess(
	w http.ResponseWriter,
	httpStatus int,
	message string,
	data interface{},
) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)

	resp := model.APIResponse{
		Code:    constant.SuccessCode,
		Message: message,
		Data:    data,
	}

	_ = json.NewEncoder(w).Encode(resp)
}

func WriteError(
	w http.ResponseWriter,
	httpStatus int,
	code string,
	message string,
) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)

	resp := model.APIResponse{
		Code:    code,
		Message: message,
	}

	_ = json.NewEncoder(w).Encode(resp)
}
