package tfjson2

import (
	"encoding/json"
	"github.com/hashicorp/terraform/terraform"
	"strings"
)

type Results map[string]Resource

type Resource struct {
	Attributes    map[string]string `json:"attributes"`
	OldAttributes map[string]string `json:"oldAttributes"`
	Name          string            `json:"name"`
	ResourceType  string            `json:"resourceType"`
	RequiresNew   bool              `json:"requiresNew"`
	WillDestroy   bool              `json:"willBeDestroyed"`
}

type JsonExporter struct {
	Plan *terraform.Plan
}

func (e *JsonExporter) Results() (Results) {
	results := Results{}

	for _, module := range e.Plan.Diff.Modules {
		var moduleName string

		for i, pathPart := range module.Path {
			if i != 0 {
				moduleName += pathPart + "."
			}
		}

		for key, diff := range module.Resources {
			attributes := make(map[string]string, len(diff.Attributes))
			oldAttributes := make(map[string]string, len(diff.Attributes))

			for attrName, attrVal := range diff.Attributes {
				attributes[attrName] = attrVal.New
				oldAttributes[attrName] = attrVal.Old
			}

			keyParts := strings.SplitN(key, ".", 2)

			results[moduleName+key] = Resource{
				Attributes:    attributes,
				OldAttributes: oldAttributes,
				RequiresNew:   diff.RequiresNew(),
				ResourceType:  keyParts[0],
				Name:          keyParts[1],
				WillDestroy:   diff.Destroy,
			}
		}
	}

	return results
}

func (e *JsonExporter) ToJson() (*string, error) {
	data, err := json.Marshal(e.Results())
	if err != nil {
		return nil, err
	}

	dataAsString := string(data)

	return &dataAsString, nil
}
