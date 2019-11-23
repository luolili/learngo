package real

import (
	"net/http"
	"net/http/httputil"
	"time"
)

type Retriever struct {
	UserAgent string
	TimeOut   time.Duration
}

//只需要实现接口 里面的方法，不需要说明是实现了哪个接口
func (r Retriever) Get(url string) string {

	resp, err := http.Get(url)

	if err != nil {
		panic("error")
	}
	result, err := httputil.DumpResponse(resp, true)
	resp.Body.Close()
	if err != nil {
		panic("error--")
	}
	return string(result)

}
