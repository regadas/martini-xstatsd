## go-xstatsd ##

```go
import(
    "github.com/go-martini/martini"
    "github.com/regadas/go-xstatsd"
    sh "github.com/regadas/martini-xstatsd"
)

func main(){
    m := martini.Classic()

    stats := statsd.New("127.0.0.1:8125", "test.prefix")

    m.Use(sh.HandlerMetrics(stats))

    m.Run()
}
```