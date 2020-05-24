package conf

type AppConfig struct {
	KafkaConf `ini:"kafka"`
	TailConf `ini:"tailLog"`
}
type KafkaConf struct {
	Address string `ini:"address"`
	Topic string `ini:"topic"`
}

type TailConf struct {
	FileName string `ini:"filename"`
}
