CURRENT_DIRECTORY := $(dir $(abspath $(firstword $(MAKEFILE_LIST))))

.PHONY: fmt
fmt:
	go fmt ./...
	go mod tidy -compat=1.17

.PHONY: test
test: fmt
	go test ./...

.PHONY: install
install: test
	go install .

guard-%:
	@# See: https://dev.to/daniel13rady/guarding-makefile-targets-1nb8
	@#$(or ${$*}, $(error $* is not set))
