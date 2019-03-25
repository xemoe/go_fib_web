package actions

import (
	"strconv"

	"github.com/gobuffalo/buffalo"
	f "github.com/xemoe/fib/fib"
)

type Message struct {
	Status string `json:"status" xml:"status" form:"status" query:"status"`
	Result string `json:"result" xml:"result" form:"result" query:"result"`
}

func FibHandler(c buffalo.Context) error {
	num, err := strconv.Atoi(c.Param("num"))

	if err != nil {
		return err
	}

	resp := &Message{
		Status: "OK",
		Result: f.Fib(num),
	}

	return c.Render(200, r.JSON(resp))
}
