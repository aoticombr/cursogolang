package api

func Ping() (int, []byte, error) {
	return 200, []byte("pong"), nil
}
