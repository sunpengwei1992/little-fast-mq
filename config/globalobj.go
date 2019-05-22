package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"little-fast-mq/siface"
	"little-fast-mq/utils"
)

/*
   存储一切有关Zinx框架的全局参数，供其他模块使用
   一些参数也可以通过 用户根据 zinx.json来配置
*/
type GlobalObj struct {
	TcpServer        siface.IServer //当前Zinx的全局Server对象
	Host             string         //当前服务器主机IP
	TcpPort          int            //当前服务器主机监听端口号
	Name             string         //当前服务器名称
	Version          string         //当前Zinx版本号
	MaxPacketSize    uint32         //都需数据包的最大值
	MaxConn          int            //当前服务器主机允许的最大链接个数
	WorkerPoolSize   uint32
	MaxWorkerTaskLen uint32
	ConfFilePath     string
	MaxMsgChanLen    uint32
}

//读取用户的配置文件
func (g *GlobalObj) Reload() {
	data, err := ioutil.ReadFile("../conf/speed-mq.json")
	if err != nil {
		panic(err)
	}
	//将json数据解析到struct中
	fmt.Printf("json :%s\n", data)
	err = json.Unmarshal(data, &GlobalObject)
	if err != nil {
		panic(err)
	}
}

/*
   提供init方法，默认加载
*/
func init() {
	//初始化GlobalObject变量，设置一些默认值
	GlobalObject = &GlobalObj{
		Name:             "little-fast-mq-ServerApp",
		Version:          "V0.4",
		TcpPort:          7777,
		Host:             utils.GetLocalIp(),
		MaxConn:          12000,
		MaxPacketSize:    4096,
		WorkerPoolSize:   10,
		ConfFilePath:     "conf/speed-mq.json",
		MaxWorkerTaskLen: 1024,
		MaxMsgChanLen:    1024,
	}

	//从配置文件中加载一些用户配置的参数
	GlobalObject.Reload()
}

/*
   定义一个全局的对象
*/
var GlobalObject *GlobalObj
