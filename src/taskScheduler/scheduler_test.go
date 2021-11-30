package scheduler

import "testing"
import "fmt"

func TestTasklet(t *testing.T) {
	tmp := NewTasklet("Test", "TestDescription", 9, 30, 60, 4, 1)
	tmp.blockTimeSlot()
	fmt.Println(tmp)
}