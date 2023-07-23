// File untuk keperluan testing(test ataupun benchmark) dipisah dengan file utama, namanya harus berakhiran _test.go, dan package-nya harus sama
// Testing atau run nya di cmd, di vscode semenetara tidak muncul waktu rata-rata run satu fungsi nya
// go test 58_UnitTest.go 58_UnitTest_benchmark_test.go -v -bench=.
package main

import "testing"

var kubus = Kubus{4}

func BenchmarkHitungLuas(b *testing.B) {
	for i := 0; i < b.N; i++ {
		kubus.Luas()
	}
}