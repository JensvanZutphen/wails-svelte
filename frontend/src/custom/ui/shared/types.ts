// Shared TypeScript interfaces extracted from page.svelte

export interface Particle {
	id: number;
	x: number;
	y: number;
	delay: number;
}

export interface Pokemon {
	id: number;
	name: string;
	image_url: string;
	types: string[];
}

// Component prop interfaces
export interface AnimatedBackgroundProps {
	particleCount?: number;
	animationSpeed?: number;
	gradientColors?: string[];
	enableParticles?: boolean;
}

export interface HeroSectionProps {
	title?: string;
	subtitle?: string;
	buttonText?: string;
	onDiscoverClick: () => void;
	isLoading?: boolean;
}

export interface PokemonCardProps {
	pokemon: Pokemon | null;
	isLoading: boolean;
	error?: string;
}

export interface WelcomeStateProps {
	isVisible: boolean;
}