package common

type Mode string

const (
	Prod Mode = "prod"
	Dev  Mode = "dev"
)

var RunMode Mode
