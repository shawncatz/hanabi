package hanabi

import (
	"fmt"

	"github.com/hashicorp/consul/api"
)

type ConsulConfigurer struct {
	bloom  string
	client *api.Client
}

func newConsul(bloom string, URL string) (*ConsulConfigurer, error) {
	var err error

	c := &ConsulConfigurer{
		bloom: bloom,
	}

	cfg := api.DefaultConfig()
	cfg.Address = URL

	c.client, err = api.NewClient(cfg)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (c *ConsulConfigurer) getString(key string) (string, error) {
	val, err := c.get(key)
	if err != nil {
		return "", err
	}
	return string(val), nil
}

func (c *ConsulConfigurer) get(key string) ([]byte, error) {
	k := "hanabi/" + c.bloom + "/" + key
	kv := c.client.KV()

	pair, _, err := kv.Get(k, nil)
	if err != nil {
		return nil, err
	}
	if pair == nil {
		return nil, fmt.Errorf("configuration key not found: %s", k)
	}

	return pair.Value, nil
}
