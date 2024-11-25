# A "Face Builder" Workflow with Temporal.io  

This repository demonstrates a simple project using Temporal.io to run a workflow that processes facial features and integrates with a microservice.  

The project includes:  
1. A Temporal server setup.  
2. A worker that listens on a task queue and executes workflows.  
3. An HTTP microservice with two endpoints.  
4. A workflow that accepts facial features as JSON input.  

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

* GET /hair
    * Query parameters: `eyes, ears,` and `mouth`.
    * If all present, returns `"black hair"`
    * If any missing, responds with `400 Bad Request`
* GET /voice
    * Query parameters: `nose` and `hair`.
    * If all present, returns `"big voice"`
    * If any missing, responds with `400 Bad Request`

---

### 4. Start the Workflow

1. Ensures you have the input JSON file in the root directory. Example `input.json`:

```json
{
    "eyes": "brown",
    "ears": 2,
    "mouth": "wide lips",
    "nose": "roman nose"
}
```

2. In the root directory start the workflow using the Temporal CLI:

```bash
temporal workflow start --type WorkflowOne --task-queue workflow-face --workflow-id workflow-face-id --input-file 'input.json'
```

* `workflow-face`: Task queue name for the worker.
* `workflow-face-id`: The ID for this workflow instance.
* `WorkflowOne`: The name of the workflow to execute

---

#### Project Structure

* `worker`: Contains the worker code ready to execute workflows.
* `microservice`: Contains the HTTP microservice.
* `input.json`: Example JSON input file for the workflow.

---

#### Notes

* This project can be executed by adding more workflows or microservice endpoints.
* For additional guidance, visit the [Temporal Documentation](https://docs.temporal.io)