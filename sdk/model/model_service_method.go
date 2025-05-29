package model

type ServiceMethod struct {
	ExampleInput  map[string]interface{} `json:"exampleInput,omitempty"`
	Id            int64                  `json:"id,omitempty"`
	InputType     string                 `json:"inputType,omitempty"`
	MethodName    string                 `json:"methodName,omitempty"`
	MethodType    string                 `json:"methodType,omitempty"`
	OperationName string                 `json:"operationName,omitempty"`
	OutputType    string                 `json:"outputType,omitempty"`
	RequestParams []RequestParam         `json:"requestParams,omitempty"`
}
