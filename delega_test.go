package delega

import (
    "testing"
)

func TestCreateMethodOnFreeProxyLists(t *testing.T) {
    provider, err := Create(FreeProxyLists)

    if err != nil {
        t.Error(err)
    }

    if provider.Name() != "Free Proxy Lists" {
        t.Error("Expected provider name is not equals.")
    }
}


func TestCreateMethodOnXiCiDaiLi(t *testing.T) {
    provider, err := Create(XiCiDaiLi)

    if err != nil {
        t.Error(err)
    }

    if provider.Name() != "Xi Ci Dai Li" {
        t.Error("Expected provider name is not equals.")
    }
}
