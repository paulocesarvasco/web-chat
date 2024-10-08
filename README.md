Here's an updated `README.md` with a reference to an architecture diagram image in the Architecture section. Be sure to replace `architecture.png` with the actual filename and location of your architecture diagram.

---

# Web-Chat Project

A real-time web chat application that uses WebSocket for seamless communication between clients. This application features a secure authentication flow to ensure that only authorized users can join the chat.

## Overview

The web-chat application allows users to join a chat room and communicate in real-time. Users authenticate by providing their credentials, which are validated by an authorization service. If authentication is successful, the connection is upgraded to a WebSocket connection for live messaging.

## Features

- **User Authentication**: Credentials are validated by an authorization service, which issues an access token.
- **WebSocket Communication**: After authentication, the HTTP connection is upgraded to WebSocket for real-time chat functionality.
- **Secure**: Only authenticated users can access the chat.

## Architecture

The overall architecture of the web-chat application includes a client interface, a server that manages WebSocket connections, and an authorization service that validates credentials. The architecture is outlined in the following diagram:

![Architecture Diagram](images/web-chat-arch.png)

1. **Client (Browser)**:
   - The client provides a simple interface for users to enter their credentials and chat messages.
   - It initially sends an HTTP request to the server with user credentials.

2. **Server**:
   - Receives HTTP requests and validates the credentials through the authorization service.
   - If the credentials are valid, an access token is issued, and the connection is upgraded to WebSocket.
   - Manages the WebSocket connections and broadcasts messages to all connected clients.

3. **Authorization Service**:
   - Responsible for validating user credentials.
   - Generates and returns an access token if credentials are valid.

## How It Works

### 1. Client Sends Credentials
   - The user enters their username and password into the client interface.
   - An HTTP POST request is sent to the server with these credentials.

### 2. Server Validates Credentials
   - The server forwards the credentials to the authorization service.
   - The authorization service verifies the credentials and, if valid, returns an access token to the server.

### 3. WebSocket Connection
   - If the access token is valid, the server upgrades the connection to WebSocket.
   - The client and server can now exchange messages in real-time using the WebSocket connection.

## Installation

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/your-username/web-chat.git
   cd web-chat
   ```

2. **Setup Environment Variables**:
   - Create a `.env` file based on `.env.example`.
   - Add your configuration values, such as authorization service URLs and API keys.

3. **Build and Run the Application**:
   ```bash
   docker-compose up --build
   ```

## Usage

1. **Start the Application**:
   - Open a web browser and navigate to `http://localhost:8080`.

2. **Log In**:
   - Enter your credentials (username and password) on the login page.
   - If your credentials are valid, you will be redirected to the chat page.

3. **Start Chatting**:
   - Send and receive messages in real-time with other users in the chat room.

## Technologies Used

- **Go**: Server and WebSocket handling.
- **JavaScript/HTML/CSS**: Client-side user interface.
- **Docker**: Containerized deployment of the application and its services.
- **WebSocket**: Real-time messaging between clients.
- **Authorization Service**: Validates credentials and manages access tokens.
