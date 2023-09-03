// Code generated by tzgen. DO NOT EDIT.
// Package model contains the types.
package tz

import "time"
import "sync"

var (
	onceAustraliaBroken_HillLocation  sync.Once
	cacheAustraliaBroken_HillLocation *time.Location
)

type AustraliaBroken_Hill struct{}

func (AustraliaBroken_Hill) Location() *time.Location {
	onceAustraliaBroken_HillLocation.Do(func() {
		loc, err := time.LoadLocation("Australia/Broken_Hill")
		if err != nil {
			panic(err)
		}
		cacheAustraliaBroken_HillLocation = loc
	})
	return cacheAustraliaBroken_HillLocation
}