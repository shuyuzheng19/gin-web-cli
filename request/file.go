package request

type FileRequest struct {
	Page    int    `form:"page"`    //第几页文件
	Keyword string `form:"keyword"` //文件的关键字
	Sort    string `form:"sort"`    //文件排序方式 size:大小排序和date:日期排序
}
