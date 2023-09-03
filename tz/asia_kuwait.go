// Code generated by tzgen. DO NOT EDIT.
// Package model contains the types.
package tz

import "time"
import "sync"

var (
	onceAsiaKuwaitLocation  sync.Once
	cacheAsiaKuwaitLocation *time.Location
)

type AsiaKuwait struct{}

func (AsiaKuwait) Location() *time.Location {
	onceAsiaKuwaitLocation.Do(func() {
		loc, err := time.LoadLocation("Asia/Kuwait")
		if err != nil {
			panic(err)
		}
		cacheAsiaKuwaitLocation = loc
	})
	return cacheAsiaKuwaitLocation
}