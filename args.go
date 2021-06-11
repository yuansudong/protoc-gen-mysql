package main

/*
DROP TABLE IF EXISTS `User`;
CREATE TABLE `User`  (
  `ID` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键自增,无符号',
  `Name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '名称',
  `DateTime` datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '日期',
  `Timestamp` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '时间戳',
  `Status` enum('one','two','three') CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT 'one' COMMENT '状态',
  `Version` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '版本',
  `CurrDate` date NOT NULL DEFAULT '2020-01-01' COMMENT '当前日期',
  `IC` bigint(20) NOT NULL DEFAULT 0 COMMENT '有符号',
  `Text` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '文本',
  `Time` time(0) NOT NULL COMMENT '时间',
  `Lat` float UNSIGNED NOT NULL DEFAULT 0 COMMENT '经度',
  `Lgt` float NOT NULL DEFAULT 0 COMMENT '纬度',
  PRIMARY KEY (`ID`) USING BTREE,
  UNIQUE INDEX `Name`(`Name`) USING BTREE COMMENT '唯一索引',
  UNIQUE INDEX `DateTime_Timestamp`(`DateTime`, `Timestamp`) USING BTREE COMMENT '符合唯一索引',
  FULLTEXT INDEX `FULL_TEXT`(`Name`) COMMENT '111',
  INDEX `normal_status_index`(`Status`) USING BTREE COMMENT '111',
  INDEX `IC_Status_Index`(`IC`, `Status`) USING BTREE COMMENT '符合索引'
) ENGINE = InnoDB AUTO_INCREMENT = 13 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用于存储用户数据' ROW_FORMAT = Dynamic;

*/

// TemplateArgs 用于描述一个模板参数
type TemplateArgs struct {
	Tables []*TemplateTable
}

// TemplateTable 用于描述一个模板表格
type TemplateTable struct {
	DBName    string   //  数据库名
	TableName string   // 表名
	Engine    string   // 引擎,默认为innodb
	AutoIncr  int64    // 自增开始的时间, 默认c从十万开始 100000
	Character string   // 编码,默认为utf8mb4
	Collate   string   // utf8mb4_general_ci
	RowFormat string   // Dynamic
	Comment   string   //  对该表的注释
	Lang      []string // 用于存储一段段的sql语句
	Code      string   //  Code 用于代码生成
}

// NewTamplateTable 用于新建立一个Table 结构
func NewTamplateTable() *TemplateTable {
	inst := new(TemplateTable)
	inst.AutoIncr = 100000
	inst.Engine = "innodb"
	inst.Character = "utf8mb4"
	inst.Collate = "utf8mb4_general_ci"
	inst.RowFormat = "Dynamic"
	inst.Comment = "这个人很恶心,什么都没写"
	return inst
}

// TemplateField 用于描述一个模板字段
type TemplateField struct {
}
