// Code generated by tzgen. DO NOT EDIT.
// Package model contains the types.
package tz

import "time"
import "sync"

var (
	oncePacificGuamLocation  sync.Once
	cachePacificGuamLocation *time.Location
)

type PacificGuam struct{}

func (PacificGuam) Location() *time.Location {
	oncePacificGuamLocation.Do(func() {
		loc, err := time.LoadLocation("Pacific/Guam")
		if err != nil {
			panic(err)
		}
		cachePacificGuamLocation = loc
	})
	return cachePacificGuamLocation
}