package registry

import (
	"mcp-server/tools"
	"mcp-server/types"
)

var Tools []types.Tool

func init() {
	Tools = append(Tools, tools.GetTextTools()...)
	Tools = append(Tools, tools.GetDatetimeTools()...)
	Tools = append(Tools, tools.GetConverterTools()...)
	Tools = append(Tools, tools.GetRandomTools()...)
}

func FindTool(name string) *types.Tool {
	for i := range Tools {
		if Tools[i].Name == name {
			return &Tools[i]
		}
	}
	return nil
}
