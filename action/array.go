package action

import (
	"fmt"
	"sort"
)

func DataTypeArray() {
	var array [3]string = [3]string{"WT", "XMM", "WYY"}
	fmt.Println(array)

	var array1 = [3]string{"WT", "XMM", "WYY"}
	fmt.Println(array1)

	var array2 = [...]int{1, 2, 3, 4}
	array2[0] = 0
	fmt.Println(array2)
}

func DataTypeIndex() {
	array := []string{"WT", "XMM", "WYY"}
	array = append(array, "XLH")

	fmt.Println(array)
	fmt.Println(array[1])
	fmt.Println(array[len(array)-1])
}

func DataTypeSlice() {
	var slices []string

	slices = append(slices, "WT")
	slices = append(slices, "XMM")
	slices = append(slices, "WYY")

	fmt.Println(slices)
}

func DataTypeMake() {
	// 只声明未复制则为nil
	var sliceNil []string
	fmt.Println(sliceNil == nil)

	// 声明一个空的切片，值不为nil
	var slices = make([]string, 0)

	fmt.Println(slices == nil)
	slices = append(slices, "WT")
	fmt.Println(slices)

	lists := make([]int, 2, 2)
	fmt.Println(lists)
}

func DataTypeArrayToSlice() {
	array := [3]string{"WT", "XMM", "WYY"}
	fmt.Println(array)

	slice := array[:]
	fmt.Println(slice)
	fmt.Println(slice[0:2])
}

func DataTypeSliceSort() {
	list := []int{4, 5, 3, 2, 7}
	fmt.Println("排序前:", list)

	sort.Ints(list)
	fmt.Println("升序排序后:", list)

	sort.Sort(sort.Reverse(sort.IntSlice(list)))
	fmt.Println("降序排序后:", list)
}

func DataTypeMap() {
	mapList := map[string]int{}
	fmt.Println(mapList)

	mapList1 := make(map[string]string, 0)
	mapList1["sex"] = "male"
	mapList1["name"] = "wt"
	fmt.Println(mapList1)
	fmt.Println(mapList1["name"])

	delete(mapList1, "name")
	fmt.Println(mapList1)
}

func DataTypeMapGetValue() {
	mapList := map[string]int{
		"age":    23,
		"height": 175,
		"weight": 75,
	}

	sex := mapList["sex"]
	fmt.Println(sex)

	value, error := mapList["sex"]
	fmt.Println(value, error)

	value1, error1 := mapList["age"]
	fmt.Println(value1, error1)
}
