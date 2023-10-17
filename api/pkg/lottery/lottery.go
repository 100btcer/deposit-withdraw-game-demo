package lottery

import (
	"math/rand"
)

type LotteryInfer interface {
	RandomType(seed int64) int
}

type Lottery struct {
	rate       map[int]int64 // type/percent  1%=100
	bucketMaps map[int]*BucketRange
	total      int64
}

type BucketRange struct {
	Start int64
	End   int64
}

func NewLottery(rate map[int]int64) LotteryInfer {
	var totalPercent int64
	for _, v := range rate {
		totalPercent += v
	}

	bucketMap := make(map[int]*BucketRange)

	start := int64(0)
	for k, v := range rate {
		bucketMap[k] = &BucketRange{
			Start: start,
			End:   start + v - 1,
		}
		start = bucketMap[k].End + 1
	}

	return &Lottery{
		bucketMaps: bucketMap,
		total:      totalPercent,
	}
}

// RandomType -1 means failure
func (l *Lottery) RandomType(seed int64) int {
	rand.Seed(seed)
	randNum := rand.Int63n(l.total)
	for k, v := range l.bucketMaps {
		if v.Start <= randNum && randNum <= v.End {
			return k
		}
	}
	return -1
}
