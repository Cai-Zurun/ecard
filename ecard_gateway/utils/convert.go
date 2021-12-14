package utils

import (
	"encoding/binary"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
	"unicode/utf16"
	"unicode/utf8"
)

/*
字符串反转：先将字符串转换成为rune（即int32）类型，再将其首位元素依次交换
*/
func ReverseString(s string) string {
	runes := []rune(s)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}

/*
心跳包（反转）、GPS_4G定位包：
将数据包中的16进制转为二进制字符串
可选择是否要字符串反转，反转达到的效果：
1字节中Bit7就对应bin[7],Bit0对应bin[0]（心跳包）
*/
func BytesToBinary(buf []byte, reverse bool) string {
	var bin string
	for _, val := range buf {
		binVal := strconv.FormatInt(int64(val), 2)
		// 如果binVal小于8位，在其前面补0，再与bin拼接
		bin += fmt.Sprint(strings.Repeat("0", 8-len(binVal)), binVal)
	}
	if reverse {
		return ReverseString(bin)
	}
	return bin
}

/*
心跳包：
通过电压等级计算电量大小
电量%=电压v（GK320电量大小就等于计算出来的电压大小）
*/
func CalculateElectQuantity(electLevel []byte) string {
	electVal := (int(electLevel[0])*256 + int(electLevel[1])) / 100
	electQuantity := fmt.Sprint(electVal)
	return electQuantity
}

func StringToFloat64(s string) (f float64) {
	f, _ = strconv.ParseFloat(s, 64)
	return f
}

/*
GPS_4G定位包、4G多围栏报警包：
将数据包中的日期转换为我们指定的时间格式
*/
func ConvertTime(buf []byte) string {
	// 观察发现，数据包中年份大小=当前年份-2000，所以我们要手动加上2000
	year := 2000 + int(buf[0])
	month := time.Month(buf[1])
	day := int(buf[2])
	hour := int(buf[3])
	min := int(buf[4])
	sec := int(buf[5])
	t := time.Date(year, month, day, hour, min, sec, 0, time.UTC)
	return t.Format("20060102.1504")
}

/*
GPS_4G定位包、4G多围栏报警包：
将数据包中的日期转换为我们指定的时间格式
为轨迹数据生成时间格式 2021-04-01
*/
func ConvertTimeForTrajectory(buf []byte) string {
	// 观察发现，数据包中年份大小=当前年份-2000，所以我们要手动加上2000
	year := 2000 + int(buf[0])
	month := time.Month(buf[1])
	day := int(buf[2])
	hour := int(buf[3])
	min := int(buf[4])
	sec := int(buf[5])
	t := time.Date(year, month, day, hour, min, sec, 0, time.UTC)
	return t.Format("2006-01-02 15:04:05")
}

func IntToString(ValueInt int) (ValueStr string) {
	ValueStr = strconv.Itoa(ValueInt)
	return
}

/*
GPS_4G定位包：
将16进制字节切片转换成10进制和
*/
func BytesToDecimal(buf []byte) float64 {
	var sum float64 = 0
	var j float64 = 0
	for i := len(buf) - 1; i >= 0; i-- {
		sum += float64(buf[i]) * math.Pow(256, j)
		j++
	}
	return sum
}

/*
将int转成两字节的[]byte，信息序列号只占两字节
*/
func IntToSeqBytes(num int) (seq []byte) {
	seq = append(seq, byte(num/256))
	seq = append(seq, byte(num%256))
	return
}

/*
校时包回复：
返回六个字节的当前时间
*/
func TimeNowToBytes() []byte {
	year := byte(time.Now().Year() - 2000)
	month := byte(time.Now().Month())
	day := byte(time.Now().Day())
	hour := byte(time.Now().Hour())
	min := byte(time.Now().Minute())
	sec := byte(time.Now().Second())
	return []byte{year, month, day, hour, min, sec}
}

/*
将由UTF16-BE编码的字节数组转成字符串
*/
func UTF16BytesToString(b []byte, o binary.ByteOrder) string {
	utf := make([]uint16, (len(b)+(2-1))/2)
	for i := 0; i+(2-1) < len(b); i += 2 {
		utf[i/2] = o.Uint16(b[i:])
	}
	if len(b)/2 < len(utf) {
		utf[len(utf)-1] = utf8.RuneError
	}
	return string(utf16.Decode(utf))
}
