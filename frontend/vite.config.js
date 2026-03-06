import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vite';
import tailwindcss from '@tailwindcss/vite';

export default defineConfig({
	plugins: [
		tailwindcss(),
		sveltekit(),
	],

	server: {
		proxy: {
			'/api': {
				target: 'http:/127.0.0.1:8080',
				changeOrigin: true,
				proxyTimeout: 10000
			},
			'/n8n': {
				target: 'http://localhost:5678',
				rewrite: (path) => path.replace(/^\/n8n/, ''),
				changeOrigin: true
			}
		}
	}
});
