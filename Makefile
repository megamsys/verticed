#Copyright (c) 2014 Megam Systems.
#
#   Licensed under the Apache License, Version 2.0 (the "License");
#   you may not use this file except in compliance with the License.
#   You may obtain a copy of the License at
#
#       http://www.apache.org/licenses/LICENSE-2.0
#
#   Unless required by applicable law or agreed to in writing, software
#   distributed under the License is distributed on an "AS IS" BASIS,
#   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
#   See the License for the specific language governing permissions and
#   limitations under the License.
###############################################################################
# Makefile to compile cib.
# lists all the dependencies for test, prod and we can run a go build aftermath.
###############################################################################

GOPATH  := $(GOPATH):$(shell pwd)/../../../../


define HG_ERROR

FATAL: you need mercurial (hg) to download megamd dependencies.
       Check README.md for details


endef

define GIT_ERROR

FATAL: you need git to download megamd dependencies.
       Check README.md for details
endef

define BZR_ERROR

FATAL: you need bazaar (bzr) to download megamd dependencies.
       Check README.md for details
endef

.PHONY: all check-path get hg git bzr get-code test

all: check-path get test

build: check-path get _go_test _megamd

# It does not support GOPATH with multiple paths.
check-path:
ifndef GOPATH
	@echo "FATAL: you must declare GOPATH environment variable, for more"
	@echo "       details, please check README.md file and/or"
	@echo "       http://golang.org/cmd/go/#GOPATH_environment_variable"
	@exit 1
endif
	@exit 0

get: hg git bzr get-code godep

hg:
	$(if $(shell hg), , $(error $(HG_ERROR)))

git:
	$(if $(shell git), , $(error $(GIT_ERROR)))

bzr:
	$(if $(shell bzr), , $(error $(BZR_ERROR)))

get-code:
	go get $(GO_EXTRAFLAGS) -u -d -t -insecure ./...

godep:
	go get $(GO_EXTRAFLAGS) github.com/tools/godep
	godep restore ./...

_go_test:
	go clean $(GO_EXTRAFLAGS) ./...
	go test $(GO_EXTRAFLAGS) ./...

_megamd:
	rm -f megamd
	go build $(GO_EXTRAFLAGS) -o megamd ./cmd/megamd


_megamdr:
	./megamd start
	rm -f megamd

_sh_tests:
	@conf/trusty/megam/megam_test.sh

test: _go_test _megamd _megamdr

_install_deadcode: git
	go get $(GO_EXTRAFLAGS) github.com/remyoudompheng/go-misc/deadcode

deadcode: _install_deadcode
	@go list ./... | sed -e 's;github.com/megamsys/megamd/;;' | xargs deadcode


# do_build builds the code. The version and commit must be passed in.
#do_build()
#    for b in ${BINS[*]}; do
#        rm -f $GOPATH_INSTALL/bin/$b
#    done
#    go get -u -f -d ./...
#    if [ $? -ne 0 ]; then
#        echo "WARNING: failed to 'go get' packages."
#    fi
#    git checkout $TARGET_BRANCH # go get switches to master, so ensure we're back.
#    version=$1
#    commit=`git rev-parse HEAD`
#    branch=`current_branch`
#    if [ $? -ne 0 ]; then
#        echo "Unable to retrieve current commit -- aborting"
#        cleanup_exit 1
#    fi
#    go install -a -ldflags="-X main.version $version -X main.branch $branch -X main.commit $commit" ./...
#    if [ $? -ne 0 ]; then
#        echo "Build failed, unable to create package -- aborting"
#        cleanup_exit 1
#    fi
#    echo "Build completed successfully."
#}
