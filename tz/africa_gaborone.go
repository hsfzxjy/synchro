// Code generated by tzgen. DO NOT EDIT.
// Package model contains the types.
package tz

import "time"
import "sync"

var (
	onceAfricaGaboroneLocation  sync.Once
	cacheAfricaGaboroneLocation *time.Location
)

type AfricaGaborone struct{}

func (AfricaGaborone) Location() *time.Location {
	onceAfricaGaboroneLocation.Do(func() {
		loc, err := time.LoadLocation("Africa/Gaborone")
		if err != nil {
			panic(err)
		}
		cacheAfricaGaboroneLocation = loc
	})
	return cacheAfricaGaboroneLocation
}