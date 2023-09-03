// Code generated by tzgen. DO NOT EDIT.
// Package model contains the types.
package tz

import "time"
import "sync"

var (
	onceAfricaDoualaLocation  sync.Once
	cacheAfricaDoualaLocation *time.Location
)

type AfricaDouala struct{}

func (AfricaDouala) Location() *time.Location {
	onceAfricaDoualaLocation.Do(func() {
		loc, err := time.LoadLocation("Africa/Douala")
		if err != nil {
			panic(err)
		}
		cacheAfricaDoualaLocation = loc
	})
	return cacheAfricaDoualaLocation
}