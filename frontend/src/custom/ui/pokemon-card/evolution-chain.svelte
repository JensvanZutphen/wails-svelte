<script lang="ts">
	import type { models } from '$lib/wailsjs/go/models';
	import { GetPokemonByID } from '$lib/wailsjs/go/app/App';

	let { evolutions, onPokemonClick }: { 
		evolutions: models.PokemonEvolution[]; 
		onPokemonClick?: (pokemon: models.Pokemon) => void;
	} = $props();

	async function handlePokemonClick(evolution: models.PokemonEvolution) {
		if (!onPokemonClick) return;
		try {
			const pokemon = await GetPokemonByID(evolution.id);
			onPokemonClick(pokemon);
		} catch (error) {
			console.error('Failed to fetch Pokemon:', error);
		}
	}

	function getEvolutionRequirement(evolution: models.PokemonEvolution): string {
		if (evolution.min_level && evolution.min_level > 0) {
			return `Level ${evolution.min_level}`;
		}
		if (evolution.item_name) {
			return evolution.item_name.replace('-', ' ');
		}
		if (evolution.held_item_name) {
			return `Hold ${evolution.held_item_name.replace('-', ' ')}`;
		}
		if (evolution.trigger_name && evolution.trigger_name !== 'level-up') {
			return evolution.trigger_name.replace('-', ' ');
		}
		return '';
	}
</script>

{#if evolutions && evolutions.length > 0}
	<div class="evolution-section">
		<h4 style="font-size: 14px; font-weight: 600; margin-bottom: 8px; color: #374151;">
			Evolution Chain
		</h4>
		<div class="evolution-chain">
			{#each evolutions as evolution, index}
				<button 
					class="evolution-item" 
					class:clickable={onPokemonClick} 
					onclick={() => handlePokemonClick(evolution)}
					disabled={!onPokemonClick}
					aria-label="View {evolution.name} details"
				>
					<img
						src={evolution.image_url}
						alt={evolution.name}
						style="width: 40px; height: 40px; object-fit: contain; opacity: 1; transition: opacity 0.3s ease-in-out;"
						loading="eager"
					/>
					<span style="font-size: 12px; font-weight: 500; color: #6B7280; text-align: center;">
						#{evolution.id.toString().padStart(3, '0')}
					</span>
					<span style="font-size: 12px; font-weight: 600; color: #374151; text-align: center; text-transform: capitalize;">
						{evolution.name}
					</span>
				</button>
				{#if index < evolutions.length - 1}
					{@const requirement = getEvolutionRequirement(evolutions[index + 1])}
					<div class="evolution-requirement">
						{#if requirement}
							<span class="requirement-text">{requirement}</span>
						{/if}
						<div class="evolution-arrow">→</div>
					</div>
				{/if}
			{/each}
		</div>
	</div>
{/if}

<style>
	.evolution-section {
		margin-top: 12px;
		padding: 12px;
		background: #F9FAFB;
		border-radius: 8px;
		border: 1px solid #E5E7EB;
	}

	.evolution-chain {
		display: flex;
		align-items: center;
		gap: 8px;
		overflow-x: auto;
		padding: 4px 0;
	}

	.evolution-item {
		display: flex;
		flex-direction: column;
		align-items: center;
		gap: 4px;
		min-width: 60px;
		padding: 8px;
		background: white;
		border-radius: 8px;
		border: 1px solid #E5E7EB;
		transition: all 0.2s ease;
		cursor: pointer;
		font: inherit;
	}

	.evolution-item:disabled {
		cursor: default;
		opacity: 0.7;
	}

	.evolution-item:hover {
		transform: translateY(-2px);
		box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
	}

	.evolution-item.clickable {
		cursor: pointer;
	}

	.evolution-item.clickable:hover {
		border-color: #3B82F6;
		background: #EFF6FF;
	}

	.evolution-requirement {
		display: flex;
		flex-direction: column;
		align-items: center;
		gap: 2px;
		margin: 0 4px;
	}

	.requirement-text {
		font-size: 10px;
		font-weight: 500;
		color: #6B7280;
		text-align: center;
		background: #F3F4F6;
		padding: 2px 6px;
		border-radius: 4px;
		white-space: nowrap;
		text-transform: capitalize;
	}

	.evolution-arrow {
		color: #9CA3AF;
		font-weight: bold;
		font-size: 16px;
	}

	@media (max-width: 640px) {
		.evolution-chain {
			gap: 6px;
		}
		
		.evolution-item {
			min-width: 50px;
			padding: 6px;
		}
		
		.evolution-arrow {
			font-size: 14px;
			margin: 0 2px;
		}
	}
</style>