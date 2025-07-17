<script lang="ts">
	import { Card, CardContent } from '$lib/components/ui/card';
	import { Badge } from '$lib/components/ui/badge';
	import { ChevronDown, ChevronUp, Sparkles } from '@lucide/svelte';
	import type { PokemonAbility } from '../shared/types';

	export let abilities: PokemonAbility[];

	let expanded: Record<number, boolean> = {};

	function formatName(name: string): string {
		return name
			.split('-')
			.map((word) => word.charAt(0).toUpperCase() + word.slice(1))
			.join(' ');
	}

	function isLong(text: string): boolean {
		return !!(text && text.length > 120);
	}

	function truncate(text: string): string {
		if (!text || text.length <= 120) return text;
		return text.substring(0, 120).trim() + '...';
	}

	function toggle(index: number) {
		expanded[index] = !expanded[index];
	}
</script>

<div class="space-y-3">
	<div class="mb-3 flex items-center gap-2">
		<Sparkles class="h-5 w-5 text-purple-600" />
		<h3 class="text-lg font-semibold text-gray-800">Abilities</h3>
	</div>
	{#each abilities as ability, index}
		<Card class="bg-muted/50">
			<CardContent class="p-3">
				<div class="mb-2 flex items-center gap-2">
					<span class="text-foreground font-medium">
						{formatName(ability.name)}
					</span>
					{#if ability.is_hidden}
						<Badge variant="secondary" class="text-xs">Hidden</Badge>
					{/if}
				</div>

				{#if ability.description}
					{#if isLong(ability.description)}
						<div class="space-y-2">
							<p class="text-muted-foreground text-sm leading-relaxed">
								{expanded[index] ? ability.description : truncate(ability.description)}
							</p>
							<button
								class="text-primary hover:text-primary/80 flex items-center gap-1 text-xs font-medium transition-colors"
								onclick={(e) => {
									e.stopPropagation();
									toggle(index);
								}}
							>
								{expanded[index] ? 'Show less' : 'Show more'}
								{#if expanded[index]}
									<ChevronUp class="h-3 w-3" />
								{:else}
									<ChevronDown class="h-3 w-3" />
								{/if}
							</button>
						</div>
					{:else}
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
