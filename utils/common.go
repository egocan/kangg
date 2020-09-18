package utils

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"kangg/models/admin"
	"reflect"
	"strconv"
)

// @Title 生成密码
// @Description create AccountAccount
// @Param	body		body 	models.AccountAccount	true		"body for AccountAccount content"
// @Success 201 {int} models.AccountAccount
// @Failure 403 body is empty
func PasswordMD5(passwd, salt string) string {
	h := md5.New()
	// 后面增加一个无意义字符串
	h.Write([]byte(passwd + salt + "@.YnO-"))
	cipherStr := h.Sum(nil)
	result := hex.EncodeToString(cipherStr)
	return result
}

// ToString 类型转换，获得string
func ToString(v interface{}) (re string) {
	re = v.(string)
	return
}

// StringsJoin 字符串拼接
func StringsJoin(strs ...string) string {
	var str string
	var b bytes.Buffer
	strsLen := len(strs)
	if strsLen == 0 {
		return str
	}
	for i := 0; i < strsLen; i++ {
		b.WriteString(strs[i])
	}
	str = b.String()
	return str

}

// ToInt64 类型转换，获得int64
func ToInt64(v interface{}) (re int64, err error) {
	switch v.(type) {
	case string:
		re, err = strconv.ParseInt(v.(string), 10, 64)
	case float64:
		re = int64(v.(float64))
	case float32:
		re = int64(v.(float32))
	case int64:
		re = v.(int64)
	case int32:
		re = v.(int64)
	default:
		err = errors.New("不能转换")
	}
	return
}

// ToSlice 转换为数组
func ToSlice(arr interface{}) []interface{} {
	v := reflect.ValueOf(arr)
	if v.Kind() != reflect.Slice {
		panic("toslice arr not slice")
	}
	l := v.Len()
	ret := make([]interface{}, l)
	for i := 0; i < l; i++ {
		ret[i] = v.Index(i).Interface()
	}
	return ret
}

type CateTree struct {
	Id    int
	Name  string
	Ename string
	Pid   int
	Sort  int
	Level int
	Son   []CateTree
}

//递归实现(返回树状结果得数据)
func CategoryTree(allCate []*admin.Category, pid int, level int) []CateTree {
	var arr []CateTree
	for _, v := range allCate {
		if pid == v.Pid {
			ctree := CateTree{}
			ctree.Id = v.Id
			ctree.Pid = v.Pid
			ctree.Name = v.Name
			ctree.Ename = v.Ename
			ctree.Sort = v.Sort
			ctree.Level = level
			sonCate := CategoryTree(allCate, v.Id, level+1)
			ctree.Son = sonCate
			arr = append(arr, ctree)
		}
	}
	return arr
}

func CategoryTreeR(allCate []*admin.Category, pid int, level int) []CateTree {
	var arr []CateTree
	for _, v := range allCate {
		if pid == v.Pid {
			ctree := CateTree{}
			ctree.Id = v.Id
			ctree.Pid = v.Pid
			ctree.Name = v.Name
			ctree.Ename = v.Ename
			ctree.Sort = v.Sort
			ctree.Level = level
			arr = append(arr, ctree)
			sonCate := CategoryTreeR(allCate, v.Id, level+1)
			arr = append(arr, sonCate...)
			//ctree.Son = sonCate
			//arr = append(arr, ctree)
		}
	}
	return arr
}

func SubString(str string,len int) string  {
	return string([]rune(str)[:len])
}



type MenuTree struct {
	Id    int
	Title  string
	Pid   int
	Sort  int
	Level int
	Url   string
	Target string
	Son   []MenuTree
}

func MenuTreeR(allCate []interface{}, pid int, level int) []MenuTree {
	var arr []MenuTree
	for _, v := range allCate {
		v1 := v.(admin.Menu)
		if pid == v1.Pid {
			ctree := MenuTree{}
			ctree.Id = v1.Id
			ctree.Pid = v1.Pid
			ctree.Title = v1.Title
			ctree.Sort = v1.Sort
			ctree.Level = level
			ctree.Url = v1.Url
			ctree.Target = v1.Target
			arr = append(arr, ctree)
			sonCate := MenuTreeR(allCate, v1.Id, level+1)
			arr = append(arr, sonCate...)
		}
	}
	return arr
}

func MenuData(allCate []interface{}, pid int, level int) []MenuTree {
	var arr []MenuTree
	for _, v := range allCate {
		v1 := v.(admin.Menu)
		if pid == v1.Pid {
			ctree := MenuTree{}
			ctree.Id = v1.Id
			ctree.Pid = v1.Pid
			ctree.Title = v1.Title
			ctree.Sort = v1.Sort
			ctree.Level = level
			ctree.Url = v1.Url
			ctree.Target = v1.Target
			sonCate := MenuData(allCate, v1.Id, level+1)
			ctree.Son = sonCate
			arr = append(arr, ctree)
		}
	}
	return arr
}
