A production-ready e-commerce platform built with Go microservices architecture, featuring gRPC for inter-service communication and GraphQL as a unified API gateway. This project demonstrates modern distributed systems design patterns and best practices for scalable backend development.
🏗️ Architecture Overview
mermaidgraph TB
    Client[Client Applications] --> GraphQL[GraphQL Gateway]
    GraphQL --> Account[Account Service]
    GraphQL --> Catalog[Catalog Service]
    GraphQL --> Order[Order Service]
    
    Account --> PG1[(PostgreSQL)]
    Order --> PG2[(PostgreSQL)]
    Catalog --> ES[(Elasticsearch)]
    
    subgraph "gRPC Communication"
        Account
        Catalog
        Order
    end
🎯 Core Principles

Backend For Frontend (BFF) pattern with GraphQL gateway
Database per service for data isolation
Protocol Buffers for strongly-typed service contracts
Event-driven architecture for loose coupling
Containerized deployment for consistency and scalability

🚀 Features
🔐 Account Management

User registration and authentication
Profile management and preferences
Role-based access control
Session management

📦 Product Catalog

Advanced product search with Elasticsearch
Category and inventory management
Real-time product availability
Faceted search and filtering

🛍️ Order Processing

Shopping cart functionality
Order lifecycle management
Payment processing integration
Order history and tracking

🔗 API Gateway

Unified GraphQL endpoint
Query optimization and batching
Request/response transformation
Rate limiting and caching

🛠️ Technology Stack
ComponentTechnologyPurposeLanguageGo 1.21+High-performance backend servicesCommunicationgRPCEfficient inter-service communicationAPI GatewayGraphQL (gqlgen)Unified client interfaceDatabasesPostgreSQL, ElasticsearchData persistence and searchContainerizationDocker, Docker ComposeDeployment and orchestrationMessage FormatProtocol BuffersService contracts and serialization
🏃 Quick Start
Prerequisites

Go 1.21+ installed
Docker and Docker Compose
Protocol Buffers compiler (protoc)
Make (optional, for convenience commands)

Installation

Clone the repository

bashgit clone https://github.com/yourusername/go-grpc-graphql-microservices.git
cd go-grpc-graphql-microservices

Install Go dependencies

bashgo mod download

Generate Protocol Buffer code

bashmake proto-gen
# Or manually:
protoc --go_out=. --go-grpc_out=. proto/*.proto

Start the services

docker-compose up -d


bash# Check service health
curl http://localhost:8080/health

# Access GraphQL Playground
open http://localhost:8080/graphql
📁 Project Structure
├── account/                    # Account microservice
│   ├── cmd/account/           # Service entry point
│   ├── internal/              # Business logic
│   ├── proto/                 # Protocol Buffer definitions
│   └── Dockerfile
├── catalog/                   # Catalog microservice
│   ├── cmd/catalog/
│   ├── internal/
│   ├── proto/
│   └── Dockerfile
├── order/                     # Order microservice
│   ├── cmd/order/
│   ├── internal/
│   ├── proto/
│   └── Dockerfile
├── graphql/                   # GraphQL gateway
│   ├── cmd/gateway/
│   ├── graph/                 # GraphQL schema and resolvers
│   ├── internal/
│   └── Dockerfile
├── shared/                    # Shared utilities and types
├── deployments/               # Deployment configurations
│   ├── docker-compose.yml
│   └── kubernetes/
├── scripts/                   # Build and utility scripts
└── docs/                     # Additional documentation

📊 API Documentation
GraphQL Schema
Access the interactive GraphQL Playground at http://localhost:8080/graphql
Sample Queries
Fetch Products:
graphqlquery GetProducts($limit: Int!) {
  products(limit: $limit) {
    id
    name
    price
    description
    category
    inStock
  }
}
Create Order:
graphqlmutation CreateOrder($input: CreateOrderInput!) {
  createOrder(input: $input) {
    id
    status
    totalAmount
    items {
      productId
      quantity
      price
    }
  }
}
gRPC Services
Each microservice exposes gRPC endpoints defined in Protocol Buffer files:

Account Service: proto/account.proto
Catalog Service: proto/catalog.proto
Order Service: proto/order.proto


GraphQL Gateway: GET /health
Individual Services: gRPC health check protocol


Development Guidelines

Follow Go coding standards and gofmt formatting
Write comprehensive tests with >80% coverage
Update documentation for API changes
Use conventional commits for commit messages


gRPC Team for the excellent RPC framework
GraphQL Foundation for the flexible query language
Go Community for the robust ecosystem
Elasticsearch for powerful search capabilities
