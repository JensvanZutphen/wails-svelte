import type { Particle } from './types.js';

export function createAnimationLoop(callback: (time: number) => void): () => void {
	let id: number;
	let start: number;

	function animate(time: number) {
		if (!start) start = time;
		callback(time - start);
		id = requestAnimationFrame(animate);
	}

	id = requestAnimationFrame(animate);
	return () => cancelAnimationFrame(id);
}

export function calculateOrbPositions(time: number) {
	const t = time * 0.001;
	
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

export class ParticleSystem {
	private count: number;

	constructor(count: number = 8) {
		this.count = count;
	}

	createParticles(): Particle[] {
		const particles: Particle[] = [];
		
		for (let i = 0; i < this.count; i++) {
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