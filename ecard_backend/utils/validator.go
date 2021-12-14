package utils

import (
    "github.com/go-playground/validator/v10"
    "sync"
)

var (
	validate *validator.Validate
	once sync.Once
)

// 单例
func GetValidator() *validator.Validate {
    once.Do(func() {
        validate = validator.New()
    })
    return validate
}
