.PHONY: generate

generate:
	@go run . | pbcopy && pbpaste
