
all: tag update_install-sh
	source .env && goreleaser

tag:
	git tag -a "${version}" -m "${msg}"
	git push origin ${version}

test:
	goreleaser --snapshot --skip-publish --rm-dist

update_install-sh:
	sed -i.bak "s/version=\"[0-9]\.[0-9]\.[0-9]\"/version=\"${version}\"/" install.sh
	rm install.sh.bak

clean:
	rm -rf dist
