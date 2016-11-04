#!/bin/bash
#########################################################################
# File Name: travis-build.sh
# Author: Fan Yang
# mail: missdeer@dfordsoft.com
# Created Time: 六  1/10 10:30:41 2015
#########################################################################

set -e

SUDO=sudo
GO_VER=go1.7.3
GO_TAR=${GO_VER}.linux-amd64.tar.gz
GO_URL="https://storage.googleapis.com/golang/${GO_TAR}"

: ${GITHUB_REPO:="missdeer/ifconfig"}
GITHUB_HOST="github.com"
GITHUB_CLONE="git://${GITHUB_HOST}/${GITHUB_REPO}"
GITHUB_URL="https://${GITHUB_HOST}/${GITHUB_PUSH-${GITHUB_REPO}}"

# if not set, ignore password
#GIT_ASKPASS="${TRAVIS_BUILD_DIR}/gh_ignore_askpass.sh"

skip() {
	echo "$@" 1>&2
	echo "Exiting..." 1>&2
	exit 0
}

abort() {
	echo "Error: $@" 1>&2
	echo "Exiting..." 1>&2
	exit 1
}

# TRAVIS_BUILD_DIR not set, exiting
[ -d "${TRAVIS_BUILD_DIR-/nonexistent}" ] || \
	abort '${TRAVIS_BUILD_DIR} not set or nonexistent.'

# check for pull-requests
[ "${TRAVIS_PULL_REQUEST}" = "false" ] || \
	skip "Not running Doxygen for pull-requests."

# check for branch name
[ "${TRAVIS_BRANCH}" = "master" ] || \
	skip "Running Doxygen only for updates on 'master' branch (current: ${TRAVIS_BRANCH})."

# check for job number
[ "${TRAVIS_JOB_NUMBER}" = "${TRAVIS_BUILD_NUMBER}.1" ] || \
	skip "Running Doxygen only on first job of build ${TRAVIS_BUILD_NUMBER} (current: ${TRAVIS_JOB_NUMBER})."

# install doxygen binary distribution
go_install()
{
	wget -O - "${GO_URL}" | \
		tar xz -C ${TMPDIR-/tmp} 
    export GOROOT=${TMPDIR-/tmp}/go
    export PATH=${TMPDIR-/tmp}/go/bin/:$PATH
}

go_run()
{
	cd "${TRAVIS_BUILD_DIR}";
    GOARCH=amd64 GOOS=linux go build -ldflags="-s" -o prebuild/ifconfig-linux-amd64
    GOARCH=386 GOOS=linux go build -ldflags="-s" -o prebuild/ifconfig-linux-386
    GOARCH=arm GOOS=linux go build -ldflags="-s" -o prebuild/ifconfig-linux-arm
    GOARCH=arm64 GOOS=linux go build -ldflags="-s" -o prebuild/ifconfig-linux-arm64
    GOARCH=ppc64 GOOS=linux go build -ldflags="-s" -o prebuild/ifconfig-linux-ppc64
    GOARCH=ppc64le GOOS=linux go build -ldflags="-s" -o prebuild/ifconfig-linux-ppc64le
    GOARCH=amd64 GOOS=dragonfly go build -ldflags="-s" -o prebuild/ifconfig-dragonfly-amd64
    GOARCH=amd64 GOOS=freebsd go build -ldflags="-s" -o prebuild/ifconfig-freebsd-amd64
    GOARCH=386 GOOS=freebsd go build -ldflags="-s" -o prebuild/ifconfig-freebsd-amd64
    GOARCH=arm GOOS=freebsd go build -ldflags="-s" -o prebuild/ifconfig-freebsd-amd64
    GOARCH=amd64 GOOS=netbsd go build -ldflags="-s" -o prebuild/ifconfig-netbsd-amd64
    GOARCH=386 GOOS=netbsd go build -ldflags="-s" -o prebuild/ifconfig-netbsd-amd64
    GOARCH=arm GOOS=netbsd go build -ldflags="-s" -o prebuild/ifconfig-netbsd-amd64
    GOARCH=amd64 GOOS=openbsd go build -ldflags="-s" -o prebuild/ifconfig-openbsd-amd64
    GOARCH=386 GOOS=openbsd go build -ldflags="-s" -o prebuild/ifconfig-openbsd-amd64
    GOARCH=arm GOOS=openbsd go build -ldflags="-s" -o prebuild/ifconfig-openbsd-amd64
    GOARCH=amd64 GOOS=darwin go build -o prebuild/ifconfig-darwin-amd64
    GOARCH=386 GOOS=darwin go build -o prebuild/ifconfig-darwin-amd64
    GOARCH=amd64 GOOS=windows go build -o prebuild/ifconfig-windows-amd64
    GOARCH=386 GOOS=windows go build -o prebuild/ifconfig-windows-386
}

prebuilt_prepare()
{
	cd "${TRAVIS_BUILD_DIR}";
	git --version
	git clone --single-branch -b prebuilt "${GITHUB_CLONE}" prebuilt
	cd prebuilt
	# setup git config (with defaults)
	git config user.name "${GIT_NAME}"
	git config user.email "${GIT_EMAIL}"
	# clean working dir
	rm -f .git/index
	git clean -df
}

prebuilt_commit() {
	cd "${TRAVIS_BUILD_DIR}/prebuilt";
    pwd
    ls
	git add --all;
    git add -f ./ifconfig-*
    git commit -m "Automatic pre build by travis at $(date)";
}

gh_setup_askpass() {
	cat > ${GIT_ASKPASS} <<EOF
#!/bin/bash
echo
exit 0
EOF
	chmod a+x "$GIT_ASKPASS"
}

prebuilt_push() {
	# check for secure variables
	[ "${TRAVIS_SECURE_ENV_VARS}" = "true" ] || \
		skip "Secure variables not available, not updating prebuilt branch."
	# check for GitHub access token
	[ "${GH_TOKEN+set}" = set ] || \
		skip "GitHub access token not available, not updating prebuilt branch."
	[ "${#GH_TOKEN}" -eq 40 ] || \
		abort "GitHub token invalid: found ${#GH_TOKEN} characters, expected 40."

	cd "${TRAVIS_BUILD_DIR}/prebuilt";
	# setup credentials (hide in "set -x" mode)
	git remote set-url --push origin "${GITHUB_URL}"
	git config credential.helper 'store'
	# ( set +x ; git config credential.username "${GH_TOKEN}" )
	( set +x ; [ -f ${HOME}/.git-credentials ] || \
			( echo "https://${GH_TOKEN}:@${GITHUB_HOST}" > ${HOME}/.git-credentials ; \
			 chmod go-rw ${HOME}/.git-credentials ) )
	# push to GitHub
	git push origin prebuilt
}

go_install
prebuilt_prepare
go_run
prebuilt_commit
prebuilt_push

