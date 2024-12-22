package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func lsp(environment string) error {
	projectInfo, err := getProjectInfo(environment)
	if err != nil {
		return err
	}
	jsonData, err := jsonFromString(projectInfo)
	if err != nil {
		return fmt.Errorf("Error parsing project information")
	}
	if err := createClangdFile(jsonData); err != nil {
		return err
	}
	return nil
}

func getProjectInfo(environment string) (string, error) {
	// Check if PlatformIO is installed
	_, err := exec.Command("pio", "--version").Output()
	if err != nil {
		return "", fmt.Errorf("please install PlatformIO and make sure it is in the PATH")
	}

	cmd := exec.Command("pio", "-f", "-c", "vim", "run", "-t", "idedata", "--environment", environment)
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("error executing PlatformIO command.\n" +
			"Check if the \"platformio.ini\" file exists and the environment name is correct.\n" +
			"Create a new project using \"pio init --ide vim\"")
	}

	return string(output), nil
}

func jsonFromString(text string) (*ProjectInfo, error) {
	start := strings.Index(text, "{")
	end := strings.LastIndex(text, "}")
	if start == -1 || end == -1 {
		return nil, fmt.Errorf("no JSON data found in the text")
	}

	jsonStr := text[start : end+1]
	var projectInfo ProjectInfo
	if err := json.NewDecoder(bytes.NewBufferString(jsonStr)).Decode(&projectInfo); err != nil {
		return nil, fmt.Errorf("error parsing JSON: %v", err)
	}

	return &projectInfo, nil
}

func createClangdFile(data *ProjectInfo) error {
	f, err := os.Create(".clangd")
	if err != nil {
		return fmt.Errorf("cannot open .clangd file for writing. Make sure you have write permissions: %v", err)
	}
	defer f.Close()
	f.WriteString("CompileFlags:\n  Add:\n    - -ferror-limit=0\n")
	for _, include := range data.Includes.Build {
		fmt.Fprintf(f, "    - -I%s\n", include)
	}
	for _, include := range data.Includes.Compatlib {
		fmt.Fprintf(f, "    - -I%s\n", include)
	}
	f.WriteString(`
Diagnostics:
  Suppress:
    - unused-includes
    - no_member
    - pp_file_not_found
    - ovl_no_viable_conversion_in_cast
    - init_conversion_failed
    - reference_bind_failed
    - fatal_too_many_errors
    - ovl_no_viable_function_in_init
    - typename_nested_not_found
    - access
    - ovl_no_viable_member_function_in_call
    - ovl_no_viable_function_in_call
    - member_function_call_bad_type
    - misc-definitions-in-headers
    - typecheck_member_reference_struct_union
    - typecheck_invalid_operands
    - -Wstring-plus-int
    - typecheck_convert_pointer_int
    - typecheck_nonviable_condition
`)
	return nil
}

type ProjectInfo struct {
	Includes struct {
		Build     []string `json:"build"`
		Compatlib []string `json:"compatlib"`
	} `json:"includes"`
}
