package utility

import (
	"github.com/go-playground/validator/v10"
	"fmt"
	"strings"
	"github.com/google/uuid"
	"strings"
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
	case "eqfield":
		return fd.Field()+" must be equal with "+fd.Param()+"."
	case "oneof":
		allowed := strings.ReplaceAll(fd.Param(), " ", ", ")
		return fmt.Sprintf("Field %s must be one of: %s", strings.ToLower(fd.StructField()), allowed)
	}
	}

	return "validasi gagal"
}

func IsUUID(s string) bool {
	_, err := uuid.Parse(s)
	return err == nil
}