
all: clean update_install-sh tag
	source .env && goreleaser

tag:
	git add -A
	git commit -m "${msg}"
	git push || true
	git tag -a "v${version}" -m "${msg}"
	git push origin "v${version}"

test:
	goreleaser --snapshot --skip-publish --rm-dist

update_install-sh:
	sed -E -i.bak "s/version=\"[0-9]+\.[0-9]+\.[0-9]+\"/version=\"${version}\"/" install.sh
	rm install.sh.bak

clean:
	rm -rf dist
