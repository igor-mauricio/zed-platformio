BOARD = lilygo-t-display-s3
TEST_BOARD = native

lsp:
	python clangGen.py $(BOARD)

build:
	pio run --environment $(BOARD)

upload:
	platformio run --target upload --environment $(BOARD)

monitor:
	platformio device monitor --environment $(BOARD)

upload_monitor:
	platformio run --target upload --target monitor --environment $(BOARD)

platformio:
	pio home

test_native:
	pio test --environment $(TEST_BOARD)
