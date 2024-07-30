package types

type Response struct {
	Code    int          `json:"code"`
	Message string       `json:"message"`
	Data    ResponseData `json:"data"`
}

type ResponseData struct {
	Name        string     `json:"name"`
	InputParam  []ParamDTO `json:"inputParam"`
	OutputParam []ParamDTO `json:"outputParam"`
}

type ParamDTO struct {
	Key        string          `json:"key"`
	DetailList []DetailItemDTO `json:"detailList"`
}

type DetailItemDTO struct {
	Comment  string `json:"comment"`
	Name     string `json:"name"`
	Required int    `json:"required"`
	Type     string `json:"type"`
}
