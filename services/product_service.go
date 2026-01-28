package services

import (
	"errors"

	"github.com/Kocannn/api-kasir/models"
	"github.com/Kocannn/api-kasir/repositories"
)

type ProductService struct {
	repo         *repositories.ProductRepository
	categoryRepo *repositories.CategoryRepository
}

func NewProductService(repo *repositories.ProductRepository, categoryRepo *repositories.CategoryRepository) *ProductService {
	return &ProductService{repo: repo, categoryRepo: categoryRepo}
}

func (s *ProductService) GetAll() ([]models.Product, error) {
	return s.repo.GetAll()
}

func (s *ProductService) Create(data *models.Product) error {
	if data.CategoryID != nil {
		_, err := s.categoryRepo.GetByID(*data.CategoryID)
		if err != nil {
			return errors.New("category tidak ditemukan")
		}
	}
	return s.repo.Create(data)
}

func (s *ProductService) GetByID(id int) (*models.Product, error) {
	return s.repo.GetByID(id)
}

func (s *ProductService) Update(product *models.Product) error {
	if product.CategoryID != nil {
		_, err := s.categoryRepo.GetByID(*product.CategoryID)
		if err != nil {
			return errors.New("category tidak ditemukan")
		}
	}
	return s.repo.Update(product)
}

func (s *ProductService) Delete(id int) error {
	return s.repo.Delete(id)
}
