// Code generated by tzgen. DO NOT EDIT.
// Package model contains the types.
package tz

import "time"
import "sync"

var (
	onceAsiaUlaanbaatarLocation  sync.Once
	cacheAsiaUlaanbaatarLocation *time.Location
)

type AsiaUlaanbaatar struct{}

func (AsiaUlaanbaatar) Location() *time.Location {
	onceAsiaUlaanbaatarLocation.Do(func() {
		loc, err := time.LoadLocation("Asia/Ulaanbaatar")
		if err != nil {
			panic(err)
		}
		cacheAsiaUlaanbaatarLocation = loc
	})
	return cacheAsiaUlaanbaatarLocation
}