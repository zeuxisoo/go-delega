package provider

import (
    "github.com/zeuxisoo/go-delega/contract"

    "net/http"
)

type XiCiDaiLi struct {
}

func(this *XiCiDaiLi) Name() string {
    return "Xi Ci Dai Li"
}

func(this *XiCiDaiLi) Fetch() (*http.Response, error) {
    return nil, nil
}

func (this *XiCiDaiLi) Result(response *http.Response) ([]contract.ProxyList, error) {
    return nil, nil
}
