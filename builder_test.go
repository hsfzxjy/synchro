package synchro_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/Code-Hex/synchro"
	"github.com/Code-Hex/synchro/tz"
)

func ExampleTime_Change() {
	utc := synchro.New[tz.UTC](2009, time.November, 10, 23, 0, 0, 0)
	c1 := utc.Change().Year(2010).Do()
	c2 := utc.Change().Year(2010).Month(time.December).Do()
	c3 := utc.Change().Year(2010).Month(time.December).Day(1).Do()
	c4 := c3.Change().Hour(1).Do()
	c5 := c3.Change().Hour(1).Minute(1).Do()
	c6 := c3.Change().Hour(1).Minute(1).Second(1).Do()
	c7 := c3.Change().Hour(1).Minute(1).Second(1).Nanosecond(123456789).Do()
	fmt.Printf("Go launched at %s\n", utc)
	fmt.Println(c1)
	fmt.Println(c2)
	fmt.Println(c3)
	fmt.Println(c4)
	fmt.Println(c5)
	fmt.Println(c6)
	fmt.Println(c7)
	// Output:
	// Go launched at 2009-11-10 23:00:00 +0000 UTC
	// 2010-11-10 23:00:00 +0000 UTC
	// 2010-12-10 23:00:00 +0000 UTC
	// 2010-12-01 23:00:00 +0000 UTC
	// 2010-12-01 01:00:00 +0000 UTC
	// 2010-12-01 01:01:00 +0000 UTC
	// 2010-12-01 01:01:01 +0000 UTC
	// 2010-12-01 01:01:01.123456789 +0000 UTC
}

func ExampleTime_Advance() {
	utc := synchro.New[tz.UTC](2009, time.November, 10, 23, 0, 0, 0)
	c1 := utc.Advance().Year(1).Do()
	c11 := utc.Advance().Year(1).Year(1).Do() // +2 years

	c2 := utc.Advance().Year(1).Month(1).Do()
	c3 := utc.Advance().Year(1).Month(1).Day(1).Do()
	c4 := c3.Advance().Hour(1).Do()
	c5 := c3.Advance().Hour(1).Minute(1).Do()
	c6 := c3.Advance().Hour(1).Minute(1).Second(1).Do()
	c7 := c3.Advance().Hour(1).Minute(1).Second(1).Nanosecond(123456789).Do()

	fmt.Printf("Go launched at %s\n", utc)
	fmt.Println(c1)
	fmt.Println(c11)
	fmt.Println()
	fmt.Println(c2)
	fmt.Println(c3)
	fmt.Println(c4)
	fmt.Println(c5)
	fmt.Println(c6)
	fmt.Println(c7)
	// Output:
	// Go launched at 2009-11-10 23:00:00 +0000 UTC
	// 2010-11-10 23:00:00 +0000 UTC
	// 2011-11-10 23:00:00 +0000 UTC
	//
	// 2010-12-10 23:00:00 +0000 UTC
	// 2010-12-11 23:00:00 +0000 UTC
	// 2010-12-12 00:00:00 +0000 UTC
	// 2010-12-12 00:01:00 +0000 UTC
	// 2010-12-12 00:01:01 +0000 UTC
	// 2010-12-12 00:01:01.123456789 +0000 UTC
}

func TestAdvance(t *testing.T) {
	utc := synchro.New[tz.UTC](2009, time.November, 10, 23, 0, 0, 0)

	t.Run("month twice", func(t *testing.T) {
		got := utc.Advance().Month(1).Month(2).Do() // +3 months
		want := synchro.New[tz.UTC](2010, time.February, 10, 23, 0, 0, 0)
		if want != got {
			t.Fatalf("- %s\n+ %s", want, got)
		}
	})
	t.Run("day twice", func(t *testing.T) {
		got := utc.Advance().Day(1).Day(2).Do() // +3 days
		want := synchro.New[tz.UTC](2009, time.November, 13, 23, 0, 0, 0)
		if want != got {
			t.Fatalf("- %s\n+ %s", want, got)
		}
	})
	t.Run("hour twice", func(t *testing.T) {
		got := utc.Advance().Hour(1).Hour(2).Do() // +3 hours
		want := synchro.New[tz.UTC](2009, time.November, 11, 2, 0, 0, 0)
		if want != got {
			t.Fatalf("- %s\n+ %s", want, got)
		}
	})
	t.Run("minute twice", func(t *testing.T) {
		got := utc.Advance().Minute(5).Minute(60).Do() // +65 minutes
		want := synchro.New[tz.UTC](2009, time.November, 11, 0, 5, 0, 0)
		if want != got {
			t.Fatalf("- %s\n+ %s", want, got)
		}
	})
	t.Run("second twice", func(t *testing.T) {
		got := utc.Advance().Second(5).Second(60).Do() // +65 seconds
		want := synchro.New[tz.UTC](2009, time.November, 10, 23, 1, 5, 0)
		if want != got {
			t.Fatalf("- %s\n+ %s", want, got)
		}
	})
	t.Run("nanosec twice", func(t *testing.T) {
		got := utc.Advance().Nanosecond(5).Nanosecond(60).Do() // +65 nanosec
		want := synchro.New[tz.UTC](2009, time.November, 10, 23, 0, 0, 65)
		if want != got {
			t.Fatalf("- %s\n+ %s", want, got)
		}
	})
}