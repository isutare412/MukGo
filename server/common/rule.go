package common

// Exp2Level function divides total exp into level. levExp indicates
// current level's exp points for level up. curExp indicates current exp points
// in current level. ratio indicates a ratio of curExp exp compared to levExp.
//
// Suppose you need 100 exp until level 2 and 200 exp until level 3. If you
// have 160 exp, your level is 2, levExp is 100, curExp is 60, ratio is 0.6.
func Exp2Level(exp int64) (level int32, levExp, curExp int64, ratio float64) {
	for i, needUntil := range expTable {
		if i == 0 {
			continue
		}
		needBefore := expTable[i-1]

		if exp < needUntil {
			level = int32(i)
			levExp = needUntil - needBefore
			curExp = exp - needBefore
			ratio = float64(curExp) / float64(levExp)
			return
		}
	}

	// has exp above top level
	level = int32(len(expTable))
	levExp = 0
	curExp = 0
	ratio = 0
	return
}

// Level2Sight returns a radius of user's sight derived from user's level.
// Returned radius is METER.
func Level2Sight(level int32) float64 {
	return 10.0 + float64(level)*10.0
}

// Exp2Sight returns a radius of user's sight derived from user's exp point.
// Returned radius is METER.
func Exp2Sight(exp int64) float64 {
	level, _, _, _ := Exp2Level(exp)
	return Level2Sight(level)
}

// ReviewExp calculates exp given for each review.
func ReviewExp() int64 {
	return 50
}
