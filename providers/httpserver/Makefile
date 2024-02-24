CAPABILITY_ID = "wrpc:httpserver"
NAME = "httpserver"
VENDOR = "tchap"
PROJECT = httpserver
VERSION = 0.1.0
REVISION = 0

NAME ?= $(PROJECT)
WASH ?= wash

par: build_x86_64_linux build_x86_64_macos
	@$(WASH) par create \
		--arch x86_64-linux \
		--binary build/x86_64-linux/httpserver \
		--capid $(CAPABILITY_ID) \
		--name $(NAME) \
		--vendor $(VENDOR) \
		--version $(VERSION) \
		--revision $(REVISION) \
		--destination build/$(NAME).par \
		--compress
	@$(WASH) par insert \
		--arch x86_64-macos \
		--binary build/x86_64-macos/httpserver \
		build/$(NAME).par.gz

build_x86_64_linux:
	GOOS=linux GOARCH=amd64 go build -o build/x86_64-linux/httpserver

build_x86_64_macos:
	GOOS=darwin GOARCH=amd64 go build -o build/x86_64-macos/httpserver

oci_url_base ?= localhost:5000/v2
oci_url      ?= $(oci_url_base)/$(PROJECT):$(VERSION)
ifeq ($(WASH_REG_USER),)
	oci_insecure := --insecure
endif

push: build/httpserver.par.gz
	$(WASH) reg push $(oci_insecure) $(oci_url) $(dest_par)

link_name ?= default

start:
	$(WASH) ctl start provider $(oci_url) \
		--host-id $(shell $(WASH) ctl get hosts -o json | jq -r ".hosts[0].id") \
		--link-name $(link_name) \
		--timeout-ms 4000
