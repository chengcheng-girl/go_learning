package benchmark

import (
	"github.com/wonderivan/logger"
	"testing"
)

//性能测试
//go test -bench=. -benchmem
func BenchmarkString(t *testing.B) {
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		logger.Info("ffff")
	}
	t.StopTimer()
}
