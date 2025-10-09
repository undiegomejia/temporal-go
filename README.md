# Temporal.io Workflow Examples in Go

[![Go Version](https://img.shields.io/badge/Go-1.21+-blue.svg)](https://golang.org/dl/)
[![Temporal SDK](https://img.shields.io/badge/Temporal-1.26.0-orange.svg)](https://docs.temporal.io/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)

A comprehensive collection of Temporal.io workflow examples demonstrating different patterns and use cases using the Go SDK. This project showcases sequential workflows, parallel execution, and Domain Specific Language (DSL) implementations.

> **Attribution**: This repository builds upon examples from the [Temporal Samples Go](https://github.com/temporalio/samples-go) repository created by the Temporal.io team.

## 🏗️ Architecture Overview

The project contains three distinct workflow examples, each demonstrating different Temporal patterns:

```
temporal-go/
├── dsl/                    # DSL-based workflow configuration
├── face/                   # Sequential workflow example
├── shipping-order/         # Parallel workflow execution
├── go.mod                  # Go module definition
└── README.md              # This file
```

### Core Components

Each workflow example follows a consistent architecture:

- **Workflow Definition**: Business logic orchestration
- **Activities**: Individual task implementations  
- **Worker**: Process that executes workflows and activities
- **Starter**: Client to initiate workflow execution
- **HTTP Service**: Mock external service integrations
- **Types**: Data structures and interfaces

## 📋 Prerequisites

Ensure your development environment includes:

- **Go 1.21+**: [Download Go](https://go.dev/dl/)
- **Temporal CLI**: [Installation Guide](https://docs.temporal.io/cli/install)
- **Git**: For cloning the repository

## 🚀 Quick Start

### 1. Clone and Setup

```bash
git clone <repository-url>
cd temporal-go
go mod download
```

### 2. Start Temporal Server

```bash
temporal server start-dev
```

The server will be available at `http://localhost:8233` with the Web UI.

### 3. Choose Your Example

Select one of the three workflow examples to run:

## 📁 Workflow Examples

### 1. DSL (Domain Specific Language) Workflow

**Location**: `dsl/`  
**Pattern**: YAML-configured workflow execution

The DSL example demonstrates how to define workflows using YAML configuration files, enabling non-developers to create and modify business processes.

#### Features
- YAML-based workflow definition
- Support for sequential and parallel execution  
- Dynamic activity binding
- Variable interpolation

#### Running DSL Example

**Terminal 1 - Start Worker**:
```bash
go run dsl/worker/main.go
```

**Terminal 2 - Start HTTP Service**:
```bash
go run shipping-order/http/main.go
```

**Terminal 3 - Execute Workflow**:
```bash
# Sequential workflow
go run dsl/starter/main.go

# Parallel workflow  
go run dsl/starter/main.go -dslConfig=dsl/workflow2.yaml
```

#### DSL Configuration Format

```yaml
variables:
  orderId: order-poc6r-dki80t-kot8t-tju56
  customerId: cust-bdc6r-odk0it-ahj8t-pwh56
  
root:
  sequence:
    elements:
     - activity:
        name: CheckInventory
        arguments:
          - orderId
          - quantity
     - parallel:
          branches:
            - sequence:
                elements:
                 - activity:
                    name: PrepareShipping
```

### 2. Face Processing Workflow

**Location**: `face/`  
**Pattern**: Sequential workflow with HTTP service integration

Demonstrates a basic sequential workflow that processes facial features through external HTTP services.

#### Features
- Sequential activity execution
- HTTP service integration
- JSON data processing
- Error handling and retries

#### Running Face Example

**Terminal 1 - Start Worker**:
```bash
go run face/worker/main.go
```

**Terminal 2 - Start HTTP Service**:
```bash
go run face/http/main.go
```

**Terminal 3 - Execute Workflow**:
```bash
go run face/start/main.go
```

#### Data Structure

```go
type FaceType struct {
    Eyes  string `json:"eyes"`
    Ears  int    `json:"ears"`
    Mouth string `json:"mouth"`
    Nose  string `json:"nose"`
    Hair  string `json:"hair"`
    Voice string `json:"voice"`
}
```

### 3. Shipping Order Workflow

**Location**: `shipping-order/`  
**Pattern**: Mixed sequential and parallel execution

Showcases an order processing workflow combining sequential validation steps with parallel fulfillment activities.

#### Features
- Sequential inventory checking
- Parallel shipping preparation and invoice generation
- Comprehensive error handling
- Real-world business process modeling

#### Running Shipping Order Example

**Terminal 1 - Start Worker**:
```bash
go run shipping-order/worker/main.go
```

**Terminal 2 - Start HTTP Service**:
```bash
go run shipping-order/http/main.go
```

**Terminal 3 - Execute Workflow**:
```bash
go run shipping-order/starter/main.go
```

#### Workflow Execution Flow

1. **Sequential**: Check inventory availability
2. **Sequential**: Send order confirmation  
3. **Parallel**: Prepare shipping + Generate invoice
4. **Sequential**: Notify shipment completion

## 🔧 Configuration

### Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `TEMPORAL_HOST_PORT` | Temporal server address | `localhost:7233` |
| `HTTP_SERVICE_PORT` | Mock service port | `9999` |

### Temporal Configuration

The workflows use standard Temporal configuration:

```go
ActivityOptions{
    ScheduleToCloseTimeout: time.Hour * 2,
    StartToCloseTimeout:    time.Hour * 600,
    HeartbeatTimeout:      time.Second * 30,
}
```

## 🧪 Testing

### Unit Tests

```bash
go test ./...
```

### Integration Tests

Ensure Temporal server is running:

```bash
temporal server start-dev --headless
go test ./... -tags=integration
```

## 📊 Monitoring

Access the Temporal Web UI at `http://localhost:8233` to:

- Monitor workflow execution
- View workflow history
- Debug failed activities
- Analyze performance metrics

## 🔍 Troubleshooting

### Common Issues

**Worker Connection Failed**
```
Error: Unable to create Temporal client
```
**Solution**: Ensure Temporal server is running on `localhost:7233`

**HTTP Service Unavailable** 
```
Error: connection refused localhost:9999
```
**Solution**: Start the corresponding HTTP service for your example

**Activity Timeout**
```
Error: activity timeout
```
**Solution**: Check HTTP service responsiveness and timeout configuration

## 🛠️ Development

### Adding New Workflows

1. Create new directory under project root
2. Implement required components:
   - `workflow.go` - Business logic
   - `activities.go` - Task implementations  
   - `types.go` - Data structures
   - `worker/main.go` - Worker process
   - `starter/main.go` - Client starter
   - `http/main.go` - Mock services (optional)

### Code Structure Guidelines

```go
// Workflow functions should follow this pattern
func MyWorkflow(ctx workflow.Context, input MyInput) (MyOutput, error) {
    // Configure activity options
    options := workflow.ActivityOptions{
        StartToCloseTimeout: time.Minute * 10,
    }
    ctx = workflow.WithActivityOptions(ctx, options)
    
    // Execute activities
    var result MyResult
    err := workflow.ExecuteActivity(ctx, MyActivity, input).Get(ctx, &result)
    
    return MyOutput{Result: result}, err
}
```

## 📚 Additional Resources

- [Temporal Documentation](https://docs.temporal.io/)
- [Go SDK Reference](https://pkg.go.dev/go.temporal.io/sdk)
- [Temporal Samples Repository](https://github.com/temporalio/samples-go)
- [Workflow Patterns](https://docs.temporal.io/workflows)

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Implement changes with tests
4. Submit a pull request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
