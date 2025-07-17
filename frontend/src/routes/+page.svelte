<script>
	import { GetRandomPokemon } from '$lib/wailsjs/go/main/App';

	let pokemonId = $state(0);

	function fetchNewPokemon() {
		pokemonId++;
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
</script>

<main>
	<h1>Pokemon Example</h1>

	<button onclick={fetchNewPokemon}> Get Random Pokemon </button>

	{#if pokemonId > 0}
		{#await getPokemon()}
			<div class="loading">Loading Pokemon...</div>
		{:then pokemon}
			{#if pokemon}
				<div class="pokemon">
					<h2>#{pokemon.id} {pokemon.name}</h2>
					<img src={pokemon.image_url} alt={pokemon.name} />
					<div class="types">
						{#each pokemon.types as type}
							<span class="type">{type}</span>
						{/each}
					</div>
				</div>
			{/if}
		{:catch error}
			<div class="error">Failed to load Pokemon</div>
		{/await}
	{/if}
</main>

<style>
	main {
		max-width: 400px;
		margin: 2rem auto;
		padding: 1rem;
		text-align: center;
		font-family: Arial, sans-serif;
	}

	button {
		background: #007acc;
		color: white;
		border: none;
		padding: 0.5rem 1rem;
		border-radius: 4px;
		cursor: pointer;
		margin-bottom: 1rem;
	}

	button:disabled {
		background: #ccc;
		cursor: not-allowed;
	}

	.loading {
		padding: 2rem;
		color: #666;
		font-style: italic;
	}

	.pokemon {
		border: 1px solid #ddd;
		border-radius: 8px;
		padding: 1rem;
		margin-top: 1rem;
	}

	.pokemon img {
		max-width: 150px;
		height: auto;
	}

	.types {
		margin-top: 0.5rem;
	}

	.type {
		background: #f0f0f0;
		padding: 0.25rem 0.5rem;
		border-radius: 4px;
		margin: 0 0.25rem;
		font-size: 0.8rem;
	}
</style>
