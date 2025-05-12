package test

import (
    "encoding/json"
    "testing"

    "github.com/busykoala/zabbix-template-generator/internal/generator"
)

func TestGenerateBasicTemplate(t *testing.T) {
    export := generator.ZabbixExport{
        Version: "6.0",
        Templates: []generator.Template{
            {
                UUID:     "test-uuid",
                Host:     "TestHost",
                Template: "TestTemplate",
                Name:     "Test Template",
            },
        },
    }

    data, err := generator.Generate(export)
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
