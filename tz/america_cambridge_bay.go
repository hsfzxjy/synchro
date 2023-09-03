// Code generated by tzgen. DO NOT EDIT.
// Package model contains the types.
package tz

import "time"
import "sync"

var (
	onceAmericaCambridge_BayLocation  sync.Once
	cacheAmericaCambridge_BayLocation *time.Location
)

type AmericaCambridge_Bay struct{}

func (AmericaCambridge_Bay) Location() *time.Location {
	onceAmericaCambridge_BayLocation.Do(func() {
		loc, err := time.LoadLocation("America/Cambridge_Bay")
		if err != nil {
			panic(err)
		}
		cacheAmericaCambridge_BayLocation = loc
	})
	return cacheAmericaCambridge_BayLocation
}