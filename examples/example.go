package main

import (
	"fmt"
	"log/slog"
	"os"

	generator "github.com/busykoala/zabbix-template-generator"
)

func main() {
	templateIdName := "TestTemplate"
	templateGroup := generator.TemplateGroup{
		UUID: generator.GenerateUUID(),
		Name: "Templates/Applications",
	}
	group := generator.Group{Name: "Templates/Applications"}
	trigger := generator.Trigger{
		UUID:       generator.GenerateUUID(),
		Expression: fmt.Sprintf("nodata(/%s/testitem,1m)=0", templateIdName),
		Name:       "Test Trigger",
		Priority:   generator.PriorityHigh,
	}
	item := generator.HTTPAgentItem{
		UUID:        generator.GenerateUUID(),
		Name:        "Test Item",
		Type:        generator.ZABBIX_ITEM_TYPE_HTTP_AGENT,
		Key:         "testitem",
		History:     "1h",
		ValueType:   generator.ItemValueTypeText,
		Description: "Test Item for HTTP query",
		URL:         "https://jsonplaceholder.typicode.com/todos/1",
		Triggers:    []generator.Trigger{trigger},
	}
	dependentTrigger := generator.Trigger{
		UUID:       generator.GenerateUUID(),
		Expression: fmt.Sprintf("nodata(/%s/testitem.dependent,1m)=0", templateIdName),
		Name:       "Test Trigger Dep",
		Priority:   generator.PriorityWarn,
	}
	dependentItem := generator.DependentItem{
		UUID:        generator.GenerateUUID(),
		Name:        "Test Item: Dep",
		Type:        generator.ZABBIX_ITEM_TYPE_DEPENDENT,
		Key:         "testitem.dependent",
		ValueType:   generator.ItemValueTypeUnsigned,
		Description: "Test Item for HTTP query: Dep",
		PreProcessing: []generator.PreprocessingStep{
			{
				Type:       "JSONPATH",
				Parameters: []string{"$.userId"},
			},
		},
		MasterItem: generator.MasterItem{
			Key: "testitem",
		},
		Triggers: []generator.Trigger{dependentTrigger},
	}

	page := generator.Page{
		Name: "Test Page",
	}
	dashboard := generator.Dashboard{
		UUID:  generator.GenerateUUID(),
		Name:  "Test Dashboard",
		Pages: []generator.Page{page},
	}

	export := generator.ZabbixExport{
		Version: "7.2",
		Templates: []generator.Template{
			{
				UUID:       generator.GenerateUUID(),
				Template:   templateIdName,
				Name:       "Test Template",
				Groups:     []generator.Group{group},
				Items:      []generator.Item{item, dependentItem},
				Dashboards: []generator.Dashboard{dashboard},
			},
		},
		TemplateGroups: []generator.TemplateGroup{templateGroup},
	}

	data, err := generator.Generate(export)
	if err != nil {
		slog.Error("generating export", "error", err)
		os.Exit(1)
	}

	err = os.WriteFile("zabbix-export.json", data, 0644)
	if err != nil {
		slog.Error("writing file", "error", err)
		os.Exit(1)
	}

	fmt.Println("✅ Zabbix export written to zabbix-export.json")
}
