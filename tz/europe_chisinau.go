// Code generated by tzgen. DO NOT EDIT.
// Package model contains the types.
package tz

import "time"
import "sync"

var (
	onceEuropeChisinauLocation  sync.Once
	cacheEuropeChisinauLocation *time.Location
)

type EuropeChisinau struct{}

func (EuropeChisinau) Location() *time.Location {
	onceEuropeChisinauLocation.Do(func() {
		loc, err := time.LoadLocation("Europe/Chisinau")
		if err != nil {
			panic(err)
		}
		cacheEuropeChisinauLocation = loc
	})
	return cacheEuropeChisinauLocation
}