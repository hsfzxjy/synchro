// Code generated by tzgen. DO NOT EDIT.
// Package model contains the types.
package tz

import "time"
import "sync"

var (
	onceAsiaBahrainLocation  sync.Once
	cacheAsiaBahrainLocation *time.Location
)

type AsiaBahrain struct{}

func (AsiaBahrain) Location() *time.Location {
	onceAsiaBahrainLocation.Do(func() {
		loc, err := time.LoadLocation("Asia/Bahrain")
		if err != nil {
			panic(err)
		}
		cacheAsiaBahrainLocation = loc
	})
	return cacheAsiaBahrainLocation
}