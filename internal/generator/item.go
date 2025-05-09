package generator

const (
	ZABBIX_ITEM_TYPE_DEPENDENT  ZabbixItemType = "DEPENDENT"
	ZABBIX_ITEM_TYPE_HTTP_AGENT ZabbixItemType = "HTTP_AGENT"
)

type Item interface {
	ValidateType() bool
}

type ZabbixItemType string

type DependentItem struct {
	UUID          string              `json:"uuid"`
	Name          string              `json:"name"`
	Type          ZabbixItemType      `json:"type"`
	Key           string              `json:"key"`
	Delay         string              `json:"delay"`
	ValueType     string              `json:"value_type"`
	Units         string              `json:"units"`
	Description   string              `json:"description"`
	PreProcessing []PreprocessingStep `json:"preprocessing"`
	MasterItem    MasterItem          `json:"master_item"`
	Tags          []Tag               `json:"tags"`
}

func (i DependentItem) ValidateType() bool {
	return i.Type == ZABBIX_ITEM_TYPE_DEPENDENT
}

type HTTPAgentItem struct {
	UUID          string              `json:"uuid"`
	Name          string              `json:"name"`
	Type          ZabbixItemType      `json:"type"`
	Key           string              `json:"key"`
	History       string              `json:"history"`
	Trends        string              `json:"trends"`
	ValueType     string              `json:"value_type"`
	Description   string              `json:"description"`
	PreProcessing []PreprocessingStep `json:"preprocessing"`
	URL           string              `json:"url"`
	RetrieveMode  string              `json:"retrieve_mode"`
	Tags          []Tag               `json:"tags"`
	Triggers      []Trigger           `json:"triggers"`
}

func (i HTTPAgentItem) ValidateType() bool {
	return i.Type == ZABBIX_ITEM_TYPE_HTTP_AGENT
}

type PreprocessingStep struct {
	Type         string   `json:"type"`
	Parameters   []string `json:"parameters"`
	ErrorHandler string   `json:"error_handler,omitempty"`
}

type MasterItem struct {
	Key string `json:"key"`
}

type Trigger struct {
	UUID        string `json:"uuid"`
	Expression  string `json:"expression"`
	Name        string `json:"name"`
	EventName   string `json:"event_name"`
	Priority    string `json:"priority"` // TODO: Make Type
	Description string `json:"description,omitempty"`
	ManualClose string `json:"manual_close,omitempty"`
	// Dependencies []Dependency `json:"dependencies"` //TODO: Make Type
	Tags []Tag `json:"tags,omitempty"`
}
