package conf

type Config struct {
	KafkaConfig `ini:"kafka"`
	EtcdConfig  `ini:"etcd"`
}

type KafkaConfig struct {
	Address     string `ini:"address"`
	ChanMaxSize int    `ini:"chan_max_size"`
}

type EtcdConfig struct {
	Endpoints   string `ini:"endpoints"`
	DialTimeout int    `ini:"dialtimeout"`
	LogAgent    string `ini:"collect_log_key"`
}

//-- unused----
type TaillogConfig struct {
	Path string `ini:"path"`
}
