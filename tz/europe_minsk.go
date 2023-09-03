// Code generated by tzgen. DO NOT EDIT.
// Package model contains the types.
package tz

import "time"
import "sync"

var (
	onceEuropeMinskLocation  sync.Once
	cacheEuropeMinskLocation *time.Location
)

type EuropeMinsk struct{}

func (EuropeMinsk) Location() *time.Location {
	onceEuropeMinskLocation.Do(func() {
		loc, err := time.LoadLocation("Europe/Minsk")
		if err != nil {
			panic(err)
		}
		cacheEuropeMinskLocation = loc
	})
	return cacheEuropeMinskLocation
}