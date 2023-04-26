GO			?= go
GOOS		?= $(shell $(GO) env GOOS)
GOARCH	?= $(shell $(GO) env GOARCH)

all: build

.PHONY: test
test:
	go version

.PHONY: get
get: test
	go mod download

.PHONY: build
build: get
	mkdir -p .build/$(GOOS)-$(GOARCH)/
	GOOS=$(GOOS) GOARCH=$(GOARCH) $(GO) build
	if [ $(GOOS) = "windows" ] ; then \
		mv ./log_total.exe ./.build/$(GOOS)-$(GOARCH)/ ; \
	else \
		mv ./log_total ./.build/$(GOOS)-$(GOARCH)/ ; \
	fi ; \

TAG	?=	$(shell git tag | tail -n1)
.PHONY: package
package:
	mkdir -p ./packages/$(TAG)/log_total-$(TAG).$(GOOS)-$(GOARCH)
	cp -r CREDITS \
		Makefile \
		LICENSE \
		README.md \
		./packages/$(TAG)/log_total-$(TAG).$(GOOS)-$(GOARCH)
	if [ $(GOOS) = "windows" ] ; then \
		cp ./.build/$(GOOS)-$(GOARCH)/log_total.exe ./packages/$(TAG)/log_total-$(TAG).$(GOOS)-$(GOARCH) ; \
	else \
		cp ./.build/$(GOOS)-$(GOARCH)/log_total ./packages/$(TAG)/log_total-$(TAG).$(GOOS)-$(GOARCH) ; \
	fi
	cd ./packages/$(TAG) ; \
	if [ $(GOOS) = "windows" ] ; then \
		zip -r log_total-$(TAG).$(GOOS)-$(GOARCH).zip ./log_total-$(TAG).$(GOOS)-$(GOARCH) ; \
	else \
		tar cvf log_total-$(TAG).$(GOOS)-$(GOARCH).tar.gz ./log_total-$(TAG).$(GOOS)-$(GOARCH) ; \
	fi ; \
	rm -r ./log_total-$(TAG).$(GOOS)-$(GOARCH)

.PHONY: package-all-with-build
package-all-with-build: get
	$(GO) tool dist list | grep 'aix\|darwin\|freebsd\|illumos\|linux\|netbsd\|openbsd\|windows' | while read line ; \
	do \
		printf GOOS= > ./.build.env ; \
		echo $$line | cut -f 1 -d "/" >> ./.build.env ; \
		printf GOARCH= >> ./.build.env ; \
		echo $$line | cut -f 2 -d "/" >> ./.build.env ; \
		. ./.build.env ; \
		make build GOOS=$$GOOS GOARCH=$$GOARCH ; \
		make package GOOS=$$GOOS GOARCH=$$GOARCH ; \
	done
	rm ./.build.env

.PHONY: clean
clean:
	-rm -r ./.build ./packages ./.build.env
