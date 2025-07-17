<script lang="ts">
	import { GetRandomPokemon } from '$lib/wailsjs/go/app/App';
	import { AnimatedBackground } from '../custom/ui/animated-background/index.js';
	import { HeroSection } from '../custom/ui/hero-section/index.js';
	import { PokemonCard } from '../custom/ui/pokemon-card/index.js';
	import { WelcomeState } from '../custom/ui/welcome-state/index.js';
	import type { Pokemon } from '../custom/ui/shared/types.js';
	import { preloadImages } from '../custom/ui/shared/utils.js';

	let pokemonPromise = $state<Promise<Pokemon> | null>(null);
	let bg: AnimatedBackground;

	async function loadPokemon(): Promise<Pokemon> {
		const pokemon = await GetRandomPokemon();
		const urls = [pokemon.image_url];
		
		if (pokemon.species?.evolutions) {
			pokemon.species.evolutions.forEach(evo => {
				if (evo.image_url) urls.push(evo.image_url);
			});
		}
		
		await preloadImages(urls);
		return pokemon;
	}

	function fetchNew() {
		if (bg) bg.triggerParticles();
		pokemonPromise = loadPokemon();
	}

	function changePokemon(pokemon: Pokemon) {
		if (bg) bg.triggerParticles();

		pokemonPromise = (async () => {
			const urls = [pokemon.image_url];
			if (pokemon.species?.evolutions) {
				pokemon.species.evolutions.forEach(evo => {
					if (evo.image_url) urls.push(evo.image_url);
				});
			}
			await preloadImages(urls);
			return pokemon;
		})();
	}
</script>

<div class="relative flex h-screen flex-col overflow-hidden">
	<AnimatedBackground bind:this={bg} />

	<main class="container relative mx-auto flex h-full max-w-7xl flex-col justify-center px-3 py-4 sm:px-4 sm:py-6 lg:px-6">
		<HeroSection onDiscoverClick={fetchNew} />

		{#if pokemonPromise}
			{#await pokemonPromise}
				<div class="text-center">
					<PokemonCard pokemon={null} isLoading={true} />
					<p class="mt-2 text-sm text-gray-500 animate-pulse">Loading...</p>
				</div>
			{:then pokemon}
				<PokemonCard pokemon={pokemon} isLoading={false} onPokemonChange={changePokemon} />
			{:catch error}
				<PokemonCard pokemon={null} isLoading={false} error={error.message} />
			{/await}
		{:else}
			<WelcomeState isVisible={true} />
		{/if}
	</main>
</div>
