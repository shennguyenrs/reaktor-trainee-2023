# Reaktor pre-assignment 2023

Here is my solution for pre-assignment. The application was built using:

- Backend: [Golang](https://github.com/golang/go), [Go Fiber](https://github.com/gofiber/fiber)
- Frontend: [SvelteKit](https://github.com/sveltejs/kit), [Tailwind CSS](https://github.com/tailwindlabs/tailwindcss)
- Database: [MongoDB](https://www.mongodb.com/home)
- DevOps: [Docker](https://www.docker.com/), [NGINX](https://www.nginx.com/), Ubuntu VPS

### What are missings if running application in the local machine?

- Cron task to fetch violation drone every 2 seconds
- Cron task to save violation to MongoDB every 10 minutes
- MongoDB database

### Prerequisites

Prepare your machine has installed:

- Golang
- [Air](https://github.com/cosmtrek/air)
- Node JS & NPM
- PNPM

Edit the `.env` file for frontend and backend to the following format:

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

### Run the application

You can run the backend by go the the `back` directory and use these commands:

```bash
go mod download
air run

```

The frontend can be run using these commands:

```bash
pnpm install
pnpm run build
node build

```
