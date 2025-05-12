package test

import (
	"encoding/json"
	"github.com/busykoala/zabbix-template-generator"
	"testing"
)

func TestGenerateBasicTemplate(t *testing.T) {
	export := zabbix_template_generator.ZabbixExport{
		Version: "6.0",
		Templates: []zabbix_template_generator.Template{
			{
				UUID:     "test-uuid",
				Host:     "TestHost",
				Template: "TestTemplate",
				Name:     "Test Template",
			},
		},
	}

	data, err := zabbix_template_generator.Generate(export)
	if err != nil {
		t.Fatalf("Generate() returned error: %v", err)
	}

	if !json.Valid(data) {
		t.Errorf("Generated output is not valid JSON:\n%s", string(data))
	}

	if len(data) == 0 {
		t.Errorf("Generated JSON is empty")
	}
}
