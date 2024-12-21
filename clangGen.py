import subprocess
import json

def main():
    output = run_pio_command()
    if output:
        json_data = extract_json(output)
        if json_data:
            create_clangd_file(json_data)
        else:
            print("No JSON data found in the output.")
    else:
        print("Failed to run PlatformIO command.")

def run_pio_command():
    command = ["pio", "-f", "-c", "vim", "run", "-t", "idedata", "--environment", "lilygo-t-display-s3"]
    try:
        result = subprocess.run(command, capture_output=True, text=True, check=True)
        return result.stdout
    except subprocess.CalledProcessError as e:
        print("Error executing PlatformIO command:")
        print(e.stderr)
        return None

def extract_json(output: str):
    try:
        start = output.find('{')
        end = output.rfind('}')
        if start != -1 and end != -1:
            json_string = output[start:end+1]
            return json.loads(json_string)
    except json.JSONDecodeError as e:
        print(f"Error parsing JSON: {e}")
    return None

def create_clangd_file(data: dict):
    includes_build = data.get("includes", {}).get("build", [])
    includes_compatlib = data.get("includes", {}).get("compatlib", [])
    # includes_toolchain = data.get("includes", {}).get("toolchain", [])
    # compile_flags = data.get("cxx_flags", [])
    # compiler_path = data.get("cxx_path", "")

    with open(".clangd", "w") as f:
        f.write("CompileFlags:\n")
        f.write("  Add:\n    - -ferror-limit=0\n")

        # for flag in compile_flags:
        #     f.write(f"    - {flag}\n")
        for include in includes_build:
            f.write(f"    - -I{include}\n")
        for include in includes_compatlib:
            f.write(f"    - -I{include}\n")
        # for include in includes_toolchain:
        #     f.write(f"    - -I{include}\n")
        # if compiler_path:
        #     f.write(f"    - --gcc-toolchain={compiler_path}\n")

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
    print(".clangd file created successfully.")



if __name__ == "__main__":
    main()
