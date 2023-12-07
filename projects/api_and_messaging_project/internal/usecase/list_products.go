package usecase

import "github.com/GuiCintra27/go/projects/api_and_messaging_project/internal/entity"

type ListProductsOutputDTO struct {
	Id string
	Name string
	Price float64
}

type ListProductsUseCase struct {
	ProductRepository entity.ProductRepository
}

func NewListProductsUseCase(productRepository entity.ProductRepository) *ListProductsUseCase {
	return &ListProductsUseCase{ProductRepository: productRepository}
}

func (u *ListProductsUseCase) Execute() ([]*ListProductsOutputDTO, error) {

	products, err := u.ProductRepository.FindAll()

	if err != nil {
		return nil, err
	}

	var productsOutput []*ListProductsOutputDTO

	for _, product := range products {
		productsOutput = append(productsOutput, &ListProductsOutputDTO{
			Id: product.Id,
			Name: product.Name,
			Price: product.Price,
		})
	}

	return productsOutput, nil
}