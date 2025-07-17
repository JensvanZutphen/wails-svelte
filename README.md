# Wails + Svelte Pokemon Example

A minimal example demonstrating how to build desktop applications with **Wails v2** and **Svelte 5**, featuring a Pokemon API integration.

## What is Wails?

Wails is a framework for building desktop applications using Go for the backend and web technologies (HTML, CSS, JavaScript) for the frontend. Think of it as Electron, but with Go instead of Node.js - resulting in smaller binaries and better performance.

## Project Structure

```
├── main.go                     # Wails app entry point
├── wails.json                  # Wails configuration
├── go.mod                      # Go module dependencies
├── internal/                   # Private application code
│   ├── app/
│   │   └── app.go             # Core app structure and Wails bindings
│   ├── models/
│   │   └── pokemon.go         # Data models and types
│   └── services/
│       └── pokemon.go         # Business logic and API calls
└── frontend/                   # Svelte frontend
    ├── src/
    │   ├── routes/
    │   │   └── +page.svelte   # Main Pokemon component
    │   └── lib/
    │       └── wailsjs/       # Auto-generated Go bindings
    └── svelte.config.js       # Svelte configuration
```

## How It Works

### Backend (Go)

The backend follows a clean architecture pattern with clear separation of concerns:

1. **App Layer** (`internal/app/`): Wails bindings and application entry points
2. **Service Layer** (`internal/services/`): Business logic and external API integration
3. **Model Layer** (`internal/models/`): Data structures and types
4. **Method Binding**: Go methods are automatically exposed to the frontend

#### Architecture Benefits:
- **Separation of Concerns**: Each layer has a single responsibility
- **Testability**: Services can be easily unit tested
- **Maintainability**: Changes to API logic don't affect app structure
- **Scalability**: Easy to add new services and models

```go
// This Go method becomes available in the frontend
func (a *App) GetRandomPokemon() (*Pokemon, error) {
    // Fetch from external API
    // Return structured data
}
```

### Frontend (Svelte)

1. **Auto-generated Bindings**: Wails creates TypeScript bindings for your Go methods
2. **Direct Method Calls**: Call Go functions directly from Svelte components
3. **Type Safety**: Full TypeScript support for Go structs and methods

```javascript
import { GetRandomPokemon } from '$lib/wailsjs/go/main/App';

// Call Go method directly
const pokemon = await GetRandomPokemon();
```

### The Bridge

Wails automatically:
- Generates TypeScript definitions from your Go structs
- Creates JavaScript bindings for your Go methods
- Handles serialization between Go and JavaScript
- Manages the communication layer

## Features Demonstrated

- **HTTP API Integration**: Fetching data from external REST APIs
- **Async/Await**: Modern JavaScript async patterns with Go backends
- **Image Preloading**: Ensuring all resources are ready before display
- **Error Handling**: Proper error boundaries and user feedback
- **Svelte 5 Runes**: Modern reactive state management
- **Clean Architecture**: Separation of concerns between API, models, and UI

## Prerequisites

- **Go 1.18+**: [Download Go](https://golang.org/dl/)
- **Node.js 16+**: [Download Node.js](https://nodejs.org/)
- **Wails CLI**: Install with `go install github.com/wailsapp/wails/v2/cmd/wails@latest`

## Setup Instructions

### 1. Clone and Install Dependencies

```bash
# Clone the repository
git clone <your-repo-url>
cd wails-svelte-pokemon

# Install Go dependencies
go mod tidy

# Install frontend dependencies
cd frontend
npm install
cd ..
```

### 2. Development Mode

```bash
# Run in development mode with hot reload
wails dev
```

This starts:
- Go backend with hot reload
- Svelte frontend with Vite dev server
- Automatic browser refresh on changes

### 3. Build for Production

```bash
# Build desktop application
wails build
```

Creates a native executable in the `build/bin/` directory.

## Key Concepts

### Method Binding

Any public method on your App struct becomes available in the frontend:

```go
// app.go
func (a *App) MyMethod(param string) (string, error) {
    return "Hello " + param, nil
}
```

```javascript
// frontend
import { MyMethod } from '$lib/wailsjs/go/main/App';
const result = await MyMethod("World");
```

### Context and Lifecycle

```go
func (a *App) startup(ctx context.Context) {
    a.ctx = ctx
    // App initialization logic
}
```

The startup method runs when the app launches, giving you access to the Wails runtime context.

### Configuration

`wails.json` controls build settings:

```json
{
  "name": "MyApp",
  "outputfilename": "MyApp",
  "frontend:install": "npm install",
  "frontend:build": "npm run build",
  "frontend:dev:watcher": "npm run dev"
}
```

## Svelte 5 Features Used

- **Runes**: `$state()` for reactive variables
- **Modern Syntax**: `onclick` instead of `on:click`
- **Async Handling**: `{#await}` blocks for loading states
- **Experimental Features**: Async compiler options enabled

## API Integration Pattern

This example shows a clean pattern for integrating external APIs:

1. **Separate API logic** into dedicated files (`api.go`)
2. **Define clear data models** (`pokemon.go`)
3. **Handle errors gracefully** with proper Go error returns
4. **Preload resources** (images) before displaying
5. **Provide loading states** for better UX

## Common Patterns

### Error Handling
```go
func (a *App) MyMethod() (*Data, error) {
    data, err := externalAPI.Fetch()
    if err != nil {
        return nil, fmt.Errorf("failed to fetch: %v", err)
    }
    return data, nil
}
```

### Loading States
```svelte
{#await myAsyncFunction()}
    <div>Loading...</div>
{:then result}
    <div>Success: {result}</div>
{:catch error}
    <div>Error: {error}</div>
{/await}
```

### Resource Preloading
```javascript
// Wait for image to load
await new Promise((resolve, reject) => {
    const img = new Image();
    img.onload = resolve;
    img.onerror = reject;
    img.src = imageUrl;
});
```

## Building and Distribution

### Development
- `wails dev` - Hot reload development
- `wails doctor` - Check system requirements
- `wails build -debug` - Debug build with DevTools

### Production
- `wails build` - Optimized production build
- `wails build -platform windows/amd64` - Cross-platform builds
- `wails build -upx` - Compress with UPX

## Why Wails + Svelte?

- **Performance**: Go backend is fast and efficient
- **Small Binaries**: No Node.js runtime needed
- **Native Feel**: True desktop application, not a web wrapper
- **Developer Experience**: Hot reload, TypeScript support, modern tooling
- **Cross Platform**: Build for Windows, macOS, and Linux from one codebase

## Next Steps

- Add more complex state management
- Implement native menus and system tray
- Add file system operations
- Integrate with system notifications
- Build custom window controls

## Resources

- [Wails Documentation](https://wails.io/)
- [Svelte 5 Documentation](https://svelte.dev/)
- [Go Documentation](https://golang.org/doc/)
- [Pokemon API](https://pokeapi.co/)

---

This example demonstrates the core concepts of Wails development. The combination of Go's performance and Svelte's reactivity creates powerful desktop applications with minimal complexity.