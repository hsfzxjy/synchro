// Code generated by tzgen. DO NOT EDIT.
// Package model contains the types.
package tz

import "time"
import "sync"

var (
	onceAfricaJubaLocation  sync.Once
	cacheAfricaJubaLocation *time.Location
)

type AfricaJuba struct{}

func (AfricaJuba) Location() *time.Location {
	onceAfricaJubaLocation.Do(func() {
		loc, err := time.LoadLocation("Africa/Juba")
		if err != nil {
			panic(err)
		}
		cacheAfricaJubaLocation = loc
	})
	return cacheAfricaJubaLocation
}