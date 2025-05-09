package test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/busykoala/zabbix-template-generator/internal/generator"
)

func TestGenerateBasicTemplate(t *testing.T) {
	group := generator.TemplateGroup{
		UUID: "uuid-group-1",
		Name: "Templates/Applications",
	}

	item := generator.Item{
		UUID:      "uuid-item-1",
		Name:      "HTTP 5xx rate",
		Type:      "DEPENDENT",
		Key:       "kubernetes.controller.client_http_requests_500.rate",
		Delay:     "0",
		ValueType: "FLOAT",
		Preprocessing: []generator.PreprocessingStep{
			{Type: "CHANGE_PER_SECOND", Parameters: []string{""}},
		},
		MasterItem: &generator.MasterItem{
			Key: "kubernetes.controller.get_metrics",
		},
	}

	template := generator.Template{
		UUID:        "uuid-template-1",
		Template:    "My Template",
		Name:        "My Template",
		Description: "Test description",
		Vendor:      generator.Vendor{Name: "Zabbix", Version: "7.0-1"},
		Groups:      []generator.TemplateGroup{group},
		Items:       []generator.Item{item},
	}

	export := generator.ZabbixExport{
		Version:        "7.0",
		TemplateGroups: []generator.TemplateGroup{group},
		Templates:      []generator.Template{template},
	}

	data, err := generator.Generate(export)
	if err != nil {
		t.Fatalf("failed to generate JSON: %v", err)
	}

	var parsed map[string]interface{}
	err = json.Unmarshal(data, &parsed)
	if err != nil {
		t.Fatalf("generated JSON is not valid: %v", err)
	}

	if _, ok := parsed["zabbix_export"]; !ok {
		t.Error("missing root key 'zabbix_export'")
	}

	t.Logf("Generated JSON:\n%s", data)
}

func TestMissingOptionalFields(t *testing.T) {
	item := generator.Item{
		UUID:      "uuid-item-2",
		Name:      "Minimal",
		Type:      "DEPENDENT",
		Key:       "some.key",
		Delay:     "1s",
		ValueType: "FLOAT",
	}

	data, err := json.Marshal(item)
	if err != nil {
		t.Fatalf("failed to marshal item: %v", err)
	}

	jsonStr := string(data)
	if strings.Contains(jsonStr, "description") {
		t.Error("unexpected field 'description' in minimal item")
	}
	if strings.Contains(jsonStr, "tags") {
		t.Error("unexpected field 'tags' in minimal item")
	}
}
