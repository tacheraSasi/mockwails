# MockWails

## About

**MockWails** is a desktop application for creating and managing mock HTTP servers. It's built with **Wails**, using **Go** for the backend and a modern frontend stack.

The application allows you to define mock servers with specific endpoints, HTTP methods, response statuses, headers, and bodies. These servers are persisted in a local SQLite database and can be started and stopped from the UI. Active servers are automatically restarted when the application launches.

## Features

- **Create and Manage Mock Servers**: Easily create, edit, and delete mock servers.
- **Customizable Responses**: Configure the endpoint, HTTP method, status code, headers, and body for each mock server.
- **Persistent Storage**: Mock server configurations are saved to a local SQLite database.
- **Automatic Restart**: Active servers are automatically restarted when the application starts.
- **Easy to Use UI**: A clean and intuitive user interface for managing your mock servers.

## Screenshots

<div align="center">

### Main Interface
![MockWails Main Interface](./assets/img1.png)

### Mock Server Creation
![Create Mock Server](./assets/img2.png)

### Server Management
![Server Management](./assets/img3.png)

### Settings & Configuration
![Settings](./assets/img4.png)

</div>

</br>

## Tech Stack

- **Go** for the backend
- **Wails** for desktop app development
- **React** for the UI
- **TypeScript** for type safety
- **Shadcn UI** for accessible, themeable components
- **Biome** for code formatting and linting

## Live Development

To run in live development mode, use:

```sh
# install dependencies
make deps

# run in development mode
make dev
```

This will start the app with hot reload for both frontend and backend. For frontend-only development in a browser (with access to Go methods), open http://localhost:34115 in your browser.

## Building

To build a redistributable, production mode package, use `make build-[os]`, where `os` is one of `windows`, `linux`, or `mac`.

## Makefile Commands

The project includes a `Makefile` to simplify common development tasks. Run `make help` to see all available commands. Here are some of a few key targets:

- `make dev`: Run the application in development mode with hot reloading.
- `make build-windows`: Build the application for Windows (64-bit).
- `make build-linux`: Build the application for Linux (64-bit).
- `make build-mac`: Build the application for macOS (universal).
- `make clean`: Clean build artifacts.
- `make deps`: Install frontend dependencies from `package.json`.
- `make lint`: Lint the frontend code using Biome.
- `make format`: Format the frontend code using Biome.
- `make check`: Check and apply automatic fixes to the frontend code using Biome.