package contract

import (
    "net/http"
)

type Provider interface {
    Name() string
    Fetch() (*http.Response, error)
    Result(*http.Response) ([]ProxyList, error)
}

type ProxyList struct {
    Ip          string
    Port        string
    Protocol    string
}
