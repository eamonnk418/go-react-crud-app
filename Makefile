.PHONY: docker-frontend
docker-frontend:
	docker build -t go-react-crud-app-frontend -f frontend/Dockerfile frontend
