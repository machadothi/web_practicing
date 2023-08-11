#!/bin/bash

# Create the directory structure
mkdir -p backend-service/api/endpoints
mkdir -p backend-service/api/middleware
mkdir -p backend-service/cmd/backend-service
mkdir -p backend-service/config
mkdir -p backend-service/db/migrations
mkdir -p backend-service/db/repository
mkdir -p backend-service/docker
mkdir -p backend-service/internal/utils
mkdir -p backend-service/models
mkdir -p backend-service/scripts
mkdir -p backend-service/tests

# Create sample files in each directory
touch backend-service/.gitignore
touch backend-service/README.md
touch backend-service/go.mod
touch backend-service/go.sum

touch backend-service/api/endpoints/user.go
touch backend-service/api/endpoints/product.go
touch backend-service/api/middleware/authentication.go
touch backend-service/api/middleware/logging.go

touch backend-service/cmd/backend-service/main.go

touch backend-service/config/config.yml
touch backend-service/config/dev.config.yml
touch backend-service/config/prod.config.yml

touch backend-service/db/migrations/001_initial_schema.up.sql
touch backend-service/db/migrations/001_initial_schema.down.sql
touch backend-service/db/repository/user.go
touch backend-service/db/repository/product.go

touch backend-service/docker/Dockerfile
touch backend-service/docker/.dockerignore

touch backend-service/internal/utils/date.go
touch backend-service/internal/utils/string.go

touch backend-service/models/user.go
touch backend-service/models/product.go

touch backend-service/scripts/deploy.sh
touch backend-service/scripts/test.sh

touch backend-service/tests/user_test.go
touch backend-service/tests/product_test.go

echo "File tree created successfully!"