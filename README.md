# Asset Tracking System using Hyperledger Fabric

A simple asset tracking system built on Hyperledger Fabric, with:
- Smart contract in Go (`assettrack.go`)
- Node.js REST API (`server.js`)

## Requirements

- Node.js (v14+)
- Go (v1.17+)
- Docker & Docker Compose
- Hyperledger Fabric binaries

## Project Structure

```
├── api/
│   ├── server.js
│   └── connection/gateway.js
└── chaincode/
    └── assettrack/assettrack.go
```

## Setup

1. Start Fabric network:
```bash
./network.sh up createChannel -ca
```

2. Deploy chaincode:
```bash
./network.sh deployCC -ccn assettrack -ccp ../chaincode/assettrack -ccl go
```

3. Start API server:
```bash
cd api
npm install
node server.js
```

## API Endpoints

- `POST /registerAsset` – Register asset
- `GET /getAsset/:id` – Get asset by ID
- `GET /allAssets` – List all assets

## Chaincode Functions

- `CreateAsset`
- `ReadAsset`
- `GetAllAssets`
# SmartFalcon
