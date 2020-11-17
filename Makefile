JSDIR=./js/

ifeq ($(OS),Windows_NT)
    BROWSER = start
else
	UNAME := $(shell uname -s)
	ifeq ($(UNAME), Linux)
		BROWSER = xdg-open
	endif
	ifeq ($(UNAME), Darwin)
		BROWSER = open
	endif
endif

.PHONY: all clean serve

all: wasm serve

wasm:
	[ -d $(JSDIR) ] || mkdir -p $(JSDIR)
	cp -f "$$(go env GOROOT)/misc/wasm/wasm_exec.js" ./js/
	GOOS=js GOARCH=wasm go build -o lib.wasm main.go

serve:
	$(BROWSER) 'http://localhost:5000'
	GOARCH=amd64 go build -o pigo-face-tracking server/main.go && ./pigo-face-tracking

clean:
	rm -f *.wasm

debug:
	@echo $(UNAME)
