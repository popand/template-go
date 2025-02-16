import { NextResponse } from 'next/server';

export const runtime = 'edge';

export async function POST(req: Request) {
  const body = await req.json();
  
  const response = await fetch('http://localhost:8080/api/openai/chat', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(body),
  });

  // Check if the response is a stream
  const contentType = response.headers.get('content-type');
  if (contentType && contentType.includes('text/event-stream')) {
    // Forward the stream
    return new Response(response.body, {
      headers: {
        'Content-Type': 'text/event-stream',
        'Cache-Control': 'no-cache',
        'Connection': 'keep-alive',
      },
    });
  }

  // Handle non-stream responses (like errors)
  const data = await response.json();
  return NextResponse.json(data, { status: response.status });
}
