package cart_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/Risuii/helpers/exception"
	"github.com/Risuii/internal/cart"
	"github.com/Risuii/models/filter"
	"github.com/Risuii/models/product"
	"github.com/Risuii/tests/mocks"
)

func TestUseCaseAddItems(t *testing.T) {
	t.Run("Add Items Success", func(t *testing.T) {

		ctx := context.TODO()
		mockData := product.Product{
			ID:         1,
			Nama:       "test",
			KodeProduk: "test",
			Kuantitas:  1,
			CreatedAt:  time.Now(),
		}

		cartRepository := new(mocks.CartRepository)
		cartRepository.On("FindByKodeProduk", mock.Anything, mock.AnythingOfType("string")).Return(mockData, exception.ErrNotFound)
		cartRepository.On("Add", mock.Anything, mock.AnythingOfType("product.Product")).Return(int64(1), nil)

		cartUseCase := cart.NewCartUseCaseImpl(
			cartRepository,
		)

		resp := cartUseCase.AddItems(ctx, mockData)

		assert.NoError(t, resp.Err())

		cartRepository.AssertExpectations(t)
	})

	t.Run("Add Items Error", func(t *testing.T) {
		ctx := context.TODO()
		mockData := product.Product{
			ID:         1,
			Nama:       "test",
			KodeProduk: "test",
			Kuantitas:  1,
			CreatedAt:  time.Now(),
		}

		cartRepository := new(mocks.CartRepository)
		cartRepository.On("FindByKodeProduk", mock.Anything, mock.AnythingOfType("string")).Return(mockData, exception.ErrNotFound)
		cartRepository.On("Add", mock.Anything, mock.AnythingOfType("product.Product")).Return(int64(0), exception.ErrInternalServer)

		cartUseCase := cart.NewCartUseCaseImpl(
			cartRepository,
		)

		resp := cartUseCase.AddItems(ctx, mockData)

		assert.Error(t, resp.Err())

		cartRepository.AssertExpectations(t)
	})

	t.Run("Add Items If Existing", func(t *testing.T) {
		ctx := context.TODO()
		mockData := product.Product{
			ID:         1,
			Nama:       "test",
			KodeProduk: "test",
			Kuantitas:  1,
			CreatedAt:  time.Now(),
		}

		cartRepository := new(mocks.CartRepository)
		cartRepository.On("FindByKodeProduk", mock.Anything, mock.AnythingOfType("string")).Return(mockData, nil)
		cartRepository.On("UpdateKuantitas", mock.Anything, mock.AnythingOfType("int64"), mock.AnythingOfType("product.Product")).Return(nil)

		cartUseCase := cart.NewCartUseCaseImpl(
			cartRepository,
		)

		resp := cartUseCase.AddItems(ctx, mockData)

		assert.NoError(t, resp.Err())

		cartRepository.AssertExpectations(t)
	})

	t.Run("Add Items If Existing But Got Error", func(t *testing.T) {
		ctx := context.TODO()
		mockData := product.Product{
			ID:         1,
			Nama:       "test",
			KodeProduk: "test",
			Kuantitas:  1,
			CreatedAt:  time.Now(),
		}

		cartRepository := new(mocks.CartRepository)
		cartRepository.On("FindByKodeProduk", mock.Anything, mock.AnythingOfType("string")).Return(mockData, nil)
		cartRepository.On("UpdateKuantitas", mock.Anything, mock.AnythingOfType("int64"), mock.AnythingOfType("product.Product")).Return(exception.ErrInternalServer)

		cartUseCase := cart.NewCartUseCaseImpl(
			cartRepository,
		)

		resp := cartUseCase.AddItems(ctx, mockData)

		assert.Error(t, resp.Err())

		cartRepository.AssertExpectations(t)
	})
}

func TestUseCaseGetItems(t *testing.T) {

	t.Run("Get All Items Success", func(t *testing.T) {
		var data []product.Product
		ctx := context.TODO()
		mockData := filter.Filter{
			Nama:      "",
			Kuantitas: 0,
		}

		cartRepository := new(mocks.CartRepository)
		cartRepository.On("FindAll").Return(data, nil)

		cartUseCase := cart.NewCartUseCaseImpl(
			cartRepository,
		)

		resp := cartUseCase.GetItems(ctx, mockData)

		assert.NoError(t, resp.Err())

		cartRepository.AssertExpectations(t)
	})

	t.Run("Get All Items Error Not Found", func(t *testing.T) {

		ctx := context.TODO()
		mockData := filter.Filter{
			Nama:      "",
			Kuantitas: 0,
		}

		cartRepository := new(mocks.CartRepository)
		cartRepository.On("FindAll").Return(nil, exception.ErrNotFound)

		cartUseCase := cart.NewCartUseCaseImpl(
			cartRepository,
		)

		resp := cartUseCase.GetItems(ctx, mockData)

		assert.Error(t, resp.Err())

		cartRepository.AssertExpectations(t)
	})

	t.Run("Get All Items Error Internal Server", func(t *testing.T) {

		ctx := context.TODO()
		mockData := filter.Filter{
			Nama:      "",
			Kuantitas: 0,
		}

		cartRepository := new(mocks.CartRepository)
		cartRepository.On("FindAll").Return(nil, exception.ErrInternalServer)

		cartUseCase := cart.NewCartUseCaseImpl(
			cartRepository,
		)

		resp := cartUseCase.GetItems(ctx, mockData)

		assert.Error(t, resp.Err())

		cartRepository.AssertExpectations(t)
	})

	t.Run("Get Items By Filter Success", func(t *testing.T) {
		var data []product.Product
		ctx := context.TODO()
		mockData := filter.Filter{
			Nama:      "test",
			Kuantitas: 1,
		}

		cartRepository := new(mocks.CartRepository)
		cartRepository.On("FindByFilter", mock.Anything, mock.AnythingOfType("filter.Filter")).Return(data, nil)

		cartUseCase := cart.NewCartUseCaseImpl(
			cartRepository,
		)

		resp := cartUseCase.GetItems(ctx, mockData)

		assert.NoError(t, resp.Err())

		cartRepository.AssertExpectations(t)
	})

	t.Run("Get Items By Filter Error Not Found", func(t *testing.T) {

		ctx := context.TODO()
		mockData := filter.Filter{
			Nama:      "test",
			Kuantitas: 1,
		}

		cartRepository := new(mocks.CartRepository)
		cartRepository.On("FindByFilter", mock.Anything, mock.AnythingOfType("filter.Filter")).Return(nil, exception.ErrNotFound)

		cartUseCase := cart.NewCartUseCaseImpl(
			cartRepository,
		)

		resp := cartUseCase.GetItems(ctx, mockData)

		assert.Error(t, resp.Err())

		cartRepository.AssertExpectations(t)
	})

	t.Run("Get Items By Filter Error Internal Server", func(t *testing.T) {
		ctx := context.TODO()
		mockData := filter.Filter{
			Nama:      "test",
			Kuantitas: 1,
		}

		cartRepository := new(mocks.CartRepository)
		cartRepository.On("FindByFilter", mock.Anything, mock.AnythingOfType("filter.Filter")).Return(nil, exception.ErrInternalServer)

		cartUseCase := cart.NewCartUseCaseImpl(
			cartRepository,
		)

		resp := cartUseCase.GetItems(ctx, mockData)

		assert.Error(t, resp.Err())

		cartRepository.AssertExpectations(t)
	})
}

func TestUseCaseDeleteItems(t *testing.T) {
	t.Run("Delete Items Success", func(t *testing.T) {
		type req struct {
			Data string
		}

		newReq := req{
			Data: "test",
		}

		ctx := context.TODO()

		cartRepository := new(mocks.CartRepository)
		cartRepository.On("FindByKodeProduk", mock.Anything, mock.AnythingOfType("string")).Return(product.Product{}, nil)
		cartRepository.On("Delete", mock.Anything, mock.AnythingOfType("int64")).Return(nil)

		cartUseCase := cart.NewCartUseCaseImpl(
			cartRepository,
		)

		resp := cartUseCase.DeleteItems(ctx, newReq.Data)

		assert.NoError(t, resp.Err())

		cartRepository.AssertExpectations(t)
	})

	t.Run("Delete Items Error", func(t *testing.T) {
		type req struct {
			Data string
		}

		newReq := req{
			Data: "test",
		}

		ctx := context.TODO()

		cartRepository := new(mocks.CartRepository)
		cartRepository.On("FindByKodeProduk", mock.Anything, mock.AnythingOfType("string")).Return(product.Product{}, nil)
		cartRepository.On("Delete", mock.Anything, mock.AnythingOfType("int64")).Return(exception.ErrInternalServer)

		cartUseCase := cart.NewCartUseCaseImpl(
			cartRepository,
		)

		resp := cartUseCase.DeleteItems(ctx, newReq.Data)

		assert.Error(t, resp.Err())

		cartRepository.AssertExpectations(t)
	})

	t.Run("Delete Items Find By Kode Produk Error Not Found", func(t *testing.T) {
		type req struct {
			Data string
		}

		newReq := req{
			Data: "test",
		}

		ctx := context.TODO()

		cartRepository := new(mocks.CartRepository)
		cartRepository.On("FindByKodeProduk", mock.Anything, mock.AnythingOfType("string")).Return(product.Product{}, exception.ErrNotFound)

		cartUseCase := cart.NewCartUseCaseImpl(
			cartRepository,
		)

		resp := cartUseCase.DeleteItems(ctx, newReq.Data)

		assert.Error(t, resp.Err())

		cartRepository.AssertExpectations(t)
	})

	t.Run("Delete Items Find By Kode Produk Error Internal Server", func(t *testing.T) {
		type req struct {
			Data string
		}

		newReq := req{
			Data: "test",
		}

		ctx := context.TODO()

		cartRepository := new(mocks.CartRepository)
		cartRepository.On("FindByKodeProduk", mock.Anything, mock.AnythingOfType("string")).Return(product.Product{}, exception.ErrInternalServer)

		cartUseCase := cart.NewCartUseCaseImpl(
			cartRepository,
		)

		resp := cartUseCase.DeleteItems(ctx, newReq.Data)

		assert.Error(t, resp.Err())

		cartRepository.AssertExpectations(t)
	})
}
