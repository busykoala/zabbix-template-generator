package generator

import (
    "encoding/json"
)

type ZabbixExport struct {
    Version        string     `json:"version"`
    TemplateGroups []TemplateGroup `json:"template_groups"`
    Templates      []Template `json:"templates"`
}

func Generate(export ZabbixExport) ([]byte, error) {
    return json.MarshalIndent(map[string]any{
        "zabbix_export": export,
    }, "", "    ")
}

