package common

// Exp2Level function divides total exp into level and residual exp. ratio
// indicates a ratio of the residual exp compared to the exp needed to level up.
//
// Suppose you need 100 exp until level 2 and 200 exp until level 3. If you
// have 160 exp, your level is 2, residual is 60, ratio is 0.6.
func Exp2Level(exp int64) (level int, residual int64, ratio float64) {
	level = len(expTable)
	for i, needUntil := range expTable {
		if i == 0 {
			continue
		}
		needBefore := expTable[i-1]

		if exp < needUntil {
			level = i
			residual = exp - needBefore

			needCur := needUntil - needBefore
			ratio = float64(residual) / float64(needCur)
			break
		}
	}
	return
}

// Level2Sight returns a radius of user's sight derived from user's level.
// Returned radius is METER.
func Level2Sight(level int) float64 {
	return 10.0 + float64(level)*10.0
}

// Exp2Sight returns a radius of user's sight derived from user's exp point.
// Returned radius is METER.
func Exp2Sight(exp int64) float64 {
	level, _, _ := Exp2Level(exp)
	return Level2Sight(level)
}
