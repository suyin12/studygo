package main

import (
	"errors"
	"fmt"
)

func queryById(id int64) (int64, error) {
	if id <= 0 {
		return id, errors.New("无效的id")
	}
	return id, nil
}

func main() {
	res, err := queryById(-1)
	fmt.Println(res)
	fmt.Println(err)
	// fmt.Errorf("查询数据库失败，err:%w", err)
}
