import adapter from '@sveltejs/adapter-node';
import tailwindcss from '@tailwindcss/vite';
import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vite';

const apiTarget = process.env.VITE_BACKEND_TARGET || 'https://alhamidcintamulya.com';

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
				target: apiTarget,
				changeOrigin: true,
				secure: false
			}
		}
	}
});
