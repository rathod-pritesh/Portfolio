import { MARKDOWN_HOMEPAGE } from '$lib/markdownHome.js';

const KNOWN_BOT_UAS = [
	'gptbot',
	'claudebot',
	'chatgpt-user',
	'claude-user',
	'perplexitybot',
	'perplexity-user',
	'google-extended',
	'applebot-extended',
	'ora-agent',
	'deepseekbot'
];

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
	const userAgent = (request.headers.get('user-agent') || '').toLowerCase();
	const isBot = KNOWN_BOT_UAS.some((bot) => userAgent.includes(bot));
	const wantsMarkdown = acceptHeader.includes('text/markdown');
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
				home: 'https://portfolio-i5x9.onrender.com/api/home',
				projects: 'https://portfolio-i5x9.onrender.com/api/projects',
				skills_data: 'https://portfolio-i5x9.onrender.com/api/skills',
				contact: 'https://portfolio-i5x9.onrender.com/api/contact'
			},
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

	// 2. Markdown Content Negotiation & Bot-UA markdown serving for homepage
	if (url.pathname === '/' && (wantsMarkdown || isBot)) {
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

	// 3. Agent-friendly 404 with markdown body when requested or for bots
	if (response.status === 404 && (wantsMarkdown || isBot)) {
		return new Response(NOT_FOUND_MARKDOWN, {
			status: 404,
			headers: {
				'Content-Type': 'text/markdown; charset=utf-8',
				Vary: 'Accept, Accept-Encoding',
				'Access-Control-Allow-Origin': '*'
			}
		});
	}

	// Ensure Vary: Accept is set
	response.headers.set('Vary', 'Accept, Accept-Encoding');

	return response;
}
