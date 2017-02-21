help:    ## Show this help.
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'
start:   ## Build and start all Docker containers
	. config/environments/development
	docker-compose up
restart: ## Stop, rebuild, and start all Docker containers
	. config/environments/development
	docker-compose stop
	docker-compose rm -f
	docker-compose up
clean:   ## Destroy all containers and the development database data
	docker-compose rm -f
	rm -rf tmp/data
