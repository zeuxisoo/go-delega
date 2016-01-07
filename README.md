# go-delega

Package go-delega is a simple proxy list fetcher and collector.

[![Travis Build Status](https://travis-ci.org/zeuxisoo/go-delega.svg?branch=master)](https://travis-ci.org/zeuxisoo/go-delega)
[![GoWalker](https://img.shields.io/badge/gowalker-reference-blue.svg)](https://gowalker.org/github.com/zeuxisoo/go-delega)
[![GoDoc](https://godoc.org/github.com/zeuxisoo/go-delega?status.png)](http://godoc.org/github.com/zeuxisoo/go-delega)

## Installation

	go get github.com/zeuxisoo/go-delega

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

## Documents

See in [GoWalker] or [GoDoc] for usage and details.

[GoDoc]: https://gowalker.org/github.com/zeuxisoo/go-delega
[GoWalker]: https://gowalker.org/github.com/zeuxisoo/go-delega

## License

MIT License