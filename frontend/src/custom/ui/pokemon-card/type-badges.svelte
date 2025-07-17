<script lang="ts">
	import { Badge } from '$lib/components/ui/badge';
	import { getTypeColor } from '../shared/utils';

	export let types: string[];
	export let size: 'sm' | 'md' | 'lg' = 'md';

	// Get size classes based on prop
	function getSizeClasses(size: string): string {
		switch (size) {
			case 'sm':
				return 'px-2 py-1 text-xs';
			case 'lg':
				return 'px-4 py-2 text-base';
			default:
				return 'px-3 py-1 text-sm';
		}
	}
</script>

<div class="flex flex-wrap justify-center gap-2">
	{#each types as type}
		<Badge 
			class="transform rounded-full font-semibold text-white shadow-lg transition-all duration-300 hover:scale-105 hover:shadow-xl {getSizeClasses(size)} {getTypeColor(type)}"
		>
			{type.charAt(0).toUpperCase() + type.slice(1)}
		</Badge>
	{/each}
</div>

<style>
	/* Responsive adjustments */
	@media (max-width: 640px) {
		.flex.flex-wrap {
			gap: 0.25rem;
		}
	}
	
	/* Accessibility: Respect reduced motion preferences */
	@media (prefers-reduced-motion: reduce) {
		:global(.transform) {
			transform: none !important;
		}
		
		:global(.transition-all) {
			transition: none !important;
		}
		
		:global(.hover\:scale-105:hover) {
			transform: none !important;
		}
	}
</style>