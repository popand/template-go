# Full-Stack AI Template with Go Backend

A modern full-stack template that combines Next.js 14 frontend with a Go backend, featuring multiple AI service integrations. This template provides a production-ready foundation for building AI-powered applications.

## Features

- **Frontend**: Next.js 14 with App Router, TypeScript, and Tailwind CSS
- **Backend**: Go with Gin framework
- **AI Services Integration**:
  - OpenAI (Chat and Audio Transcription)
  - Anthropic (Chat)
  - Replicate (Image Generation)
  - Deepgram (Audio Processing)
- **Additional Features**:
  - Firebase Authentication
  - Real-time streaming responses
  - CORS configuration
  - Environment variables management

## Prerequisites

- Node.js 18+ and npm
- Go 1.21+
- API keys for the services you plan to use

## Getting Started

1. Clone the repository:
```bash
git clone https://github.com/popand/template-go.git
cd template-go
```

2. Install frontend dependencies:
```bash
npm install
```

3. Set up environment variables:
   Create a `.env` file in the root directory with the following:
```env
PORT=8080
OPENAI_API_KEY=your_openai_api_key
ANTHROPIC_API_KEY=your_anthropic_api_key
REPLICATE_API_TOKEN=your_replicate_api_token
DEEPGRAM_API_KEY=your_deepgram_api_key
```

## Running the Application

1. Start the Go backend (in one terminal):
```bash
cd backend
go run main.go
```
The backend will be available at `http://localhost:8080`

2. Start the Next.js frontend (in another terminal):
```bash
npm run dev
```
The frontend will be available at `http://localhost:3000`

## API Endpoints

### OpenAI
- `POST /api/openai/chat` - Chat completion with streaming
- `POST /api/openai/transcribe` - Audio transcription

### Anthropic
- `POST /api/anthropic/chat` - Chat completion with Claude

### Replicate
- `POST /api/replicate/generate-image` - Image generation with Stable Diffusion

### Deepgram
- `GET /api/deepgram` - Get Deepgram API key for client-side usage

## Project Structure

```
.
├── backend/                 # Go backend
│   ├── api/                # API handlers
│   └── main.go             # Main server file
├── src/
│   ├── app/               # Next.js pages and API routes
│   ├── components/        # React components
│   └── lib/               # Utilities, hooks, and contexts
├── public/                # Static files
└── paths/                 # Project templates
```

## Development

- The frontend uses the Next.js 14 App Router for routing
- The backend uses Gin for API routing and middleware
- All API responses are properly typed with TypeScript
- The Go backend supports streaming responses for AI chat endpoints

## Testing API Endpoints

You can test the API endpoints using curl:

```bash
# Test OpenAI chat
curl -X POST http://localhost:8080/api/openai/chat \
  -H "Content-Type: application/json" \
  -d '{"messages": [{"role": "user", "content": "Hello!"}]}'

# Test Deepgram key
curl http://localhost:8080/api/deepgram
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

MIT License - feel free to use this template for any project.