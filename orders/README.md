# Order Workflow with Parallel Activities in Temporal.io  

This project demonstrates a simple project using Temporal.io to run a workflow that processes a product order and integrates with a microservice.  

The project includes:  
1. A Temporal server setup.  
2. A worker that listens on a task queue and executes workflows.  
3. An HTTP microservice with five endpoints.  

---

## Prerequisites  

Ensure you have the following installed:  

- [Go](https://go.dev/) (for building and running the worker and microservice).  
- [Temporal CLI](https://docs.temporal.io/cli) (to interact with Temporal).  

---

## Steps  

### 1. Start the Temporal Server  

Run the Temporal server:  

```bash
temporal server start-dev
```

---

### 2. Run the Worker

The worker is located in the `worker` directory. It is ready to execute workflows when started.

1. Navigate to the `worker` directory:

```bash
cd worker
```

2. Start the worker by running the `main.go` file:

```bash
go run main.go
```

---

### 3. Run the Microservice

The microservice is located in the `microservice` directory. It provides two HTTP endpoints fo facial feature processing.

1. Navigate to the `microservice` directory:

```bash
cd ../microservice
```

2. Start the microservice by running the `main.go` file:

```bash
go run main.go
```

#### Microservice Endopints

* GET /check-inventory
    * Query parameters: `orderId,` and `quantity`.
    * If all present, returns `200 Status Ok, true`
    * If any missing, responds with `400 Bad Request`
* GET /send-confirmation
    * Query parameters: `orderId` and `customerId`.
    * If all present, returns `200 Status Ok, true`
    * If any missing, responds with `400 Bad Request`
* GET /prepare-shipping
    * Query parameters: `orderId` and `shippingAddress`.
    * If all present, returns `200 Status Ok, true`
    * If any missing, responds with `400 Bad Request`
* GET /generate-invoice
    * Query parameters: `orderId`.
    * If all present, returns `200 Status Ok, invoceId`
    * If any missing, responds with `400 Bad Request`
* GET /notify-shipment
    * Query parameters: `orderId` and `invoiceId`.
    * If all present, returns `200 Status Ok, invoceId`
    * If any missing, responds with `400 Bad Request`

---

### 4. Start the Workflow

1. Navigate to the `starter` directory:

```bash
cd ../starter
```

2. Start the workflow starter by running the `main.go` file:

```bash
go run main.go
```

#### Project Structure

* `worker`: Contains the worker code ready to execute workflows.
* `microservice`: Contains the HTTP microservice.
* `starter`: Contains the starter code ready to execute.

---

#### Notes

* This project can be executed by adding more workflows or microservice endpoints.
* For additional guidance, visit the [Temporal Documentation](https://docs.temporal.io)
