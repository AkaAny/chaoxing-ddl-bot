package response

import "encoding/json"

type ResponseEntity struct {
	StatusCode int
	Body       BaseResponse
}

type BaseResponse struct {
	Error int         `json:"error"`
	Cache bool        `json:"cache"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data"`
}

func MakeSuccess(data interface{}) BaseResponse {
	return BaseResponse{
		Error: 0,
		Msg:   "success",
		Data:  data,
	}
}

func (br BaseResponse) GetData(v interface{}) error {
	rawData, err := json.Marshal(br.Data)
	if err != nil {
		return err
	}
	return json.Unmarshal(rawData, v)
}
