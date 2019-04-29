.PHONY : docker

# Docker Build
docker: Dockerfile
	docker build -t gomod-docker:latest .