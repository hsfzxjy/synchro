// Code generated by tzgen. DO NOT EDIT.
// Package model contains the types.
package tz

import "time"
import "sync"

var (
	onceAmericaLos_AngelesLocation  sync.Once
	cacheAmericaLos_AngelesLocation *time.Location
)

type AmericaLos_Angeles struct{}

func (AmericaLos_Angeles) Location() *time.Location {
	onceAmericaLos_AngelesLocation.Do(func() {
		loc, err := time.LoadLocation("America/Los_Angeles")
		if err != nil {
			panic(err)
		}
		cacheAmericaLos_AngelesLocation = loc
	})
	return cacheAmericaLos_AngelesLocation
}