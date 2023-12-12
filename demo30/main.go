package main

func main() {
	escape5()
}

// func escape1() *int {
// 	var a int = 1
// 	return &a
// }

// func escape2() {
// 	s := make([]int, 0, 10000)
// 	for index := range s {
// 		s = append(s, index)
// 	}
// }

// func escape3() {
// 	number := 10
// 	s := make([]int, number)
// 	fmt.Println(s)
// }

// func escape4() {
// 	fmt.Println(111)
// }

func escape5() func() int {
	x := 1
	return func() int {
		return x
	}
}
