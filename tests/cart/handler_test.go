package cart_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/Risuii/helpers/exception"
	"github.com/Risuii/helpers/response"
	"github.com/Risuii/internal/cart"
	"github.com/Risuii/models/filter"
	"github.com/Risuii/models/product"
	"github.com/Risuii/tests/mocks"
)

func TestHandler_AddItems(t *testing.T) {
	t.Run("Add Items Success", func(t *testing.T) {
		data := product.Product{
			ID:         1,
			Nama:       "test",
			KodeProduk: "test-01",
			Kuantitas:  1,
		}

		resp := response.Success(response.StatusCreated, data)

		newReq, err := json.Marshal(data)
		if err != nil {
			t.Error(err)
			return
		}

		validate := validator.New()
		cartUseCase := new(mocks.CartUseCase)
		cartUseCase.On("AddItems", mock.Anything, mock.AnythingOfType("product.Product")).Return(resp)

		cartHandler := cart.CartHandler{
			Validate: validate,
			UseCase:  cartUseCase,
		}

		r := httptest.NewRequest(http.MethodPost, "/just/for/testing", bytes.NewReader(newReq))
		recorder := httptest.NewRecorder()

		handler := http.HandlerFunc(cartHandler.AddItems)
		handler.ServeHTTP(recorder, r)

		rb := response.ResponseImpl{}
		if err := json.NewDecoder(recorder.Body).Decode(&rb); err != nil {
			t.Error(err)
			return
		}

		assert.Equal(t, response.StatusCreated, rb.Status)
		assert.NotNil(t, rb.Data)
	})

	t.Run("Add Items Error Entity", func(t *testing.T) {
		validator := validator.New()
		cartUseCase := new(mocks.CartUseCase)

		cartHandler := cart.CartHandler{
			Validate: validator,
			UseCase:  cartUseCase,
		}

		r := httptest.NewRequest(http.MethodPost, "/just/for/testing", nil)
		recorder := httptest.NewRecorder()

		handler := http.HandlerFunc(cartHandler.AddItems)
		handler.ServeHTTP(recorder, r)

		rb := response.ResponseImpl{}
		if err := json.NewDecoder(recorder.Body).Decode(&rb); err != nil {
			t.Error(err)
			return
		}

		assert.Equal(t, response.StatusUnprocessableEntity, rb.Status)
		assert.Nil(t, rb.Data)
	})

	t.Run("Add Items Error Bad Request", func(t *testing.T) {
		type invalidReq struct {
			Data string
		}

		mockData := invalidReq{
			Data: "Error",
		}

		resp := response.Error(response.StatusBadRequest, exception.ErrBadRequest)

		newReq, err := json.Marshal(mockData)
		if err != nil {
			t.Error(err)
			return
		}

		validate := validator.New()
		cartUseCase := new(mocks.CartUseCase)
		cartUseCase.On("AddItems", mock.Anything, mock.AnythingOfType("product.Product")).Return(resp)

		cartHandler := cart.CartHandler{
			Validate: validate,
			UseCase:  cartUseCase,
		}

		r := httptest.NewRequest(http.MethodPost, "/just/for/testing", bytes.NewReader(newReq))
		recorder := httptest.NewRecorder()

		handler := http.HandlerFunc(cartHandler.AddItems)
		handler.ServeHTTP(recorder, r)

		rb := response.ResponseImpl{}
		if err := json.NewDecoder(recorder.Body).Decode(&rb); err != nil {
			t.Error(err)
			return
		}

		assert.Equal(t, response.StatusBadRequest, rb.Status)
		assert.Nil(t, rb.Data)
	})
}

func TestHandler_GetItems(t *testing.T) {
	t.Run("Get Items Success", func(t *testing.T) {
		mockDataRes := []product.Product{
			{
				ID:         1,
				Nama:       "test",
				KodeProduk: "test",
				Kuantitas:  1,
			},
		}

		mockDataReq := filter.Filter{
			Nama:      "test",
			Kuantitas: 1,
		}

		resp := response.Success(response.StatusOK, mockDataRes)

		newReq, err := json.Marshal(mockDataReq)
		if err != nil {
			t.Error(err)
			return
		}

		cartUseCase := new(mocks.CartUseCase)
		cartUseCase.On("GetItems", mock.Anything, mock.AnythingOfType("filter.Filter")).Return(resp)

		cartHandler := cart.CartHandler{
			UseCase: cartUseCase,
		}

		r := httptest.NewRequest(http.MethodGet, "/just/for/testing", bytes.NewReader(newReq))
		recorder := httptest.NewRecorder()

		handler := http.HandlerFunc(cartHandler.GetItems)
		handler.ServeHTTP(recorder, r)

		rb := response.ResponseImpl{}

		if err := json.NewDecoder(recorder.Body).Decode(&rb); err != nil {
			t.Error(err)
			return
		}

		assert.Equal(t, response.StatusOK, rb.Status)
		assert.NotNil(t, rb.Data)
	})
}
