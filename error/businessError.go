package businessError

type Err struct {
	Code int `json:"code"`
	Message string `json:"message"`
}

func New(code int, msg ...string) *Err  {
	message := ""
	if len(msg) <= 0 {
		message = GetMsg(code)
	} else {
		message = msg[0]
	}
	return &Err{
		Code:    code,
		Message: message,
	}
}

func (b *Err) Error() string  {
	return b.Message
}



