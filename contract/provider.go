package contract

import (
    "net/http"
)

type Provider interface{
    Name() string
    Fetch() (*http.Response, error)
    Result() string
}
