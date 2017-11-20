package hanabi

type Configurer interface {
	getString(key string) (string, error)
	get(key string) ([]byte, error)
}
