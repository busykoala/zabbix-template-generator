package zabbix_template_generator

type ValueMap struct {
	UUID     string    `json:"uuid"`
	Name     string    `json:"name"`
	Mappings []Mapping `json:"mappings"`
}

type Mapping struct {
	Value    string `json:"value"`
	NewValue string `json:"newvalue"`
}
