// Shared utility functions extracted from page.svelte

export const typeColors: Record<string, string> = {
	normal: 'bg-gray-400',
	fire: 'bg-red-500',
	water: 'bg-blue-500',
	electric: 'bg-yellow-400',
	grass: 'bg-green-500',
	ice: 'bg-blue-200',
	fighting: 'bg-red-700',
	poison: 'bg-purple-500',
	ground: 'bg-yellow-600',
	flying: 'bg-indigo-400',
	psychic: 'bg-pink-500',
	bug: 'bg-green-400',
	rock: 'bg-yellow-800',
	ghost: 'bg-purple-700',
	dragon: 'bg-indigo-700',
	dark: 'bg-gray-800',
	steel: 'bg-gray-500',
	fairy: 'bg-pink-300'
};

export function getTypeColor(type: string): string {
	return typeColors[type.toLowerCase() as keyof typeof typeColors] || 'bg-gray-500';
}

// Utility for creating particles with unique IDs
export function createParticles(count: number): Array<{id: number; x: number; y: number; delay: number}> {
	return Array.from({ length: count }, (_, i) => ({
		id: Date.now() + i,
		x: Math.random() * 100,
		y: Math.random() * 100,
		delay: i * 80
	}));
}

// Utility for preloading images
export async function preloadImage(src: string): Promise<void> {
	return new Promise((resolve, reject) => {
		if (!src) {
			resolve();
			return;
		}
		const img = new Image();
		img.onload = () => resolve();
		img.onerror = () => resolve(); // Don't reject, just resolve to prevent blocking
		img.src = src;
	});
}

// Utility for preloading multiple images
export async function preloadImages(srcs: string[]): Promise<void> {
	const validSrcs = srcs.filter(src => src && src.trim() !== '');
	if (validSrcs.length === 0) return;
	
	await Promise.allSettled(validSrcs.map(src => preloadImage(src)));
}