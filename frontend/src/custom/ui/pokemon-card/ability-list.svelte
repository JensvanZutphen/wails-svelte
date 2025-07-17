<script lang="ts">
	import { Card, CardContent } from '$lib/components/ui/card';
	import { Badge } from '$lib/components/ui/badge';
	import { ChevronDown, ChevronUp, Sparkles } from '@lucide/svelte';
	import type { PokemonAbility } from '../shared/types';

	export let abilities: PokemonAbility[];

	// Track expanded state for each ability
	let expandedStates: Record<number, boolean> = {};

	// Format ability name for display
	function formatAbilityName(name: string): string {
		return name
			.split('-')
			.map((word) => word.charAt(0).toUpperCase() + word.slice(1))
			.join(' ');
	}

	// Check if description is long enough to warrant expand/collapse
	function isLongDescription(description: string, threshold: number = 120): boolean {
		return !!(description && description.length > threshold);
	}

	// Get truncated text for long descriptions
	function getTruncatedText(description: string, maxLength: number = 120): string {
		if (!description || description.length <= maxLength) return description;
		return description.substring(0, maxLength).trim() + '...';
	}

	// Toggle expanded state
	function toggleExpanded(index: number) {
		expandedStates[index] = !expandedStates[index];
	}
</script>

<div class="space-y-3">
	<div class="flex items-center gap-2 mb-3">
		<Sparkles class="h-5 w-5 text-purple-600" />
		<h3 class="text-lg font-semibold text-gray-800">Abilities</h3>
	</div>
	{#each abilities as ability, index}
		<Card class="bg-muted/50">
			<CardContent class="p-3">
				<div class="mb-2 flex items-center gap-2">
					<span class="text-foreground font-medium">
						{formatAbilityName(ability.name)}
					</span>
					{#if ability.is_hidden}
						<Badge variant="secondary" class="text-xs">Hidden</Badge>
					{/if}
				</div>

				{#if ability.description}
					{#if isLongDescription(ability.description)}
						<!-- Long description with expand/collapse -->
						<div class="space-y-2">
							<p class="text-muted-foreground text-sm leading-relaxed">
								{expandedStates[index]
									? ability.description
									: getTruncatedText(ability.description)}
							</p>
							<button
								class="text-primary hover:text-primary/80 flex items-center gap-1 text-xs font-medium transition-colors"
								onclick={(e) => {
									e.stopPropagation();
									toggleExpanded(index);
								}}
							>
								{expandedStates[index] ? 'Show less' : 'Show more'}
								{#if expandedStates[index]}
									<ChevronUp class="h-3 w-3" />
								{:else}
									<ChevronDown class="h-3 w-3" />
								{/if}
							</button>
						</div>
					{:else}
						<!-- Short description - no expand needed -->
						<p class="text-muted-foreground text-sm leading-relaxed">
							{ability.description}
						</p>
					{/if}
				{:else}
					<p class="text-muted-foreground text-sm italic">Description not available</p>
				{/if}
			</CardContent>
		</Card>
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
