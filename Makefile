BINARY_NAME=app
.PHONY: push

push:
	git push -uf origin main
	git push -uf gadam main
