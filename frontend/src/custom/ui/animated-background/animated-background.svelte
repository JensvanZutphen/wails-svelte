<script lang="ts">
	import { Star } from '@lucide/svelte';
	import { sineInOut, bounceOut, elasticOut } from 'svelte/easing';
	import { scale, fly } from 'svelte/transition';
	import type { AnimatedBackgroundProps, Particle } from '../shared/types.js';
	import { createAnimationLoop, calculateOrbPositions, ParticleSystem } from '../shared/animations.js';

	let {
		particleCount = 8,
		animationSpeed = 1,
		gradientColors = ['from-blue-50 via-indigo-50 to-purple-50'],
		enableParticles = true
	}: AnimatedBackgroundProps = $props();

	// Svelte 5 runes
	let mounted = $state(false);
	let particles = $state<Particle[]>([]);
	let animationTime = $state(0);

	// Create particle system instance
	const particleSystem = new ParticleSystem(particleCount);

	// Derived values for background orbs using runes
	let orbPositions = $derived(calculateOrbPositions(animationTime * animationSpeed));

	// Effect to handle animation loop
	$effect(() => {
		if (!mounted) return;
		
		const cleanup = createAnimationLoop((time) => {
			animationTime = time;
		});
		
		return cleanup;
	});

	// Effect to handle mounting
	$effect(() => {
		mounted = true;
	});

	// Function to trigger particle animation (called from parent)
	export function triggerParticles(): void {
		if (!enableParticles) return;
		
		const newParticles = particleSystem.createParticles();
		particles = newParticles;

		// Remove particles after animation
		setTimeout(() => {
			particles = [];
		}, 2000);
	}
</script>

<!-- Background with animated gradient -->
<div class="pointer-events-none absolute inset-0 overflow-hidden bg-gradient-to-br {gradientColors.join(' ')}">
	<!-- Animated background elements -->
	<div class="pointer-events-none absolute inset-0 overflow-hidden">
		{#if mounted}
			<div
				class="absolute h-80 w-80 rounded-full bg-gradient-to-br from-blue-400/20 to-purple-400/20 blur-3xl"
				style="transform: translate({orbPositions.orb1.x}px, {orbPositions.orb1.y}px) scale({orbPositions.orb1.scale})"
			></div>
			<div
				class="absolute h-80 w-80 rounded-full bg-gradient-to-br from-pink-400/20 to-yellow-400/20 blur-3xl"
				style="transform: translate({orbPositions.orb2.x}px, {orbPositions.orb2.y}px) scale({orbPositions.orb2.scale})"
			></div>
			<div
				class="absolute left-1/2 top-1/2 h-96 w-96 -translate-x-1/2 -translate-y-1/2 rounded-full bg-gradient-to-br from-green-400/10 to-blue-400/10 blur-3xl"
				style="transform: translate(calc(-50% + {orbPositions.orb3.x}px), calc(-50% + {orbPositions.orb3.y}px)) scale({orbPositions.orb3.scale})"
			></div>
		{/if}
	</div>

	<!-- Floating particles -->
	{#if enableParticles}
		<div class="pointer-events-none absolute inset-0">
			{#each particles as particle (particle.id)}
				<div
					in:fly={{
						y: 60 + Math.random() * 40,
						x: (Math.random() - 0.5) * 100,
						duration: 1200,
						delay: particle.delay,
						easing: elasticOut
					}}
					out:scale={{
						duration: 600,
						delay: Math.random() * 200,
						easing: sineInOut,
						start: 1.2
					}}
					class="absolute"
					style="left: {particle.x}%; top: {particle.y}%;"
				>
					<div
						in:scale={{
							duration: 800,
							delay: particle.delay + 200,
							easing: bounceOut,
							start: 0
						}}
					>
						<Star class="h-4 w-4 fill-current text-yellow-400 drop-shadow-sm" />
					</div>
				</div>
			{/each}
		</div>
	{/if}
</div>