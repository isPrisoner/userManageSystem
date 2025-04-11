package utils

type Message struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data,omitempty"`
	Msg    string      `json:"msg,omitempty"`
}

func Ok(data interface{}, msg string) *Message {
	return &Message{
		Status: 200,
		Data:   data,
		Msg:    msg,
	}
}

func Lose(data interface{}, msg string) *Message {
	return &Message{
		Status: 404,
		Data:   data,
		Msg:    msg,
	}
}
