package service

import (
	"context"

	"github.com/uqichi/go-discount-grpc/proto"
)

type DiscountService struct{}

func NewDiscountService() *DiscountService {
	return &DiscountService{}
}

func (s *DiscountService) GetOne(context.Context, *pb.GetOneRequest) (*pb.GetOneResponse, error) {
	return nil, nil
}

func (s *DiscountService) CreateForAll(context.Context, *pb.CreateForAllRequest) (*pb.CreateResponse, error) {
	return nil, nil
}

func (s *DiscountService) CreateForOne(context.Context, *pb.CreateForOneRequest) (*pb.CreateResponse, error) {
	return nil, nil
}

func (s *DiscountService) CreateForSet(context.Context, *pb.CreateForSetRequest) (*pb.CreateResponse, error) {
	return nil, nil
}

func (s *DiscountService) CreateForBuyAForB(context.Context, *pb.CreateForBuyAForBRequest) (*pb.CreateResponse, error) {
	return nil, nil
}

func (s *DiscountService) CreateForBuyXGetOne(context.Context, *pb.CreateForBuyXGetRequest) (*pb.CreateResponse, error) {
	return nil, nil
}

func (s *DiscountService) CreateForBuyXGetMore(context.Context, *pb.CreateForBuyXGetRequest) (*pb.CreateResponse, error) {
	return nil, nil
}

func (s *DiscountService) CreateForBuyXGetTotal(context.Context, *pb.CreateForBuyXGetRequest) (*pb.CreateResponse, error) {
	return nil, nil
}
