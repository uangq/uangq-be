package configs

const (
	address = "localhost:6379"
)

type RedisConfiguration struct {
	Address string
}

var RedisConfig = RedisConfiguration{
	Address: address,
}
