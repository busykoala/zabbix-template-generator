package generator

type TemplateGroup struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
}

type Vendor struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type Macro struct {
	Macro       string `json:"macro"`
	Value       string `json:"value,omitempty"`
	Description string `json:"description,omitempty"`
}

type MasterItem struct {
	Key string `json:"key"`
}

type ItemTag struct {
	Tag   string `json:"tag"`
	Value string `json:"value"`
}

type PreprocessingStep struct {
	Type         string   `json:"type"`
	Parameters   []string `json:"parameters"`
	ErrorHandler string   `json:"error_handler,omitempty"`
}

type TriggerTag struct {
	Tag   string `json:"tag"`
	Value string `json:"value"`
}

type Trigger struct {
	UUID        string       `json:"uuid"`
	Expression  string       `json:"expression"`
	Name        string       `json:"name"`
	EventName   string       `json:"event_name"`
	Priority    string       `json:"priority"`
	Description string       `json:"description,omitempty"`
	Tags        []TriggerTag `json:"tags,omitempty"`
}

type Item struct {
	UUID          string              `json:"uuid"`
	Name          string              `json:"name"`
	Type          string              `json:"type"`
	Key           string              `json:"key"`
	Delay         string              `json:"delay"`
	ValueType     string              `json:"value_type"`
	Description   string              `json:"description,omitempty"`
	Units         string              `json:"units,omitempty"`
	History       string              `json:"history,omitempty"`
	Trends        string              `json:"trends,omitempty"`
	Timeout       string              `json:"timeout,omitempty"`
	URL           string              `json:"url,omitempty"`
	Headers       []map[string]string `json:"headers,omitempty"`
	Preprocessing []PreprocessingStep `json:"preprocessing,omitempty"`
	MasterItem    *MasterItem         `json:"master_item,omitempty"`
	Tags          []ItemTag           `json:"tags,omitempty"`
	Triggers      []Trigger           `json:"triggers,omitempty"`
	ValueMap      *ValueMapRef        `json:"valuemap,omitempty"`
}

type ValueMapRef struct {
	Name string `json:"name"`
}

type Template struct {
	UUID        string          `json:"uuid"`
	Template    string          `json:"template"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Vendor      Vendor          `json:"vendor"`
	Groups      []TemplateGroup `json:"groups"`
	Items       []Item          `json:"items"`
	Macros      []Macro         `json:"macros,omitempty"`
	Tags        []ItemTag       `json:"tags,omitempty"`
}
