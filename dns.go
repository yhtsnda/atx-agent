package main

import (
	"context"
	"log"
	"net"
	"time"
)

type dnsSmartClient struct {
	dialer *net.Dialer
}

func newDnsSmartClient() *dnsSmartClient {
	return &dnsSmartClient{
		dialer: &net.Dialer{
			Timeout:   3 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		},
	}
}
func (c *dnsSmartClient) Dial(ctx context.Context, network, address string) (conn net.Conn, err error) {
	dns1 := getProperty("net.dns1")
	if dns1 == "" {
		// 国内DNS列表: https://www.zhihu.com/question/32229915
		dns1 = "114.114.114.114"
	}
	log.Println("dns resolve", dns1)
	return c.dialer.DialContext(ctx, "udp", dns1+":53")
}

func init() {
	net.DefaultResolver = &net.Resolver{
		PreferGo: true,
		Dial:     newDnsSmartClient().Dial,
	}
}
