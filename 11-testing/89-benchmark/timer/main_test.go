package timer

import "testing"

func BenchmarkFoo1(b *testing.B) {
	expensiveSetup()
	// 大きな時間がかかる操作の後に呼び出す
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		functionUnderTest()
	}
}

// loopの場合はこちらの方が良い
func BenchmarkFoo2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// 準備時間はタイマーを切る
		b.StopTimer()
		expensiveSetup()
		b.StartTimer()
		functionUnderTest()
	}
}

func functionUnderTest() {
}

func expensiveSetup() {
}
