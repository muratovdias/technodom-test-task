IMAGE_NAME=test
CONTAINER_NAME=test-container
PORT=8787

build:
	sudo docker build -t $(IMAGE_NAME) .

run-container:
	sudo docker run --rm -it --name $(CONTAINER_NAME) -p $(PORT):$(PORT) $(IMAGE_NAME)
