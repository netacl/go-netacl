package netacl

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strconv"
)

type MinecraftProxies struct {
	Owned []MinecraftProxy `json:"owned"`

	// map keys are type 'username (email)'
	Managed map[string][]MinecraftProxy `json:"managed"`
}

type MinecraftProxy struct {
	Domain string `json:"domain"`

	Address string `json:"address"`
	Port    int16  `json:"port"`
}

func (c *APICLient) GetProxies() (*MinecraftProxies, error) {
	raw, err := c.Request("/minecraft-proxy", http.MethodGet, Application_json, nil)
	if err != nil {
		return nil, err
	}
	type payload struct {
		Owned map[string]struct {
			ProxyTo string `json:"proxy_to"`
		} `json:"owned"`
		Managed map[string]map[string]struct {
			ProxyTo string `json:"proxy_to"`
		} `json:"managed"`
	}
	pl := payload{}
	if err := json.Unmarshal(raw, &pl); err != nil {
		return nil, err
	}

	data := &MinecraftProxies{}
	for domain, proxy := range pl.Owned {
		ip, port, err := net.SplitHostPort(proxy.ProxyTo)
		if err != nil {
			return nil, err
		}
		i, _ := strconv.Atoi(port)
		data.Owned = append(data.Owned, MinecraftProxy{Domain: domain, Address: ip, Port: int16(i)})
	}
	for account, domains := range pl.Managed {
		if data.Managed[account] == nil {
			data.Managed[account] = []MinecraftProxy{}
		}
		for domain, proxy := range domains {
			ip, port, err := net.SplitHostPort(proxy.ProxyTo)
			if err != nil {
				return nil, err
			}
			i, _ := strconv.Atoi(port)
			data.Managed[account] = append(data.Owned, MinecraftProxy{Domain: domain, Address: ip, Port: int16(i)})
		}
	}
	return data, nil
}

func (c *APICLient) NewProxy(p *MinecraftProxy) error {
	if p == nil {
		return fmt.Errorf("proxy cannot be nil")
	}
	_, err := c.Request(fmt.Sprintf("/minecraft-proxy/%v", p.Domain), http.MethodPost, Application_json, map[string]string{"proxy_to": fmt.Sprintf("%v:%v", p.Address, p.Port)})
	return err
}

func (c *APICLient) DeleteProxy(name string) error {
	_, err := c.Request(fmt.Sprintf("/minecraft-proxy/%v", name), http.MethodDelete, Application_json, nil)
	return err
}
