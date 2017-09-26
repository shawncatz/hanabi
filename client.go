package hanabi

type Client struct {
	Bloom  string
	Consul string
	Config *Config
	Rabbit string
}

func NewClient(bloom, consul string) (*Client, error) {
	c := &Client{
		Bloom:  bloom,
		Consul: consul,
	}

	return c, nil
}

func (c *Client) Send(event string, data []byte) error {
	return nil
}