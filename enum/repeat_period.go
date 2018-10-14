package enum

type RepeatPeriodType uint8

const (
	Daily = iota + 1
	Weekly
	Monthly
	Yearly
)
