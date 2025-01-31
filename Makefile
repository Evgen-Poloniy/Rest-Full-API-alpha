Image_Name = go-api-image
Container_Name = go-api

Flag_Mode = -it
Flag_Port = -p 8080:8080

all: build up

build:
	docker-compose build
.PHONY: dcbuild

up:
	docker-compose up
.PHONY: up

down:
	docker-compose down
.PHONY: down

list:
	@echo "*------------------------------------------------------------------------------------------------------*" 
	@docker images 
	@echo "*------------------------------------------------------------------------------------------------------*" 
	@echo ""
	@echo "*------------------------------------------------------------------------------------------------------*" 
	@docker ps -a 
	@echo "*------------------------------------------------------------------------------------------------------*" 
.PHONY: list

clean:
	docker rm $(Container_Name)
	docker rmi $(Image_Name)
.PHONY: clean

app:
	go build -o app ./src/app

delete:
	docker rmi -f $(docker images -aq) 