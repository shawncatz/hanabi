package hanabi_test

import (
	"testing"

	"fmt"
	"github.com/shawncatz/hanabi"
)

func TestNewClient(t *testing.T) {
	client, err := hanabi.NewClient("local", "127.0.0.1:8500")
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("%#v\n", client)
	fmt.Printf("%#v\n", client.Config)
}
