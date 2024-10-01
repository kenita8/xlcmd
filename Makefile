# Copyright 2024 kenita8
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#	http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

COMMANDS := $(notdir $(wildcard cmd/*))
VERSION := $(shell cat VERSION)

.PHONY: build
build: $(COMMANDS)

.PHONY: $(COMMANDS)
$(COMMANDS):
	GOOS=linux GOARCH=amd64 EXT= COMMAND=$@ VERSION=$(VERSION) make buildsub
	GOOS=windows GOARCH=amd64 EXT=.exe COMMAND=$@ VERSION=$(VERSION) make buildsub

.PHONY: buildsub
buildsub:
	go build -ldflags "-X main.Version=$(VERSION)" -o ./bin/$(GOOS)/$(COMMAND)$(EXT) ./cmd/$(COMMAND)

.PHONY: xlcmd
xlcmd: notice clean build
	mkdir dst | true
	VERSION=$(VERSION) cd bin/windows; zip -r ../../dst/xlcmd-$(VERSION).windows-amd64.zip *
	VERSION=$(VERSION) zip -r ./dst/xlcmd-$(VERSION).windows-amd64.zip LICENSE NOTICE examples
	chmod u+x examples/example_linux.sh
	VERSION=$(VERSION) cd bin/linux; tar --numeric-owner --owner=0 --group=0 -cvf ../../dst/xlcmd-$(VERSION).linux-amd64.tar *
	VERSION=$(VERSION) tar --numeric-owner --owner=0 --group=0 -rvf ./dst/xlcmd-$(VERSION).linux-amd64.tar LICENSE NOTICE examples
	VERSION=$(VERSION) gzip ./dst/xlcmd-$(VERSION).linux-amd64.tar

.PHONY: clean
clean:
	rm -rf bin/*
	rm -rf dst/xlcmd*

.PHONY: notice
notice:
	gocredits > NOTICE
