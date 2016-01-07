# go-delega

Package go-delega is a simple proxy list fetcher and collector.

## Usage

Simple usage:

    package main

    import (
        "fmt"
        "strings"

        "github.com/zeuxisoo/go-delega"
    )

    func main() {
        provider, _  := delega.Create(delega.XiCiDaiLi)
        response, _  := provider.Fetch()
        proxyList, _ := provider.Result(response)

        for _, proxy := range proxyList {
            fmt.Printf("%s://%s:%s\n", strings.ToLower(proxy.Protocol), proxy.Ip, proxy.Port)
        }
    }
