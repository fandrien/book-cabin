package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/fandrien/book-cabin/constant"
	"github.com/fandrien/book-cabin/model"
	"github.com/fandrien/book-cabin/response"
	"github.com/fandrien/book-cabin/service"
	"github.com/fandrien/book-cabin/validator"
)

type SearchHandler struct {
	searchService *service.SearchService
}

func NewSearchHandler(
	searchService *service.SearchService,
) *SearchHandler {
	return &SearchHandler{
		searchService: searchService,
	}
}

func (h *SearchHandler) Search(
	w http.ResponseWriter,
	r *http.Request,
) {

	var req model.SearchRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {

		response.WriteError(
			w,
			http.StatusBadRequest,
			constant.ErrInvalidRequest,
			"invalid request body",
		)
		return
	}

	if err := validator.ValidateSearchRequest(req); err != nil {
		response.WriteError(
			w,
			http.StatusBadRequest,
			constant.ErrInvalidRequest,
			err.Error(),
		)
		return
	}

	flights, err := h.searchService.Search(
		r.Context(),
		req,
	)

	if err != nil {

		log.Printf("search failed: %v", err)

		response.WriteError(
			w,
			http.StatusInternalServerError,
			constant.ErrInternalServer,
			"internal server error",
		)
		return
	}

	response.WriteSuccess(
		w,
		http.StatusOK,
		"Flights retrieved successfully",
		flights,
	)
}
