package main

import (
	"fmt"
	"strconv"
)

func genValueBy(key int) string {
	return strconv.Itoa(key)
}

func fillMap(m map[int]string, length int) {
	if length < 0 {
		length = 0
	}
	var count int
	for count < length {
		m[count] = genValueBy(count)
		count++
	}
}

func main() {
	max := 10
	map1 := make(map[int]string)
	fillMap(map1, max)

	fmt.Printf("Map: [\n  ")
	for k, v := range map1 {
		fmt.Printf("%d:%s ", k, v)
	}
	fmt.Printf("\n]\n\n")

	fmt.Printf("Map: %v\n\n", map1)

	// 生成第一个快照。
	snapshot1 := fmt.Sprint(map1)
	fmt.Printf("Snapshot 1: %s\n", snapshot1)
	// 修改 map1：增加一些新的键值对。
	for i := max + 1; i <= max+3; i++ {
		key1 := i
		map1[key1] = genValueBy(key1)
	}
	// 修改 map1：改动一个已有的键值对。
	map1[max-5] = "0"
	// 生成第二个快照。
	snapshot2 := fmt.Sprint(map1)
	fmt.Printf("Snapshot 2: %s\n", snapshot2)
	// 直接比较两个快照。
	fmt.Printf("Snapshot 2 > Snapshot 1 ? %v\n", snapshot2 > snapshot1)

}
