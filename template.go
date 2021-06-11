package main

import (
	"bytes"
	"fmt"
	"log"
	"strings"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/yuansudong/gengo"
	"github.com/yuansudong/gengo/descriptor"
	plugin "github.com/yuansudong/gengo/plugin"
	"github.com/yuansudong/protoc-gen-mysql/pboptions"
	"google.golang.org/protobuf/proto"
)

//  DBName    string //  数据库名
// 	TableName string // 表名
// 	Engine    string // 引擎,默认为innodb
// 	AutoIncr  int64  // 自增开始的时间, 默认c从十万开始 100000
// 	Character string // 编码,默认为utf8mb4
// 	Collate   string // utf8mb4_general_ci
// 	RowFormat string // Dynamic
// 	Comment   string //  对该表的注释
//  DROP TABLE IF EXISTS {{ .TableName }};
const sqlTemplate = `
CREATE TABLE {{ .TableName }}  (
  {{ .Code }}
) ENGINE = {{ .Engine }} AUTO_INCREMENT = {{ .AutoIncr }} CHARACTER SET = {{ .Character }} COLLATE = {{ .Collate }} COMMENT = '{{ .Comment }}' ROW_FORMAT = {{ .RowFormat }};
`

// GoTemplate 用于生成相关的模板
func (g *Generator) GoTemplate(file *gengo.File) *plugin.CodeGeneratorResponse_File {
	// var err error
	// buf := bytes.NewBuffer(make([]byte, 0, 40960))
	rspFile := new(plugin.CodeGeneratorResponse_File)
	args := new(TemplateArgs)
	for _, msg := range file.Messages {
		if proto.HasExtension(msg.GetOptions(), pboptions.E_Message) {
			args.Tables = append(args.Tables, g.GoMessage(file, msg))
		}
	}
	g.CreateTable(args)
	return rspFile
}

// GoMessage 用于检测消息扩展,并得到相应的
func (g *Generator) GoMessage(file *gengo.File, msg *gengo.Message) (table *TemplateTable) {
	table = NewTamplateTable()
	optMsg := proto.GetExtension(msg.GetOptions(), pboptions.E_Message).(*pboptions.MysqlMessage)
	table.DBName = optMsg.Database
	table.Comment = optMsg.Comment
	table.TableName = msg.GetName()
	if optMsg.Character != "" {
		table.Character = optMsg.Character
	}
	if optMsg.Collate != "" {
		table.Collate = optMsg.Collate
	}
	if optMsg.Engine != "" {
		table.Engine = optMsg.Engine
	}
	//  检索字段
	for _, field := range msg.Fields {
		table.Lang = append(table.Lang, g.FieldProto2SQL(file, field))
	}
	if len(table.Lang) == 0 {
		return
	}
	//  增加索引
	for _, index := range optMsg.Indexs {
		// 判断索引类型
		switch index.IndexType {
		case pboptions.Type_T_PKI: //  索引类型为主键
			table.Lang = append(
				table.Lang,
				fmt.Sprintf(
					"PRIMARY KEY (%s) USING BTREE COMMENT '%s'",
					strings.Join(index.IndexField, ","),
					index.Comment))
		case pboptions.Type_T_FTI: //  全文检索索引
			table.Lang = append(
				table.Lang,
				fmt.Sprintf(
					"FULLTEXT INDEX %s_Fulltext_Index(%s) USING BTREE COMMENT '%s'",
					strings.Join(index.IndexField, "_"),
					strings.Join(index.IndexField, ","),
					index.Comment,
				),
			)
		case pboptions.Type_T_NII: //  普通索引
			table.Lang = append(
				table.Lang,
				fmt.Sprintf(
					"INDEX %s_Normal_Index(%s) USING BTREE COMMENT '%s'",
					strings.Join(index.IndexField, "_"),
					strings.Join(index.IndexField, ","),
					index.Comment,
				),
			)
		case pboptions.Type_T_UII: //  唯一索引
			table.Lang = append(
				table.Lang,
				fmt.Sprintf(
					"UNIQUE INDEX %s_Unique_Index(%s) USING BTREE COMMENT '%s'",
					strings.Join(index.IndexField, "_"),
					strings.Join(index.IndexField, ","),
					index.Comment,
				),
			)
		}
	}
	return table
}

// FieldProto2SQL 用于返回一个go的类型
func (g *Generator) FieldProto2SQL(rootFile *gengo.File, field *gengo.Field) string {
	typ := ""
	optField := new(pboptions.MysqlField)
	if proto.HasExtension(field.GetOptions(), pboptions.E_Field) {
		optField = proto.GetExtension(field.GetOptions(), pboptions.E_Field).(*pboptions.MysqlField)
	}
	switch *field.Type {
	case descriptor.FieldDescriptorProto_TYPE_DOUBLE:
		typ = " DECIMAL(10,16) NOT NULL"
		if optField.AutoIncr {
			typ = typ + " " + "AUTO_INCREMENT"
		} else {
			typ = typ + " " + "DEFAULT 0"
		}
	case descriptor.FieldDescriptorProto_TYPE_FLOAT:
		typ = " DECIMAL(10,8) NOT NULL"
		if optField.AutoIncr {
			typ = typ + " " + "AUTO_INCREMENT"
		} else {
			typ = typ + " " + "DEFAULT 0"
		}
	case descriptor.FieldDescriptorProto_TYPE_INT64:
		typ = " BIGINT(20) NOT NULL"
		if optField.AutoIncr {
			typ = typ + " " + "AUTO_INCREMENT"
		} else {
			typ = typ + " " + "DEFAULT 0"
		}
	case descriptor.FieldDescriptorProto_TYPE_UINT64:
		typ = " BIGINT(20) UNSIGNED NOT NULL"
		if optField.AutoIncr {
			typ = typ + " " + "AUTO_INCREMENT"
		} else {
			typ = typ + " " + "DEFAULT 0"
		}
	case descriptor.FieldDescriptorProto_TYPE_INT32:
		typ = " INT NOT NULL"
		if optField.AutoIncr {
			typ = typ + " " + "AUTO_INCREMENT"
		} else {
			typ = typ + " " + "DEFAULT 0"
		}
	case descriptor.FieldDescriptorProto_TYPE_UINT32:
		typ = " INT UNSIGNED NOT NULL"
		if optField.AutoIncr {
			typ = typ + " " + "AUTO_INCREMENT"
		} else {
			typ = typ + " " + "DEFAULT 0"
		}
	case descriptor.FieldDescriptorProto_TYPE_FIXED64:
		typ = " BIGINT(20) UNSIGNED NOT NULL"
		if optField.AutoIncr {
			typ = typ + " " + "AUTO_INCREMENT"
		} else {
			typ = typ + " " + "DEFAULT 0"
		}
	case descriptor.FieldDescriptorProto_TYPE_FIXED32:
		typ = " INT NOT NULL"
		if optField.AutoIncr {
			typ = typ + " " + "AUTO_INCREMENT"
		} else {
			typ = typ + " " + "DEFAULT 0"
		}
	case descriptor.FieldDescriptorProto_TYPE_BOOL:
		typ = " TINYINT(1) UNSIGNED NOT NULL DEFAULT 0"
	case descriptor.FieldDescriptorProto_TYPE_STRING:
		typ = " VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL"
		if optField.AutoIncr {
			typ = typ + " " + "AUTO_INCREMENT "
		} else {
			typ = typ + " " + "DEFAULT '' "
		}
	case descriptor.FieldDescriptorProto_TYPE_BYTES:
		typ = " BLOB NOT NULL "
	case descriptor.FieldDescriptorProto_TYPE_SFIXED32:
		typ = " INT NOT NULL "
		if optField.AutoIncr {
			typ = typ + " " + "AUTO_INCREMENT"
		} else {
			typ = typ + " " + "DEFAULT 0"
		}
	case descriptor.FieldDescriptorProto_TYPE_SFIXED64:
		typ = " BIGINT(20) NOT NULL"
		if optField.AutoIncr {
			typ = typ + " " + "AUTO_INCREMENT"
		} else {
			typ = typ + " " + "DEFAULT 0"
		}
	case descriptor.FieldDescriptorProto_TYPE_SINT32:
		typ = " INT NOT NULL"
		if optField.AutoIncr {
			typ = typ + " " + "AUTO_INCREMENT"
		} else {
			typ = typ + " " + "DEFAULT 0"
		}
	case descriptor.FieldDescriptorProto_TYPE_SINT64:
		typ = " BIGINT(20) NOT NULL"
		if optField.AutoIncr {
			typ = typ + " " + "AUTO_INCREMENT"
		} else {
			typ = typ + " " + "DEFAULT 0"
		}
	case descriptor.FieldDescriptorProto_TYPE_ENUM:
		typ = "INT NOT NULL DEFAULT 0"
	default:
		typ = "TEXT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci "
	}
	if field.GetLabel() == descriptor.FieldDescriptorProto_LABEL_REPEATED {
		typ = "TEXT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci "
	}
	return fmt.Sprintf("%s %s COMMENT '%s'", field.GetName(), typ, optField.Comment)
}

//  CreateTable 用于执行一个sql语句
func (g *Generator) CreateTable(args *TemplateArgs) {
	for _, table := range args.Tables {
		if len(table.Lang) == 0 {
			continue
		}
		table.Comment += " yuansudong,服务于掌沃无限时制作 "
		b := bytes.NewBuffer(make([]byte, 0, 40960))
		t := template.New("sql_template")
		t1, err := t.Parse(sqlTemplate)
		if err != nil {
			log.Fatalln(err.Error())
		}

		table.Code = strings.Join(table.Lang, ",\r\n")
		if err := t1.Execute(b, table); err != nil {
			log.Fatalln(err.Error())
		}
		db, err := xorm.NewEngine("mysql", fmt.Sprintf(
			"%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			g.User,
			g.Password,
			g.Host,
			g.Port,
			table.DBName,
		))
		if err != nil {
			log.Fatalln(err.Error())
		}
		if _, err := db.Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s", table.TableName)); err != nil {
			log.Fatalln(err.Error())
		}
		if _, err := db.Exec(b.String()); err != nil {
			log.Fatalln(err.Error(), b.String())
		}
		db.Close()
	}
}
