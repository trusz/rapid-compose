
run:
	@scripts/make-run.sh	

build:
	@scripts/make-build.sh

install:
	@scripts/make-install.sh

uninstall:
	@rm ~/go/bin/rc

release:
	@scripts/make-release.sh

