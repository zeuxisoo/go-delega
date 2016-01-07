// Provider contract is a set of interfaces or structs that the services used by the library
package contract

import (
    "net/http"
)

// Provider is an interface that make sure all providers are contain same methods
type Provider interface {
    Name() string
    Fetch() (*http.Response, error)
    Result(*http.Response) ([]ProxyList, error)
}

// ProxyList represents a collection of proxy matching defined format
type ProxyList struct {
    Ip          string
    Port        string
    Protocol    string
    Country     string
}
