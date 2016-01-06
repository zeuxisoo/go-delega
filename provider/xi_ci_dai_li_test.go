package provider

import (
    "testing"
)

func TestNameMethod(t *testing.T) {
    proxyProvider := &XiCiDaiLi{}

    if proxyProvider.Name() != "Xi Ci Dai Li" {
        t.Errorf("Expected name should be Free Proxy Lists, current is %s", proxyProvider.Name())
    }
}

func TestFetchMethod(t *testing.T) {
    proxyProvider := &XiCiDaiLi{}
    resp, err     := proxyProvider.Fetch()

    if err != nil {
        t.Error(err)
    }

    if resp.StatusCode != 200 {
        t.Errorf("Expected StatusCode=200, current StatusCode=%v", resp.StatusCode)
    }
}

func TestResultMethod(t *testing.T) {
    proxyProvider := &XiCiDaiLi{}
    resp, _       := proxyProvider.Fetch()

    proxyList, err := proxyProvider.Result(resp)

    if err != nil {
        t.Error(err)
    }

    if len(proxyList) != 100 {
        t.Error("Expected proxy list should be equals 100 rows")
    }

    proxy := proxyList[0]

    if proxy.Ip == "" || proxy.Port == "" || proxy.Protocol == "" {
        t.Error("The first proxy should not be empty in IP, Port and Protocol fields")
    }
}
