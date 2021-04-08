package admin

import (
	"strconv"
	"strings"
)

// 定义一个DbId类型，64位整型，改写json方法
type Juint64 uint64

// 这里主要注意反序列化时需要去掉两边的引号
func (d *Juint64) UnmarshalJSON(data []byte) (err error) {
    s := strings.Trim(string(data), "\"")
    intNum, err := strconv.ParseUint(s, 10, 64)
    if err != nil {
        return err
    }
    *d = Juint64(intNum)
    return
}

// 这里需要手动加上引号，不会自动生成
func (d *Juint64) MarshalJSON() ([]byte, error) {
    return ([]byte)("\"" + strconv.FormatUint(uint64(*d), 10) + "\""), nil
}
// 这里需要手动加上引号，不会自动生成
func (d *Juint64) ToString() (string) {
    return string( ([]byte)("\"" + strconv.FormatUint(uint64(*d), 10) + "\""))
}