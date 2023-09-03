// Code generated by tzgen. DO NOT EDIT.
// Package model contains the types.
package tz

import "time"
import "sync"

var (
	onceAfricaBlantyreLocation  sync.Once
	cacheAfricaBlantyreLocation *time.Location
)

type AfricaBlantyre struct{}

func (AfricaBlantyre) Location() *time.Location {
	onceAfricaBlantyreLocation.Do(func() {
		loc, err := time.LoadLocation("Africa/Blantyre")
		if err != nil {
			panic(err)
		}
		cacheAfricaBlantyreLocation = loc
	})
	return cacheAfricaBlantyreLocation
}