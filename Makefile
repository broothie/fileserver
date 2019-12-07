
all: clean tag release

release:
	source .env && goreleaser

tag:
	git add -A
	git commit -m "${m}"
	git push || true
	git tag -a "v${v}" -m "${msg}"
	git push origin "v${v}"

tag.latest:
	git tag | cat

test:
	goreleaser --snapshot --skip-publish --rm-dist

clean:
	rm -rf dist
