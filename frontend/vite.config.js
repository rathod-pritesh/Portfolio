import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vite';
import tailwindcss from '@tailwindcss/vite';
import { API_BASE } from "./src/lib/config/api";

export default defineConfig({
	plugins: [
		tailwindcss(),
		sveltekit(),
	],

	server: {
		proxy: {
			'/api': {
				target: `${API_BASE}`,
				changeOrigin: true,
				proxyTimeout: 10000
			}
		}
	}
});
