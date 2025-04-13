# To-Do Planner

A Go application that collects tasks from multiple providers and schedules them among five developers, aiming to complete all tasks in the fewest total weeks.

## Features

1. **List Tasks / Providers**  
   View tasks and providers via a command-line interface (CLI) or REST API.  
2. **Add Provider**  
   Register a new provider with its API URL from the CLI or through an HTTP endpoint.  
3. **Schedule**  
   Gather tasks from both providers, then assign them to developers based on capacity.  
   - Each developer works 45 hours/week.  
   - Developers have different capacity factors (e.g., DEV5 can handle 5x as much in an hour vs DEV1).  
   - Outputs how many **weeks** total are needed to finish all tasks.

---

## Running the API

1. **Build** the HTTP server:

   ```bash
   go build -o bin/api ./cmd/api
   ```

2. **Run** the server:

   ```bash
   ./bin/api
   ```
   By default, it listens on port `:3000`. You can verify it’s running with:

   ```bash
   curl http://localhost:3000/
   ```

3. **Interact with the API**:
   - `GET /tasks` – List all tasks.
   - `GET /providers` – List all providers.
   - `POST /providers` – Add a provider (JSON body).
   - `GET /schedule` – Generate and retrieve the task schedule.

---

## Running the CLI

1. **Build** the CLI:

   ```bash
   go build -o bin/cli ./cmd/cli
   ```

2. **Show available commands** (help menu):

   ```bash
   ./bin/cli
   ```

3. **Examples**:
   - **List Tasks**:
     ```bash
     ./bin/cli list tasks
     ```
   - **List Providers**:
     ```bash
     ./bin/cli list providers
     ```
   - **Add a Provider** (JSON data):
     ```bash
     ./bin/cli add provider --json='{"name":"ProviderX","api_url":"/api/provider-x"}'
     ```
   - **Load Tasks**:
     ```bash
     ./bin/cli load task
     ```

---

## Quick Start / Testing
<img width="450" alt="index" src="https://github.com/user-attachments/assets/4717047c-3026-4e41-9492-df29b0751dd2" />

1. **Start the Database**  
   Make sure you have `docker` and `docker-compose` installed, then run:
   ```bash
   docker-compose up -d
   ```
   This brings up the database container in the background.

2. **Run the API**  
   After building (as above), run:
   ```bash
   ./bin/api
   ```
   Confirm it’s accessible on <http://localhost:3000/>.

3. **Run the CLI**  
   In a separate terminal, after building the CLI, you can execute commands like:
   ```bash
   ./bin/cli load task
   ```
   This will fetch tasks from the providers and store them as needed.

4. **Open the UI**  
   In your browser, open `ui/index.html` (e.g., by double-clicking it or serving it via a static server) to see the results.  
   This minimal page will communicate with the API at <http://localhost:3000/> to display tasks or scheduling info.

That’s it! You’re now able to test loading tasks, viewing them, and scheduling work among the developers.
