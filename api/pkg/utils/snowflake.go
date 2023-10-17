package utils

//
//import (
//	"errors"
//	"sync"
//	"time"
//)
//
//const (
//	workerBits  uint8 = 10                      //机器码位数
//	numberBits  uint8 = 12                      //序列号位数
//	workerMax   int64 = -1 ^ (-1 << workerBits) //机器码最大值（即1023）
//	numberMax   int64 = -1 ^ (-1 << numberBits) //序列号最大值（即4095）
//	timeShift   uint8 = workerBits + numberBits //时间戳偏移量
//	workerShift uint8 = numberBits              //机器码偏移量
//	epoch       int64 = 1656856144640           //起始常量时间戳（毫秒）,此处选取的时间是2022-07-03 21:49:04
//)
//
//type Worker struct {
//	mu        sync.Mutex
//	timeStamp int64
//	workerId  int64
//	number    int64
//}
//
//func NewWorker(workerId int64) (*Worker, error) {
//	if workerId < 0 || workerId > workerMax {
//		return nil, errors.New("WorkerId超过了限制！")
//	}
//	return &Worker{
//		timeStamp: 0,
//		workerId:  workerId,
//		number:    0,
//	}, nil
//}
//
//func (w *Worker) NextId() int64 {
//	w.mu.Lock()
//	defer w.mu.Unlock()
//	//当前时间的毫秒时间戳
//	now := time.Now().UnixNano() / 1e6
//	//如果时间戳与当前时间相同，则增加序列号
//	if w.timeStamp == now {
//		w.number++
//		//如果序列号超过了最大值，则更新时间戳
//		if w.number > numberMax {
//			for now <= w.timeStamp {
//				now = time.Now().UnixNano() / 1e6
//			}
//		}
//	} else { //如果时间戳与当前时间不同，则直接更新时间戳
//		w.number = 0
//		w.timeStamp = now
//	}
//	//ID由时间戳、机器编码、序列号组成
//	ID := (now-epoch)<<timeShift | (w.workerId << workerShift) | (w.number)
//	return ID
//}
import (
	"fmt"
	_ "github.com/onsi/gomega"
	"go.uber.org/zap"
	"mm-ndj/config"
	"strconv"
	"time"

	"github.com/sony/sonyflake"
)

var (
	sonyFlake     *sonyflake.Sonyflake
	sonyMachineID uint16
)

func getMachineID() (uint16, error) {
	return sonyMachineID, nil
}

// 需传入当前的机器ID

func SonyFlakeInit(startTime string, machineId uint16) (err error) {
	sonyMachineID = machineId
	var st time.Time
	st, err = time.Parse("2006-01-02", startTime)
	if err != nil {
		return err
	}
	settings := sonyflake.Settings{
		StartTime: st,
		MachineID: getMachineID,
	}
	sonyFlake = sonyflake.NewSonyflake(settings)
	return
}

// GenID 生成id  ： 主要是  userID missionID 和 spaceID projectID
func GenID() (id int, err error) {
	if sonyFlake == nil {
		err = fmt.Errorf("sony flake not inited")
		config.Logger.Error("GenID()", zap.Error(err))
		return
	}
	genid, err := sonyFlake.NextID()
	if err != nil {
		config.Logger.Error("sonyFlake.NextID()", zap.Error(err))
		return
	}
	id = int(genid)
	return
}

func ValidId(id string) bool {
	if len(id) != 16 {
		return false
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return false
	}
	if idInt <= 0 {
		return false
	}
	return true
}
