PIO_ENV = lilygo-t-display-s3
TEST_PIO_ENV = native

PYTHON := $(shell command -v python3 || command -v python || echo "PYTHON_NOT_FOUND")
ifeq ($(PYTHON),PYTHON_NOT_FOUND)
    $(error Python not found in PATH)
endif

lsp:
	$(PYTHON) clangGen.py $(PIO_ENV)

build:
	pio run --environment $(PIO_ENV)

upload:
	pio run --target upload --environment $(PIO_ENV)

monitor:
	pio device monitor --environment $(PIO_ENV)

upload_monitor:
	pio run --target upload --target monitor --environment $(PIO_ENV)

platformio:
	pio home

pio_test:
	pio test --environment $(TEST_PIO_ENV)
