package task

import (
	_type "BasicOA/type"
	"github.com/robfig/cron"
)

func Run() {
	c := cron.New()
	if err := c.AddFunc("*/1 * * * *", func() {
		Channel <- _type.Respond{
			Code:    1,
			Message: "在线人数",
			Data:    len(clients),
		}
	}); err != nil {
		return
	}
	c.Start()
}
