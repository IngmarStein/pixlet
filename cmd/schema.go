package cmd

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"tidbyt.dev/pixlet/runtime"
	"tidbyt.dev/pixlet/tools"
)

var (
	schemaOutput string
)

func init() {
	SchemaCmd.Flags().StringVarP(&schemaOutput, "output", "o", "", "Path for schema")
}

var SchemaCmd = &cobra.Command{
	Use:   "schema [path]",
	Short: "Print the configuration schema for a Pixlet app",
	Args:  cobra.MinimumNArgs(1),
	RunE:  schema,
	Long: `Determine the configuration schema for a Pixlet app.

The path argument should be the path to the Pixlet app to run. The
app can be a single file with the .star extension, or a directory
containing multiple Starlark files and resources. The output is in
JSON format.
	`,
}

func schema(cmd *cobra.Command, args []string) error {
	path := args[0]

	// check if path exists, and whether it is a directory or a file
	info, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("failed to stat %s: %w", path, err)
	}

	var fs fs.FS
	if info.IsDir() {
		fs = os.DirFS(path)
	} else {
		if !strings.HasSuffix(path, ".star") {
			return fmt.Errorf("script file must have suffix .star: %s", path)
		}

		fs = tools.NewSingleFileFS(path)
	}

	applet, err := runtime.NewAppletFromFS(filepath.Base(path), fs)
	if err != nil {
		return fmt.Errorf("failed to load applet: %w", err)
	}

	if schemaOutput == "" || schemaOutput == "-" {
		buf, err := json.MarshalIndent(applet.Schema, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to marshal JSON: %w", err)
		}
		fmt.Println(string(buf))
	} else {
		err = os.WriteFile(schemaOutput, applet.SchemaJSON, 0644)
		if err != nil {
			return fmt.Errorf("failed to write schema to file: %w", err)
		}
	}

	return nil
}
