pkgs      = $(shell go list ./... | grep -v /tests | grep -v /vendor/ | grep -v /common/)
datetime	= $(shell date +%s)
datetimeFormat	= $(shell date +"%Y-%m-%d %H:%M:%S")

build-all:
	make build-frontend && make build-backend
build-frontend:
	cd frontend && npm run build:staging

build-backend:
	@echo "Building Go Lambda function"
	@cd backend && gox -os="linux" -arch="amd64" -output="ametory-pm"  

bundle:
	@echo "Creating bundle"
	zip ametory-pm.zip -r frontend/package.json frontend/package-lock.json frontend/build backend/ametory-pm templates


deploy-staging:
	rsync -a backend/ametory-pm ametory@103.172.205.9:/home/ametory/ametory-pm/backend/ametory-pm-$(datetime) -v --stats --progress
	rsync -a templates ametory@103.172.205.9:/home/ametory/ametory-pm -v --stats --progress
	rsync -a frontend/build ametory@103.172.205.9:/home/ametory/ametory-pm/frontend -v --stats --progress
	ssh ametory@103.172.205.9 "cd /home/ametory/ametory-pm/backend && sudo service project stop && sudo unlink ametory-pm && sudo ln -s ametory-pm-$(datetime) ametory-pm && sudo service project start"