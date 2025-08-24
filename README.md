# README

## About

**A modern Wails desktop app template** featuring a modern stack:
- **React** for UI
- **TypeScript** for type safety
- **Shadcn UI** for accessible, themeable components
- **Biome** for code formatting and linting
- **Wails** for desktop app development
- **Go** for backend

https://github.com/user-attachments/assets/3ca52ea6-9049-45bb-85b4-08715b96e1b3

</br>





Wails is a framework for building desktop applications using Go for the backend and modern web technologies for the frontend. It enables you to create fast, native-feeling desktop apps with a seamless bridge between Go and JavaScript/TypeScript. Check official docs for more info: https://wails.io/docs/introduction

You can configure the project by editing `wails.json`. More information about the project settings can be found
here: https://wails.io/docs/reference/project-config

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

## Key Frontend Libraries

This project utilizes a modern frontend stack:

- [**React**](https://react.dev/): A JavaScript library for building user interfaces.
- [**TypeScript**](https://www.typescriptlang.org/): A typed superset of JavaScript that compiles to plain JavaScript.
- [**Vite**](https://vitejs.dev/): A fast build tool and development server.
- [**Tailwind CSS**](https://tailwindcss.com/): A utility-first CSS framework for rapid UI development.
- [**Shadcn UI**](https://ui.shadcn.com/): A collection of re-usable UI components built with Radix UI and Tailwind CSS.
- [**Biome**](https://biomejs.dev/): A fast formatter and linter for web projects.
- [**Radix UI**](https://www.radix-ui.com/): Primitives for building accessible design systems and web applications.

### Common Shadcn UI Commands

Use the following commands (from the `frontend` directory) to add new Shadcn UI components:

```sh
npx shadcn-ui@latest add button
npx shadcn-ui@latest add input
npx shadcn-ui@latest add card
# ...add any other supported component
```

See the [Shadcn UI documentation](https://ui.shadcn.com/docs/components) for a full list of available components and usage instructions.

## Library Versions

Below are the main library versions currently in use:

| Library                          | Version    |
| -------------------------------- | ---------- |
| react                            | ^18.3.1    |
| react-dom                        | ^18.3.1    |
| typescript                       | ^5.8.3     |
| vite                             | ^6.3.5     |
| @vitejs/plugin-react             | ^4.5.0     |
| tailwindcss                      | ^4.1.8     |
| @tailwindcss/vite                | ^4.1.8     |
| tw-animate-css                   | ^1.3.3     |
| class-variance-authority         | ^0.7.1     |
| clsx                             | ^2.1.1     |
| tailwind-merge                   | ^3.3.0     |
| @radix-ui/react-aspect-ratio     | ^1.1.7     |
| @radix-ui/react-dropdown-menu    | ^2.1.15    |
| @radix-ui/react-slot             | ^1.2.3     |
| lucide-react                     | ^0.511.0   |
| next-themes                      | ^0.4.6     |
| sonner                           | ^2.0.5     |
| @biomejs/biome                   | ^1.9.4     |


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
