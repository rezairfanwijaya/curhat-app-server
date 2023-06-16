package helper

import "github.com/go-playground/validator/v10"

type responseAPI struct {
	Meta meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type meta struct {
	Status   int    `json:"status"`
	Messsage string `json:"message"`
}

func GenerateResponseAPI(status int, message string, data interface{}) *responseAPI {
	return &responseAPI{
		Meta: meta{
			Status:   status,
			Messsage: message,
		},
		Data: data,
	}
}

func GenerateErrorBinding(err error) (errorBinding []string) {
	for _, e := range err.(validator.ValidationErrors) {
		errorBinding = append(errorBinding, e.Error())
	}

	return errorBinding
}
