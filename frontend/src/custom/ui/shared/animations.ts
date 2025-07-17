// Animation utilities and particle system extracted from page.svelte

import type { Particle } from './types.js';

export interface OrbPosition {
	x: number;
	y: number;
	scale: number;
}

// Animation loop utility
export function createAnimationLoop(callback: (time: number) => void): () => void {
	let animationId: number;
	let isRunning = true;
	
	function animate() {
		if (!isRunning) return;
		callback(Date.now());
		animationId = requestAnimationFrame(animate);
	}
	
	animate();
	
	// Return cleanup function
	return () => {
		isRunning = false;
		if (animationId) {
			cancelAnimationFrame(animationId);
		}
	};
}

// Calculate orb positions based on animation time
export function calculateOrbPositions(animationTime: number): {
	orb1: OrbPosition;
	orb2: OrbPosition;
	orb3: OrbPosition;
} {
	return {
		orb1: {
			x: Math.sin(animationTime * 0.001) * 50 - 160,
			y: Math.cos(animationTime * 0.001) * 30 - 160,
			scale: 1 + Math.sin(animationTime * 0.002) * 0.2
		},
		orb2: {
			x: Math.sin(animationTime * 0.0008) * 40 - 160,
			y: Math.sin(animationTime * 0.0016) * 25 - 160,
			scale: 1 + Math.cos(animationTime * 0.0015) * 0.15
		},
		orb3: {
			x: Math.sin(animationTime * 0.0006) * 20,
			y: Math.cos(animationTime * 0.0012) * 15,
			scale: 1 + Math.sin(animationTime * 0.001) * 0.1
		}
	};
}

// Particle system class for managing floating particles
export class ParticleSystem {
	private particles: Particle[] = [];
	private particleCount: number;
	
	constructor(particleCount: number = 8) {
		this.particleCount = particleCount;
	}
	
	createParticles(): Particle[] {
		this.particles = Array.from({ length: this.particleCount }, (_, i) => ({
			id: Date.now() + i,
			x: Math.random() * 100,
			y: Math.random() * 100,
			delay: i * 80
		}));
		
		return this.particles;
	}
	
	getParticles(): Particle[] {
		return this.particles;
	}
	
	clearParticles(): void {
		this.particles = [];
	}
	
	setParticleCount(count: number): void {
		this.particleCount = count;
	}
}