# example: make run CHAPTER_NUM=4
test:
	@echo "🟢 Running chapter tests"
	@if [ -z "$(CHAPTER_NUM)" ]; then \
		CHAPTER_NUM=$$(ls -d sec-* | tail -n 1 | cut -d '-' -f 2); \
		echo "🔍 No chapter provided, running last chapter: $$CHAPTER_NUM"; \
		go test -v ./chapter-$$CHAPTER_NUM/...; \
	else \
		echo "🔍 Running chapter: $(CHAPTER_NUM)"; \
		go test -v ./chapter-$(CHAPTER_NUM)/...; \
	fi

# run all tests
test-all:
	@echo "🟢 Running tests..."
	go test -v ./...

# example: make run CHAPTER_NUM=1
# if chapter has after folder, run the code in that folder
run:
	@echo "🏁 Running code..."
	@if [ -d "chapter-$(CHAPTER_NUM)/after" ]; then \
		echo "has after folder"; \
		go run $(shell ls -1 chapter-$(CHAPTER_NUM)/after/*.go | grep -v _test.go); \
	else \
		go run $(shell ls -1 chapter-$(CHAPTER_NUM)/*.go | grep -v _test.go); \
	fi


help:
	@echo "📖 Available commands:"
	@echo "  make run CHAPTER_NUM=1"
	@echo "  make test CHAPTER_NUM=1"
	@echo "  make test-all"
	@echo "  make help"
