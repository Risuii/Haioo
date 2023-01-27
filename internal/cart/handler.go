package cart

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"

	"github.com/Risuii/helpers/exception"
	"github.com/Risuii/helpers/response"
	"github.com/Risuii/models/filter"
	"github.com/Risuii/models/product"
)

type CartHandler struct {
	Validate *validator.Validate
	UseCase  CartUseCase
}

func NewCartHandler(router *mux.Router, validate *validator.Validate, usecase CartUseCase) {
	handler := CartHandler{
		Validate: validate,
		UseCase:  usecase,
	}

	api := router.PathPrefix("/cart").Subrouter()

	api.HandleFunc("/items", handler.AddItems).Methods(http.MethodPost)
	api.HandleFunc("/items", handler.GetItems).Methods(http.MethodGet)
	api.HandleFunc("/items", handler.DeleteItems).Methods(http.MethodDelete)
}

func (handler *CartHandler) AddItems(w http.ResponseWriter, r *http.Request) {
	var res response.Response
	var userInput product.Product

	ctx := r.Context()

	if err := json.NewDecoder(r.Body).Decode(&userInput); err != nil {
		res = response.Error(response.StatusUnprocessableEntity, exception.ErrUnprocessableEntity)
		res.JSON(w)
		return
	}

	err := handler.Validate.StructCtx(ctx, userInput)
	if err != nil {
		res = response.Error(response.StatusBadRequest, exception.ErrBadRequest)
		res.JSON(w)
		return
	}

	res = handler.UseCase.AddItems(ctx, userInput)

	res.JSON(w)
}

func (handler *CartHandler) GetItems(w http.ResponseWriter, r *http.Request) {
	var res response.Response
	var userInput filter.Filter

	ctx := r.Context()

	if err := json.NewDecoder(r.Body).Decode(&userInput); err != nil {
		res = response.Error(response.StatusUnprocessableEntity, exception.ErrUnprocessableEntity)
		res.JSON(w)
		return
	}

	res = handler.UseCase.GetItems(ctx, userInput)

	res.JSON(w)
}

func (handler *CartHandler) DeleteItems(w http.ResponseWriter, r *http.Request) {
	var res response.Response
	var userInput product.Product

	ctx := r.Context()

	if err := json.NewDecoder(r.Body).Decode(&userInput); err != nil {
		res = response.Error(response.StatusUnprocessableEntity, exception.ErrUnprocessableEntity)
		res.JSON(w)
		return
	}

	res = handler.UseCase.DeleteItems(ctx, userInput.KodeProduk)

	res.JSON(w)
}
