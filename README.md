# URL Shortener
A simple URL shortener application with a Go backend and React frontend.

<p align="center">
  <img src="/frontend/public/Go-Logo_Blue.png" alt="Golang" height="100">
  <img src="/frontend/public/logo192.png" alt="React" height="100">
</p>

## Setup Instructions

### Backend (Go)

```bash
#Clone the repository
git clone https://github.com/your-username/url-shortener.git 

#Navigate to the backend directory
cd url-shortener/backend 

#Initialize the Go module (if not already done)
go mod init

#Install dependencies
go mod tidy

#Run the Go server
go run .
```
The backend should now be running on ```http://localhost:8080```.

### Frontend (React)
```bash

#Navigate to the frontend directory
cd url-shortener/frontend 

#Install dependencies
npm install

#Run the React development server
npm start
```
The frontend should now be running on ```http://localhost:3000```.

### Backend
- **Go**: The backend is powered by Go (Golang).
- **Gin**: A web framework for Go to handle HTTP requests and routing.
- **Gorm**: ORM for Go that helps interact with (in-memory) database.
- **CORS**: Package for handling Cross-Origin Resource Sharing (CORS) requests.

### Frontend
- **React**: A JavaScript library for building user interfaces.
- **Material-UI (MUI)**: A popular React component library for building modern, responsive web applications.
- **Axios**: Promise-based HTTP client for making requests from the React frontend to the Go backend.

## Features

- Shorten URLs.
- Display a list of shortened URLs.
- Delete shortened URLs.

## Endpoints
- **GET /***: Get a list of shortened URLs.
- **POST /shorten**: Shorten a given URL
- **DELETE /delete/:shortenedUrls**: Delete a shortened URL.

## Libraries and tools
- **Gin**: For building Go API [github.com/gin-gonic/gin](https://github.com/gin-gonic/gin)
- Gorm: ORM for Go [github.com/jinzhu/gorm](https://www.github.com/jinzhu/gorm).
- CORS: For handling cross-origin requests [github.com/gin-contrib/cors](https://www.github.com/gin-contrib/cors).
- [Material-UI (MUI)](https://mui.com/material-ui/all-components/): For building the React frontend UI (@mui/material).
- [Axios](https://axios-http.com/docs/intro): For making HTTP requests from React to the Go backend (axios).
