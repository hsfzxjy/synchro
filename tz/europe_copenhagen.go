// Code generated by tzgen. DO NOT EDIT.
// Package model contains the types.
package tz

import "time"
import "sync"

var (
	onceEuropeCopenhagenLocation  sync.Once
	cacheEuropeCopenhagenLocation *time.Location
)

type EuropeCopenhagen struct{}

func (EuropeCopenhagen) Location() *time.Location {
	onceEuropeCopenhagenLocation.Do(func() {
		loc, err := time.LoadLocation("Europe/Copenhagen")
		if err != nil {
			panic(err)
		}
		cacheEuropeCopenhagenLocation = loc
	})
	return cacheEuropeCopenhagenLocation
}