# Reaktor pre-assignment 2023

![reaktor-trainee](https://user-images.githubusercontent.com/17680161/208032439-7ddd0545-a369-4ffe-bded-0fab8760d206.jpg)

Here is my solution for pre-assignment. The application was built using:

- Backend: [Golang](https://github.com/golang/go), [Go Fiber](https://github.com/gofiber/fiber)
- Frontend: [SvelteKit](https://github.com/sveltejs/kit), [Tailwind CSS](https://github.com/tailwindlabs/tailwindcss)
- Database: [MongoDB](https://www.mongodb.com/home)
- DevOps: [Docker](https://www.docker.com/), [NGINX](https://www.nginx.com/), Ubuntu VPS

### Prerequisites

Prepare your machine has installed:

- Docker & Docker Compose

You will need to create new directory named `mongodb`. Then, it is necessary to creating `.env` file for frontend, backend, and database to the following format:

```text
// front/.env

PUBLIC_API_URL=

```

```text
// back/.env

MONGO_URI=
DB_NAME=
VIOLATION_LIST=
DRONES_API=https://assignments.reaktor.com/birdnest/drones
PILOT_API=https://assignments.reaktor.com/birdnest/pilots
```

```text
// mongodb/.env

MONGO_INITDB_ROOT_USERNAME=
MONGO_INITDB_ROOT_PASSWORD=
MONGO_INITDB_DATABASE=
```

### Run the application

You can run the application using the script `run-prod`.

```bash
sh ./run-prod.sh

```

To clean the production container, You can use the script `clean-prod`.

```bash
sh ./clean-prod.sh

```
