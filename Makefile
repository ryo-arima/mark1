s:
	git add .
	commit-emoji
	git push origin develop

build-bin:
	tool/main.sh build

build-deb:
	tool/main.sh build-deb

push-deb:
	tool/main.sh push-deb

build-rpm:
	tool/main.sh build-rpm

push-rpm:
	tool/main.sh push-rpm

build-deb-local:
	./tool/main.sh build-deb-local

build-rpm-local:
	./tool/main.sh build-rpm-local

build-container-up:
	./tool/main.sh build-container-up

update-readme:
	./tool/main.sh update-readme
