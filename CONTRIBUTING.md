# Contributing to NetMapper

## Dev Environment Setup

**Requirements:** Go 1.21+, Git, make

1. Fork the repository on GitHub: [github.com/mdryaan/netmapper](https://github.com/mdryaan/netmapper)

2. Clone your fork (not the original):

```bash
git clone https://github.com/your_username/netmapper.git
cd netmapper
go mod tidy
make build
./netmapper --help
```

3. Add the upstream remote so you can pull future changes:

```bash
git remote add upstream https://github.com/mdryaan/netmapper.git
```

Run against the example configs to verify your build:

```bash
./netmapper validate --config examples/simple.yaml
./netmapper analyze  --config examples/complex.yaml
./netmapper visualize --config examples/simple.yaml
```

---

## Adding a New Command

1. Create `cmd/<name>.go` in the `cmd` package
2. Define a `*cobra.Command` var
3. Call `rootCmd.AddCommand(<name>Cmd)` in `init()`
4. Use `parser.Load(cfgFile)` and `graph.Build(topology)` to get the graph
5. Use `pkg/output` for all terminal output — never call `fmt.Println` directly in commands

```go
package cmd

import (
    "github.com/mdryaan/netmapper/pkg/output"
    "github.com/mdryaan/netmapper/pkg/parser"
    "github.com/spf13/cobra"
)

var myCmd = &cobra.Command{
    Use:   "mycommand",
    Short: "Short description",
    RunE: func(cmd *cobra.Command, args []string) error {
        if cfgFile == "" {
            output.Fatal("--config flag is required")
        }
        topology, err := parser.Load(cfgFile)
        if err != nil {
            output.Fatal(err.Error())
        }
        // your logic here
        return nil
    },
}

func init() {
    rootCmd.AddCommand(myCmd)
}
```

---

## Adding a New Exporter

1. Create `pkg/exporter/<format>_exporter.go`
2. Define a struct that implements the `Exporter` interface:

```go
type Exporter interface {
    Export(report *models.AnalysisReport) ([]byte, error)
}
```

3. Register it in `pkg/exporter/exporter.go` inside the `New()` switch

```go
case "xml":
    return &XMLExporter{}, nil
```

---

## Adding a New Node Type

1. Add the constant to `internal/models/node.go`:

```go
const (
    NodeTypeLoadBalancer NodeType = "load-balancer"
)
```

2. Update the `IsValid()` method on `Node` to include the new type
3. Add a display symbol to `nodeSymbols` in `pkg/visualizer/ascii.go`

---

## PR Guidelines

- One logical change per PR
- Run `go mod tidy` and `make build` before submitting
- PR title follows conventional commits: `feat(scope): description`
- All exported functions must have a purpose that's obvious from their name — no docstrings needed
- Include an example of the command output in the PR description when adding a new command

---

## Code Style Rules

- Zero comments — names must speak for themselves
- No `fmt.Println` in commands — use `pkg/output` functions
- No global mutable state beyond cobra flag vars in `cmd/`
- Strong types everywhere — avoid `interface{}` and `map[string]interface{}`
- Return errors up; call `output.Fatal` only at the command layer
- Keep functions short — if a function needs more than ~30 lines, split it
- All exported types go in `internal/models/` — never define domain types inside `pkg/`
