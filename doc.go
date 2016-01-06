/*
Package go-delega is a simple proxy list fetcher and collector.

Simple usage:

    func main() {
        provider, _  := delega.Create(delega.XiCiDaiLi)
        resp, _      := provider.Fetch()
        proxyList, _ := provider.Result(resp)

        fmt.Println(proxyList)
    }
*/
package delega
