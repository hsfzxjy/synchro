// Code generated by tzgen. DO NOT EDIT.
// Package model contains the types.
package tz

import "time"
import "sync"

var (
	onceAntarcticaMcMurdoLocation  sync.Once
	cacheAntarcticaMcMurdoLocation *time.Location
)

type AntarcticaMcMurdo struct{}

func (AntarcticaMcMurdo) Location() *time.Location {
	onceAntarcticaMcMurdoLocation.Do(func() {
		loc, err := time.LoadLocation("Antarctica/McMurdo")
		if err != nil {
			panic(err)
		}
		cacheAntarcticaMcMurdoLocation = loc
	})
	return cacheAntarcticaMcMurdoLocation
}