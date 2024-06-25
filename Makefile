# run tests
test:
	@echo "🟢 Running tests..."
	go test -v ./...

# run node
# example: make run CHAPTER_NUM=1
run:
	@echo "🏁 Running code..."
	go run chapter_$(CHAPTER_NUM)/after/after.go

help:
	@echo "📖 Available commands:"
	@echo "  make run"
	@echo "  make test"
	@echo "  make help"
