compile:
	git stash -u
	gox -output "build/{{.Dir}}_{{.OS}}_{{.Arch}}"

version:
	git stash -u
	sed -i "s/[[:digit:]]\+\.[[:digit:]]\+\.[[:digit:]]\+/$(v)/g" cmd/version.go
	git add -A
	git commit -m "feat(version) : "$(v)
	git tag v$(v) master

fmt:
	find ! -path "./vendor/*" -name "*.go" -exec go fmt {} \;

checker:
	gometalinter -D gotype --vendor --deadline=240s -e '_string' -j 5 ./...

test-all:
	./test.sh

test-package:
	go test -race -cover -coverprofile=/tmp/yogo github.com/sakib0hasan/yogo/$(pkg)
	go tool cover -html=/tmp/yogo
