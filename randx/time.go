package randx

import "time"

// 一天内随机时间
func Time() time.Time {
	now := time.Now()
	return TimeRange(now.Add(time.Hour*-24), now)
}

// 随机时间
func TimeRange(min, max time.Time) time.Time {
	if max.After(min) {
		return time.Unix(0, Int64Range(min.UnixNano(), max.UnixNano()))
	}
	return min
}

// 生成日期
func Date() string {
	now := time.Now().Unix()
	diff := Int64Range(1, now)
	return time.Unix(diff, 0).Format("2006-01-02")
}
