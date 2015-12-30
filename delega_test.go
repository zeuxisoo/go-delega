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

func TestFreeProxyListParseMethod(t *testing.T) {
    provider, _ := Create(FreeProxyLists)
    resp, _ := provider.Fetch()

    proxyList, err := provider.Result(resp)

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
