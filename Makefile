# Copyright 2009 The Go Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.
SOURCE_DIR := $(shell pwd)
OUTPUT_DIR := $(SOURCE_DIR)/build

.PHONY: all

all:oce ocecli
	@mkdir -p $(OUTPUT_DIR)
	@echo  "build success.\n"
	@echo "Installed files list."
	@echo "  $(OCE_EXEC_FILE)"
	@echo "  $(OCECLI_EXEC_FILE)"
	@echo "Now you can start your OCEChain trip"
	

OCE_EXEC_FILE := $(SOURCE_DIR)/build/oce
OCE_SOURCE_FILES :=  $(SOURCE_DIR)/oce/main.go $(SOURCE_DIR)/oce/utils.go
.PHONY: oce
oce:$(OCE_EXEC)
	go build -o $(OCE_EXEC_FILE) $(OCE_SOURCE_FILES)
	@echo "output file: $(OCE_EXEC_FILE)"

OCECLI_EXEC_FILE := $(SOURCE_DIR)/build/ocecli
OCECLI_SOURCE_FILES :=  $(SOURCE_DIR)/ocecli/main.go
.PHONY: ocecli
ocecli:$(OCECLI_EXEC)
	go build -o $(OCECLI_EXEC_FILE) $(OCECLI_SOURCE_FILES)
	@echo "output file: $(OCECLI_EXEC_FILE)"