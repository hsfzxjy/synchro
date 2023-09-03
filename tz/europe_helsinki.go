// Code generated by tzgen. DO NOT EDIT.
// Package model contains the types.
package tz

import "time"
import "sync"

var (
	onceEuropeHelsinkiLocation  sync.Once
	cacheEuropeHelsinkiLocation *time.Location
)

type EuropeHelsinki struct{}

func (EuropeHelsinki) Location() *time.Location {
	onceEuropeHelsinkiLocation.Do(func() {
		loc, err := time.LoadLocation("Europe/Helsinki")
		if err != nil {
			panic(err)
		}
		cacheEuropeHelsinkiLocation = loc
	})
	return cacheEuropeHelsinkiLocation
}