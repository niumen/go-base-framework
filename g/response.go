package g

const (
	Success     int    = 0
	SuccessMsg  string = "success"
	FailedMsg  string = "failed"
	BadRequest  int    = 40001
	InternalErr int    = 50001
)

type ListResult struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Result  List   `json:"result"`
}

type Result struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

type List struct {
	List     interface{} `json:"list"`
	Total    int         `json:"total"`
	PageSize int         `json:"pageSize"`
}
