import { defineConfig } from 'vite';
import path from 'path';
import customHmrPlugin from './assets/vite-hmr-plugin';

export default defineConfig({
	plugins: [customHmrPlugin()],
	build: {
		outDir: 'dist',
		emptyOutDir: true,
		manifest: true, // Generate a manifest file for asset mapping
		rollupOptions: {
			input: {
				main: path.resolve(__dirname, 'assets/js/main.js'),
			},
			output: {
				entryFileNames: 'js/[name]-[hash].js',
				chunkFileNames: 'js/[name]-[hash].js',
				assetFileNames: 'assets/[name]-[hash].[ext]',
			},
		},
	},
	server: {
		hmr: true,
	},
	css: {
		postcss: true,
	},
});
