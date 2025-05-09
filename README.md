# Zabbix Template Generator

## Usage

```go
package main

import (
    "fmt"
    "github.com/yourusername/zabbix-template-generator/internal/generator"
)

func main() {
    // Initialize the generator
    gen := generator.NewGenerator()

    // Generate a Zabbix template
    template, err := gen.GenerateTemplate("MyTemplate")
    if err != nil {
        fmt.Println("Error generating template:", err)
        return
    }

    // Print the generated template
    fmt.Println("Generated Zabbix Template:")
    fmt.Println(template)
}
```

## Test

```bash
go test ./...
```
