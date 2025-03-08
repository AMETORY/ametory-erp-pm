build-all:
	make build-frontend && make build-backend
build-frontend:
	cd frontend && npm run build

build-backend:
	@echo "Building Go Lambda function"
	@cd backend && gox -os="linux" -arch="amd64" -output="ametory-pm"  


	