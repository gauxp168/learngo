package main

import (
	"fmt"
	"gopkg.in/ini.v1"
	"simqo.com/mygospace/learngo/example/logagent/conf"
	"simqo.com/mygospace/learngo/example/logagent/etcd"
	"simqo.com/mygospace/learngo/example/logagent/kafka"
	"simqo.com/mygospace/learngo/example/logagent/taillog"
	"simqo.com/mygospace/learngo/example/logagent/utils"
	"sync"
	"time"
)

// logagent 入口文件

var (
	cfg = new(conf.AppConf)
)

func main() {
	// 1. 加载配置文件
	err := ini.MapTo(cfg, "./conf/config.ini")
	if err != nil {
		fmt.Printf("load ini failed, error:%v\n", err)
		return
	}
	// 2. 初始化kafka连接
	err = kafka.Init([]string{cfg.KafkaConf.Address}, cfg.KafkaConf.ChanMaxSize)
	if err != nil {
		fmt.Printf("init kafka failed, error:%v\n", err)
		return
	}
	fmt.Println("init kafka successs")
	// 3. 初始化etcd
	err = etcd.Init(cfg.EtcdConf.Address, time.Duration(cfg.EtcdConf.Timeout)*time.Second)
	if err != nil {
		fmt.Printf("init etcd failed, error:%v\n", err)
		return
	}
	fmt.Println("init etcd success")
	// 为了实现每个logagent都拉去自己独有的配置，所以要以自己的IP地址作为区分
	ip, err := utils.GetOutboundIP()
	if err != nil {
		panic(err)
	}
	etcdConfKey := fmt.Sprintf(cfg.EtcdConf.Key, ip)
	logEntryConf, err := etcd.GetConf(etcdConfKey)
	if err != nil {
		fmt.Printf("etcd get conf failed, error:%v\n", err)
		return
	}
	// 拍一个哨兵去监视收集箱的变化
	for key, value := range logEntryConf {
		fmt.Printf("index:%v value:%v\n", key, value)
	}
	// 3. 收集日志发往 kafka
	taillog.Init(logEntryConf)
	newConfChan := taillog.NewConfChan()
	var wg sync.WaitGroup
	wg.Add(1)
	go etcd.WatchConf(etcdConfKey, newConfChan)
	wg.Wait()
}
