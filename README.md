## Instructions:

#### Start server:
* go run main.go
* Should see  "Server is running on port 8000..."

#### Create post:
* curl -X POST http://localhost:8000/posts -d '{"id": "1", "title": "First Post", "content": "This is the first post"}' -H "Content-Type: application/json"


#### List all posts:
* curl http://localhost:8000/posts

#### To call a specific post:
* curl http://localhost:8000/posts/1

#### Update an existing post: 
* curl -X PUT http://localhost:8000/posts/1 -d '{"title": "Updated Title", "content": "Updated content"}' -H "Content-Type: application/json"

#### Delete a post:
* curl -X DELETE http://localhost:8000/posts/1


#### Potential Updates:
* Persistent Data Storage: Right now, all posts are stored in memory (the posts slice), which means they are lost when the server is restarted. To persist data, you can:

    * Integrate a database like SQLite or PostgreSQL.
    * Use a file-based storage mechanism (e.g., JSON files or CSV).

* Input Validation: Add validation for user input to ensure all required fields (like Title and Content) are provided and valid.

* Error Handling: Improve error handling by using Go's error management capabilities. Ensure that all errors are returned with appropriate status codes (e.g., 400 Bad Request, 404 Not Found, etc.).

* Authentication: If you want to make the API more secure, you can implement user authentication (e.g., JWT tokens).


