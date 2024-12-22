# Welcome to the platformio-zed wiki!

## Quick Start

1. Install the PlatformIO CLI

#### Linux/macOS

```bash
curl -fsSL -o get-platformio.py https://raw.githubusercontent.com/platformio/platformio-core-installer/master/get-platformio.py
```
[official docs](https://docs.platformio.org/en/latest/core/installation/methods/installer-script.html)

2. Install zed-platformio package

```bash
go install github.com/igor-mauricio/zed-platformio@latest
```

3. Run the lsp command to update the files in clangd

```bash
zed-platformio lsp
```



## Available commands:

```bash
build       Build the project
completion  Generate the autocompletion script for the specified shell
help        Help about any command
home        Open Platformio Home
lsp         Configure Clangd for LSP
monitor     Monitor the serial port
test        Run test suite
upload      Upload firmware to the board
```

### uninstall

```bash
rm -f $(which zed-platformio)
```
