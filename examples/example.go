package main

import (
    "fmt"
    "log/slog"
    "os"

    generator "github.com/busykoala/zabbix-template-generator"
)

func main() {
    templateGroup := generator.TemplateGroup{
        UUID: "a571c0d144b14fd4a87a9d9b2aa9fcd6",
        Name: "Templates/Applications",
    }
    group := generator.Group{Name: "Templates/Applications"}
    item := generator.HTTPAgentItem{
        UUID:        "b571c0d144b14fd4a87a9d9b2aa9fcd6",
        Name:        "Test Item",
        Type:        generator.ZABBIX_ITEM_TYPE_HTTP_AGENT,
        Key:         "testitem",
        History:     "1h",
        ValueType:   "TEXT",
        Description: "Test Item for HTTP query",
        URL:         "https://jsonplaceholder.typicode.com/todos/1",
    }
    dependentItem := generator.DependentItem{
        UUID:        "c571c0d144b14fd4a87a9d9b2aa9fcd6",
        Name:        "Test Item: Dep",
        Type:        generator.ZABBIX_ITEM_TYPE_DEPENDENT,
        Key:         "testitem.dependent",
        ValueType:   "TEXT",
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
    }

    export := generator.ZabbixExport{
        Version: "7.2",
        Templates: []generator.Template{
            {
                UUID:     "e9f4a7c3b1fd41d08c89bfc58e91f5d4",
                Template: "TestTemplate",
                Name:     "Test Template",
                Groups:   []generator.Group{group},
                Items:    []generator.Item{item, dependentItem},
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
