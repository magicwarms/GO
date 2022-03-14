package service

import (
	"belajar-unit-test/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: categoryRepository}

func TestCategoryServices_Get(t *testing.T) {
	// program mock nya
	categoryRepository.Mock.On("FindById", "1").Return(nil)
	category, err := categoryService.Get("1")

	assert.Nil(t, category)
	assert.NotNil(t, err)
}
