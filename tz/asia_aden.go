// Code generated by tzgen. DO NOT EDIT.
// Package model contains the types.
package tz

import "time"
import "sync"

var (
	onceAsiaAdenLocation  sync.Once
	cacheAsiaAdenLocation *time.Location
)

type AsiaAden struct{}

func (AsiaAden) Location() *time.Location {
	onceAsiaAdenLocation.Do(func() {
		loc, err := time.LoadLocation("Asia/Aden")
		if err != nil {
			panic(err)
		}
		cacheAsiaAdenLocation = loc
	})
	return cacheAsiaAdenLocation
}