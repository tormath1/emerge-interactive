BUILD_DIR=$(PWD)/build
APP=emerge-interactive

build: 
	@mkdir $(BUILD_DIR)
	@go build -i -o $(BUILD_DIR)/$(APP) main.go

clean:
	@rm -r $(BUILD_DIR)

install: build
	@cp $(BUILD_DIR)/$(APP) $(GOPATH)/bin

uninstall: clean
	@rm $(GOPATH)/bin/$(APP)
