package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/uqichi/go-discount-grpc/models"
)

func TestDiscountService_CreateForAll(t *testing.T) {
	assert := assert.New(t)
	db, _ := models.NewDB()
	s := NewDiscountService(db)
	_, err := s.CreateForAll(context.Background(), nil)
	assert.NoError(err)
}
