package hanabi

type Client struct {
	Config *Config
}

func NewClient(bloom, consul string) (*Client, error) {
	c := &Config{
		Bloom:  bloom,
		Consul: consul,
	}

	err := c.load()
	if err != nil {
		return nil, err
	}

	x := &Client{
		Config: c,
	}

	return x, nil
}

func (c *Client) Send(event string, data []byte) error {
	return nil
}
