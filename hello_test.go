package main

import (
	"testing"
	"github.com/prashantv/gostub"
	"time"
	"github.com/smartystreets/assertions"
)

func TestTodayOfMonth(t *testing.T) {
	mockDate := time.Date(2000, 10, 11, 0, 0, 0, 0, time.UTC)

	stubs := gostub.Stub(&now, mockDate)
	defer stubs.Reset()

	assertions.ShouldEqual(TodayOfMonth(), 11)
}
