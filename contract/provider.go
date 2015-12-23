package contract

type Provider interface{
    Name() string
    Fetch() string
    Result() string
}
