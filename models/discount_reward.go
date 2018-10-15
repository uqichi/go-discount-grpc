package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
)

type DiscountReward struct {
	ID uuid.UUID `db:"id"`

	DiscountID        uuid.UUID `db:"discount_id"`
	TotalFixedAmount  uint      `db:"total_fixed_amount"`
	TotalPercentPoint float32   `db:"total_percent_point"`

	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

// String is not required by pop and may be deleted
func (d DiscountReward) String() string {
	jd, _ := json.Marshal(d)
	return string(jd)
}

// DiscountRewards is not required by pop and may be deleted
type DiscountRewards []DiscountReward

// String is not required by pop and may be deleted
func (d DiscountRewards) String() string {
	jd, _ := json.Marshal(d)
	return string(jd)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (d *DiscountReward) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (d *DiscountReward) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (d *DiscountReward) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
