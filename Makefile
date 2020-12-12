build:
		docker-compose build app

# test: TODO

development:
		docker-compose up -d app

restart:
		docker-compose stop app
		docker-compose up -d app

clean:
		docker-compose down

rebuild:
		docker-compose rm
		docker-compose build app
