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
