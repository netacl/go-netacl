package netacl

import (
	"log"
	"testing"
)

func TestGetMinecraftProxies(t *testing.T) {
	setTestKey(t)

	proxies, err := c.GetProxies()
	if err != nil {
		t.Error(err)
	}
	log.Println(proxies)
}

func TestCreateProxy(t *testing.T) {
	setTestKey(t)

	err := c.NewProxy(&MinecraftProxy{Domain: "", Address: "", Port: int16(25565)})
	if err != nil {
		t.Error(err)
	}
}
