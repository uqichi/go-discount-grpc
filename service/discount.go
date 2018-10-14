package service

import (
	"context"

	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/uqichi/go-discount-grpc/enum"
	"github.com/uqichi/go-discount-grpc/models"
	"github.com/uqichi/go-discount-grpc/proto"
)

type DiscountService struct {
	db *pop.Connection
}

func NewDiscountService(db *pop.Connection) *DiscountService {
	return &DiscountService{db}
}

func (s *DiscountService) GetOne(context.Context, *pb.GetOneRequest) (*pb.GetOneResponse, error) {
	return nil, nil
}

func (s *DiscountService) CreateForAll(context.Context, *pb.CreateForAllRequest) (*pb.CreateResponse, error) {
	// TODO: dynamic settings
	conds := models.DiscountConditions{models.DiscountCondition{
		CategoryID:               uuid.UUID{},
		MinTotalPurchaseQuantity: 1,
		MinTotalPurchaseAmount:   0,
		IsMember:                 false,
		IsFirstPurchase:          false,
	}}
	rewds := models.DiscountRewards{models.DiscountReward{
		TotalFixedAmount:  0,
		TotalPercentPoint: 10.0,
	}}
	discount := models.Discount{
		TenantID:          uuid.Must(uuid.NewV4()),
		Title:             "registration",
		Type:              enum.TIME_SALE,
		MaxAvailableCount: 99,
		StartTime:         time.Now(),
		EndTime:           time.Now(),
		RepeatPeriodType:  0,
		RepeatPeriodValue: 0,
		Conditions:        conds,
		Rewards:           rewds,
	}
	err := s.db.Eager().Create(&discount)
	return nil, err
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
