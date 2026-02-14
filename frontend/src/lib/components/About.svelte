<script>
	import Section from './Section.svelte';

// 	const about = {
// 	title: 'About Me',
// 	description:
// 		"I'm an aspiring Full Stack Developer who enjoys turning ideas into functional, user-friendly web applications. I focus on writing clean, maintainable code and designing intuitive interfaces while continuously learning modern technologies to improve my craft.",
// 	highlights: [
// 		{ label: '4+', description: 'Years of Hands-On Learning' },
// 		{ label: '5+', description: 'Projects Built' },
// 		{ label: 'Full-Stack', description: 'Development Skills' },
// 		{ label: 'Always', description: 'Learning & Improving' }
// 	]
// };
	export let about;
</script>

<Section id="about">
	<div class="space-y-12">
		<div class="text-center space-y-4">
			<h2 class="text-4xl md:text-5xl font-bold bg-gradient-to-r from-primary to-secondary gradient-text text-primary underline decoration-[#6366f1]">
				{about.title}
			</h2>
			<div class="w-20 h-1 bg-gradient-to-r from-primary to-secondary mx-auto"></div>
		</div>

		<div class="max-w-3xl mx-auto">
			<p class="text-lg md:text-xl text-gray-300 leading-relaxed text-center">
				{about.description}
			</p>
		</div>

		<div class="grid grid-cols-2 md:grid-cols-4 gap-8 pt-8">
			{#each about.highlights as highlight}
				<div class="text-center space-y-2 p-6 rounded-lg bg-dark/50 border border-gray-800 hover:border-primary transition-colors duration-300">
					<h3 class="text-3xl md:text-4xl font-bold text-primary">{highlight.label}</h3>
					<p class="text-gray-400 text-sm md:text-base">{highlight.description}</p>
				</div>
			{/each}
		</div>
	</div>
</Section>
