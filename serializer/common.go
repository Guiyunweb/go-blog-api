package serializer

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// DataList 基础列表结构
type DataList struct {
	Content       interface{} `json:"content"`
	TotalElements int64       `json:"totalElements"`
}

func BuildListResponse(content interface{}, totalElements int64) Response {
	return Response{
		Data: DataList{
			Content:       content,
			TotalElements: totalElements,
		},
	}
}
