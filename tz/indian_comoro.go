// Code generated by tzgen. DO NOT EDIT.
// Package model contains the types.
package tz

import "time"
import "sync"

var (
	onceIndianComoroLocation  sync.Once
	cacheIndianComoroLocation *time.Location
)

type IndianComoro struct{}

func (IndianComoro) Location() *time.Location {
	onceIndianComoroLocation.Do(func() {
		loc, err := time.LoadLocation("Indian/Comoro")
		if err != nil {
			panic(err)
		}
		cacheIndianComoroLocation = loc
	})
	return cacheIndianComoroLocation
}