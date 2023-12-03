package usecase

import "github.com/GuiCintra27/go/api_and_messaging/internal/entity"

type CreateProductInputDTO struct {
	Name  string `json:"name"`
	Price float64 `json:"price"`
}

type CreateProductOutputDTO struct {
	Id string
	Name string
	Price float64
}

type CreateProductUseCase struct {
	ProductRepository entity.ProductRepository
}

func NewCreateProductUseCase(productRepository entity.ProductRepository) *CreateProductUseCase {
	return &CreateProductUseCase{ProductRepository: productRepository}
}

func (u *CreateProductUseCase) Execute(input CreateProductInputDTO) (*CreateProductOutputDTO, error) {
	product := entity.NewProduct(input.Name, input.Price)

	err := u.ProductRepository.Create(product)

	if err != nil {
		return nil, err
	}

	return &CreateProductOutputDTO{
		Id: product.Id,
		Name: product.Name,
		Price: product.Price,
	}, nil
}