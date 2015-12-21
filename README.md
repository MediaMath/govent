### govent - cli and library for sending events to the graphite events api

#### Command line utility
To install:

```bash
$ go install github.com/MediaMath/govent
```

To use:

```bash
$ export GRAPHITE_URL=https://example.com/events/ 
$ export GRAPHITE_USER=foo 
$ export GRAPHITE_PASSWORD=bar 
$ govent --tag go.write.me.an.event.build --what what.aint.no.country "my data is fo realz"
```

#### Go library

To get:

```bash
$ go get github.com/MediaMath/govent/graphite
```

To use:

```go
import "github.com/MediaMath/govent/graphite"
g := graphite.New("foo", "bar", "https://example.com/events/")
g.Publish(graphite.NewEvent("what.aint.no.country", "my data is fo realz", "go.write.me.an.event.build"))
```
