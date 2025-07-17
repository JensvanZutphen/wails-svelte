export interface Particle {
	id: number;
	x: number;
	y: number;
	delay: number;
}

import { models } from '$lib/wailsjs/go/models';

export type Pokemon = models.Pokemon;
export type PokemonStat = models.PokemonStat;
export type PokemonAbility = models.PokemonAbility;
export type PokemonSpecies = models.PokemonSpecies;
export type PokemonEvolution = models.PokemonEvolution;

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
}

export interface PokemonCardProps {
	pokemon: Pokemon | null;
	isLoading: boolean;
	error?: string;
}

export interface PokemonFlipCardProps {
	pokemon: Pokemon | null;
	isLoading: boolean;
	error?: string;
	isFlipped?: boolean;
	onFlip?: () => void;
	onPokemonChange?: (pokemon: Pokemon) => void;
}

export interface WelcomeStateProps {
	isVisible: boolean;
}