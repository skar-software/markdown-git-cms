import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vite';

export default defineConfig({
	server: {
		host: true,
		port: 4173,
		proxy: {
			'/api': {
				target: 'http://back:3000',  // 'back' is the service name of your backend in docker-compose
				changeOrigin: true,
			}
		}
	},
	preview: {
		host: true,
		port: 4173,
		proxy: {
			'/api': {
				target: 'http://back:3000',
				changeOrigin: true,
			}
		}
	},
	plugins: [sveltekit()]
});