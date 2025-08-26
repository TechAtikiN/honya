## API Documentation ⚙️
The API documentation for Honya Books is provided below. This documentation outlines the available endpoints, request parameters, and response formats for interacting with the backend services.

### Base URL 🔗
The base URL for all API endpoints is:
```
http://localhost:8080/api
```

### Swagger Documentation 📄
For an interactive experience, you can access the Swagger UI for the API documentation at:
```
http://localhost:8080/docs
```

### Response Format ✅
All API responses follow a consistent JSON format with appropriate HTTP status codes:
- **200**: Success
- **201**: Created successfully
- **400**: Bad request (validation errors)
- **404**: Resource not found
- **409**: Conflict (duplicate resources)
- **500**: Internal server error

---

### Endpoints

#### 1. Books 📚

##### **GET /books**
Retrieve a paginated list of books with advanced filtering, sorting, and search capabilities.

**Query Parameters:**
- `query` (string, optional): Search books by title, author, or description
- `offset` (integer, optional): Pagination offset (default: 0)
- `limit` (integer, optional): Number of books per page (default: 10)
- `category` (string, optional): Filter by category (fiction, non_fiction, science, history, fantasy, mystery, thriller, cooking, travel, classics)
- `publication_year` (integer, optional): Filter by publication year
- `rating` (number, optional): Minimum rating filter
- `pages` (integer, optional): Filter by number of pages
- `sort` (string, optional): Sort by field (title, publication_year, rating)

**Response:** Returns a paginated list of books with metadata including total count and pagination info.

---
##### **GET /books/{id}**
Retrieve detailed information about a specific book including all its reviews.

**Path Parameters:**
- `id` (UUID, required): Book ID

**Response:** Returns complete book details with associated reviews.

##### **POST /books**
Create a new book entry with optional cover image upload.

**Content Type:** `multipart/form-data`

**Form Data:**
- `title` (string, required): Book title
- `description` (string, optional): Book description
- `category` (string, required): Book category
- `publication_year` (integer, required): Publication year
- `rating` (number, required): Book rating (0-5)
- `pages` (integer, required): Number of pages
- `isbn` (string, required): ISBN number (must be unique)
- `author_name` (string, required): Author name
- `image` (file, optional): Book cover image

**Response:** Returns the created book with generated ID and image URL.


---

##### **PATCH /books/{id}**
Update an existing book's details. Supports both JSON and form-data requests.

**Path Parameters:**
- `id` (UUID, required): Book ID

**Request Body/Form Data:**
- All book fields are optional except ISBN (cannot be updated)
- Supports partial updates

**Response:** Returns the updated book information.

--- 

##### **DELETE /books/{id}**
Delete a book and all its associated reviews.

**Path Parameters:**
- `id` (UUID, required): Book ID

**Response:** Confirmation message of successful deletion.

---

#### 2. Reviews 📝

##### **GET /reviews**
Retrieve a list of all reviews across all books with pagination and search capabilities.

**Query Parameters:**
- `query` (string, optional): Search query to filter reviews
- `offset` (integer, optional): Pagination offset (default: 0)
- `limit` (integer, optional): Number of reviews per page (default: 10)

##### **GET /reviews/{id}**
Get detailed information about a specific review.

**Path Parameters:**
- `id` (UUID, required): Review ID

##### **GET /books/{book_id}/reviews**
Retrieve all reviews for a specific book with pagination and search capabilities.

**Path Parameters:**
- `book_id` (UUID, required): Book ID

**Query Parameters:**
- `query` (string, optional): Search query to filter reviews
- `offset` (integer, optional): Pagination offset (default: 0)
- `limit` (integer, optional): Number of reviews per page (default: 10)

##### **POST /reviews**
Add a new review for a book.

**Request Body:**
```json
{
  "book_id": "uuid",
  "name": "Reviewer name (required)",
  "email": "reviewer@email.com (required, valid email)",
  "content": "Review content (required)"
}
```

##### **PATCH /reviews/{id}**
Update an existing review.

**Path Parameters:**
- `id` (UUID, required): Review ID

**Request Body:**
```json
{
  "name": "Updated reviewer name (optional)",
  "email": "updated@email.com (optional, valid email)",
  "content": "Updated review content (optional)"
}
```

**Note:** All fields are optional for partial updates.

##### **DELETE /reviews/{id}**
Delete a specific review.

**Path Parameters:**
- `id` (UUID, required): Review ID

---

#### 3. Dashboard Analytics 📊

##### **GET /dashboard/books-data**
Get aggregated statistical data for books with various filtering options for analytics visualization.

**Query Parameters:**
- `filter_by` (string, optional): Group data by field (category, author, rating, publication_year)

**Response:** Returns aggregated data suitable for charts and analytics dashboards.

##### **GET /dashboard/reviews-data**
Get top reviewers data showing most active users by review count.

**Query Parameters:**
- `limit` (integer, optional): Number of top reviewers to return (default: 10)

**Response:** Returns list of top reviewers with their review counts.

---

#### 4. URL Processing 🔗

##### **POST /url/process-url**
Process URLs to get redirection paths, canonical URLs, or both for link cleanup and validation.

**Request Body:**
```json
{
  "url": "https://example.com/some-url",
  "operation": "redirection|canonical|all"
}
```

**Operations:**
- `redirection`: Get the final redirect URL after following all redirects
- `canonical`: Extract the canonical URL from the page
- `all`: Perform both redirection and canonical URL extraction

**Response:**
```json
{
  "processed_url": "https://final-processed-url.com"
}
```

---

### Seeding Data
1. Using Makefile
```
make seed
```

2. Using API
##### **POST /seed**
Seed the database with sample data.

**Response:** Returns a message indicating the status of the seeding process for books and reviews.

---

### Schema Definitions
The schema definitions for the API requests and responses can be found in the [SCHEMA.md](./SCHEMA.md) file.