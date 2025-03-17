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


	