package cli

import (
	"flag"
	"fmt"

	"github.com/hardfist/gopack/pkg/bundler"
)

// CLI represents the command line interface
type CLI struct {
	bundler *bundler.Bundler
}

// New creates a new CLI instance
func New() *CLI {
	return &CLI{
		bundler: bundler.New(),
	}
}

// Run executes the CLI application
func (c *CLI) Run() error {
	var (
		entryFile  = flag.String("entry", "", "Entry file to bundle")
		outputDir  = flag.String("output", "./dist", "Output directory")
		outputFile = flag.String("name", "bundle.js", "Output file name")
		mode       = flag.String("mode", "development", "Build mode (development/production)")
		help       = flag.Bool("help", false, "Show help")
	)

	flag.Parse()

	if *help {
		c.showHelp()
		return nil
	}

	// If no flag is provided but there are args, use the first arg as entry
	if *entryFile == "" && len(flag.Args()) > 0 {
		*entryFile = flag.Args()[0]
	}

	if *entryFile == "" {
		c.showHelp()
		return fmt.Errorf("entry file is required")
	}

	// Configure bundler
	c.bundler.SetOutputDir(*outputDir)
	c.bundler.SetConfig(&bundler.Config{
		OutputFile: *outputFile,
		Mode:       *mode,
	})

	// Run bundling
	return c.bundler.Bundle(*entryFile)
}

func (c *CLI) showHelp() {
	fmt.Println(`Gopack - Webpack in Go

Usage:
  gopack [options] <entry-file>
  gopack --entry <entry-file> [options]

Options:
  --entry <file>     Entry file to bundle
  --output <dir>     Output directory (default: ./dist)
  --name <file>      Output file name (default: bundle.js)
  --mode <mode>      Build mode: development/production (default: development)
  --help             Show this help

Examples:
  gopack app.js
  gopack --entry src/index.js --output build --name app.bundle.js
  gopack --entry main.js --mode production`)
}
