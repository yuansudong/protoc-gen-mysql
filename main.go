package main

import (
	"os"

	"github.com/golang/glog"
	"github.com/yuansudong/gengo"
)

// PluginName 用于描述一个插件名称
const PluginName = "protoc-gen-mysql:"

func main() {
	startFromStdin()
}

// startFromStdin 用于开始从标准输入
func startFromStdin() {
	in := os.Stdin
	req, err := gengo.GetRequest(in)
	if err != nil {
		gengo.WriteError(err)
		return
	}
	reg := gengo.NewRegistry()
	if err = reg.Load(req); err != nil {
		gengo.WriteError(err)
		return
	}
	var targets []*gengo.File
	for _, target := range req.FileToGenerate {
		f, err := reg.LookupFile(target)
		if err != nil {
			glog.Fatal(err)
		}
		targets = append(targets, f)
	}
	g := New(reg)

	g.GetReqParam(req)
	out, err := g.Generate(targets)
	if err != nil {
		gengo.WriteError(err)
		return
	}
	gengo.WriteFiles(out)
}
