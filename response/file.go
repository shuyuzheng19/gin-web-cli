package response

type FileResponse struct {
	Id        int    `json:"id"`      //文件ID
	Name      string `json:"name"`    //文件名
	CreatedAt myTime `json:"dateStr"` //上传日期
	Suffix    string `json:"suffix"`  //文件后缀
	Size      int64  `json:"size"`    //文件大小
	Md5       string `json:"md5"`     //文件md5
	Url       string `json:"url"`     //文件url
}

type SimpleFileResponse struct {
	Status  string `json:"status"`  //是否上传成功
	Message string `json:"message"` //成功或失败的原因
	Name    string `json:"name"`    //文件的名称
	Create  string `json:"create"`  //文件上传的日期
	Url     string `json:"url"`     //上传成功后返回的url
}
