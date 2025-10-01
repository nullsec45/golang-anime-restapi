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

func CreateResponseError(message string) ResponseError [string]{
	return ResponseError[string]{
		Code : 400,
		Message: message,	
		// Error: "",
	}
}

func CreateResponseErrorData(message string, error map[string]string) ResponseError [string]{
	return ResponseError[string]{
		Code : 422,
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