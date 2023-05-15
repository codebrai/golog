# https://docs.docker.com/engine/reference/commandline/docker/
# Docker base commands

local.docker.up:
	docker-compose up --detach
local.docker.up.build:
	docker-compose up --build --detach
local.docker.stop:
	docker-compose stop
local.docker.down:
	docker-compose down
local.docker.ps:
	docker-compose ps
local.docker.clean:
	docker-compose down --rmi all --volumes