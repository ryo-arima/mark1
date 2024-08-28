s:
	git add .
	commit-emoji
	git push origin develop

build:
	tool/main.sh build
