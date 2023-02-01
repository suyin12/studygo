package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	// demo1()
	// demo2()
	// demo3()
	// demo4()
	// demo5()
	// demo6()
	// demo7()
	demo8()
}

func demo1() {
	//首先声明
	map1 := make(map[string]int, 10)
	//赋值
	map1["sujianhui"] = 20
	map1["xiaojuan"] = 18
	map1["xiaojuan1"] = 18
	map1["xiaojuan2"] = 18
	map1["xiaojuan3"] = 18
	map1["xiaojuan4"] = 18
	map1["xiaojuan5"] = 18
	map1["xiaojuan6"] = 18
	map1["xiaojuan7"] = 18
	map1["xiaojuan8"] = 18
	map1["xiaojuan9"] = 18
	map1["xiaojuan10"] = 18
	map1["xiaojuan11"] = 18
	map1["xiaojuan12"] = 18
	fmt.Println(map1)
	fmt.Println(len(map1))
}

func demo2() {
	//声明并赋值
	map2 := map[string]string{
		"username": "sujianhui",
		"password": "123456",
	}
	fmt.Println(map2)

	// 如果key存在ok为true,v为对应的值；不存在ok为false,v为值类型的零值
	value, key := map2["username"]
	if key {
		fmt.Println(value)
	} else {
		fmt.Println("查无此人")
	}
}

// 循环遍历
func demo3() {
	map3 := map[string]int{
		"su":   10,
		"jian": 18,
		"hui":  20,
		"xiao": 25,
		"juan": 28,
	}
	for key, value := range map3 {
		fmt.Printf("key %v, value %d\n", key, value)
	}

	//删除delete(map,key)
	delete(map3, "su")
	for key, value := range map3 {
		fmt.Printf("key %v, value %d\n", key, value)
	}

	dd := fmt.Sprintf("stu%02d", 20)
	fmt.Println(dd)
}

// 排序后遍历
func demo4() {
	rand.Seed(time.Now().UnixNano()) //初始化随机数种子

	var scoreMap = make(map[string]int, 200)

	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i) //生成stu开头的字符串
		value := rand.Intn(100)          //生成0~99的随机整数
		scoreMap[key] = value
	}
	//取出map中的所有key存入切片keys
	var keys = make([]string, 0, 200)
	for key := range scoreMap {
		keys = append(keys, key)
	}
	//对切片进行排序
	sort.Strings(keys)
	//按照排序后的key遍历map
	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}
}

// 元素为map类型的切片
func demo5() {
	var mapSlice = make([]map[string]string, 3)
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}
	fmt.Println("after init")
	// 对切片中的map元素进行初始化
	mapSlice[0] = make(map[string]string, 10)
	mapSlice[0]["name"] = "小王子"
	mapSlice[0]["password"] = "123456"
	mapSlice[0]["address"] = "沙河"
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}
}

// 值为切片类型的map
func demo6() {
	var sliceMap = make(map[string][]string, 3)
	fmt.Println(sliceMap)
	fmt.Println("after init")
	key := "中国"
	value, ok := sliceMap[key]
	if !ok {
		value = make([]string, 0, 2)
	}
	value = append(value, "北京", "上海", "深圳")
	sliceMap[key] = value
	fmt.Println(sliceMap)
}

// 值为切片类型的map
func demo7() {
	map1 := make(map[string][]string, 10)
	fmt.Println(map1)
	key := "中国"
	value, k := map1[key]
	if k {
		fmt.Println(value)
	} else {
		value = make([]string, 0, 10)
	}
	value = append(value, "北京", "上海", "深圳")
	fmt.Println(value)
	map1[key] = value
	fmt.Println(map1)
}

// 元素为map类型的切片
func demo8() {
	slice1 := make([]map[string]string, 3)
	fmt.Println(slice1)
	slice1[0] = make(map[string]string, 10)
	slice1[0]["su"] = "粟"
	slice1[0]["jian"] = "建"
	slice1[0]["hui"] = "晖"
	fmt.Println(slice1)
	for index, value := range slice1 {
		fmt.Printf("index:%d value:%v\n", index, value)
	}
}
