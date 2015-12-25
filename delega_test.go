package delega

import (
    "fmt"
    "testing"
)

func TestCreateMethod(t *testing.T) {
    provider, err := Create(FreeProxyLists)

    if err != nil {
        t.Error(err)
    }

    if provider.Name() != "Free Proxy Lists" {
        t.Error("Expected provider name is not equals.")
    }
}

func TestFreeProxyListFetchMethod(t *testing.T) {
    provider, _ := Create(FreeProxyLists)

    resp, err := provider.Fetch()

    if err != nil {
        t.Error(err)
    }

    if resp.StatusCode != 200 {
        t.Errorf("Expected StatusCode=200, current StatusCode=%v", resp.StatusCode)
    }
}
