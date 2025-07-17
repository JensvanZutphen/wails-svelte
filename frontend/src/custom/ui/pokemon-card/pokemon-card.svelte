<script lang="ts">
	import { Card, CardContent } from '$lib/components/ui/card';
	import { Alert, AlertDescription } from '$lib/components/ui/alert';
	import { Skeleton } from '$lib/components/ui/skeleton';
	import { Zap } from '@lucide/svelte';
	import { bounceOut } from 'svelte/easing';
	import { fly, scale } from 'svelte/transition';
	import type { PokemonFlipCardProps } from '../shared/types.js';

	import StatDisplay from './stat-display.svelte';
	import AbilityList from './ability-list.svelte';
	import PhysicalStats from './physical-stats.svelte';
	import SpeciesDescription from './species-description.svelte';
	import EvolutionChain from './evolution-chain.svelte';

	let { pokemon, isLoading, error, isFlipped = false, onFlip, onPokemonChange }: PokemonFlipCardProps = $props();

	const typeColors: Record<string, string> = {
		normal: '#9CA3AF', fire: '#EF4444', water: '#3B82F6', electric: '#FACC15',
		grass: '#10B981', ice: '#BFDBFE', fighting: '#B91C1C', poison: '#A855F7',
		ground: '#CA8A04', flying: '#818CF8', psychic: '#EC4899', bug: '#4ADE80',
		rock: '#92400E', ghost: '#7C3AED', dragon: '#4338CA', dark: '#1F2937',
		steel: '#6B7280', fairy: '#F9A8D4'
	};

	let flipped = $state(isFlipped);
	let animating = $state(false);

	function flip() {
		if (!pokemon || animating) return;
		flipped = !flipped;
		animating = true;
		onFlip?.();
		setTimeout(() => animating = false, 500);
	}

	function handleKey(e: KeyboardEvent) {
		if (e.key === 'Enter' || e.key === ' ') {
			e.preventDefault();
			flip();
		}
	}

	function getTypeHex(type: string): string {
		return typeColors[type.toLowerCase()] || '#6B7280';
	}
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
	<div class="mx-auto w-full max-w-[400px]" style="perspective:1000px;">
		<div
			style="position: relative; width: 100%; height: 500px; transform-style: preserve-3d; transition: transform 0.5s; cursor: pointer; {flipped ? 'transform: rotateY(180deg);' : ''}"
			onclick={flip}
			onkeydown={handleKey}
			tabindex="0"
			role="button"
			aria-label={flipped ? `Click to flip back to ${pokemon.name} image` : `Click to flip and see detailed stats for ${pokemon.name}`}
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
					style="width: 200px; height: 200px; object-fit: contain; margin-bottom: 16px; opacity: 1; transition: opacity 0.3s ease-in-out;"
					loading="eager"
				/>
				<div style="display: flex; gap: 8px; flex-wrap: wrap; justify-content: center;">
					{#each pokemon.types as type}
						<span
							style="background: {getTypeHex(type)}; color: white; padding: 6px 12px; border-radius: 16px; font-size: 14px; font-weight: 600;"
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
							style="background: {getTypeHex(type)}; color: white; padding: 4px 8px; border-radius: 12px; font-size: 12px; font-weight: 600;"
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

				<!-- Evolution Chain -->
				{#if pokemon.species && pokemon.species.evolutions && pokemon.species.evolutions.length > 0}
					<EvolutionChain
						evolutions={pokemon.species.evolutions}
						onPokemonClick={onPokemonChange}
					/>
				{/if}
			</div>
		</div>
	</div>


{/if}
