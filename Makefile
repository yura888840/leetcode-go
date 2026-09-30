.PHONY: test one fmt vet

test:
	go test ./... -count=1 -timeout=30s

# Example: make one P=p0001_two_sum
one:
	@test -n "$(P)" || (echo 'Usage: make one P=p0001_two_sum'; exit 1)
	go test ./problems/$(P) -v -count=1 -timeout=30s

fmt:
	go fmt ./...

vet:
	go vet ./...

play:
	go test ./problems/$(P) -run '^TestPlayground$$' -v -count=1