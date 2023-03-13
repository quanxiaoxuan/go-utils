package sqlx

const (
	String     = "string"
	Varchar    = "varchar"
	Varchar255 = "varchar(255)"
	Text       = "text"
	Int        = "int"
	Smallint   = "smallint"
	Bigint     = "bigint"
	Int2       = "int2"
	Int4       = "int4"
	Int8       = "int8"
	Int64      = "int64"
	Float4     = "float4"
	Numeric    = "numeric"
	Numeric6   = "numeric(10,6)"
	Numeric2   = "numeric(10,2)"
	Time       = "time.Time"
	Timestamp  = "timestamp"
	Date       = "date"
	Bool       = "bool"
)

// java类常用类型
const (
	JavaString     = "String"
	JavaInteger    = "Integer"
	JavaLong       = "Long"
	JavaDate       = "Date"
	JavaBigDecimal = "BigDecimal"
	JavaBoolean    = "Boolean"
)

var Pg2GoTypeMap = initPg2GoTypeMap()
var Pg2GormTypeMap = initPg2GormTypeMap()
var Pg2JavaTypeMap = initPg2JavaTypeMap()

// PG-Go类型映射
func initPg2GoTypeMap() map[string]string {
	var typeMap = make(map[string]string)
	typeMap[Varchar] = String
	typeMap[Text] = String
	typeMap[Int2] = Int
	typeMap[Int4] = Int
	typeMap[Int8] = Int64
	typeMap[Timestamp] = Time
	typeMap[Date] = Time
	typeMap[Float4] = Float4
	typeMap[Numeric] = Float4
	typeMap[Bool] = Bool
	return typeMap
}

// PG-Gorm类型映射
func initPg2GormTypeMap() map[string]string {
	var typeMap = make(map[string]string)
	typeMap[Varchar] = Varchar255
	typeMap[Text] = Text
	typeMap[Int2] = Smallint
	typeMap[Int4] = Int
	typeMap[Int8] = Bigint
	typeMap[Timestamp] = Timestamp
	typeMap[Date] = Date
	typeMap[Float4] = Numeric6
	typeMap[Numeric] = Numeric2
	typeMap[Bool] = Bool
	return typeMap
}

// PG-java类型映射
func initPg2JavaTypeMap() map[string]string {
	var typeMap = make(map[string]string)
	typeMap[Varchar] = JavaString
	typeMap[Text] = JavaString
	typeMap[Int2] = JavaInteger
	typeMap[Int4] = JavaLong
	typeMap[Int8] = JavaLong
	typeMap[Timestamp] = JavaDate
	typeMap[Date] = JavaDate
	typeMap[Float4] = JavaBigDecimal
	typeMap[Numeric] = JavaBigDecimal
	typeMap[Bool] = JavaBoolean
	return typeMap
}
