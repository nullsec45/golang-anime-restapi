package dto

type ResponseSuccess[T any] struct {
	Code    int         `json:"code"`	
	Message string      `json:"message"`
	Data    *T 		    `json:"data,omitempty"`
	Meta    *PageMeta	`json:"meta,omitempty"`
}

type ResponseError[T any] struct {
	Code     int         `json:"code"`	
	Message  string      `json:"message"`
	Error    interface{} `json:"error,omitempty"`
}

func CreateResponseError(code int,message string) ResponseError [string]{
	return ResponseError[string]{
		Code : code,
		Message: message,	
		// Error: "",
	}
}

func CreateResponseErrorData(code int, message string, error map[string]string) ResponseError [string]{
	return ResponseError[string]{
		Code : code,
		Message: message,	
		Error: error,
	}
}

func CreateResponseSuccess(message string) ResponseSuccess [string]{
	return ResponseSuccess[string]{
		Code : 200,
		Message: message,	
		// Data: data,
	}
}

func CreateResponseSuccessWithData[T any](message string, data T) ResponseSuccess [T]{
	return ResponseSuccess[T]{
		Code : 200,
		Message: message,	
		Data: &data,
	}
}

func CreateResponseSuccessWithDataPagination[T any](message string, p Paginated[T]) ResponseSuccess[[]T] {
	return ResponseSuccess[[]T]{
		Code:    200,
		Message: message,
		Data:    &p.Data,  
		Meta:    &p.Meta,
	}
}