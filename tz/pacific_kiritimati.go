// Code generated by tzgen. DO NOT EDIT.
// Package model contains the types.
package tz

import "time"
import "sync"

var (
	oncePacificKiritimatiLocation  sync.Once
	cachePacificKiritimatiLocation *time.Location
)

type PacificKiritimati struct{}

func (PacificKiritimati) Location() *time.Location {
	oncePacificKiritimatiLocation.Do(func() {
		loc, err := time.LoadLocation("Pacific/Kiritimati")
		if err != nil {
			panic(err)
		}
		cachePacificKiritimatiLocation = loc
	})
	return cachePacificKiritimatiLocation
}