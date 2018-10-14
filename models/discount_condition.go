package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
)

type DiscountCondition struct {
	ID uuid.UUID `db:"id"`

	DiscountID               uuid.UUID `db:"discount_id"`
	CategoryID               uuid.UUID `db:"category_id"`
	MinTotalPurchaseQuantity uint      `db:"min_total_purchase_quantity"`
	MinTotalPurchaseAmount   uint      `db:"min_total_purchase_amount"`
	IsMember                 bool      `db:"is_member"`
	IsFirstPurchase          bool      `db:"is_first_purchase"`

	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

// String is not required by pop and may be deleted
func (d DiscountCondition) String() string {
	jd, _ := json.Marshal(d)
	return string(jd)
}

// DiscountConditions is not required by pop and may be deleted
type DiscountConditions []DiscountCondition

// String is not required by pop and may be deleted
func (d DiscountConditions) String() string {
	jd, _ := json.Marshal(d)
	return string(jd)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (d *DiscountCondition) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (d *DiscountCondition) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (d *DiscountCondition) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
