package tfjson2

import (
	"encoding/json"
	"github.com/hashicorp/terraform/terraform"
)

type Results map[string]Resource

type Resource struct {
	WillDestroy   bool              `json:"willBeDestroyed"`
	Attributes    map[string]string `json:"attributes"`
	OldAttributes map[string]string `json:"oldAttributes"`
}

type JsonExporter struct {
	Plan *terraform.Plan
}

func (e *JsonExporter) Results() (Results) {
	results := Results{}

	for _, module := range e.Plan.Diff.Modules {
		for key, diff := range module.Resources {
			attributes := make(map[string]string, len(diff.Attributes))
			oldAttributes := make(map[string]string, len(diff.Attributes))

			for attrName, attrVal := range diff.Attributes {
				attributes[attrName] = attrVal.New
				oldAttributes[attrName] = attrVal.Old
			}

			results[key] = Resource{
				WillDestroy: diff.Destroy,
				Attributes: attributes,
				OldAttributes: oldAttributes,
			}
		}
	}

	return results
}

func (e *JsonExporter) ToJson() (*string, error) {
	data, err := json.Marshal( e.Results() )
	if err != nil {
		return nil, err
	}

	dataAsString := string( data )

	return &dataAsString, nil
}