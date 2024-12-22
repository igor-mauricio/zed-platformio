import subprocess
import json
import sys
from dataclasses import dataclass
from typing import Optional

@dataclass
class Error:
    message: str = ""
    data: dict|None = None

def main() -> int:
    if len(sys.argv) < 2:
        print("Usage: python3 clangGen.py <environment>")
        return 1
    environment = sys.argv[1]
    project_info, err = get_project_info(environment)
    if err != None:
        print(err.message)
        return 1
    json_data, err = json_from_string(project_info)
    if err != None:
        print("Error parsing project information.")
        return 1
    err = create_clangd_file(json_data)
    if err != None:
        print(err.message)
        return 1
    print(".clangd file created successfully.")
    return 0

def get_project_info(environment: str) -> tuple[str, Optional[Error]]:
    cmd = ["pio", "--version"]
    try:
        subprocess.run(cmd, capture_output=True, text=True, check=True)
    except:
        return "", Error("Please install PlatformIO and make sure it is in the PATH.")
    command = ["pio", "-f", "-c", "vim", "run", "-t", "idedata", "--environment", environment]
    try:
        result = subprocess.run(command, capture_output=True, text=True, check=True).stdout
    except:
        return "", Error("Error executing PlatformIO command.\nCheck if the \"platformio.ini\" file exists and the environment name is correct.")
    return result, None

def json_from_string(text: str) -> tuple[dict, Optional[Error]]:
    start = text.find('{')
    end = text.rfind('}')
    if start == -1 or end == -1:
        return {}, Error("No JSON data found in the text.")
    json_string:dict
    try:
       json_string = json.loads(text[start:end+1])
    except:
        return {}, Error("Error parsing JSON.")
    return json_string, None

def create_clangd_file(data: dict) -> Optional[Error]:
    includes_build = data.get("includes", {}).get("build", [])
    includes_compatlib = data.get("includes", {}).get("compatlib", [])
    try:
        f = open(".clangd", "w")
    except:
        return Error("Cannot open .clangd file for writing. Make sure you have write permissions.")
    f.write("CompileFlags:\n  Add:\n    - -ferror-limit=0\n")
    for include in includes_build:
        f.write(f"    - -I{include}\n")
    for include in includes_compatlib:
        f.write(f"    - -I{include}\n")
    f.write("""
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
    - typecheck_nonviable_condition\n""")
    return None

if __name__ == "__main__":
    exit(main())
