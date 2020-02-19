package conf

type Config struct {
	KafkaConfig `ini:"kafka"`
	EtcdConfig  `ini:"etcd"`
}

type KafkaConfig struct {
	Address string `ini:"address"`
	Topic   string `ini:"topic"`
}

type EtcdConfig struct {
	Endpoints   string `ini:"endpoints"`
	DialTimeout int    `ini:"dialtimeout"`
}

//-- unused----
type TaillogConfig struct {
	Path string `ini:"path"`
}
