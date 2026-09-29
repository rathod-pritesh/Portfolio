import { MARKDOWN_HOMEPAGE } from '$lib/markdownHome.js';

const NOT_FOUND_MARKDOWN = `# 404 Not Found

The requested resource was not found on this domain.

## Recovery Resources for AI Agents & Developers
- [Homepage](https://priteshrathod.vercel.app/)
- [LLMs Index](https://priteshrathod.vercel.app/llms.txt)
- [OpenAPI Spec](https://priteshrathod.vercel.app/openapi.json)
- [Sitemap](https://priteshrathod.vercel.app/sitemap.xml)
- [API Documentation](https://priteshrathod.vercel.app/docs)
- [Agent Auth Guide](https://priteshrathod.vercel.app/auth.md)
`;

/** @type {import('@sveltejs/kit').Handle} */
export async function handle({ event, resolve }) {
	const { url, request } = event;
	const acceptHeader = request.headers.get('accept') || '';
	const isAgentMode = url.searchParams.get('mode') === 'agent';

	// 1. Agent Mode View (?mode=agent)
	if (url.pathname === '/' && isAgentMode) {
		const agentView = {
			'@context': 'https://schema.org',
			name: 'Pritesh Rathod Developer Portfolio',
			mode: 'agent',
			endpoints: {
				openapi: 'https://priteshrathod.vercel.app/openapi.json',
				openapi_yaml: 'https://priteshrathod.vercel.app/api/openapi.yaml',
				api_catalog: 'https://priteshrathod.vercel.app/.well-known/api-catalog',
				auth_guide: 'https://priteshrathod.vercel.app/auth.md',
				mcp_server: 'https://priteshrathod.vercel.app/.well-known/mcp/server-card.json',
				agent_card: 'https://priteshrathod.vercel.app/.well-known/agent-card.json',
				skills: 'https://priteshrathod.vercel.app/SKILL.md',
				health: 'https://priteshrathod.vercel.app/health',
				home: 'https://priteshrathod.vercel.app/api/home',
				projects: 'https://priteshrathod.vercel.app/api/projects',
				skills_data: 'https://priteshrathod.vercel.app/api/skills',
				contact: 'https://priteshrathod.vercel.app/api/contact'
			},
			servers: [
				{
					url: 'https://priteshrathod.vercel.app',
					description: 'Primary Edge API Server'
				},
				{
					url: 'https://portfolio-i5x9.onrender.com',
					description: 'Direct Backend Service'
				}
			],
			capabilities: [
				'FastAPI Backend Engineering',
				'Golang REST APIs',
				'LangChain & LangGraph RAG Workflows',
				'Computer Vision (YOLO/OpenCV)',
				'Model Context Protocol (MCP)'
			],
			contact: {
				email: 'rathodpritesh0712@gmail.com',
				github: 'https://github.com/rathod-pritesh',
				linkedin: 'https://www.linkedin.com/in/rathodpritesh/'
			}
		};

		return new Response(JSON.stringify(agentView, null, 2), {
			status: 200,
			headers: {
				'Content-Type': 'application/json; charset=utf-8',
				Vary: 'Accept, Accept-Encoding',
				'Access-Control-Allow-Origin': '*'
			}
		});
	}

	// 2. Explicit Markdown Negotiation (only if client explicitly asks for markdown and not HTML)
	if (url.pathname === '/' && acceptHeader.includes('text/markdown') && !acceptHeader.includes('text/html')) {
		return new Response(MARKDOWN_HOMEPAGE, {
			status: 200,
			headers: {
				'Content-Type': 'text/markdown; charset=utf-8',
				Vary: 'Accept, Accept-Encoding',
				'Access-Control-Allow-Origin': '*'
			}
		});
	}

	const response = await resolve(event);

	// 3. JSON error responses for API paths or clients requesting JSON (RFC 7807 Problem Details)
	const isApiRequest = url.pathname.startsWith('/api/') || url.pathname === '/health';
	const prefersJson = acceptHeader.includes('application/json') || acceptHeader.includes('application/problem+json');

	if (response.status >= 400 && (isApiRequest || (prefersJson && !acceptHeader.includes('text/html')))) {
		const problem = {
			type: `https://priteshrathod.vercel.app/docs/errors#${response.status}`,
			title: response.status === 404 ? 'Not Found' : response.status === 405 ? 'Method Not Allowed' : 'API Error',
			status: response.status,
			detail: response.status === 404
				? `The requested API endpoint '${url.pathname}' was not found on this server.`
				: `An error occurred processing the request to '${url.pathname}'.`,
			instance: url.pathname,
			recovery_resources: {
				openapi: 'https://priteshrathod.vercel.app/openapi.json',
				docs: 'https://priteshrathod.vercel.app/docs',
				health: 'https://priteshrathod.vercel.app/health'
			}
		};

		return new Response(JSON.stringify(problem, null, 2), {
			status: response.status,
			headers: {
				'Content-Type': 'application/problem+json; charset=utf-8',
				'Access-Control-Allow-Origin': '*',
				'Access-Control-Allow-Methods': 'GET, POST, OPTIONS',
				'Access-Control-Allow-Headers': 'Content-Type, Authorization, Accept',
				Vary: 'Accept, Accept-Encoding'
			}
		});
	}

	// 4. Agent-friendly 404 with markdown body if client specifically requested markdown
	if (response.status === 404 && acceptHeader.includes('text/markdown') && !acceptHeader.includes('text/html')) {
		return new Response(NOT_FOUND_MARKDOWN, {
			status: 404,
			headers: {
				'Content-Type': 'text/markdown; charset=utf-8',
				Vary: 'Accept, Accept-Encoding',
				'Access-Control-Allow-Origin': '*'
			}
		});
	}

	// Ensure Vary header is set so CDNs respect content negotiation
	response.headers.set('Vary', 'Accept, Accept-Encoding');

	return response;
}
