package lottery

import (
	"fmt"
	"github.com/go-redis/redis"
	"math/rand"
	"music-nft/internal/constant"
	"sync"
	"testing"
	"time"
)

var (
	mx sync.Mutex
	wg sync.WaitGroup
)

func TestLottery(t *testing.T) {
	store := redis.NewClient(&redis.Options{
		Addr:         "127.0.0.1:6379",
		DB:           0,
		PoolSize:     100,
		MinIdleConns: 10,
		Password:     "",
	})

	rate := map[int]int64{
		1:  300,
		2:  1450,
		3:  1450,
		4:  300,
		5:  1450,
		6:  300,
		7:  1450,
		8:  300,
		9:  1450,
		10: 1450,
		11: 50,
		12: 50,
	}

	//for i := 1; i <= 200; i++ {
	//	rate[i] = 50
	//}
	key := fmt.Sprintf(constant.LotteryPrefix, 1, 1, 1)
	//stock := int64(2500) // 出售数量

	//var totalPercent int64 // 总权重
	//for _, v := range rate {
	//	totalPercent += v
	//}
	//
	//// 根据发行数量按概率计算新权重
	//stockPercent := decimal.NewFromInt(stock).
	//	Div(decimal.NewFromInt(totalPercent)).RoundFloor(2)
	//newRateMap := make(map[int]int64)
	//for k, v := range rate {
	//	newRateMap[k] = decimal.NewFromInt(v).Mul(stockPercent).Floor().IntPart()
	//}
	//store.Del(store.Context(), key)
	//var sum int64
	//for k, v := range newRateMap {
	//	sum += v
	//	//fmt.Println(k, v)
	//	if err := store.HMSet(store.Context(), key, k, v).Err(); err != nil {
	//		panic(err)
	//	}
	//}

	var max int64
	for _, v := range rate {
		if v > max {
			max = v
		}
	}
	var maxKeys []int
	for k, v := range rate {
		if v == max {
			maxKeys = append(maxKeys, k)
		}
	}

	l := NewLottery(rate)

	result := make(map[int]int64)
	for j := 0; j < 10; j++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			for i := 0; i < 10; i++ {
				res := l.RandomType(time.Now().UnixNano())
				resp, err := store.Eval(store.Context(), `
				local stock = tonumber(redis.call('hGet', KEYS[1], ARGV[1]))
				if stock <= 0 then
					return 2;
				end
				redis.call('hIncrBy', KEYS[1], ARGV[1], -1);
				return 1;`, []string{key}, res).Int64()
				if err != nil {
					panic(err)
				}
				if resp == 2 {
					res = maxKeys[rand.Intn(len(maxKeys)-1)]
					fmt.Println(resp, res)
				}
				mx.Lock()
				if res == 11 || res == 12 {
					fmt.Println((j+1)*(i+1), res)
				}
				result[res]++
				mx.Unlock()
			}
		}(j)
	}
	wg.Wait()
	fmt.Println(result)

	var cc int64
	for _, v := range result {
		cc += v
	}
	fmt.Println(cc)

}

func TestNewLottery(t *testing.T) {
	rate := map[int]int64{
		1:  300,
		2:  1450,
		3:  1450,
		4:  300,
		5:  1450,
		6:  300,
		7:  1450,
		8:  300,
		9:  1450,
		10: 1450,
		11: 50,
		12: 50,
	}

	//stock := int64(1000) // 此次发售库存
	l := NewLottery(rate)
	result := make(map[int]int64)
	for j := 0; j < 10; j++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			for i := 0; i < 1; i++ {
				res := l.RandomType(time.Now().UnixNano())
				mx.Lock()
				if res == 11 || res == 12 {
					fmt.Println((j+1)*(i+1), res)
				}
				result[res]++
				mx.Unlock()
			}
		}(j)
	}
	wg.Wait()
	fmt.Println(result)

	var cc int64
	for _, v := range result {
		cc += v
	}
	fmt.Println(cc)
}
