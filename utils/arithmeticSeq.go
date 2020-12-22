package utils

import (
	"fmt"
	"strconv"
	"strings"
)

type ArithmeticSeqInfo struct {
	Lev   int32
	Value int32
}

type ArithmeticSeq struct {
	mapData    map[int32]*ArithmeticSeqInfo
	dataList   []*ArithmeticSeqInfo
	strData    string
	FirstValue int32
	MaxLev     int32
}

func NewArithmeticSeq(data string) *ArithmeticSeq {
	seq := new(ArithmeticSeq)
	if !seq.init(data) {
		return nil
	}
	return seq
}

func (as *ArithmeticSeq) init(data string) bool {
	as.strData = data
	as.mapData = make(map[int32]*ArithmeticSeqInfo)
	as.dataList = make([]*ArithmeticSeqInfo, 0)
	if len(data) == 0 {
		return true
	}

	groupList := strings.Split(data, ",")
	for _, v := range groupList {
		group := strings.Split(v, ":")
		if len(group) != 2 {
			return false
		}
		lev, err := strconv.Atoi(group[0])
		if nil != err {
			continue
		}
		value, err := strconv.Atoi(group[1])
		if nil != err {
			fmt.Printf("data:%v group:%s error.\n", data, group)
			continue
		}
		info := &ArithmeticSeqInfo{Lev: int32(lev), Value: int32(value)}
		as.dataList = append(as.dataList, info)
		if info.Lev > as.MaxLev {
			as.MaxLev = info.Lev
		}
	}

	for lev :=1; lev <=int(as.MaxLev); lev++ {
		value := as.getCalcValue(int32(lev))
		info := &ArithmeticSeqInfo{Lev: int32(lev), Value: int32(value)}
		as.mapData[info.Lev] = info
	}
	as.FirstValue = as.getCalcValue(1)

	return true
}

func (as *ArithmeticSeq) GetValue(lev int32) int32 {
	if lev > as.MaxLev {
		return 0
	}
	var value int32
	v, ok := as.mapData[lev]
	if !ok {
		value = as.getCalcValue(lev)
	} else {
		value = v.Value
	}

	return value
}

//返回计算后相对应的值
//返回-1 为无效值
func (as *ArithmeticSeq) getCalcValue(lev int32) int32 {
	if lev > as.MaxLev {
		return -1
	}

	count := len(as.dataList)
	if count == 0 {
		return 0
	}

	if 1 == count {
		return as.dataList[0].Value
	}

	first := as.dataList[0]
	last := as.dataList[count-1]

	if lev <= first.Lev {
		return calValue(as.dataList[0], as.dataList[1], lev)
	}
	if lev >= last.Lev {
		return calValue(as.dataList[count-2], as.dataList[count-1], lev)
	}

	for i := 0; i < count-1; i++ {
		start := as.dataList[i]
		end := as.dataList[i+1]

		if lev == start.Lev {
			return start.Value
		}
		if lev == end.Lev {
			return end.Value
		}
		//在该区间范围内，通过斜率进行计算
		if lev > start.Lev && lev < end.Lev {
			if end.Lev != start.Lev {
				return calValue(start, end, lev)
			}
		}
	}
	return 0
}

func calValue(start, end *ArithmeticSeqInfo, lev int32) int32 {
	k := float64(end.Value-start.Value) / float64(end.Lev-start.Lev+0.0)
	res := k*float64(lev-start.Lev) + float64(start.Value)
	return (int32)(res + 0.5)
}

func (as ArithmeticSeq) GetFirstValue() int32 {
	return as.FirstValue
}
