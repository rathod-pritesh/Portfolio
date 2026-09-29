import { MARKDOWN_HOMEPAGE } from '$lib/markdownHome.js';

/** @type {import('./$types').RequestHandler} */
export function GET() {
	return new Response(MARKDOWN_HOMEPAGE, {
		status: 200,
		headers: {
			'Content-Type': 'text/markdown; charset=utf-8',
			'Vary': 'Accept, Accept-Encoding',
			'Access-Control-Allow-Origin': '*',
			'Cache-Control': 'public, max-age=3600'
		}
	});
}
