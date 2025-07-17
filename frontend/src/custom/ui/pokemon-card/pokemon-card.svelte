<script lang="ts">
	import { Card, CardContent, CardHeader, CardTitle } from '$lib/components/ui/card';
	import { Badge } from '$lib/components/ui/badge';
	import { Alert, AlertDescription } from '$lib/components/ui/alert';
	import { Skeleton } from '$lib/components/ui/skeleton';
	import { Zap, Star, Heart } from '@lucide/svelte';
	import { bounceOut, elasticOut } from 'svelte/easing';
	import { fly, scale } from 'svelte/transition';
	import type { PokemonFlipCardProps } from '../shared/types.js';
	import { getTypeColor } from '../shared/utils.js';

	// Type color mapping for inline styles
	const typeColorMap: Record<string, string> = {
		normal: '#9CA3AF',
		fire: '#EF4444',
		water: '#3B82F6',
		electric: '#FACC15',
		grass: '#10B981',
		ice: '#BFDBFE',
		fighting: '#B91C1C',
		poison: '#A855F7',
		ground: '#CA8A04',
		flying: '#818CF8',
		psychic: '#EC4899',
		bug: '#4ADE80',
		rock: '#92400E',
		ghost: '#7C3AED',
		dragon: '#4338CA',
		dark: '#1F2937',
		steel: '#6B7280',
		fairy: '#F9A8D4'
	};

	function getTypeColorHex(type: string): string {
		return typeColorMap[type.toLowerCase()] || '#6B7280';
	}
	import {
		createFlipState,
		toggleFlipState,
		setAnimationComplete,
		createFlipAnimation
	} from '../shared/animations.js';

	// Import back-side components
	import StatDisplay from './stat-display.svelte';
	import AbilityList from './ability-list.svelte';
	import PhysicalStats from './physical-stats.svelte';
	import SpeciesDescription from './species-description.svelte';

	let { pokemon, isLoading, error, isFlipped = false, onFlip }: PokemonFlipCardProps = $props();

	// Flip state management
	let flipState = $state(createFlipState(isFlipped));
	let animationConfig = createFlipAnimation();

	// Handle flip interaction
	function handleFlip() {
		if (!pokemon || flipState.isAnimating) return;

		flipState = toggleFlipState(flipState);
		onFlip?.();

		// Reset animation state after duration
		setTimeout(() => {
			flipState = setAnimationComplete(flipState);
		}, animationConfig.duration);
	}

	// Handle keyboard interaction
	function handleKeydown(event: KeyboardEvent) {
		if (event.key === 'Enter' || event.key === ' ') {
			event.preventDefault();
			handleFlip();
		}
	}

	// Sync external isFlipped prop with internal state (disabled for now to prevent auto-unflip)
	// $effect(() => {
	// 	if (isFlipped !== flipState.isFlipped && !flipState.isAnimating) {
	// 		flipState = { ...flipState, isFlipped };
	// 	}
	// });
</script>

{#if isLoading}
	<!-- Enhanced Loading State -->
	<div
		class="mx-auto w-full max-w-xs sm:max-w-sm lg:max-w-md"
		in:scale={{ duration: 400, easing: bounceOut }}
	>
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
	<!-- Pokemon Flip Card -->
	<div class="w-full max-w-[400px] mx-auto" style="perspective:1000px;">
		<div
			style="position: relative; width: 100%; height: 500px; transform-style: preserve-3d; transition: transform 0.5s; cursor: pointer; {flipState.isFlipped
				? 'transform: rotateY(180deg);'
				: ''}"
			onclick={handleFlip}
			onkeydown={handleKeydown}
			tabindex="0"
			role="button"
			aria-label={flipState.isFlipped
				? `Click to flip back to ${pokemon.name} image`
				: `Click to flip and see detailed stats for ${pokemon.name}`}
		>
			<!-- Front Side -->
			<div
				style="position: absolute; width: 100%; height: 100%; backface-visibility: hidden; background: white; border-radius: 16px; box-shadow: 0 8px 32px rgba(0,0,0,0.1); padding: 24px; display: flex; flex-direction: column; align-items: center; justify-content: center;"
			>
				<h2 style="color: #333; margin-bottom: 16px; font-size: 24px; font-weight: bold;">
					#{pokemon.id.toString().padStart(3, '0')}
					{pokemon.name.charAt(0).toUpperCase() + pokemon.name.slice(1)}
				</h2>
				<img
					src={pokemon.image_url}
					alt={pokemon.name}
					style="width: 200px; height: 200px; object-fit: contain; margin-bottom: 16px;"
				/>
				<div style="display: flex; gap: 8px; flex-wrap: wrap; justify-content: center;">
					{#each pokemon.types as type}
						<span
							style="background: {getTypeColorHex(type)}; color: white; padding: 6px 12px; border-radius: 16px; font-size: 14px; font-weight: 600;"
						>
							{type.charAt(0).toUpperCase() + type.slice(1)}
						</span>
					{/each}
				</div>
				<p style="margin-top: 16px; color: #666; font-size: 14px; text-align: center;">
					Click to see detailed stats!
				</p>
			</div>

			<!-- Back Side -->
			<div
				style="position: absolute; width: 100%; height: 100%; backface-visibility: hidden; transform: rotateY(180deg); background: white; border-radius: 16px; box-shadow: 0 8px 32px rgba(0,0,0,0.1); padding: 16px; overflow-y: auto;"
			>
				<h3
					style="text-align: center; margin-bottom: 16px; color: #333; font-size: 20px; font-weight: bold;"
				>
					#{pokemon.id.toString().padStart(3, '0')}
					{pokemon.name.charAt(0).toUpperCase() + pokemon.name.slice(1)}
				</h3>

				<!-- Type Badges -->
				<div
					style="display: flex; gap: 8px; flex-wrap: wrap; justify-content: center; margin-bottom: 16px;"
				>
					{#each pokemon.types as type}
						<span
							style="background: {getTypeColorHex(type)}; color: white; padding: 4px 8px; border-radius: 12px; font-size: 12px; font-weight: 600;"
						>
							{type.charAt(0).toUpperCase() + type.slice(1)}
						</span>
					{/each}
				</div>

				<!-- Physical Stats -->
				{#if pokemon.height && pokemon.weight}
					<PhysicalStats height={pokemon.height} weight={pokemon.weight} />
				{/if}

				<!-- Base Stats -->
				{#if pokemon.stats && pokemon.stats.length > 0}
					<StatDisplay stats={pokemon.stats} />
				{/if}

				<!-- Abilities -->
				{#if pokemon.abilities && pokemon.abilities.length > 0}
					<AbilityList abilities={pokemon.abilities} />
				{/if}

				<!-- Species Description -->
				{#if pokemon.species}
					<SpeciesDescription species={pokemon.species} />
				{/if}
			</div>
		</div>
	</div>

	<!-- Original complex flip card (hidden) -->
	<div style="display: none;">
		<div
			class="flip-card-container"
			onclick={handleFlip}
			onkeydown={handleKeydown}
			tabindex="0"
			role="button"
			aria-label={flipState.isFlipped
				? `Click to flip back to ${pokemon.name} image`
				: `Click to flip and see detailed stats for ${pokemon.name}`}
			aria-pressed={flipState.isFlipped}
			aria-describedby="flip-card-instructions"
			style="--front-display: {flipState.isFlipped
				? 'none'
				: 'block'}; --back-display: {flipState.isFlipped ? 'block' : 'none'};"
		>
			<div class="flip-card-inner {flipState.isFlipped ? 'flipped' : ''}">
				<!-- Front Side - Original Pokemon Card -->
				<div class="flip-card-front">
					<Card
						class="hover:shadow-3xl h-full cursor-pointer overflow-hidden border-0 bg-white/90 shadow-2xl backdrop-blur-sm transition-all duration-500"
					>
						<div class="relative h-full">
							<!-- Card Header with Gradient -->
							<CardHeader
								class="relative overflow-hidden bg-gradient-to-r from-blue-500/10 to-purple-500/10 p-3 sm:p-6"
							>
								<div
									class="absolute inset-0 bg-gradient-to-r from-blue-500/5 to-purple-500/5"
								></div>
								<CardTitle class="relative text-center">
									<div class="mb-1 flex items-center justify-center gap-1 sm:mb-2 sm:gap-2">
										<span
											class="bg-gradient-to-r from-blue-600 to-purple-600 bg-clip-text text-lg font-bold text-transparent sm:text-xl lg:text-2xl"
										>
											#{pokemon.id.toString().padStart(3, '0')}
										</span>
									</div>
									<h2
										class="bg-gradient-to-r from-gray-800 to-gray-600 bg-clip-text text-xl font-bold capitalize text-transparent sm:text-2xl lg:text-3xl"
									>
										{pokemon.name}
									</h2>
								</CardTitle>
							</CardHeader>

							<!-- Pokemon Image with Effects -->
							<CardContent class="p-4 text-center sm:p-6">
								<div class="group relative mb-3 sm:mb-4">
									<div
										class="absolute inset-0 rounded-2xl bg-gradient-to-r from-blue-400/20 to-purple-400/20 blur-lg transition-all duration-500 group-hover:blur-xl sm:rounded-3xl sm:blur-xl sm:group-hover:blur-2xl"
									></div>
									<div
										class="relative rounded-2xl bg-gradient-to-br from-gray-50 to-gray-100 p-3 shadow-inner sm:rounded-3xl sm:p-4 lg:p-6"
									>
										<img
											src={pokemon.image_url}
											alt={pokemon.name}
											class="mx-auto h-24 w-24 object-contain drop-shadow-lg transition-all duration-500 hover:scale-110 hover:drop-shadow-2xl sm:h-32 sm:w-32 lg:h-40 lg:w-40"
										/>
									</div>
								</div>

								<!-- Type Badges with Custom Colors -->
								<div class="mb-3 flex flex-wrap justify-center gap-1 sm:mb-4 sm:gap-2">
									{#each pokemon.types as type}
										<Badge
											class="transform rounded-full px-2 py-1 text-xs font-semibold text-white shadow-lg transition-all duration-300 hover:scale-105 hover:shadow-xl sm:px-3 sm:text-sm {getTypeColor(
												type
											)}"
										>
											{type.charAt(0).toUpperCase() + type.slice(1)}
										</Badge>
									{/each}
								</div>

								<!-- Decorative Elements -->
								<div class="flex items-center justify-center gap-1 text-pink-400 sm:gap-2">
									<Heart class="h-3 w-3 animate-pulse fill-current sm:h-4 sm:w-4" />
									<Star
										class="h-2 w-2 animate-pulse fill-current sm:h-3 sm:w-3"
										style="animation-delay: 0.2s;"
									/>
									<Heart
										class="h-3 w-3 animate-pulse fill-current sm:h-4 sm:w-4"
										style="animation-delay: 0.4s;"
									/>
								</div>
							</CardContent>
						</div>
					</Card>
				</div>

				<!-- Back Side - Detailed Pokemon Information -->
				<div class="flip-card-back">
					<Card
						class="h-full cursor-pointer overflow-hidden border-0 bg-white/90 shadow-2xl backdrop-blur-sm"
					>
						<div class="relative h-full">
							<!-- Back Card Header -->
							<CardHeader
								class="relative overflow-hidden bg-gradient-to-r from-purple-500/10 to-blue-500/10 p-3 sm:p-4"
							>
								<CardTitle class="text-center">
									<div class="mb-1 flex items-center justify-center gap-1 sm:mb-2 sm:gap-2">
										<span
											class="bg-gradient-to-r from-purple-600 to-blue-600 bg-clip-text text-lg font-bold text-transparent sm:text-xl"
										>
											#{pokemon.id.toString().padStart(3, '0')}
										</span>
									</div>
									<h2
										class="bg-gradient-to-r from-gray-800 to-gray-600 bg-clip-text text-lg font-bold capitalize text-transparent sm:text-xl"
									>
										{pokemon.name}
									</h2>
								</CardTitle>
							</CardHeader>

							<!-- Detailed Pokemon Information -->
							<CardContent class="card-content overflow-y-auto p-3 sm:p-4">
								<div class="space-y-4">
									<!-- Type Badges -->
									<div class="flex flex-wrap justify-center gap-1 sm:gap-2">
										{#each pokemon.types as type}
											<Badge
												class="rounded-full px-2 py-1 text-xs font-semibold text-white shadow-md sm:px-3 sm:text-sm {getTypeColor(
													type
												)}"
											>
												{type.charAt(0).toUpperCase() + type.slice(1)}
											</Badge>
										{/each}
									</div>

									<!-- Physical Stats -->
									{#if pokemon.height && pokemon.weight}
										<PhysicalStats height={pokemon.height} weight={pokemon.weight} />
									{/if}

									<!-- Base Stats -->
									{#if pokemon.stats && pokemon.stats.length > 0}
										<StatDisplay stats={pokemon.stats} />
									{/if}

									<!-- Abilities -->
									{#if pokemon.abilities && pokemon.abilities.length > 0}
										<AbilityList abilities={pokemon.abilities} />
									{/if}

									<!-- Species Description -->
									{#if pokemon.species}
										<SpeciesDescription species={pokemon.species} />
									{/if}
								</div>
							</CardContent>
						</div>
					</Card>
				</div>
			</div>
		</div>
	</div>
{/if}

<style>
	.flip-card-container {
		perspective: 1000px;
		width: 100%;
		height: auto;
	}

	.flip-card-inner {
		position: relative;
		width: 100%;
		height: 100%;
		text-align: center;
		transition: transform 500ms cubic-bezier(0.4, 0, 0.2, 1);
		transform-style: preserve-3d;
	}

	.flip-card-inner.flipped {
		transform: rotateY(180deg);
	}

	.flip-card-front,
	.flip-card-back {
		position: absolute;
		width: 100%;
		height: 100%;
		backface-visibility: hidden;
		border-radius: 1.5rem;
		overflow: hidden;
	}

	.flip-card-back {
		transform: rotateY(180deg);
	}

	/* Accessibility: Respect reduced motion preferences */
	@media (prefers-reduced-motion: reduce) {
		.flip-card-inner {
			transition: none;
		}

		.flip-card-inner.flipped {
			transform: none;
		}

		.flip-card-front {
			display: var(--front-display, block);
		}

		.flip-card-back {
			display: var(--back-display, none);
			transform: none;
			position: relative;
		}
	}

	/* Responsive adjustments */
	@media (max-width: 640px) {
		.flip-card-back {
			font-size: 0.875rem;
		}
	}
</style>
