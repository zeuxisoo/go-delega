package delega

import (
    "errors"

    "github.com/zeuxisoo/go-delega/contract"
    "github.com/zeuxisoo/go-delega/provider"
)

const (
    FreeProxyLists  = iota
    OtherProxyLists
)

type Delega struct {

}

func Create(providerType int) (contract.Provider, error) {
    switch(providerType) {
        case FreeProxyLists:
            return new(provider.FreeProxyLists), nil
        default:
            return nil, errors.New("Invalid provider type")
    }
}
