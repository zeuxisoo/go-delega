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
