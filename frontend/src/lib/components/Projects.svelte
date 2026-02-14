<script>
	import Section from './Section.svelte';

	// const projects = {
	// 	title: 'Featured Projects',
	// 	items: [
	// 		{
	// 			id: 1,
	// 			title: 'Personal Voice Assistant AI',
	// 			description:
	// 				'Web-based AI voice assistant that understands natural language commands, detects user intent using a BERT-based NLP model, and performs real-time actions through a fast SvelteKit interface with a FastAPI backend.',
	// 			image: '/projects/ai-voice-assistant.webp',
	// 			github: 'https://github.com/rathod-pritesh/ai-voice-assistant',
	// 			technologies: [
	// 				'SvelteKit',
	// 				'FastAPI',
	// 				'Python',
	// 				'NLP',
	// 				'all-MiniLM-L6-v2',
	// 				'Web Speech API'
	// 			],
	// 			gradient: 'from-indigo-500 to-purple-600'
	// 		},
	// 		{
	// 			id: 2,
	// 			title: 'BuildTrack – Construction Management System',
	// 			description:
	// 				'A role-based construction management system that digitizes project planning, attendance tracking, cost estimation, and reporting. Designed to improve transparency and efficiency across admins, members, and clients.',
	// 			image: '/projects/buildtrack.webp',
	// 			github: 'https://github.com/rathod-pritesh/BuildTrack-construction-management-system',
	// 			technologies: [
	// 				'Flask',
	// 				'Python',
	// 				'MySQL',
	// 				'Bootstrap'
	// 			],
	// 			gradient: 'from-orange-500 to-amber-500'
	// 		},
	// 		{
	// 			id: 3,
	// 			title: 'Netflix Watchlist API',
	// 			description:
	// 				'RESTful API built with Go and MongoDB for managing a movie watchlist. Supports CRUD operations, watch status updates, and ObjectID handling using a clean MVC-style architecture.',
	// 			image: '/projects/api-image.webp',
	// 			github: 'https://github.com/rathod-pritesh/go-mongodb-crud-api',
	// 			technologies: [
	// 				'Go',
	// 				'MongoDB',
	// 				'Gorilla Mux',
	// 				'REST API'
	// 			],
	// 			gradient: 'from-red-500 to-rose-600'
	// 		},
	// 		{
	// 			id: 4,
	// 			title: 'AI Image Generator',
	// 			description:
	// 				'AI-powered image generation web app that converts text prompts into images using the Pollinations AI API. Features a modern SvelteKit frontend and a FastAPI backend for seamless prompt handling.',
	// 			image: '/projects/ai-image-generator.webp',
	// 			github: 'https://github.com/rathod-pritesh/ai-image-genetator',
	// 			technologies: [
	// 				'SvelteKit',
	// 				'FastAPI',
	// 				'Pollinations AI',
	// 				'Base64 Images',
	// 				'REST API'
	// 			],
	// 			gradient: 'from-cyan-500 to-blue-600'
	// 		}
	// 	]
	// };

	export let projects
</script>

<Section id="projects">
	<div class="space-y-12">
		<div class="text-center space-y-4">
			<h2 class="text-4xl md:text-5xl font-bold bg-gradient-to-r from-primary to-secondary gradient-text text-primary underline decoration-[#6366f1]">
				Featured Projects
			</h2>
			<div class="w-20 h-1 bg-gradient-to-r from-primary to-secondary mx-auto"></div>
		</div>

		<div class="grid md:grid-cols-2 gap-8">
			{#each projects as project}
				<div class="group bg-dark/50 border border-gray-800 rounded-xl overflow-hidden hover:border-primary transition-all duration-300 hover:shadow-xl hover:shadow-primary/20 hover:scale-[1.02]">
					<div class="relative h-64 overflow-hidden">
						<img
							src={project.image}
							alt={project.title}
							class="w-full h-full object-cover group-hover:scale-110 transition-transform duration-500"
						/>
						<div class="absolute inset-0 bg-gradient-to-t {project.gradient} opacity-0 group-hover:opacity-30 transition-opacity duration-300"></div>
					</div>
					<div class="p-6 space-y-4">
						<h3 class="text-2xl font-bold text-white group-hover:text-primary transition-colors">
							{project.title}
						</h3>
						<p class="text-gray-400 leading-relaxed">
							{project.description}
						</p>
						<div class="flex flex-wrap gap-2">
							{#each project.technologies as tech}
								<span class="px-3 py-1 bg-primary/20 border border-primary/30 rounded-full text-xs font-medium">
									{tech}
								</span>
							{/each}
						</div>
						<div class="flex gap-4 pt-4">
							{#if project.github}
								<a 
									href={project.github}
									target="_blank"
									rel="noopener noreferrer"
									class="px-4 py-2 bg-gray-800 border border-gray-700
				   rounded-lg text-sm hover:border-primary
				   hover:text-primary transition"
								>
								<i class="fa-brands fa-github"></i>
									GitHub
								</a>
							{/if}
						</div>
					</div>
				</div>
			{/each}
		</div>
	</div>
</Section>
