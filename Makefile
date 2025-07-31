.PHONY: modules repo fetch flatten

REPO_URL := https://github.com/github/gitignore
BRANCH := main

EMBED_DIR := source
DEST_DIR := source/repo

repo: fetch flatten

# fetch and unzip repo contents
fetch:
	@echo "Downloading $(REPO_URL) at branch $(BRANCH)..."
	rm -rf $(DEST_DIR)
	mkdir -p $(DEST_DIR)
	curl -L $(REPO_URL)/archive/refs/heads/$(BRANCH).zip -o /tmp/repo.zip
	unzip -oq /tmp/repo.zip -d $(DEST_DIR)
	rm /tmp/repo.zip

# traverse repo contents and "flatten" symlinks
flatten:
	@echo "Flattening symlinks in $(DEST_DIR)..."
	@find $(DEST_DIR) -type l | while read symlink; do \
		target=$$(readlink -f $$symlink); \
		if [ -f "$$target" ]; then \
			echo "Flattening symlink: $$symlink -> $$target"; \
			rm "$$symlink"; \
			cp "$$target" "$$symlink"; \
		else \
			echo "Skipping non-file symlink: $$symlink -> $$target"; \
		fi \
	done

modules:
	go mod tidy
	go mod vendor

