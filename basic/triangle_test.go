package main

import "testing"

/**
测试的时候，必须选中basic.go和test文件
*/
func TestTriangle(t *testing.T) {

	//定义测试数据
	tests := []struct{ a, b, c int }{ //开始这个括号必须和}在同一行，否则报错
		{3, 4, 5},
		{5, 12, 13},
		{3000, 4000, 5000},
	}

	for _, tt := range tests {
		if actual := calcTriangle(tt.a, tt.b); actual != tt.c {
			t.Errorf("calcTriangle(%d,%d),got %d ,expected %d", tt.a, tt.b, actual, tt.c)
		}
	}
}

//性能测试:BenchmarkTest-4   	2000000000	         0.70 ns/op
/**
运行了 2000000000 时间0.70 ns/op每个操作
*/
func BenchmarkTest(b *testing.B) {
	a, d, c := 3000, 4000, 5000
	for i := 0; i < b.N; i++ {
		if actual := calcTriangle(a, d); actual != c {
			b.Errorf("calcTriangle(%d,%d),got %d ,expected %d", a, d, actual, c)
		}
	}

}
