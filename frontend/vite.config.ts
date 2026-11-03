import adapter from '@sveltejs/adapter-node';
import tailwindcss from '@tailwindcss/vite';
import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vite';

export default defineConfig({
	plugins: [
		tailwindcss(),
		sveltekit({
			adapter: adapter(),
			alias: {
				'$lib': 'src/lib',
				'$lib/*': 'src/lib/*'
			}
		})
	],
	server: {
		proxy: {
			'/api': {
				target: 'http://127.0.0.1:8088',
				changeOrigin: true
			}
		}
	}
});
