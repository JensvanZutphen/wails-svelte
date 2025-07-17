<script lang="ts">
	import { Heart, Sword, Shield, Zap, ShieldCheck, Wind, TrendingUp } from '@lucide/svelte';
	import type { PokemonStat } from '../shared/types';

	export let stats: PokemonStat[];

	function getPercent(stat: number): number {
		return Math.min((stat / 255) * 100, 100);
	}

	function formatName(name: string): string {
		return name.split('-').map(word => word.charAt(0).toUpperCase() + word.slice(1)).join(' ');
	}

	function getColors(stat: number) {
		if (stat >= 120) {
			return {
				bg: 'bg-green-100/50',
				fill: 'bg-gradient-to-r from-green-400 to-green-500',
				badge: 'bg-gradient-to-r from-green-400 to-green-500 text-white'
			};
		}
		if (stat >= 80) {
			return {
				bg: 'bg-blue-100/50',
				fill: 'bg-gradient-to-r from-blue-400 to-blue-500',
				badge: 'bg-gradient-to-r from-blue-400 to-blue-500 text-white'
			};
		}
		if (stat >= 50) {
			return {
				bg: 'bg-yellow-100/50',
				fill: 'bg-gradient-to-r from-yellow-400 to-yellow-500',
				badge: 'bg-gradient-to-r from-yellow-400 to-yellow-500 text-white'
			};
		}
		return {
			bg: 'bg-red-100/50',
			fill: 'bg-gradient-to-r from-red-400 to-red-500',
			badge: 'bg-gradient-to-r from-red-400 to-red-500 text-white'
		};
	}

	function getStyle(name: string) {
		const styles: Record<string, { color: string; icon: any }> = {
			hp: { color: 'text-red-600', icon: Heart },
			attack: { color: 'text-orange-600', icon: Sword },
			defense: { color: 'text-blue-600', icon: Shield },
			'special-attack': { color: 'text-purple-600', icon: Zap },
			'special-defense': { color: 'text-indigo-600', icon: ShieldCheck },
			speed: { color: 'text-green-600', icon: Wind }
		};
		return styles[name] || { color: 'text-gray-600', icon: Heart };
	}
</script>

<div class="space-y-3">
	<div class="flex items-center gap-2 mb-3">
		<TrendingUp class="h-5 w-5 text-slate-600" />
		<h3 class="text-lg font-semibold text-gray-800">Base Stats</h3>
	</div>
	{#each stats as stat}
		{@const colors = getColors(stat.base_stat)}
		{@const style = getStyle(stat.name)}
		<div class="flex items-center gap-3 p-2">
			<div class="flex w-20 items-center justify-start gap-1">
				<svelte:component this={style.icon} class="h-4 w-4 {style.color}" />
				<span class="text-sm font-medium {style.color}">
					{formatName(stat.name)}
				</span>
			</div>
			<div class="flex flex-1 items-center gap-2">
				<div class="relative flex-1">
					<div class="h-3 w-full rounded-full {colors.bg} overflow-hidden">
						<div
							class="h-full transition-all duration-500 ease-out {colors.fill} shadow-sm"
							style="width: {getPercent(stat.base_stat)}%"
						></div>
					</div>
				</div>
				<div class="flex h-6 w-12 items-center justify-center rounded-full text-xs font-bold {colors.badge} shadow-sm">
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
