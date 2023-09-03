// Code generated by tzgen. DO NOT EDIT.
// Package model contains the types.
package tz

import "time"
import "sync"

var (
	onceEuropeIstanbulLocation  sync.Once
	cacheEuropeIstanbulLocation *time.Location
)

type EuropeIstanbul struct{}

func (EuropeIstanbul) Location() *time.Location {
	onceEuropeIstanbulLocation.Do(func() {
		loc, err := time.LoadLocation("Europe/Istanbul")
		if err != nil {
			panic(err)
		}
		cacheEuropeIstanbulLocation = loc
	})
	return cacheEuropeIstanbulLocation
}