package zabbix_template_generator

import (
	"encoding/json"
)

func Generate(export ZabbixExport) ([]byte, error) {
	return json.Marshal(map[string]any{
		"zabbix_export": export,
	})
}
