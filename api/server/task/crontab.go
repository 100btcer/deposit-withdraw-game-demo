package task

import (
	"fmt"
	"time"
)

func CrontabForTest() {
	fmt.Printf("Crontab for test %d\n", time.Now().Unix())
}
