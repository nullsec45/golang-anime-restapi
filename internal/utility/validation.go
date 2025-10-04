package utility

import (
	"github.com/go-playground/validator/v10"
	"fmt"
	"strings"
	"github.com/google/uuid"
)

func Validate[T any](data T)map[string]string{
	err :=  validator.New().Struct(data)
	res := map[string]string{}
	if err != nil {
		for _, v := range err.(validator.ValidationErrors) {
			res[v.StructField()] = TranslateTag(v)
		}
	}

	return res
}

func TranslateTag(fd validator.FieldError) string {
	switch fd.ActualTag(){
	case "required" :
		return fmt.Sprintf("Field %s is required", strings.ToLower(fd.StructField()))
	case "min" :
		return fmt.Sprintf("Field %s size minimum", strings.ToLower(fd.StructField()), fd.Param())
	case "unique" :
		return fmt.Sprintf("Field %s must be unique", strings.ToLower(fd.StructField()))
	case "email" :
		return fmt.Sprintf("Field %s must be email valid", strings.ToLower(fd.StructField()))
	}

	return "validasi gagal"
}

func IsUUID(s string) bool {
	_, err := uuid.Parse(s)
	return err == nil
}