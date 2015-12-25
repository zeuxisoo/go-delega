package provider

import (
    "github.com/parnurzeal/gorequest"

    "net/http"
)

type FreeProxyLists struct {
}

func (this *FreeProxyLists) Name() string {
    return "Free Proxy Lists"
}

func (this *FreeProxyLists) Fetch() (*http.Response, error) {
    request := gorequest.New()

    resp, _, errs := request.
        Get("http://www.freeproxylists.net/").
        Set("Host", "www.freeproxylists.net").
        Set("Origin", "http://www.freeproxylists.net").
        Set("Referer", "http://www.freeproxylists.net/").
        Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8").
        Set("Accept-Language", "en-US,en;q=0.8").
        Set("Cache-Control", "max-age=0").
        Set("DNT", "1").
        Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/47.0.2526.106 Safari/537.36").
        End()

    if errs != nil {
        return nil, errs[0]
    }

    return resp, nil
}

func (this *FreeProxyLists) Result() string {
    return "Result string"
}
