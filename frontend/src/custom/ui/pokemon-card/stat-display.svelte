<script lang="ts">
	import type { PokemonStat } from '../shared/types';

	export let stats: PokemonStat[];

	// Calculate percentage for stat bars (assuming max stat of 255)
	function getStatPercentage(baseStat: number): number {
		return Math.min((baseStat / 255) * 100, 100);
	}

	// Format stat name for display
	function formatStatName(name: string): string {
		return name
			.split('-')
			.map(word => word.charAt(0).toUpperCase() + word.slice(1))
			.join(' ');
	}

	// Get color class based on stat value
	function getStatColor(baseStat: number): string {
		if (baseStat >= 120) return 'bg-green-500';
		if (baseStat >= 80) return 'bg-yellow-500';
		if (baseStat >= 50) return 'bg-orange-500';
		return 'bg-red-500';
	}
</script>

<div class="space-y-3">
	<h3 class="text-lg font-semibold text-gray-800 mb-3">Base Stats</h3>
	{#each stats as stat}
		<div class="flex items-center gap-3">
			<div class="w-20 text-sm font-medium text-gray-600 text-right">
				{formatStatName(stat.name)}
			</div>
			<div class="flex-1 flex items-center gap-2">
				<div class="flex-1 bg-gray-200 rounded-full h-2 overflow-hidden">
					<div
						class="h-full transition-all duration-300 {getStatColor(stat.base_stat)}"
						style="width: {getStatPercentage(stat.base_stat)}%"
					></div>
				</div>
				<div class="w-8 text-sm font-bold text-gray-700 text-right">
					{stat.base_stat}
				</div>
			</div>
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