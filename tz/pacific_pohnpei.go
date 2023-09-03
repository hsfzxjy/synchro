// Code generated by tzgen. DO NOT EDIT.
// Package model contains the types.
package tz

import "time"
import "sync"

var (
	oncePacificPohnpeiLocation  sync.Once
	cachePacificPohnpeiLocation *time.Location
)

type PacificPohnpei struct{}

func (PacificPohnpei) Location() *time.Location {
	oncePacificPohnpeiLocation.Do(func() {
		loc, err := time.LoadLocation("Pacific/Pohnpei")
		if err != nil {
			panic(err)
		}
		cachePacificPohnpeiLocation = loc
	})
	return cachePacificPohnpeiLocation
}