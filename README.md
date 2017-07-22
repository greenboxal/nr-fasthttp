# nr-fasthttp

NewRelic instrumentation helpers for valyala/fasthttp.

## Usage

### Instrumenting external segments:

```go
txn := newrelicApp.StartTransaction("test", nil, nil)
defer txn.End()

req := fasthttp.AcquireRequest()
defer fasthttp.ReleaseRequest(req)

resp := fasthttp.AcquireResponse()
defer fasthttp.ReleaseResponse(resp)

proxyClient := &fasthttp.HostClient{
  Addr: "localhost:8080",
}

nrfasthttp.Do(proxyClient, txn, req, res)
```

###  Instrumenting web transactions

With buaazp/fasthttprouter:
```go
router.Handle("GET", "/me", nrfasthttp.WrapHandler(newrelicApp, "/me", yourHandler)

router.NotFound = nrfasthttp.WrapHandler(newrelicApp, "404", yourHandler)
```

Using fastthttp only:
```go
fasthttp.Serve(listener, nrfasthttp.WrapHandler(newrelicApp, "Request", yourHandler))
```

For complete documentation, check [here](https://godoc.org/github.com/greenboxal/nr-fasthttp). The API is based on newrelic/go-agent API so usage is very similiar.

## License

See [here](LICENSE).

