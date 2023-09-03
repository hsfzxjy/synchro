// Code generated by tzgen. DO NOT EDIT.
// Package model contains the types.
package tz

import "time"
import "sync"

var (
	onceAfricaBujumburaLocation  sync.Once
	cacheAfricaBujumburaLocation *time.Location
)

type AfricaBujumbura struct{}

func (AfricaBujumbura) Location() *time.Location {
	onceAfricaBujumburaLocation.Do(func() {
		loc, err := time.LoadLocation("Africa/Bujumbura")
		if err != nil {
			panic(err)
		}
		cacheAfricaBujumburaLocation = loc
	})
	return cacheAfricaBujumburaLocation
}