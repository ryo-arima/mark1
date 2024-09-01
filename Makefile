s:
	git add .
	commit-emoji
	git push origin develop

build-bin:
	tool/main.sh build

build-deb:
	tool/deb/build.sh run

build-rpm:
	tool/rpm/build.sh run

build-deb-local:
	./tool/main.sh build-deb-local

build-rpm-local:
	./tool/main.sh build-rpm-local
