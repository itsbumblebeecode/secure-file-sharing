# Define the path to the SFS executable using the `which` command
GO=$(shell which go)

# Specify the source code directory
SOURCE_CODE_DIRECTORY=$(shell pwd)/app

# Define the directory where the compiled executable will be placed
DIST_DIRECTORY=$(shell pwd)/dist

# Name of the executable file
EXECUTEABLE_NAME=sfs

# Full path to the executable file
EXECUTEABLE_PATH=$(DIST_DIRECTORY)/$(EXECUTEABLE_NAME)

# Build the SFS application
build:
	$(GO) build -o $(EXECUTEABLE_PATH) $(SOURCE_CODE_DIRECTORY)

# Run the compiled application
run: build
	$(EXECUTEABLE_PATH) $(ARGS)

# Start the development server using npm
dev:
	npm run dev
