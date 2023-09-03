// Code generated by tzgen. DO NOT EDIT.
// Package model contains the types.
package tz

import "time"
import "sync"

var (
	onceAmericaSt_JohnsLocation  sync.Once
	cacheAmericaSt_JohnsLocation *time.Location
)

type AmericaSt_Johns struct{}

func (AmericaSt_Johns) Location() *time.Location {
	onceAmericaSt_JohnsLocation.Do(func() {
		loc, err := time.LoadLocation("America/St_Johns")
		if err != nil {
			panic(err)
		}
		cacheAmericaSt_JohnsLocation = loc
	})
	return cacheAmericaSt_JohnsLocation
}