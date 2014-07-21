package statsd

import (
	"fmt"
	"github.com/go-martini/martini"
	"github.com/regadas/go-xstatsd"
	"net"
	"net/http"
	"strings"
	"time"
)

func HandlerMetrics(s *statsd.Statsd) martini.Handler {

	return func(rw http.ResponseWriter, r *http.Request, c martini.Context) {
		reqStartTime := time.Now()
		res := rw.(martini.ResponseWriter)

		c.Next()

		reqFinishTime := time.Now()
		duration := int64(reqFinishTime.Sub(reqStartTime) / time.Millisecond)

		sendStats := func(method string, path string, status int, duration int64) {
			path = strings.TrimSuffix(path, "/")
			statAll := fmt.Sprintf("request.%s", method)
			statReq := fmt.Sprintf(
				"request%s.%s",
				strings.Replace(path, "/", ".", -1),
				method,
			)
			statStatus := fmt.Sprintf("response.%d", status)

			s.Client.WithConnection(func(conn *net.Conn) {
				s.TimingRaw(conn, statAll, duration)
				s.TimingRaw(conn, statReq, duration)
				s.IncrementRaw(conn, statStatus)
			})
		}

		go sendStats(r.Method, r.URL.Path, res.Status(), duration)
	}
}
