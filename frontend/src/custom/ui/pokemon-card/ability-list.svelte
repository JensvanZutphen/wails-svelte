<script lang="ts">
	import type { PokemonAbility } from '../shared/types';

	export let abilities: PokemonAbility[];

	// Format ability name for display
	function formatAbilityName(name: string): string {
		return name
			.split('-')
			.map(word => word.charAt(0).toUpperCase() + word.slice(1))
			.join(' ');
	}

	// Truncate long descriptions for better layout
	function truncateDescription(description: string, maxLength: number = 120): string {
		if (!description || description.length <= maxLength) return description;
		return description.substring(0, maxLength).trim() + '...';
	}
</script>

<div class="space-y-3">
	<h3 class="text-lg font-semibold text-gray-800 mb-3">Abilities</h3>
	{#each abilities as ability}
		<div class="bg-gray-50 rounded-lg p-3 border border-gray-200">
			<div class="flex items-center gap-2 mb-2">
				<span class="font-medium text-gray-800">
					{formatAbilityName(ability.name)}
				</span>
				{#if ability.is_hidden}
					<span class="px-2 py-1 text-xs bg-purple-100 text-purple-700 rounded-full font-medium">
						Hidden
					</span>
				{/if}
			</div>
			{#if ability.description}
				<p class="text-sm text-gray-600 leading-relaxed">
					{truncateDescription(ability.description)}
				</p>
			{:else}
				<p class="text-sm text-gray-400 italic">
					Description not available
				</p>
			{/if}
		</div>
	{/each}
</div>

<style>
	/* Responsive adjustments */
	@media (max-width: 640px) {
		h3 {
			font-size: 1rem;
			margin-bottom: 0.75rem;
		}
	}
</style>