<script>
	import Section from './Section.svelte';
	export let about; 
</script>

<Section id="about">
	<div class="space-y-12">

		<!-- Heading -->
		<div class="text-center space-y-4">
			<h2 class="text-4xl md:text-5xl font-bold bg-gradient-to-r from-primary to-secondary gradient-text underline decoration-[#94a3b8]">
				{about.title}
			</h2>
			<div class="w-20 h-1 bg-gradient-to-r from-primary to-secondary mx-auto rounded-full"></div>
		</div>

		<!-- Description -->
		<div class="max-w-3xl mx-auto">
			<p class="text-lg md:text-xl text-gray-300 leading-relaxed text-center">
				{about.description}
			</p>
		</div>

		<!-- Highlight stats -->
		<div class="grid grid-cols-2 md:grid-cols-4 gap-3 pt-8">
			{#each about.highlights as highlight}
				<div class="text-center space-y-1 p-3 rounded-lg bg-dark/50 border border-gray-800 hover:border-primary transition-colors duration-300">
					<h3 class="text-sm md:text-base font-bold text-primary leading-tight">{highlight.label}</h3>
					<p class="text-gray-400 text-xs md:text-sm leading-snug">{highlight.description}</p>
				</div>
			{/each}
		</div>

	</div>
</Section>