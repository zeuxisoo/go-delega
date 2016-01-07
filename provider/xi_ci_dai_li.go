package provider

import (
    "github.com/parnurzeal/gorequest"
    "github.com/PuerkitoBio/goquery"

    "github.com/zeuxisoo/go-delega/contract"

    "net/http"
    "strings"
)

// A XiCiDaiLi is an object to provider set of defined methods for client.
type XiCiDaiLi struct {
}

// Get the current provider name
func(this *XiCiDaiLi) Name() string {
    return "Xi Ci Dai Li"
}

// Make request to the target service and return the http response
func(this *XiCiDaiLi) Fetch() (*http.Response, error) {
    request := gorequest.New()

    resp, _, errs := request.
        Get("http://www.xicidaili.com/wn/").
        Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/47.0.2526.106 Safari/537.36").
        End()

    if errs != nil {
        return nil, errs[0]
    }

    return resp, nil
}

// Get the collection of the ProxyList by parse the http response
func (this *XiCiDaiLi) Result(response *http.Response) ([]contract.ProxyList, error) {
    doc, err := goquery.NewDocumentFromResponse(response)

    if err != nil {
        return nil, err
    }

    var proxyList []contract.ProxyList

    proxyTable := doc.Find("table#ip_list")
    proxyRows  := proxyTable.Find("tr:nth-child(n+2)")

    proxyRows.Each(func(i int, s *goquery.Selection) {
        country  := s.Find("td:nth-child(2) img").AttrOr("alt", "n/a")
        ip       := s.Find("td:nth-child(3)").Text()
        port     := s.Find("td:nth-child(4)").Text()
        protocol := s.Find("td:nth-child(7)").Text()

        if ip != "" && port != "" {
            proxyList = append(proxyList, contract.ProxyList{
                Ip      : ip,
                Port    : port,
                Protocol: protocol,
                Country : strings.ToLower(country),
            })
        }
    })

    return proxyList, nil
}
