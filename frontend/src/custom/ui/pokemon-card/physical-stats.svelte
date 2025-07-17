<script lang="ts">
	import { Card, CardContent, CardHeader } from '$lib/components/ui/card';
	import { Ruler, Weight, Activity } from '@lucide/svelte';

	export let height: number; // in decimeters
	export let weight: number; // in hectograms

	// Convert height from decimeters to meters and feet
	function formatHeight(decimeters: number): { meters: string; feet: string } {
		const meters = (decimeters / 10).toFixed(1);
		const totalInches = decimeters * 3.937;
		const feet = Math.floor(totalInches / 12);
		const inches = Math.round(totalInches % 12);

		return {
			meters: `${meters} m`,
			feet: `${feet}'${inches}"`
		};
	}

	// Convert weight from hectograms to kilograms and pounds
	function formatWeight(hectograms: number): { kg: string; lbs: string } {
		const kg = (hectograms / 10).toFixed(1);
		const lbs = (hectograms * 0.220462).toFixed(1);

		return {
			kg: `${kg} kg`,
			lbs: `${lbs} lbs`
		};
	}

	$: formattedHeight = formatHeight(height);
	$: formattedWeight = formatWeight(weight);
</script>

<div class="space-y-4">
	<div class="flex items-center gap-2 mb-3">
		<Activity class="h-5 w-5 text-emerald-600" />
		<h3 class="text-lg font-semibold text-gray-800">Physical Stats</h3>
	</div>

	<div class="grid grid-cols-2 gap-4">
		<!-- Height -->
		<Card class="border-blue-200 bg-blue-50/50">
			<CardHeader class="pb-2">
				<div class="flex items-center gap-2">
					<Ruler class="h-4 w-4 text-blue-600" />
					<span class="text-sm font-medium text-blue-800">Height</span>
				</div>
			</CardHeader>
			<CardContent class="pt-0">
				<div class="space-y-1">
					<div class="text-lg font-bold text-blue-900">
						{formattedHeight.meters}
					</div>
					<div class="text-sm text-blue-600">
						{formattedHeight.feet}
					</div>
				</div>
			</CardContent>
		</Card>

		<!-- Weight -->
		<Card class="border-green-200 bg-green-50/50">
			<CardHeader class="pb-2">
				<div class="flex items-center gap-2">
					<Weight class="h-4 w-4 text-green-600" />
					<span class="text-sm font-medium text-green-800">Weight</span>
				</div>
			</CardHeader>
			<CardContent class="pt-0">
				<div class="space-y-1">
					<div class="text-lg font-bold text-green-900">
						{formattedWeight.kg}
					</div>
					<div class="text-sm text-green-600">
						{formattedWeight.lbs}
					</div>
				</div>
			</CardContent>
		</Card>
	</div>
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
