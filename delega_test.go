package delega

import (
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
