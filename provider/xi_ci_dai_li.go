package provider

import (
    "github.com/parnurzeal/gorequest"

    "github.com/zeuxisoo/go-delega/contract"

    "net/http"
)

type XiCiDaiLi struct {
}

func(this *XiCiDaiLi) Name() string {
    return "Xi Ci Dai Li"
}

func(this *XiCiDaiLi) Fetch() (*http.Response, error) {
    request := gorequest.New()

    resp, _, errs := request.
        Get("http://www.xicidaili.com/").
        Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/47.0.2526.106 Safari/537.36").
        End()

    if errs != nil {
        return nil, errs[0]
    }

    return resp, nil
}

func (this *XiCiDaiLi) Result(response *http.Response) ([]contract.ProxyList, error) {
    return nil, nil
}
