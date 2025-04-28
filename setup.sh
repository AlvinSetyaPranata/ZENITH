#!/bin/bash

# Exit immediately if a command exits with a non-zero status
set -e

echo "Starting project setup..."

# Step 1: Go into the docker directory and start Elasticsearch and Kibana
echo "Starting Docker containers for Elasticsearch and Kibana..."

docker-compose up -d
cd ..

# Step 2: Backend Setup
echo "Setting up Go backend..."
cd backend

# (optional) Initialize go mod if needed
# go mod init your-module-name

# (optional) Download Go modules
go mod tidy

cd ..

# Step 3: Frontend Setup
echo "Setting up frontend..."
cd frontend

# Install frontend dependencies (if you use npm)
npm install

cd ..

echo "Setup complete!"
