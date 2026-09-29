<script>
	import { page } from '$app/state';
</script>

<div class="min-h-screen flex flex-col items-center justify-center p-6 text-center text-gray-200">
	<h1 class="text-4xl font-bold text-primary mb-4">{page.status} - Page Not Found</h1>
	<p class="max-w-md text-gray-400 mb-6">
		The requested resource could not be found on this domain.
	</p>
	<div class="bg-gray-900 border border-gray-800 rounded-lg p-5 text-left max-w-lg w-full mb-6 font-mono text-sm">
		<p class="text-secondary font-bold mb-3"># Recovery Resources for AI Agents & Developers:</p>
		<ul class="space-y-1 text-gray-300">
			<li>- Homepage: <a href="/" class="text-primary hover:underline">https://priteshrathod.vercel.app/</a></li>
			<li>- LLMs Index: <a href="/llms.txt" class="text-primary hover:underline">https://priteshrathod.vercel.app/llms.txt</a></li>
			<li>- OpenAPI Spec: <a href="/openapi.json" class="text-primary hover:underline">https://priteshrathod.vercel.app/openapi.json</a></li>
			<li>- API Docs: <a href="/docs" class="text-primary hover:underline">https://priteshrathod.vercel.app/docs</a></li>
			<li>- Sitemap: <a href="/sitemap.xml" class="text-primary hover:underline">https://priteshrathod.vercel.app/sitemap.xml</a></li>
		</ul>
	</div>
	<a href="/" class="px-6 py-2 rounded-lg bg-primary text-black font-semibold hover:opacity-90">
		Back to Homepage
	</a>
</div>
