
# Notes Management System

## Overview
This project is a secure and scalable RESTful API for managing notes. It allows users to create, read, update, delete, share, and search notes. I chose GoLang for this task due to its efficiency in handling concurrent requests and its strong standard library, which is ideal for building fast and reliable web servers.

## Installation

1. Clone the repository:
   ```bash
   git clone [https://www.github.com/ankit-pn/notesmanagement]
   cd notesmanagement
   ```
2. Install dependencies:
   ```bash
   go mod tidy && go mod install
   ```
3. Run the server:
   ```bash
   go run .
   ```

## Usage

1. **Register a New User**: Send a POST request to `http://localhost:8080/api/auth/signup` with a JSON body like:
   ```json
   {
       "Username": "username",
       "Password": "password"
   }
   ```

2. **Login**: POST to `http://localhost:8080/api/auth/login` with:
   ```json
   {
       "Username": "username",
       "Password": "password"
   }
   ```
   Use the token and UserId received in the response for further requests.

3. **Creating Notes**: POST to `http://localhost:8080/api/notes` with:
   ```json
   {
       "Title": "title",
       "Content": "content"
   }
   ```

4. **Editing Notes**: PUT to `http://localhost:8080/api/notes/:id`.

5. **Sharing Notes**: POST to `http://localhost:8080/api/notes/:id/share`.

6. **Searching Notes**: GET from `http://localhost:8080/api/search?q=:query`.

7. **Deleting Notes**: DELETE to `http://localhost:8080/api/notes/:id`.

## Technical Details

- **Authentication**: Includes endpoints for user signup and login.
- **Note Management**: CRUD operations for notes.
- **Search**: Currently using SQL queries. However, a more efficient method way is by using ElasticSearch. I have previously used elastic search for a project, and I am familiar with its implementation. Here is link of that project - [https://github.com/ankit-pn/dataIndexer].
- **Rate Limiting**: Implemented using a bucket-based method. An advanced implementation can be done with redis. Here is link of project that I have previously used redis for - [https://github.com/ankit-pn/VideoOCR.git].

## Future Enhancements

- **Testing**: Unit and integration tests are planned to be implemented soon to ensure API reliability.
- **Advanced Search**: Integration with ElasticSearch for optimized text search. I have previously used elastic search for a project, and I am familiar with its implementation. Here is link of that project - [https://github.com/ankit-pn/dataIndexer].
- **Improved Rate Limiting**: Using Redis for more efficient rate limiting. I have previously used Redis for a project, and I am familiar with its implementation. Here is link of that project - [https://github.com/ankit-pn/VideoOCR.git].


