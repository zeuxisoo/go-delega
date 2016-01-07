package provider

import (
    "testing"
)

func TestFreeProxyListsNameMethod(t *testing.T) {
    proxyProvider := &FreeProxyLists{}

    if proxyProvider.Name() != "Free Proxy Lists" {
        t.Errorf("Expected name should be Free Proxy Lists, current is %s", proxyProvider.Name())
    }
}

func TestFreeProxyListsFetchMethod(t *testing.T) {
    proxyProvider := &FreeProxyLists{}
    resp, err     := proxyProvider.Fetch()

    if err != nil {
        t.Error(err)
    }

    if resp.StatusCode != 200 {
        t.Errorf("Expected StatusCode=200, current StatusCode=%v", resp.StatusCode)
    }
}

func TestFreeProxyListsResultMethod(t *testing.T) {
    proxyProvider := &FreeProxyLists{}
    resp, _       := proxyProvider.Fetch()

    proxyList, err := proxyProvider.Result(resp)

    if err != nil {
        t.Error(err)
    }

    if len(proxyList) != 50 {
        t.Error("Expected proxy list should be equals 50 rows")
    }

    proxy := proxyList[0]

    if proxy.Ip == "" || proxy.Port == "" || proxy.Protocol == "" {
        t.Error("The first proxy should not be empty in IP, Port and Protocol fields")
    }
}
