package randx

import "hash/crc32"

const (
	LetterLower    = "abcdefghijklmnopqrstuvwxyz"
	LetterUpper    = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	CharLower      = "abcdefghijklmnopqrstuvwxyz1234567890"
	CharAll        = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	NUMBER         = "1234567890"
	SPECIAL        = "!@#$%&*+-=?"
	ProvinceSimple = "京,津,冀,晋,蒙,辽,吉,黑,沪,苏,浙,皖,闽,赣,鲁,豫,鄂,湘,粤,桂,琼,渝,川,贵,云,藏,陕,甘,青,宁,新,台,港,澳"
	ProvinceName   = "北京市,天津市,河北省,山西省,内蒙古自治区,辽宁省,吉林省,黑龙江省,上海市,江苏省,浙江省,安徽省,福建省,江西省,山东省,河南省,湖北省,湖南省,广东省,广西壮族自治区,海南省,重庆市,四川省,贵州省,云南省,西藏自治区,陕西省,甘肃省,青海省,宁夏回族自治区,新疆维吾尔自治区,台湾省,香港特别行政区,澳门特别行政区"
	HubeiCityName  = "武汉,黄石,十堰,宜昌,襄阳,鄂州,荆门,孝感,荆州,黄冈,咸宁,随州,恩施,仙桃,潜江,天门,神农架"
	HubeiPostcode  = "420100,420200,420300,420500,420600,420700,420800,420900,421000,421100,421200,421300,422800,429004,429005,429006,429021"
	FamilyNameCn   = "赵,钱,孙,李,周,吴,黄,高,郑,王,冯,陈,蒋,沈,韩,杨,朱,秦,许,何,吕,张,孔,曹"
	NumberCn       = "一,二,三,四,五,六,七,八,九,十,百,千,万,亿"
	Animal         = "鼠,牛,虎,兔,龙,蛇,马,羊,猴,鸡,狗,猪"
	PhonePrefix    = "358"
)

// 随机种子
func Seed() int64 {
	return int64(crc32.ChecksumIEEE([]byte(UUID())))
}

// 随机bool
func Bool() bool {
	return IntRange(0, 1) == 0
}
