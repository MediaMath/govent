### gevent - cli and library for sending events to the graphite events api

```bash
GRAPHITE_URL=https://example.com/events/ GRAPHITE_PASSWORD=foo GRAHPITE_USERNAME=bar gevent --tag go.write.me.an.event.build --what what.aint.no.country "my data is fo realz"
```

```go
g := graphite.New("foo", "bar", "https://example.com/events/")
g.Publish(graphite.NewEvent("what.aint.no.country", "my data is fo realz", "go.write.me.an.event.build"))
```
