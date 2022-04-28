test:
	go test -cover ./...

install:
	go install
	dpm completion zsh > ${HOME}/.local/share/zinit/completions/_dpm
