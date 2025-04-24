build:
	docker buildx build --platform=amd64 -t edr3x/go-chi-template:latest --build-arg TARGETOS=linux --build-arg TARGETARCH=amd64 . --load
