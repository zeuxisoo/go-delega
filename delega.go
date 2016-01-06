package delega

import (
    "errors"

    "github.com/zeuxisoo/go-delega/contract"
    "github.com/zeuxisoo/go-delega/provider"
)

const (
    FreeProxyLists  = iota
    XiCiDaiLi
)

type Delega struct {

}

func Create(providerType int) (contract.Provider, error) {
    switch(providerType) {
        case FreeProxyLists:
            return new(provider.FreeProxyLists), nil
        case XiCiDaiLi:
            return new(provider.XiCiDaiLi), nil
        default:
            return nil, errors.New("Invalid provider type")
    }
}
