<script lang="ts">
	import { Card, CardContent, CardHeader, CardTitle } from '$lib/components/ui/card';
	import { Badge } from '$lib/components/ui/badge';
	import { Alert, AlertDescription } from '$lib/components/ui/alert';
	import { Skeleton } from '$lib/components/ui/skeleton';
	import { Zap, Star, Heart } from '@lucide/svelte';
	import { bounceOut, elasticOut } from 'svelte/easing';
	import { fly, scale } from 'svelte/transition';
	import type { PokemonCardProps } from '../shared/types.js';
	import { getTypeColor } from '../shared/utils.js';

	let { pokemon, isLoading, error }: PokemonCardProps = $props();
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

{#if isLoading}
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
{:else if error}
	<!-- Enhanced Error State -->
	<div class="mx-auto w-full max-w-xs sm:max-w-sm" in:fly={{ y: 20, duration: 400 }}>
		<Alert class="border-red-200 bg-red-50/80 shadow-lg backdrop-blur-sm">
			<Zap class="h-4 w-4 text-red-500 sm:h-5 sm:w-5" />
			<AlertDescription class="text-sm font-medium text-red-700 sm:text-base">
				Oops! Something went wrong while discovering your Pokémon: {error}
			</AlertDescription>
		</Alert>
	</div>
{:else if pokemon}
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