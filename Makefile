
all: tag
	goreleaser

tag:
	git commit -am "${msg}"
	git push
	git tag -a "v${version}" -m "${msg}"
	git push origin v${version}

test:
	goreleaser --snapshot --skip-publish --rm-dist
