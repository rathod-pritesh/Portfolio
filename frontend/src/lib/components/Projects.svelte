<script>
	import Section from './Section.svelte';
	export let projects;
</script>

<Section id="projects">
	<div class="space-y-20">

		<!-- Heading -->
		<div class="text-center space-y-4">
			<h2 class="text-4xl md:text-5xl font-bold bg-linear-to-r from-primary to-secondary gradient-text underline decoration-secondary">
				Featured Projects
			</h2>
			<div class="w-20 h-1 bg-linear-to-r from-primary to-secondary mx-auto rounded-full"></div>
		</div>

		<!-- Project rows -->
		<div class="space-y-6">
			{#each projects as project, i}
				<!-- Alternates: even = image left, odd = image right -->
				<div class="project-row group relative flex flex-col {i % 2 === 0 ? 'md:flex-row' : 'md:flex-row-reverse'} 
				            gap-0 overflow-hidden rounded-2xl border border-gray-800
				            hover:border-primary/50 transition-all duration-500
				            hover:shadow-2xl hover:shadow-primary/10">

					<!-- Accent bar (left edge for even, right edge for odd) -->
					<div class="absolute top-0 bottom-0 {i % 2 === 0 ? 'left-0' : 'right-0'} w-[3px]
					            bg-linear-to-b from-primary via-secondary to-transparent
					            opacity-0 group-hover:opacity-100 transition-opacity duration-500 z-10"></div>

					<!-- Image pane -->
					<div class="relative w-full md:w-[45%] shrink-0 overflow-hidden" style="min-height: 260px;">
						<img
							src={project.image}
							alt={project.title}
							class="absolute inset-0 w-full h-full object-cover
							       scale-100 group-hover:scale-105 transition-transform duration-700 ease-out"
						/>
						<!-- gradient overlay matching project.gradient -->
						<div class="absolute inset-0 bg-linear-to-b {project.gradient} opacity-40 group-hover:opacity-20 transition-opacity duration-500"></div>

						<!-- Project number watermark -->
						<span class="absolute bottom-3 {i % 2 === 0 ? 'right-4' : 'left-4'}
						             text-6xl font-black text-white/10 select-none leading-none
						             group-hover:text-white/20 transition-colors duration-500">
							{String(i + 1).padStart(2, '0')}
						</span>
					</div>

					<!-- Divider line -->
					<div class="hidden md:block w-px self-stretch bg-gray-800 group-hover:bg-primary/30 transition-colors duration-500 shrink-0"></div>

					<!-- Content pane -->
					<div class="flex flex-col justify-center gap-5 p-8 flex-1 bg-dark/60 backdrop-blur-sm">

						<!-- Index label + title -->
						<div class="space-y-1">
							<span class="text-xs font-mono tracking-[0.2em] text-primary/70 uppercase">
								Project / {String(i + 1).padStart(2, '0')}
							</span>
							<h3 class="text-2xl md:text-3xl font-bold text-white
							           group-hover:text-primary transition-colors duration-300 leading-tight">
								{project.title}
							</h3>
						</div>

						<!-- Separator -->
						<div class="w-10 h-px bg-linear-to-r from-primary to-secondary
						            scale-x-50 origin-left group-hover:scale-x-100 transition-transform duration-500"></div>

						<!-- Description -->
						<p class="text-gray-400 leading-relaxed text-sm max-w-lg">
							{project.description}
						</p>

						<!-- Tech badges -->
						<div class="flex flex-wrap gap-2">
							{#each project.technologies as tech}
								<span class="px-3 py-1 bg-primary/10 border border-primary/25 rounded-full
								             text-xs font-medium text-gray-300
								             hover:bg-primary/20 hover:border-primary/50 transition-all duration-200">
									{tech}
								</span>
							{/each}
						</div>

						<!-- Links -->
						<div class="flex gap-3 pt-1">
							{#if project.github}
								<a
									href={project.github}
									target="_blank"
									rel="noopener noreferrer"
									class="btn-link inline-flex items-center gap-2 px-4 py-2
									       bg-gray-800 border border-gray-700 rounded-lg
									       text-sm text-gray-300
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
									class="btn-link inline-flex items-center gap-2 px-4 py-2
									       bg-gray-800 border border-gray-700 rounded-lg
									       text-sm text-gray-300
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

<style>
	.project-row {
		background: linear-gradient(135deg, rgba(255,255,255,0.015) 0%, transparent 60%);
	}
</style>