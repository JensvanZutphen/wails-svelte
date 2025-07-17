// Animation utilities for flip card transitions and background effects

import type { Particle } from './types.js';

export interface FlipState {
	isFlipped: boolean;
	isAnimating: boolean;
}

export interface FlipAnimationOptions {
	duration?: number;
	easing?: string;
	respectReducedMotion?: boolean;
}

// Default animation options
export const DEFAULT_FLIP_OPTIONS: Required<FlipAnimationOptions> = {
	duration: 500,
	easing: 'cubic-bezier(0.4, 0.0, 0.2, 1)',
	respectReducedMotion: true
};

// Check if user prefers reduced motion
export function prefersReducedMotion(): boolean {
	if (typeof window === 'undefined') return false;
	return window.matchMedia('(prefers-reduced-motion: reduce)').matches;
}

// Create flip animation configuration
export function createFlipAnimation(options: FlipAnimationOptions = {}): Required<FlipAnimationOptions> {
	const config = { ...DEFAULT_FLIP_OPTIONS, ...options };
	
	// Respect reduced motion preference
	if (config.respectReducedMotion && prefersReducedMotion()) {
		config.duration = 0;
	}
	
	return config;
}

// Flip state management utilities
export function createFlipState(initialFlipped = false): FlipState {
	return {
		isFlipped: initialFlipped,
		isAnimating: false
	};
}

export function toggleFlipState(state: FlipState): FlipState {
	return {
		...state,
		isFlipped: !state.isFlipped,
		isAnimating: true
	};
}

export function setAnimationComplete(state: FlipState): FlipState {
	return {
		...state,
		isAnimating: false
	};
}

// CSS transform utilities for flip animation
export function getFlipTransform(isFlipped: boolean): string {
	return `rotateY(${isFlipped ? 180 : 0}deg)`;
}

export function getCardSideVisibility(isFlipped: boolean, isFront: boolean): string {
	const shouldShow = isFront ? !isFlipped : isFlipped;
	return shouldShow ? 'visible' : 'hidden';
}

// Background animation utilities
export interface OrbPosition {
	x: number;
	y: number;
	scale: number;
}

export interface OrbPositions {
	orb1: OrbPosition;
	orb2: OrbPosition;
	orb3: OrbPosition;
}

// Create animation loop
export function createAnimationLoop(callback: (time: number) => void): () => void {
	let animationId: number;
	let startTime: number;

	function animate(currentTime: number) {
		if (!startTime) startTime = currentTime;
		const elapsed = currentTime - startTime;
		
		callback(elapsed);
		animationId = requestAnimationFrame(animate);
	}

	animationId = requestAnimationFrame(animate);

	return () => {
		if (animationId) {
			cancelAnimationFrame(animationId);
		}
	};
}

// Calculate orb positions for background animation
export function calculateOrbPositions(time: number): OrbPositions {
	const t = time * 0.001; // Convert to seconds
	
	return {
		orb1: {
			x: Math.sin(t * 0.5) * 100,
			y: Math.cos(t * 0.3) * 80,
			scale: 0.8 + Math.sin(t * 0.7) * 0.2
		},
		orb2: {
			x: Math.cos(t * 0.4) * 120,
			y: Math.sin(t * 0.6) * 90,
			scale: 0.9 + Math.cos(t * 0.5) * 0.1
		},
		orb3: {
			x: Math.sin(t * 0.3) * 60,
			y: Math.cos(t * 0.8) * 70,
			scale: 0.7 + Math.sin(t * 0.9) * 0.3
		}
	};
}

// Particle system for background effects
export class ParticleSystem {
	private particleCount: number;

	constructor(particleCount: number = 8) {
		this.particleCount = particleCount;
	}

	createParticles(): Particle[] {
		const particles: Particle[] = [];
		
		for (let i = 0; i < this.particleCount; i++) {
			particles.push({
				id: Date.now() + i,
				x: Math.random() * 100,
				y: 60 + Math.random() * 30,
				delay: Math.random() * 500
			});
		}
		
		return particles;
	}
}