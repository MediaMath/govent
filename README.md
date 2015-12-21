### govent - cli and library for sending events to the graphite events api

#### Installation

```bash
$ go get github.com/MediaMath/govent
```

#### Usage

```bash
$ export GRAPHITE_URL=https://example.com/events/
$ export GRAPHITE_USERNAME=foo
$ export GRAPHITE_PASSWORD=bar
$ govent --tag go.write.me.an.event.build --what what.aint.no.country "my data is fo realz"
```

```go
import "github.com/MediaMath/govent/graphite"

g := graphite.New("foo", "bar", "https://example.com/events/")
g.Publish(graphite.NewEvent("what.aint.no.country", "my data is fo realz", "go.write.me.an.event.build"))
```
