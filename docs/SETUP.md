## SETUP GUIDE ⚒️
Below is a step-by-step guide to set up the Frontend (Next.js) and Backend (Golang + Fiber) for Honya Books.

### Prerequisites
Make sure you have the following installed:
- [Node.js](https://nodejs.org/en/download)
- [PNPM](https://pnpm.io/installation) (for frontend)
- [Go](https://go.dev/dl/) 
- [PostgreSQL](https://www.postgresql.org/download/)
- [Docker](https://www.docker.com/get-started) (optional, for running application in containers)

### Environment Variables
> Duplicate the `.env.example` file in both `frontend` and `backend` directories and rename it to `.env`. Make sure you have the following environment variables set up:

**Frontend**
- [ ] `BACKEND_API_URL`

**Backend**
- [ ] `DATABASE_URL`: URL of the PostgreSQL database
- [ ] `SERVER_PORT`: Port the server will run on
- [ ] `LOG_STACK`: Stack the logs will be stored in
- [ ] `LOG_RETENTION`: Retention period for the logs

AWS Credentials
- [ ] `AWS_BUCKET_NAME`: Name of the AWS bucket
- [ ] `AWS_REGION`: Region of the AWS bucket
- [ ] `AWS_ACCESS_KEY_ID`: Access key ID for the AWS bucket
- [ ] `AWS_SECRET_ACCESS_KEY`: Secret access key for the AWS bucket

URL Cleanup Original Domain
- [ ] `URL_CLEANUP_ORIGINAL_DOMAIN`: Original domain of the URL for cleanup

---

### Run the Application
#### 1. Clone the Repository
```bash
git clone https://github.com/TechAtikiN/honya.git
cd honya
```

---

#### 2. Install Dependencies
1. Using Makefile
```bash
make install
```
> `make install` will install both frontend and backend dependencies.

2. Custom Install
```bash
cd frontend
pnpm install # Install frontend dependencies
```
```bash
cd backend
go mod download # Install backend dependencies
```

---

#### 3. Run the Application
1. Using Makefile
```bash
make run # Run the application
```

2. Using Docker
```bash
make docker-up # Run the application in containers
```

3. Custom Run

**Frontend**
```bash
cd frontend
pnpm dev # Run the frontend
```
**Backend**
```bash
cd backend
go run main.go # Run the backend
```

---

#### 4. Access the Application
- Frontend: `http://localhost:3000`
- Backend: `http://localhost:8080`

Swagger UI: `http://localhost:8080/swagger/`

### API Documentation
The API documentation can be found at [API.md](./API.md) file.





