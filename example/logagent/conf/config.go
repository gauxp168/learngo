package conf

type AppConf struct {
	KafkaConf `ini:"kafka"`
	EtcdConf `ini:"etcd"`
}

type EtcdConf struct {
	Address string `ini:"address"`
	Key string `ini:"collect_log_key"`
	Timeout int `ini:"timeout"`
}

type KafkaConf struct {
	Address string `ini:"address"`
	ChanMaxSize int `ini:"chan_max_size"`
}
