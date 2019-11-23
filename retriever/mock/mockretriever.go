package mock

/**
alt+enter 实现接口
*/
type Retriever struct {
	Contents string
}

func (r Retriever) Post(url string, form map[string]string) string {
	r.Contents = form["contents"]
	return "ok"
}

//只要实现了这接口
func (r Retriever) Get(url string) string {
	return r.Contents

}
