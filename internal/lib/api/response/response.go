package response

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error,omitemty"`
}

type Message struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

func Error(msg string) Response {
	return Response{
		Status: "Error",
		Error:  msg,
	}
}

func OK() Response {
	return Response{
		Status: "Success",
	}
}
