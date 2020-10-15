package common

import (
	"testing"
)

func TestExpBoundary(t *testing.T) {
	type expTestSet struct {
		exp      int64
		level    int
		residual int64
	}

	var samples []expTestSet
	for i := range expTable {
		if i == 0 {
			continue
		}
		samples = append(samples, expTestSet{
			expTable[i-1],
			i,
			0,
		})
	}

	for _, s := range samples {
		level, residual, _ := Exp2Level(s.exp)
		if level != s.level || residual != s.residual {
			t.Errorf("Exp2Level(%v) = %v, %v, _; want %v, %v, _",
				s.exp, level, residual, s.level, s.residual)
		}
	}
}
