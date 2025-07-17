export const typeColors: Record<string, string> = {
	normal: 'bg-gray-400', fire: 'bg-red-500', water: 'bg-blue-500', electric: 'bg-yellow-400',
	grass: 'bg-green-500', ice: 'bg-blue-200', fighting: 'bg-red-700', poison: 'bg-purple-500',
	ground: 'bg-yellow-600', flying: 'bg-indigo-400', psychic: 'bg-pink-500', bug: 'bg-green-400',
	rock: 'bg-yellow-800', ghost: 'bg-purple-700', dragon: 'bg-indigo-700', dark: 'bg-gray-800',
	steel: 'bg-gray-500', fairy: 'bg-pink-300'
};

export function getTypeColor(type: string): string {
	return typeColors[type.toLowerCase()] || 'bg-gray-500';
}

export function createParticles(count: number) {
	return Array.from({ length: count }, (_, i) => ({
		id: Date.now() + i,
		x: Math.random() * 100,
		y: Math.random() * 100,
		delay: i * 80
	}));
}

export async function preloadImage(src: string): Promise<void> {
	return new Promise((resolve) => {
		if (!src) {
			resolve();
			return;
		}
		const img = new Image();
		img.onload = () => resolve();
		img.onerror = () => resolve();
		img.src = src;
	});
}

export async function preloadImages(srcs: string[]): Promise<void> {
	const urls = srcs.filter(src => src && src.trim() !== '');
	if (urls.length === 0) return;
	
	await Promise.allSettled(urls.map(src => preloadImage(src)));
}