## SETUP GUIDE ⚒️
Below is a step-by-step guide to set up the Frontend (Next.js) and Backend (Golang + Fiber) for Honya Books.

### Prerequisites
Make sure you have the following installed:
- [Node.js](https://nodejs.org/en/download)
- [Pnpm](https://pnpm.io/installation) (for frontend)
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

#### 2. Install Dependencies
**Frontend**
> Install dependencies using either Makefile or PNPM.
1. Using Makefile
```bash
make install-fe
```
2. Using PNPM
> Note: This project uses yarn as the package manager.
```bash
cd frontend
pnpm install
```

**Backend**
> Install dependencies using either Makefile or Go.
1. Using Makefile
```bash
make install-be
```
2. Using Go
```bash
cd backend
go mod download
```

#### 3. Run the Application
**Frontend**
```bash
cd frontend
pnpm dev
```
**Backend**
```bash
cd backend
go run main.go
```

#### 4. Access the Application
- Frontend: `http://localhost:3000`
- Backend: `http://localhost:8080`

### API Documentation
The API documentation can be found at [API.md](./API.md) file.





