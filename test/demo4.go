package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	str := ComputeAbnTimeLimit("W", TimeDifference("2020-01-07 09:43:32", ""))
	fmt.Println("结果：------", str)
}

func ComputeAbnTimeLimit(work, timeLimit string) string {
	var (
		strTime []string
		strI    int
		val     float64
		flval   float64
		strArr  []string
		flstr   string
		hint    int
	)
	//timelimit: 24h23m21s	还有小于1小时的。。
	strTime = strings.Split(timeLimit, "h")
	fmt.Println("截取的h：", strTime)
	strI, _ = strconv.Atoi(strTime[0])
	//val 三种值：小于1，等于1，大于1
	// strI 取整的时候。
	val = float64(strI) / float64(24)
	fmt.Println("除以24：", val)
	if val > 1 {
		flval, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", val), 64) //取浮点数2位小数
		flstr = strconv.FormatFloat(flval, 'f', -1, 64)             //浮点数转换为string
		strArr = strings.Split(flstr, ".")
		//整除是长度为1，直接返回strArr的值，不整除时长度为2
		if len(strArr) == 1 {
			//return strArr[0] + "d"
			return strArr[0] + WorkOrRest(work)
		}
		//当长度为2时，取第二位字符，并转为整数来进行判断是否大于50
		hint, _ = strconv.Atoi(strArr[1])
		//fmt.Println("hint:", hint)
		//如果小于50说明是1天，大于是1.5天
		if hint == 5 {
			//return strArr[0] + ".5"
			return strArr[0] + ".5" + WorkOrRest(work)
		} else if hint < 50 {
			//小于50，直接返回小数点前面的数字
			return strArr[0] + WorkOrRest(work)
		} else {
			//大于50，返回小数点 数字小数点+5
			//return strArr[0] + ".5"
			return strArr[0] + ".5" + WorkOrRest(work)
		}
	} else if val == 1 {
		return "1"
	} else {
		//当小于1天时，就直接返回耗时时间
		//strings.Split(timeLimit,"m")
		fmt.Println("最后转换的时间为----", timeLimit)
		hstr := strings.Replace(timeLimit, "h", "小时", -1)
		mstr := strings.Replace(hstr, "m", "分", -1)
		sstr := strings.Replace(mstr, "s", "秒", -1)
		return sstr
	}

}

// 计算时间差
func TimeDifference(endTime, startTime string) string {
	t1, _ := time.ParseInLocation("2006-01-02 15:04:05", startTime, time.Local)
	t2, _ := time.ParseInLocation("2006-01-02 15:04:05", endTime, time.Local)
	return t2.Sub(t1).String()
}

// 判断是工作日还是自然日
func WorkOrRest(str string) (work string) {
	if str == "W" {
		return "工作日"
	}
	if str == "N" {
		return "自然日"
	}
	return
}
