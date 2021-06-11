package main

import (
	"log"
	"strings"

	"github.com/yuansudong/gengo"
	plugin "github.com/yuansudong/gengo/plugin"
)

// Generator 生成结构
type Generator struct {
	reg *gengo.Registry
	// Host mysql服务地址
	Host string
	// Port mysql服务端口
	Port string
	// User mysql用户
	User string
	// Password mysql密码
	Password string
}

// Method 用于描述一个服务名称
type Method struct {
	Name string
	TReq string
	TRsp string
}

// Service 用于描述一个服务
type Service struct {
	Name    string
	Methods []Method
}

// Args 用于描述一个参数
type Args struct {
	DateTime    string
	IsHave      bool
	PackageName string
	Services    []Service
	Imports     map[string]bool
}

// GetReqParam 用于获取请求参数
func (g *Generator) GetReqParam(req *plugin.CodeGeneratorRequest) {
	//user=hfdy;password=hfdy2020;host=mysql-models.hfdy.com;port=3307
	param := strings.Join(strings.Fields(req.GetParameter()), "")
	paramArr := strings.Split(param, ";")
	for _, args := range paramArr {
		argsArr := strings.Split(args, "=")
		if len(argsArr) != 2 {
			log.Fatalln("args is invalid synatx ", args)
		}
		key, val := argsArr[0], argsArr[1]
		switch key {
		case "user":
			g.User = val
		case "password":
			g.Password = val
		case "host":
			g.Host = val
		case "port":
			g.Port = val
		}
	}
	if g.User == "" {
		log.Fatalln("user args is invalid")
	}
	if g.Port == "" {
		log.Fatalln("port args is invalid")
	}
	if g.Host == "" {
		log.Fatalln("host args is invalid")
	}
	if g.Password == "" {
		log.Fatalln("password args is invalid")
	}
}

// New returns a new generator which generates grpc gateway files.
func New(reg *gengo.Registry) *Generator {
	return &Generator{
		reg: reg,
	}
}

// Generate 用于执行生成操作
func (g *Generator) Generate(targets []*gengo.File) ([]*plugin.CodeGeneratorResponse_File, error) {
	var files []*plugin.CodeGeneratorResponse_File
	for _, file := range targets {
		g.GoTemplate(file)
	}
	return files, nil
}
