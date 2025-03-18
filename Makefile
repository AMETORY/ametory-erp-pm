build-all:
	make build-frontend && make build-backend && make bundle
build-frontend:
	cd frontend && npm run build:staging

build-backend:
	@echo "Building Go Lambda function"
	@cd backend && gox -os="linux" -arch="amd64" -output="ametory-pm"  

bundle:
	@echo "Creating bundle"
	zip ametory-pm.zip -r frontend/package.json frontend/package-lock.json frontend/build backend/ametory-pm templates


deploy-staging:build
	rsync -a ametory-pm ametory@103.172.205.9:/home/ametory/ametory-pm/ametory-pm-$(datetime) -v --stats --progress
	rsync -a template ametory@103.172.205.9:/home/ametory/ametory-pm -v --stats --progress
	ssh ametory@103.172.205.9 "cd /home/ametory/ametory-pm && sudo service ametory-pm stop && sudo unlink ametory-pm && sudo ln -s ametory-pm-$(datetime) ametory-pm && sudo service ametory-pm start"
	make discord-notif stage=STAGING-BE