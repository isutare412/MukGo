package common

import (
	"testing"
)

func TestExpBoundary(t *testing.T) {
	type expTestSet struct {
		exp      int64
		level    int32
		residual int64
	}

	var samples []expTestSet
	for i := range expTable {
		if i == 0 {
			continue
		}
		samples = append(samples, expTestSet{
			expTable[i-1], int32(i), 0,
		})
	}

	for _, s := range samples {
		level, _, curExp, _ := Exp2Level(s.exp)
		if level != s.level || curExp != s.residual {
			t.Errorf("Exp2Level(%v) = %v, %v, _; want %v, %v, _",
				s.exp, level, curExp, s.level, s.residual)
		}
	}
}

func TestLargeExp(t *testing.T) {
	maxInt := int64(^uint64(0) >> 1)
	level, levExp, curExp, ratio := Exp2Level(maxInt)
	if level != int32(len(expTable)) || levExp != 0 || curExp != 0 || ratio != 0 {
		t.Errorf("Exp2Level(%v) = %v, %v, %v, %v; want %v, %v, %v, %v",
			maxInt, level, levExp, curExp, ratio, len(expTable), 0, 0, 0)
	}
}
