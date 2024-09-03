s:
	git add .
	commit-emoji
	git push origin develop

build-bin:
	tool/main.sh build

build-deb:
	tool/main.sh build-deb

build-rpm:
	tool/main.sh build-rpm

build-deb-local:
	./tool/main.sh build-deb-local

build-rpm-local:
	./tool/main.sh build-rpm-local

build-container-up:
	./tool/main.sh build-container-up
