package provider

import (
    "github.com/parnurzeal/gorequest"
    "github.com/PuerkitoBio/goquery"

    "github.com/zeuxisoo/go-delega/contract"

    "regexp"
    "net/http"
    "net/url"
    "path/filepath"
    "strings"
)

// A FreeProxyLists is an object to provider set of defined methods for client.
type FreeProxyLists struct {
}

// Get the current provider name
func (this *FreeProxyLists) Name() string {
    return "Free Proxy Lists"
}

// Make request to the target service and return the http response
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

// Get the collection of the ProxyList by parse the http response
func (this *FreeProxyLists) Result(response *http.Response) ([]contract.ProxyList, error) {
    doc, err := goquery.NewDocumentFromResponse(response)

    if err != nil {
        return nil, err
    }

    var proxyList []contract.ProxyList

    proxyTable := doc.Find("table.DataGrid")
    proxyRows  := proxyTable.Find("tr:nth-child(n+2)")

    proxyRows.Each(func(i int, s *goquery.Selection) {
        ip       := s.Find("td:nth-child(1) script").Text()
        port     := s.Find("td:nth-child(2)").Text()
        protocol := s.Find("td:nth-child(3)").Text()
        country  := s.Find("td:nth-child(5) img").AttrOr("src", "img/na.gif")

        if ip != "" && port != "" {
            proxyList = append(proxyList, contract.ProxyList{
                Ip      : decodeIp(ip),
                Port    : port,
                Protocol: protocol,
                Country : strings.ToLower(fileName(country)),
            })
        }
    })

    return proxyList, nil
}

func decodeIp(encodeString string) (string) {
    regex   := regexp.MustCompile(`IPDecode\(\"(.*?)\"\)`)
    matches := regex.FindStringSubmatch(encodeString)
    result  := matches[1]

    decodedString, _ := url.QueryUnescape(result)

    reader := strings.NewReader(decodedString)
    doc, _ := goquery.NewDocumentFromReader(reader)

    return doc.Find("a").Text()
}

func fileName(file_path string) (string) {
    base_name := filepath.Base(file_path)

    return strings.TrimSuffix(base_name, filepath.Ext(base_name))
}
