package hanabi

type Config struct {
	Bloom  string // Name of the bloom to which we are connected
	Consul string // URL to connect to Consul

	Rabbit string // URL to connect to RabbitMQ

	cfg *ConsulConfigurer
	txp *RabbitTransporter
}

func (c *Config) load() (err error) {
	c.cfg, err = newConsul(c.Bloom, c.Consul)
	if err != nil {
		return err
	}

	c.Rabbit, err = c.cfg.getString("rabbit")
	if err != nil {
		return err
	}

	c.txp, err = newRabbit(c.Rabbit)
	if err != nil {
		return err
	}

	return nil
}
