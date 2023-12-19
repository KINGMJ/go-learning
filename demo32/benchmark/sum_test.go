package benchmark

import "testing"

func BenchmarkMutexSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// 被测试的代码片段或函数
		result := MutexSum()
		// 防止编译器优化掉函数调用
		_ = result
	}
}

func BenchmarkAtomicSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// 被测试的代码片段或函数
		result := AtomicSum()
		// 防止编译器优化掉函数调用
		_ = result
	}
}
