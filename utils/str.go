package utils

import (
	"bytes"
	"fmt"
	"math/rand"
	"net/url"
	"sort"
	"strconv"
	"time"
)

const (
	KC_RAND_KIND_NUM   = 0 // 纯数字
	KC_RAND_KIND_LOWER = 1 // 小写字母
	KC_RAND_KIND_UPPER = 2 // 大写字母
	KC_RAND_KIND_ALL   = 3 // 数字、大小写字母
)

//URLToMap
func URLToMap(s, key string) string {
	u, err := url.Parse(s)
	if err != nil {
		fmt.Println("URLToMap =>", err.Error())
		return ""
	}
	m, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		fmt.Println("URLToMap =>", err.Error())
	}

	if vaccount, ok := m["account"]; ok {
		return vaccount[0]
	}
	return ""
}

//StrToInt
func StrToInt(str string) int {
	idata, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("StrToInt =>", err.Error())
		return 0
	}
	return idata
}

func StrToInt64(s string) int64 {
	return int64(StrToInt(s))
}

//IntToStr
func IntToStr(i int) string {
	return fmt.Sprintf("%d", i)
}

func IntersToStrs(idata []interface{}) []string {
	var mdata []string
	for _, v := range idata {
		s, ok := v.(string)
		if ok {
			mdata = append(mdata, s)
		}
	}
	return mdata
}

// 随机字符串
func Krand(size int, kind int) []byte {
	ikind, kinds, result := kind, [][]int{{10, 48}, {26, 97}, {26, 65}}, make([]byte, size)
	is_all := kind > 2 || kind < 0
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		if is_all { // random ikind
			ikind = rand.Intn(3)
		}
		scope, base := kinds[ikind][0], kinds[ikind][1]
		result[i] = uint8(base + rand.Intn(scope))
	}
	return result
}

func GetOrdeID(i int) string {
	if i == 0 {
		i = 6
	}
	return fmt.Sprintf("%d", NstimeUnix()) + string(Krand(7, KC_RAND_KIND_NUM))
}

//[]string 转 []int
func StrSlicetoInt(sdata []string, stype int) ([]int, []string) {
	var ballis []int
	var balls []string
	if len(sdata) == 0 {
		return nil, nil
	}
	for _, r := range sdata {
		bnum, _ := strconv.Atoi(r)
		ballis = append(ballis, bnum)
	}
	//返回没有排序的数据
	if stype == 1 {
		return ballis, nil
	}

	sort.Ints(ballis)

	for _, ri := range ballis {
		bnum := strconv.Itoa(ri)
		balls = append(balls, bnum)
	}
	return ballis, balls
}

func TransHtmlJson(data []byte) []byte {
	data = bytes.Replace(data, []byte(`u0026`), []byte("&"), -1)
	data = bytes.Replace(data, []byte(`u003c`), []byte("<"), -1)
	data = bytes.Replace(data, []byte(`u003e`), []byte(">"), -1)
	return data
}
