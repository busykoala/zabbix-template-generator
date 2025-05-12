package zabbix_template_generator

type Macro struct {
	Macro       string `json:"macro"`
	Value       string `json:"value,omitempty"`
	Description string `json:"description,omitempty"`
}
