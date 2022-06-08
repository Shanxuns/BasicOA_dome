package task

import _type "BasicOA/type"

func SubmitTask(code int, message interface{}, task int, data interface{}) {
	go func() {
		Channel <- _type.Respond{
			Code:    code,
			Message: message,
			Task:    task,
			Data:    data,
		}
	}()
}
