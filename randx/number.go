package randx

import (
	"math"
	"math/rand"
	"strconv"
)

// 随机整数
func Int() int {
	return IntRange(0, math.MaxInt)
}

// 随机整数
func Int64() int64 {
	return Int64Range(0, math.MaxInt64)
}

// 随机浮点数
func Float64() float64 {
	return Float64Range(math.SmallestNonzeroFloat64, math.MaxFloat64, 6)
}

// 随机整数
func IntRange(min, max int) int {
	if min == max {
		return min
	}
	rand.Seed(Seed())
	return rand.Intn(1+max-min) + min
}

// 随机整数
func Int64Range(min, max int64) int64 {
	if min == max {
		return min
	}
	rand.Seed(Seed())
	return rand.Int63n(1+max-min) + min
}

// 随机浮点数
func Float64Range(min, max float64, prec int) float64 {
	if min == max {
		return min
	}
	rand.Seed(Seed())
	float := rand.Float64()*(max-min) + min
	float, _ = strconv.ParseFloat(strconv.FormatFloat(float, 'f', prec, 64), 64)
	return float
}
