# pprofui
Ripped from [cockroachdb](https://github.com/cockroachdb/cockroach/tree/master/pkg/server/debug/pprofui)

Embedded pprof webui for existing webservices, [probably a bit dangerous](https://mmcloughlin.com/posts/your-pprof-is-showing).

For an example use, see [example/example.go](https://github.com/polynomialspace/pprofui/blob/master/example/example.go).

The example creates a test load when visiting localhost:8080, and has the pprofui exposed from a second httpmux on localhost:8081.

You may have to refresh a few times for the data to populate, but it should look something like this:
![pprofui screenshot](https://raw.githubusercontent.com/polynomialspace/pprofui/master/example/screenshot.png)
