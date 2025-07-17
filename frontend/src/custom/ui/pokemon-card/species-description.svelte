<script lang="ts">
	import { Card, CardContent, CardHeader } from '$lib/components/ui/card';
	import { BookOpen } from '@lucide/svelte';
	import type { PokemonSpecies } from '../shared/types';

	export let species: PokemonSpecies;

	// Clean up the description text (remove form feeds and extra whitespace)
	function cleanDescription(description: string): string {
		if (!description) return '';
		return description
			.replace(/\f/g, ' ') // Replace form feed characters with spaces
			.replace(/\n/g, ' ') // Replace newlines with spaces
			.replace(/\s+/g, ' ') // Replace multiple spaces with single space
			.trim();
	}

	// Check if description is available and not empty
	function hasDescription(species: PokemonSpecies): boolean {
		return !!(species && species.description && species.description.trim().length > 0);
	}

	$: cleanedDescription = cleanDescription(species?.description || '');
	$: showDescription = hasDescription(species);
</script>

<div class="space-y-3">
	<div class="flex items-center gap-2 mb-3">
		<BookOpen class="h-4 w-4 text-indigo-600" />
		<h3 class="text-lg font-semibold text-gray-800">Description</h3>
	</div>
	
	{#if showDescription}
		<Card class="bg-indigo-50/50 border-indigo-200">
			<CardContent class="p-4">
				<p class="text-sm text-foreground leading-relaxed">
					{cleanedDescription}
				</p>
			</CardContent>
		</Card>
	{:else}
		<Card class="bg-muted/50">
			<CardContent class="p-4">
				<p class="text-sm text-muted-foreground italic text-center">
					Species description not available
				</p>
			</CardContent>
		</Card>
	{/if}
</div>

<style>
	/* Responsive adjustments */
	@media (max-width: 640px) {
		h3 {
			font-size: 1rem;
		}
	}
</style>