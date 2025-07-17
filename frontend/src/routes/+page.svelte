<script lang="ts">
	import { GetRandomPokemon } from '$lib/wailsjs/go/app/App';
	import { Button } from '$lib/components/ui/button';
	import { Card, CardContent, CardHeader, CardTitle } from '$lib/components/ui/card';
	import { Badge } from '$lib/components/ui/badge';
	import { Alert, AlertDescription } from '$lib/components/ui/alert';
	import { Skeleton } from '$lib/components/ui/skeleton';
	import { Sparkles, Zap, Star, Heart, Shuffle } from '@lucide/svelte';
	import { sineInOut, bounceOut, elasticOut } from 'svelte/easing';
	import { fade, fly, scale } from 'svelte/transition';

	// TypeScript interfaces
	interface Particle {
		id: number;
		x: number;
		y: number;
		delay: number;
	}

	interface Pokemon {
		id: number;
		name: string;
		image_url: string;
		types: string[];
	}

	// Svelte 5 runes
	let pokemonId = $state(0);
	let isAnimating = $state(false);
	let mounted = $state(false);
	let particles = $state<Particle[]>([]);
	let animationTime = $state(0);

	// Derived values for background orbs using runes
	let orb1Position = $derived({
		x: Math.sin(animationTime * 0.001) * 50 - 160,
		y: Math.cos(animationTime * 0.001) * 30 - 160,
		scale: 1 + Math.sin(animationTime * 0.002) * 0.2
	});

	let orb2Position = $derived({
		x: Math.sin(animationTime * 0.0008) * 40 - 160,
		y: Math.sin(animationTime * 0.0016) * 25 - 160,
		scale: 1 + Math.cos(animationTime * 0.0015) * 0.15
	});

	let orb3Position = $derived({
		x: Math.sin(animationTime * 0.0006) * 20,
		y: Math.cos(animationTime * 0.0012) * 15,
		scale: 1 + Math.sin(animationTime * 0.001) * 0.1
	});

	// Effect to handle animation loop
	$effect(() => {
		if (!mounted) return;
		
		let animationId: number;
		function animate() {
			animationTime = Date.now();
			animationId = requestAnimationFrame(animate);
		}
		animate();
		
		return () => {
			if (animationId) {
				cancelAnimationFrame(animationId);
			}
		};
	});

	// Effect to handle mounting
	$effect(() => {
		mounted = true;
	});

	function fetchNewPokemon() {
		isAnimating = true;
		pokemonId++;

		// Create new particles with unique IDs
		const newParticles = Array.from({ length: 8 }, (_, i) => ({
			id: Date.now() + i,
			x: Math.random() * 100,
			y: Math.random() * 100,
			delay: i * 80
		}));
		
		particles = newParticles;

		setTimeout(() => {
			isAnimating = false;
		}, 600);

		// Remove particles after animation
		setTimeout(() => {
			particles = [];
		}, 2000);
	}

	async function getPokemon() {
		if (pokemonId === 0) return null;
		const pokemon = await GetRandomPokemon();

		// Wait for image to load before resolving
		if (pokemon?.image_url) {
			await new Promise((resolve, reject) => {
				const img = new Image();
				img.onload = resolve;
				img.onerror = reject;
				img.src = pokemon.image_url;
			});
		}

		return pokemon;
	}

	const typeColors = {
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

	function getTypeColor(type: string): string {
		return typeColors[type.toLowerCase() as keyof typeof typeColors] || 'bg-gray-500';
	}

	// Debug inspection
	$inspect('Pokemon ID:', pokemonId);
	$inspect('Is Animating:', isAnimating);
	$inspect('Particles:', particles.length);
</script>

<style>
	.pokemon-card-3d {
		width: 100%;
		height: auto;
	}

	.pokemon-card-3d::part(inner) {
		border-radius: 1.5rem;
		overflow: hidden;
	}

	.pokemon-card-3d::part(rotate) {
		transform-style: preserve-3d;
	}

	/* Responsive rotation limits */
	@media (max-width: 640px) {
		.pokemon-card-3d {
			--atropos-rotate-x-max: 8deg;
			--atropos-rotate-y-max: 8deg;
		}
	}

	@media (min-width: 641px) and (max-width: 1024px) {
		.pokemon-card-3d {
			--atropos-rotate-x-max: 12deg;
			--atropos-rotate-y-max: 12deg;
		}
	}

	@media (min-width: 1025px) {
		.pokemon-card-3d {
			--atropos-rotate-x-max: 15deg;
			--atropos-rotate-y-max: 15deg;
		}
	}

	/* Accessibility: Respect reduced motion preferences */
	@media (prefers-reduced-motion: reduce) {
		.pokemon-card-3d {
			--atropos-rotate-x-max: 0deg;
			--atropos-rotate-y-max: 0deg;
			--atropos-duration: 0ms;
		}
	}
</style>

<!-- Background with animated gradient -->
<div class="relative flex h-screen flex-col overflow-hidden bg-gradient-to-br from-blue-50 via-indigo-50 to-purple-50">
	<!-- Animated background elements -->
	<div class="pointer-events-none absolute inset-0 overflow-hidden">
		{#if mounted}
			<div
				class="absolute h-80 w-80 rounded-full bg-gradient-to-br from-blue-400/20 to-purple-400/20 blur-3xl"
				style="transform: translate({orb1Position.x}px, {orb1Position.y}px) scale({orb1Position.scale})"
			></div>
			<div
				class="absolute h-80 w-80 rounded-full bg-gradient-to-br from-pink-400/20 to-yellow-400/20 blur-3xl"
				style="transform: translate({orb2Position.x}px, {orb2Position.y}px) scale({orb2Position.scale})"
			></div>
			<div
				class="absolute left-1/2 top-1/2 h-96 w-96 -translate-x-1/2 -translate-y-1/2 rounded-full bg-gradient-to-br from-green-400/10 to-blue-400/10 blur-3xl"
				style="transform: translate(calc(-50% + {orb3Position.x}px), calc(-50% + {orb3Position.y}px)) scale({orb3Position.scale})"
			></div>
		{/if}
	</div>

	<!-- Floating particles -->
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

	<main class="container relative mx-auto flex h-full max-w-7xl flex-col justify-center px-3 py-4 sm:px-4 sm:py-6 lg:px-6">
		<!-- Hero Section -->
		<div class="mb-6 text-center sm:mb-8">
			{#if mounted}
				<div class="mb-3 inline-flex items-center gap-1 sm:mb-4 sm:gap-2" in:scale={{ duration: 800, easing: elasticOut }}>
					<Sparkles class="h-4 w-4 text-yellow-500 sm:h-6 sm:w-6" />
					<h1 class="bg-gradient-to-r from-blue-600 via-purple-600 to-pink-600 bg-clip-text text-2xl font-bold text-transparent sm:text-4xl lg:text-5xl">
						PokéDiscover
					</h1>
					<Sparkles class="h-4 w-4 text-pink-500 sm:h-6 sm:w-6" />
				</div>
				<p class="mx-auto mb-4 max-w-2xl text-sm text-gray-600 sm:mb-6 sm:text-lg lg:text-xl" in:fade={{ delay: 300, duration: 600 }}>
					Embark on a magical journey to discover amazing Pokémon from across the universe
				</p>

				<!-- Enhanced Button -->
				<div class="relative inline-block" in:fly={{ y: 30, delay: 600, duration: 600, easing: bounceOut }}>
					<Button
						onclick={fetchNewPokemon}
						class="relative transform rounded-full bg-gradient-to-r from-blue-500 to-purple-600 px-4 py-2 text-sm font-semibold text-white shadow-lg transition-all duration-300 hover:scale-105 hover:from-blue-600 hover:to-purple-700 hover:shadow-xl sm:px-6 sm:py-3 sm:text-lg {isAnimating ? 'animate-pulse' : ''}"
						disabled={isAnimating}
					>
						<div class="flex items-center gap-1 sm:gap-2">
							<Shuffle class="h-4 w-4 sm:h-5 sm:w-5 {isAnimating ? 'animate-spin' : ''}" />
							<span class="hidden sm:inline">{isAnimating ? 'Discovering...' : 'Discover New Pokémon'}</span>
							<span class="sm:hidden">{isAnimating ? 'Finding...' : 'Discover'}</span>
							<Zap class="h-4 w-4 text-yellow-300 sm:h-5 sm:w-5" />
						</div>
					</Button>

					<!-- Button glow effect -->
					<div class="absolute inset-0 -z-10 rounded-full bg-gradient-to-r from-blue-500 to-purple-600 opacity-30 blur"></div>
				</div>
			{/if}
		</div>

		<!-- Pokemon Display Area -->
		{#if pokemonId > 0}
			{#await getPokemon()}
				<!-- Enhanced Loading State -->
				<div class="mx-auto w-full max-w-xs sm:max-w-sm lg:max-w-md" in:scale={{ duration: 400, easing: bounceOut }}>
					<Card class="overflow-hidden border-0 bg-white/80 shadow-2xl backdrop-blur-sm">
						<CardContent class="p-4 sm:p-6">
							<div class="space-y-3 text-center sm:space-y-4">
								<div class="relative">
									<Skeleton class="mx-auto h-6 w-24 rounded-full sm:h-8 sm:w-32" />
								</div>
								<div class="relative">
									<Skeleton class="mx-auto h-24 w-24 rounded-2xl sm:h-32 sm:w-32 lg:h-40 lg:w-40" />
								</div>
								<div class="flex justify-center gap-1 sm:gap-2">
									<Skeleton class="h-5 w-12 rounded-full sm:h-6 sm:w-16" />
									<Skeleton class="h-5 w-12 rounded-full sm:h-6 sm:w-16" />
								</div>
								<div class="flex justify-center">
									<Heart class="h-4 w-4 animate-pulse text-pink-400 sm:h-5 sm:w-5" />
								</div>
							</div>
						</CardContent>
					</Card>
				</div>
			{:then pokemon}
				{#if pokemon}
					<!-- Enhanced Pokemon Card with Atropos 3D -->
					<div class="mx-auto w-full max-w-xs sm:max-w-sm lg:max-w-md" in:scale={{ duration: 600, easing: elasticOut }}>
						<atropos-component
							class="pokemon-card-3d"
							active-offset="40"
							rotate-x-max="15"
							rotate-y-max="15"
							duration="300"
							shadow="true"
							shadow-offset="30"
							shadow-scale="1.05"
							highlight="true"
							rotate-touch="scroll-y"
						>
							<Card class="overflow-hidden border-0 bg-white/90 shadow-2xl backdrop-blur-sm transition-all duration-500 hover:shadow-3xl">
							<div class="relative">
								<!-- Card Header with Gradient -->
								<CardHeader class="relative overflow-hidden bg-gradient-to-r from-blue-500/10 to-purple-500/10 p-3 sm:p-6">
									<div class="absolute inset-0 bg-gradient-to-r from-blue-500/5 to-purple-500/5" data-atropos-offset="-2"></div>
									<CardTitle class="relative text-center">
										<div class="mb-1 flex items-center justify-center gap-1 sm:mb-2 sm:gap-2">
											<span class="bg-gradient-to-r from-blue-600 to-purple-600 bg-clip-text text-lg font-bold text-transparent sm:text-xl lg:text-2xl">
												#{pokemon.id.toString().padStart(3, '0')}
											</span>
										</div>
										<h2 class="bg-gradient-to-r from-gray-800 to-gray-600 bg-clip-text text-xl font-bold capitalize text-transparent sm:text-2xl lg:text-3xl">
											{pokemon.name}
										</h2>
									</CardTitle>
								</CardHeader>

								<!-- Pokemon Image with Effects -->
								<CardContent class="p-4 text-center sm:p-6">
									<div class="group relative mb-3 sm:mb-4">
										<div class="absolute inset-0 rounded-2xl bg-gradient-to-r from-blue-400/20 to-purple-400/20 blur-lg transition-all duration-500 group-hover:blur-xl sm:rounded-3xl sm:blur-xl sm:group-hover:blur-2xl" data-atropos-offset="-3"></div>
										<div class="relative rounded-2xl bg-gradient-to-br from-gray-50 to-gray-100 p-3 shadow-inner sm:rounded-3xl sm:p-4 lg:p-6" data-atropos-offset="-1">
											<img
												src={pokemon.image_url}
												alt={pokemon.name}
												class="mx-auto h-24 w-24 object-contain drop-shadow-lg transition-all duration-500 hover:scale-110 hover:drop-shadow-2xl sm:h-32 sm:w-32 lg:h-40 lg:w-40"
												data-atropos-offset="5"
											/>
										</div>
									</div>

									<!-- Type Badges with Custom Colors -->
									<div class="mb-3 flex flex-wrap justify-center gap-1 sm:mb-4 sm:gap-2" data-atropos-offset="2">
										{#each pokemon.types as type}
											<Badge class="transform rounded-full px-2 py-1 text-xs font-semibold text-white shadow-lg transition-all duration-300 hover:scale-105 hover:shadow-xl sm:px-3 sm:text-sm {getTypeColor(type)}">
												{type.charAt(0).toUpperCase() + type.slice(1)}
											</Badge>
										{/each}
									</div>

									<!-- Decorative Elements -->
									<div class="flex items-center justify-center gap-1 text-pink-400 sm:gap-2">
										<Heart class="h-3 w-3 animate-pulse fill-current sm:h-4 sm:w-4" />
										<Star class="h-2 w-2 animate-pulse fill-current sm:h-3 sm:w-3" style="animation-delay: 0.2s;" />
										<Heart class="h-3 w-3 animate-pulse fill-current sm:h-4 sm:w-4" style="animation-delay: 0.4s;" />
									</div>
								</CardContent>
							</div>
						</Card>
						</atropos-component>
					</div>
				{/if}
			{:catch error}
				<!-- Enhanced Error State -->
				<div class="mx-auto w-full max-w-xs sm:max-w-sm" in:fly={{ y: 20, duration: 400 }}>
					<Alert class="border-red-200 bg-red-50/80 shadow-lg backdrop-blur-sm">
						<Zap class="h-4 w-4 text-red-500 sm:h-5 sm:w-5" />
						<AlertDescription class="text-sm font-medium text-red-700 sm:text-base">
							Oops! Something went wrong while discovering your Pokémon: {error.message || 'Unknown error'}
						</AlertDescription>
					</Alert>
				</div>
			{/await}
		{:else}
			<!-- Welcome State -->
			{#if mounted}
				<div class="mx-auto w-full max-w-xs text-center sm:max-w-sm" in:scale={{ delay: 900, duration: 600, easing: bounceOut }}>
					<div class="rounded-2xl border border-white/20 bg-white/60 p-4 shadow-xl backdrop-blur-sm sm:rounded-3xl sm:p-6">
						<div class="mb-3 sm:mb-4">
							<Star class="mx-auto h-8 w-8 animate-bounce fill-current text-yellow-400 sm:h-12 sm:w-12" />
						</div>
						<h3 class="mb-2 text-lg font-bold text-gray-800 sm:text-xl">Ready for Adventure?</h3>
						<p class="text-sm text-gray-600 sm:text-base">
							Click the button above to discover your first Pokémon and begin your journey!
						</p>
					</div>
				</div>
			{/if}
		{/if}
	</main>
</div>