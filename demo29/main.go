package main

import (
	"fmt"
	"sort"
)

func main() {
	demo5()
}

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s: %d", p.Name, p.Age)
}

type Persons []Person

// 实现 sort.Interface 的三个方法
func (a Persons) Len() int {
	return len(a)
}

func (a Persons) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a Persons) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}

func demo1() {
	peoples := Persons{
		{"Bob", 21},
		{"John", 42},
		{"Mike", 18},
		{"Kitty", 18},
		{"Rose", 28},
	}
	fmt.Printf("未排序的数据：%v\n", peoples)
	fmt.Printf("IS Sorted：%t\n", sort.IsSorted(peoples))
	// sort.Sort(peoples)
	sort.Stable(peoples)
	fmt.Printf("已排好的数据：%v\n", peoples)
	fmt.Printf("IS Sorted：%t\n", sort.IsSorted(peoples))

	sort.Sort(sort.Reverse(peoples))
	fmt.Printf("倒序排序的结果：%v\n", peoples)
}

func demo2() {
	peoples := Persons{
		{"Bob", 21},
		{"John", 42},
		{"Mike", 18},
		{"Kitty", 18},
		{"Rose", 28},
	}
	sort.Sort(peoples)
	fmt.Println(peoples)
	index := sort.Search(len(peoples), func(i int) bool {
		return peoples[i].Age >= 30
	})
	fmt.Println(index)
}

// ----------- (●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●) ------------

func demo3() {
	x := 36
	data := []int{1, 21, 23, 66}
	index := sort.Search(len(data), func(i int) bool {
		return data[i] >= x
	})
	if index < len(data) && data[index] == x {
		fmt.Printf("找到了55，索引为 %d\n", index)
	} else {
		// 插入值到找到的位置
		data = append(data, 0)             // 扩展切片，添加一个临时元素
		copy(data[index+1:], data[index:]) // 向后移动元素
		data[index] = x                    // 插入值到指定位置
		fmt.Println(index)                 // 3
		fmt.Println(data)                  // [1 21 23 36 66]
	}
}

// ----------- (●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●) ------------

func demo4() {
	persons := []Person{
		{"Bob", 21},
		{"John", 42},
		{"Mike", 18},
		{"Kitty", 18},
		{"Rose", 28},
	}
	// 排序函数
	less := func(i, j int) bool {
		return persons[i].Age < persons[j].Age
	}
	sort.Slice(persons, less)
	sort.SliceStable(persons, less)

	fmt.Printf("已排好的数据：%v\n", persons)                              // 排好的数据：[Mike: 18 Kitty: 18 Bob: 21 Rose: 28 John: 42]
	fmt.Printf("IS Sorted：%t\n", sort.SliceIsSorted(persons, less)) // IS Sorted：true

	// 查找年龄等于 20 的学生的索引位置，判断条件必须为 >=
	f := func(i int) bool {
		return persons[i].Age >= 20
	}
	fmt.Println(sort.Search(len(persons), f)) // 2
}

// ----------- (●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●) ------------
func demo5() {
	intLists := []int{2, 3, 1, 8, 5, 4, 9}
	// 方式一：
	sort.Ints(intLists)
	// 方式二：
	sort.Sort(sort.IntSlice(intLists))
	// 方式三：
	sort.IntSlice(intLists).Sort()
	// 方式四：
	sort.Slice(intLists, func(i, j int) bool {
		return intLists[i] < intLists[j]
	})
	fmt.Printf("已排好的数据：%v\n", intLists)
	// 判断是否为递增排序：
	// 方式一：
	fmt.Println(sort.IntsAreSorted(intLists)) // true
	// 方式二：
	fmt.Println(sort.IsSorted(sort.IntSlice(intLists))) // true
}

func demo6() {
	intLists := []int{2, 3, 1, 8, 5, 4, 9}
	// 方式一：
	sort.Sort(sort.Reverse(sort.IntSlice(intLists)))
	// 方式二：
	sort.Slice(intLists, func(i, j int) bool {
		return intLists[i] > intLists[j]
	})
	fmt.Printf("已排好的数据：%v\n", intLists)
}

// 在递增序列中搜索
func demo7() {
	intLists := []int{2, 3, 1, 8, 5, 4, 9}
	// 方式一：
	fmt.Println(sort.SearchInts(intLists, 2)) // 4
	// 方式二：
	fmt.Println((sort.IntSlice(intLists)).Search(2)) // 4
}
