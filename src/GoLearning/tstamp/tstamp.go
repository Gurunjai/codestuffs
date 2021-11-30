package tstamp

import (
	"fmt"
	"time"
)

func timeStamp(gid int) string {
	gtd := int64((gid >> 12) & 0xffffffff)
	t := time.Unix(gtd, 0).UTC()
	
	return fmt.Sprintf("%s %02d, %d @ %02d:%02d:%02d UTC",
		t.Month(), t.Day(), t.Year(),
		t.Hour(), t.Minute(), t.Second())
}