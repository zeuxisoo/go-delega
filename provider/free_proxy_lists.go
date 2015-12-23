package provider

type FreeProxyLists struct {
}

func (this *FreeProxyLists) Name() string {
    return "Free Proxy Lists"
}

func (this *FreeProxyLists) Fetch() string {
    return "Fetch string"
}

func (this *FreeProxyLists) Result() string {
    return "Result string"
}
