package configs

const (
	serverHost    = "localhost"
	serverPort    = "8000"
	serverAddress = serverHost + ":" + serverPort
)

type ServerConfig struct {
	Host    string
	Port    string
	Address string
}

type TcpPoolConfiguration struct {
	Size int
}

var (
	HttpServerConfig = ServerConfig{
		Host:    serverHost,
		Port:    serverPort,
		Address: serverAddress,
	}
	SessionJwtSecret = []byte("eXSM9gFS3UzGe91c2lSYd5W65nfj3NG/Kyv5QwtKytFoayiwjvuDOlP6UrylUdg8jwZ1N0GYx/J+6x5p+Kb3924hsanlfhnw7dGhkaIk1UV6sKUTPYY/nO5yZkJwcrNOTiHdNpNI9fOjOAhbh0X7BBujx+jOWW8E2XolDShZJeL+lnq9V2CaOpJo1K4ZxAUnQnQWPwYwSUE9W3WDtEFahOeY5uclFhgx3Cy2EUNWkHW6QV37/jYM7ypQj124rhsPFYuk7frJRHuk2s4vWUGG93T/nVN2fvHCQDUufYcy6UrKDRj7ZVGvlcufxXaxkmBnW20XocRFtFKRTDMEieR6/g==")
)
