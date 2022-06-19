.PHONY: all run build

all: run

run:
    npm run dev & go run .
build:
    npm run build