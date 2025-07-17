<script>
	import { GetRandomPokemon } from '$lib/wailsjs/go/app/App';
	import { Button } from '$lib/components/ui/button';
	import { Card, CardContent, CardHeader, CardTitle } from '$lib/components/ui/card';
	import { Badge } from '$lib/components/ui/badge';
	import { Alert, AlertDescription } from '$lib/components/ui/alert';
	import { Skeleton } from '$lib/components/ui/skeleton';

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

<main class="container mx-auto max-w-2xl space-y-6 px-4 py-8">
	<h1 class="text-center text-3xl font-bold">Pokemon Example</h1>

	<div class="flex justify-center">
		<Button onclick={fetchNewPokemon}>Get Random Pokemon</Button>
	</div>

	{#if pokemonId > 0}
		{#await getPokemon()}
			<div class="space-y-4">
				<Skeleton class="h-8 w-48" />
				<Skeleton class="h-32 w-32 rounded-lg" />
				<div class="flex gap-2">
					<Skeleton class="h-6 w-16" />
					<Skeleton class="h-6 w-16" />
				</div>
			</div>
		{:then pokemon}
			{#if pokemon}
				<Card>
					<CardHeader>
						<CardTitle>#{pokemon.id} {pokemon.name}</CardTitle>
					</CardHeader>
					<CardContent>
						<img src={pokemon.image_url} alt={pokemon.name} />
						<div class="flex gap-2">
							{#each pokemon.types as type}
								<Badge variant="secondary">{type}</Badge>
							{/each}
						</div>
					</CardContent>
				</Card>
			{/if}
		{:catch error}
			<Alert variant="destructive">
				<AlertDescription
					>Failed to load Pokemon: {error.message || 'Unknown error'}</AlertDescription
				>
			</Alert>
		{/await}
	{/if}
</main>
