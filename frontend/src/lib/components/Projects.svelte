<script>
	import Section from './Section.svelte';
	export let projects; 
</script>

<Section id="projects">
	<div class="space-y-12">

		<!-- Heading -->
		<div class="text-center space-y-4">
			<h2 class="text-4xl md:text-5xl font-bold bg-gradient-to-r from-primary to-secondary gradient-text underline decoration-[#94a3b8]">
				Featured Projects
			</h2>
			<div class="w-20 h-1 bg-gradient-to-r from-primary to-secondary mx-auto rounded-full"></div>
		</div>

		<!-- Project cards-->
		<div class="grid md:grid-cols-2 gap-8">
			{#each projects as project}
				<div class="group bg-dark/50 border border-gray-800 rounded-xl overflow-hidden
				            hover:border-primary hover:shadow-xl hover:shadow-primary/15 hover:scale-[1.02]
				            transition-all duration-300">

					<!-- Project image -->
					<div class="relative h-64 overflow-hidden">
						<img
							src={project.image}
							alt={project.title}
							class="w-full h-full object-cover group-hover:scale-110 transition-transform duration-500"
						/>
						<div class="absolute inset-0 bg-gradient-to-t {project.gradient} opacity-0 group-hover:opacity-30 transition-opacity duration-300"></div>
					</div>

					<!-- Project info -->
					<div class="p-6 space-y-4">
						<h3 class="text-2xl font-bold text-white group-hover:text-primary transition-colors duration-300">
							{project.title}
						</h3>
						<p class="text-gray-400 leading-relaxed text-sm">
							{project.description}
						</p>

						<!-- Tech badges -->
						<div class="flex flex-wrap gap-2">
							{#each project.technologies as tech}
								<span class="px-3 py-1 bg-primary/10 border border-primary/25 rounded-full text-xs font-medium text-gray-300">
									{tech}
								</span>
							{/each}
						</div>

						<!-- Links -->
						<div class="flex gap-3 pt-2">
							{#if project.github}
								<a
									href={project.github}
									target="_blank"
									rel="noopener noreferrer"
									class="btn-link flex items-center gap-2 px-4 py-2 bg-gray-800 border border-gray-700 rounded-lg text-sm text-gray-300
									       hover:border-primary hover:text-primary transition-all duration-300"
								>
									<i class="fa-brands fa-github"></i>
									GitHub
								</a>
							{/if}
							{#if project.liveUrl}
								<a
									href={project.liveUrl}
									target="_blank"
									rel="noopener noreferrer"
									class="btn-link flex items-center gap-2 px-4 py-2 bg-gray-800 border border-gray-700 rounded-lg text-sm text-gray-300
									       hover:border-primary hover:text-primary transition-all duration-300"
								>
									<i class="fa-solid fa-arrow-up-right-from-square"></i>
									Live Demo
								</a>
							{/if}
						</div>
					</div>

				</div>
			{/each}
		</div>

	</div>
</Section>