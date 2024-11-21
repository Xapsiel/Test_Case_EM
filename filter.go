package EffectiveMobile

import "time"

type Filter struct {
	Song  string
	Group string
	Since time.Time
	To    time.Time
}
