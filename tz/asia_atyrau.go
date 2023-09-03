// Code generated by tzgen. DO NOT EDIT.
// Package model contains the types.
package tz

import "time"
import "sync"

var (
	onceAsiaAtyrauLocation  sync.Once
	cacheAsiaAtyrauLocation *time.Location
)

type AsiaAtyrau struct{}

func (AsiaAtyrau) Location() *time.Location {
	onceAsiaAtyrauLocation.Do(func() {
		loc, err := time.LoadLocation("Asia/Atyrau")
		if err != nil {
			panic(err)
		}
		cacheAsiaAtyrauLocation = loc
	})
	return cacheAsiaAtyrauLocation
}