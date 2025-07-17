<script lang="ts">
	import { Heart, Sword, Shield, Zap, ShieldCheck, Wind, BarChart3 } from '@lucide/svelte';
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
			.map((word) => word.charAt(0).toUpperCase() + word.slice(1))
			.join(' ');
	}

	// Get color classes based on stat value
	function getStatColors(baseStat: number): {
		progressBg: string;
		progressFill: string;
		badgeClass: string;
	} {
		if (baseStat >= 120) {
			return {
				progressBg: 'bg-green-100/50',
				progressFill: 'bg-gradient-to-r from-green-400 to-green-500',
				badgeClass: 'bg-gradient-to-r from-green-400 to-green-500 text-white'
			};
		}
		if (baseStat >= 80) {
			return {
				progressBg: 'bg-blue-100/50',
				progressFill: 'bg-gradient-to-r from-blue-400 to-blue-500',
				badgeClass: 'bg-gradient-to-r from-blue-400 to-blue-500 text-white'
			};
		}
		if (baseStat >= 50) {
			return {
				progressBg: 'bg-yellow-100/50',
				progressFill: 'bg-gradient-to-r from-yellow-400 to-yellow-500',
				badgeClass: 'bg-gradient-to-r from-yellow-400 to-yellow-500 text-white'
			};
		}
		return {
			progressBg: 'bg-red-100/50',
			progressFill: 'bg-gradient-to-r from-red-400 to-red-500',
			badgeClass: 'bg-gradient-to-r from-red-400 to-red-500 text-white'
		};
	}

	// Get stat-specific colors and icons
	function getStatSpecificStyle(statName: string): { color: string; icon: any } {
		const styleMap: Record<string, { color: string; icon: any }> = {
			hp: { color: 'text-red-600', icon: Heart },
			attack: { color: 'text-orange-600', icon: Sword },
			defense: { color: 'text-blue-600', icon: Shield },
			'special-attack': { color: 'text-purple-600', icon: Zap },
			'special-defense': { color: 'text-indigo-600', icon: ShieldCheck },
			speed: { color: 'text-green-600', icon: Wind }
		};
		return styleMap[statName] || { color: 'text-gray-600', icon: Heart };
	}
</script>

<div class="space-y-3">
	<div class="flex items-center gap-2 mb-3">
		<BarChart3 class="h-5 w-5 text-slate-600" />
		<h3 class="text-lg font-semibold text-gray-800">Base Stats</h3>
	</div>
	{#each stats as stat}
		{@const colors = getStatColors(stat.base_stat)}
		{@const style = getStatSpecificStyle(stat.name)}
		<div class="flex items-center gap-3 p-2">
			<div class="flex w-20 items-center justify-end gap-1">
				<svelte:component this={style.icon} class="h-4 w-4 {style.color}" />
				<span class="text-sm font-medium {style.color}">
					{formatStatName(stat.name)}
				</span>
			</div>
			<div class="flex flex-1 items-center gap-2">
				<div class="relative flex-1">
					<!-- Custom colored progress bar -->
					<div class="h-3 w-full rounded-full {colors.progressBg} overflow-hidden">
						<div
							class="h-full transition-all duration-500 ease-out {colors.progressFill} shadow-sm"
							style="width: {getStatPercentage(stat.base_stat)}%"
						></div>
					</div>
				</div>
				<!-- Custom colored badge -->
				<div
					class="flex h-6 w-12 items-center justify-center rounded-full text-xs font-bold {colors.badgeClass} shadow-sm"
				>
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
