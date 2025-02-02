Image_Name = go-api-image
Container_Name = go-api

Flag_Mode = -it
Flag_Port = -p 8080:8080

all: build up

build:
	docker-compose build
.PHONY: build

up:
	docker-compose up
.PHONY: up

down:
	docker-compose down
.PHONY: down

start:
	docker start go-api
.PHONY: start

stop:
	docker stop go-api
.PHONY: stop

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

comp:
	go build -o ./bin64/DBMS ./src/app

run:
	@./bin64/DBMS

delete:
	docker rmi -f $(docker images -aq)
.PHONY: delete
