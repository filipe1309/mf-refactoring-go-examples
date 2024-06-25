# run tests
test:
	@echo "🟢 Running tests..."
	go test -v ./...

# run node
run:
	@echo "🏁 Running code..."
	go run chapter_1/after/after.go

help:
	@echo "📖 Available commands:"
	@echo "  make run"
	@echo "  make test"
	@echo "  make help"
