
all: tag
	goreleaser

tag:
	git add -A
	git commit -m "${msg}"
	git push
	git tag -a "v${version}" -m "${msg}"
	git push origin v${version}

test:
	goreleaser --snapshot --skip-publish --rm-dist
