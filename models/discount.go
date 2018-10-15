package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/uqichi/go-discount-grpc/enum"
)

type Discount struct {
	ID uuid.UUID `db:"id"`

	TenantID          uuid.UUID             `db:"tenant_id"`
	Title             string                `db:"title"`
	Type              enum.Type             `db:"type"`
	MaxAvailableCount uint                  `db:"max_available_count"`
	StartTime         time.Time             `db:"start_time"`
	EndTime           time.Time             `db:"end_time"`
	RepeatPeriodType  enum.RepeatPeriodType `db:"repeat_period_type"`
	RepeatPeriodValue uint                  `db:"repeat_period_value"`

	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`

	Conditions DiscountConditions `has_many:"discount_conditions" order_by:"created_at desc"`
	Rewards    DiscountRewards    `has_many:"discount_rewards" order_by:"created_at desc"`
}

// String is not required by pop and may be deleted
func (d Discount) String() string {
	jd, _ := json.Marshal(d)
	return string(jd)
}

// Discounts is not required by pop and may be deleted
type Discounts []Discount

// String is not required by pop and may be deleted
func (d Discounts) String() string {
	jd, _ := json.Marshal(d)
	return string(jd)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (d *Discount) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (d *Discount) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (d *Discount) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
