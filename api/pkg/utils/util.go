package utils

import (
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"mm-ndj/config"
	"mm-ndj/model"
	"os"
	"strconv"
	"strings"
	"time"

	"go.uber.org/zap"
)

type ty int

const (
	datas ty = iota + 1
)

func StructToJson(collection []interface{}) string {
	for _, instance := range collection {
		// 判断是否是结构体
		marshal, err := json.Marshal(instance)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(marshal))
		//return fmt.Sprintf("%v", string(marshal))
	}

	return ""
}

var AlphanumericSet = []rune{
	'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
	'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
	'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
}

func GenInviteCode(uid string, l int) string {
	if l > 16 {
		return ""
	}
	sum := md5.Sum([]byte(uid))
	var code []rune
	for i := 0; i < l; i++ {
		idx := sum[i] % byte(len(AlphanumericSet))
		code = append(code, AlphanumericSet[idx])
	}
	return string(code)
}
func GenStarCode() string {
	id, err := GenID()
	if err != nil {
		//如果雪花算法错误 就使用纳秒 足够了
		config.Logger.Error("sonyFlake.NextID()", zap.Error(err))
		nano := time.Now().UnixNano()
		nanostr := strconv.Itoa(int(nano))
		GenInviteCode(nanostr, 6)
		return nanostr
	}
	snow := strconv.Itoa(id)
	return GenInviteCode(snow, 6)
}

func PageHelper(pageNum, pageSize int, order string, total int) model.PageInfo {
	var page model.PageInfo
	var lastPage int
	if pageNum == 0 {
		pageNum = 1
	}
	if pageSize == 0 {
		pageSize = 10
	}
	page.Total = total
	if total == 0 {
		lastPage = 1
	} else {
		if total%pageSize > 0 {
			lastPage = total/pageSize + 1
		} else {
			lastPage = total / pageSize
		}
	}
	page.LastPage = lastPage
	if pageNum < 1 {
		page.PageNum = 1
	} else if pageNum > lastPage {
		page.PageNum = pageNum
	} else {
		page.PageNum = pageNum
	}
	//单页最大20
	if pageSize > 20 {
		page.PageSize = 20
	} else if pageSize < 3 {
		//单页最小3
		page.PageSize = 3
	} else {
		page.PageSize = pageSize
	}
	//offset
	page.Offset = (page.PageNum - 1) * page.PageSize
	//order
	if order == "ASC" || order == "asc" {
		page.Order = "ASC"
	} else {
		page.Order = "DESC"
	}

	return page

}

func MapToJson(result interface{}) string {
	// map转 json str
	jsonBytes, _ := json.Marshal(result)
	jsonStr := string(jsonBytes)
	return jsonStr
}
func Unique(arr []int) []int {
	var arr_len int = len(arr) - 1
	for ; arr_len > 0; arr_len-- {
		// 拿最后项与前面的各项逐个(自后向前)进行比较
		for j := arr_len - 1; j >= 0; j-- {
			if arr[arr_len] == arr[j] {
				arr = append(arr[:arr_len], arr[arr_len+1:]...)
				break
			}
		}

		/*
		   // 或拿最后项与前面的各项逐个(自前向后)进行比较
		   for j := 0; j < arr_len; j++ {
		     if arr[arr_len] == arr[j] {
		     	// fmt.Printf("arr_len=%d equals j=%d\n ", arr[arr_len], arr[j])
		     	// 如果存在重复项，则将重复项删除，并重新给数组赋值
		       arr = append(arr[:arr_len], arr[arr_len + 1:]...)
		       break
		     }
		   }
		*/
	}
	return arr
}

// 这个利用map的去重算法会乱序
func UniqueAddress(addresslist []string) []string {
	m := make(map[string]string)
	for _, s := range addresslist {
		m[s] = s
	}
	var ret []string
	for _, v := range m {
		ret = append(ret, v)
	}
	return ret
}

//func StructToJson(instance struct{}) string {
//	// 判断是否是结构体
//	//json.MarshalIndent()
//	marshal, err := json.Marshal(instance)
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println(string(marshal))
//	return fmt.Sprintf("%v", string(marshal))
//}

//func GenSnowFlake(machineId int, datacenterId int) int {
//	var lastTimeStamp int64
//	var sn int
//	// 如果想让时间戳范围更长，也可以减去一个日期
//	curTimeStamp := time.Now().UnixNano() / 1000000
//
//	if curTimeStamp == lastTimeStamp {
//		// 2的12次方 -1 = 4095，每毫秒可产生4095个ID
//		if sn > 4095 {
//			time.Sleep(time.Millisecond)
//			curTimeStamp = time.Now().UnixNano() / 1000000
//			sn = 0
//		}
//	} else {
//		sn = 0
//	}
//	sn++
//	lastTimeStamp = curTimeStamp
//	// 应为时间戳后面有22位，所以向左移动22位
//	curTimeStamp = curTimeStamp << 22
//	machineId = machineId << 17
//	datacenterId = datacenterId << 12
//	// 通过与运算把各个部位连接在一起
//	return int(curTimeStamp) | machineId | datacenterId | sn
//}

func GetParticipateID(char string) string {
	/*
			考虑下这个用秒还是纳秒 竞品 sograph 是用的秒
		   这里stargate使用 秒 加 一个随机字符串 + 一个任务类型标识符
	*/
	unix := time.Now().Unix()
	unixnano := time.Now().UnixNano()

	s := strconv.FormatInt(unixnano, 10)
	code := GenInviteCode(s, 5)

	timestamp := strconv.Itoa(int(unix))

	pid := timestamp + code + "-" + char

	return pid
}
func GetRewardID(i int) string {
	//考虑下这个用秒还是纳秒 竞品 sograph 是用的秒
	//unix := time.Now().UnixNano()
	//timestamp := strconv.Itoa(int(unix))
	suffix := strconv.Itoa(i)
	unix := time.Now().Unix()
	unixnano := time.Now().UnixNano()

	s := strconv.FormatInt(unixnano, 10)
	code := GenInviteCode(s, 5)
	timestamp := strconv.Itoa(int(unix))

	rid := timestamp + code + "-" + suffix
	fmt.Println(rid)
	return rid
}
func Randbuilder() {
	rand.Seed(time.Now().UnixNano())
}

func ExistFile(filePath string) bool {
	_, err := os.Stat(filePath)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false

}

func QueryStringValidater(queryMap map[string]string) error {
	for k, v := range queryMap {
		if v == "" || v == "undefined" {
			msg := fmt.Sprintf("queryString: %v is absent or 'undefined' ,please check !", k)
			newerr := errors.New(msg)
			config.Logger.Error("QueryStringValidater() ", zap.Error(newerr))
			return newerr
		}
	}
	return nil
}

func DaoMapToIdSlice(list []map[string]interface{}, keyName string) (ids []int) {
	var idSlice []int

	for _, info := range list {
		switch value := info[keyName].(type) {
		case int64:
			idSlice = append(idSlice, int(value))
		case int:
			idSlice = append(idSlice, value)
		case string:
			id, _ := strconv.Atoi(value)
			idSlice = append(idSlice, id)
		}
	}
	return idSlice
}

func DaoMapToIdStringSlice(list []map[string]interface{}, keyName string) (ids []string) {
	var idStrSlice []string

	for _, info := range list {
		switch value := info[keyName].(type) {
		case int64:
			idStrSlice = append(idStrSlice, fmt.Sprint(value))
		case int:
			idStrSlice = append(idStrSlice, fmt.Sprint(value))
		case string:
			idStrSlice = append(idStrSlice, value)
		}
	}
	return idStrSlice
}

func GetUnixToDateShort(timestamp int64) string {
	t := time.Unix(timestamp, 0)
	return t.Format("2006-01-02")
}

func GetUnixToDateLong(timestamp int64) string {
	t := time.Unix(timestamp, 0)
	return t.Format("2006-01-02 15:04:05")
}

func GetTimeToDateShort(date time.Time) string {
	return date.Format("2006-01-02")
}

func StringLongToTime(date string) time.Time {
	loc, _ := time.LoadLocation("Local")
	dateTime, err := time.ParseInLocation("2006-01-02 15:04:05", date, loc)
	if err != nil {
		config.Logger.Error("StringToTime err", zap.Error(err))
	}
	return dateTime
}

type StatusType int

// 创建订单状态
const (
	OrderTy    = 7   //7天
	OrderTw    = 15  //15天
	OrderMonth = 30  //一个月
	Order3M    = 90  //三个月
	Order6m    = 180 //半年
)

// 执行动作
const (
	DaysStatus1 StatusType = iota + 1 //执行类型1
	DaysStatus2
	DaysStatus3
	DaysStatus4
	DaysStatus5
)

// 时间加减
func DateLate(day int) (int64, int) {
	var d int
	var insert int64
	now := time.Now()
	if day == int(DaysStatus1) {
		d = OrderTy
	}
	if day == int(DaysStatus2) {
		d = OrderTw
	}
	if day == int(DaysStatus3) {
		d = OrderMonth
	}
	if day == int(DaysStatus4) {
		d = Order3M
	}
	if day == int(DaysStatus5) {
		d = Order6m
	}
	late := fmt.Sprintf("%dh", 24*d)
	dd, _ := time.ParseDuration(late)
	insert = now.Add(dd).Unix()
	return insert, d
}

func GetUnixToDateChShort(timestamp int64) string {
	t := time.Unix(timestamp, 0)
	return t.Format("2006年01月02日")
}

// GetTokenUserId 获取用户id
func GetTokenUserId(c *gin.Context) (int64, error) {
	userId := c.Keys["userId"]

	if userId == "" {
		return 0, nil
	}
	return userId.(int64), nil
}

func GetEmailText(languageType, code string) string {
	pwd, err := os.Getwd()
	if err != nil {
		config.Logger.Error("os.Getwd", zap.Error(err))
		return ""
	}
	dataStr, err := os.ReadFile(pwd + "/model/send_verify_code_" + languageType + ".html")
	if err != nil {
		config.Logger.Error("os.ReadFile", zap.Error(err))
		return ""
	}
	retStr := strings.ReplaceAll(string(dataStr), "verifyCode", code)
	return retStr
}
