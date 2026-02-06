package types

type Tool struct {
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Parameters  ToolParameter `json:"parameters"`
}

type ParameterType string

const (
	TypeString  ParameterType = "string"
	TypeNumber  ParameterType = "number"
	TypeBoolean ParameterType = "boolean"
	TypeObject  ParameterType = "object"
)

type ToolParameter struct {
	Type       ParameterType                    `json:"type"`
	Properties map[string]ToolParameterProperty `json:"properties"`
	Required   []string                         `json:"required"`
}

type ToolParameterProperty struct {
	Type        ParameterType `json:"type"`
	Description *string       `json:"description,omitempty"`
	Enum        []interface{} `json:"enum,omitempty"`
}

type ToolInvokeRequest struct {
	Name      string                 `json:"name"`
	Arguments map[string]interface{} `json:"arguments"`
}

type ToolInvokeResponse struct {
	Result interface{} `json:"result"`
}
