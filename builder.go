package zabbix_template_generator

import (
	"encoding/json"
)

func Generate(export ZabbixExport) ([]byte, error) {
	return json.MarshalIndent(map[string]any{
		"zabbix_export": export,
	}, "", "    ")
}
