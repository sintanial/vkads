package vkads

import "fmt"

type ApiError struct {
	Code    string                 `json:"code"`
	Message string                 `json:"message"`
	Extra   map[string]interface{} `json:"extra"`
}

func (a *ApiError) Error() string {
	return fmt.Sprintf("Vkads api error: %s {code=%s, extra=%+v}", a.Message, a.Code, a.Extra)
}

type FailedResponse struct {
	Error map[string]interface{} `json:"error"`
}
