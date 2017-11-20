package hanabi

type Transporter interface {
	log(level int, message string) error
	send(key string, value string) error
}
