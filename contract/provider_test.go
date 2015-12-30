package contract

import (
    "testing"
)

func TestProxyList(t *testing.T) {
    proxyList := ProxyList{
        Ip: "127.0.0.1",
        Port: "80",
        Protocol: "http",
    }

    if proxyList.Ip != "127.0.0.1" {
        t.Errorf("In proxy list, The ip should be 127.0.0.1, current is %s", proxyList.Ip)
    }

    if proxyList.Port != "80" {
        t.Errorf("In proxy list, The port should be 80, current is %s", proxyList.Port)
    }

    if proxyList.Protocol != "http" {
        t.Errorf("In proxy list, The port should be http, current is %s", proxyList.Protocol)
    }
}
