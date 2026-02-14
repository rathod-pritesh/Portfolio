<script>
  import Section from "./Section.svelte";

  // const skills = {
  //   title: "Skills & Expertise",
  //   categories: [
  //     {
  //       name: "Frontend Development",
  //       icon: "💻",
  //       skills: [
  //         { name: "Svelte", img: "/skills/Svelte.png" },
  //         { name: "BootStrap", img: "/skills/Bootstrap.png" },
  //         { name: "TailwindCSS", img: "/skills/Tailwind.png" },
  //         { name: "JavaScript", img: "/skills/JavaScript.png" },
  //       ],
  //     },
  //     {
  //       name: "Backend Development",
  //       icon: "⚙️",
  //       skills: [
  //         { name: "Golang", img: "/skills/Go.png" },
  //         { name: "Python", img: "/skills/Python.png" },
  //         { name: "Java", img: "/skills/Java.png" },
  //         { name: "MySQL", img: "/skills/MySQL.png" },
  //         { name: "MongoDB", img: "/skills/MongoDB.png" },
  //         { name: "Flask", img: "/skills/Flask.png" },
  //         { name: "FastAPI", img: "/skills/FastAPI.png" },
  //       ],
  //     },
  //     {
  //       name: "UI/UX Design",
  //       icon: "🎨",
  //       skills: [{ name: "Figma", img: "/skills/Figma.png" }],
  //     },
  //     {
  //       name: "Tools & Others",
  //       icon: "🛠️",
  //       skills: [
  //         { name: "Git", img: "/skills/Git.png" },
  //         { name: "GitHub", img: "/skills/GitHub.png" },
  //         { name: "n8n", img: "/skills/n8n.png" },
  //       ],
  //     },
  //   ],
  // };
  export let skills

  let activeTab;

  $: if (skills?.categories?.length && !activeTab) {
    activeTab = skills.categories[0].name;
  }
</script>

<Section id="skills">
  <div class="space-y-12">
    <div class="text-center space-y-4">
      <h2
        class="text-4xl md:text-5xl font-bold bg-gradient-to-r from-primary to-secondary gradient-text text-primary underline decoration-[#6366f1]"
      >
        {skills.title}
      </h2>
      <div
        class="w-20 h-1 bg-gradient-to-r from-primary to-secondary mx-auto"
      ></div>
    </div>
		<!-- Tabs -->
		<div
			class="flex flex-wrap justify-center gap-4 mb-10"
		>
			{#each skills.categories as category}
				<button
					class="px-6 py-2 rounded-full border transition-all duration-300
					{activeTab === category.name
						? 'bg-gradient-to-r from-primary to-secondary text-white border-primary'
						: 'bg-dark/40 text-gray-300 border-gray-700 hover:border-primary'}"
					on:click={() => activeTab = category.name}
				>
					{category.icon} {category.name}
				</button>
			{/each}
		</div>

		<!-- Active tab content -->
		{#each skills.categories as category} 
			{#if activeTab === category.name}
				<div
					class="bg-dark/50 border border-gray-800 rounded-xl p-8 hover:border-primary transition-all duration-300 animate-fadeIn"
				>

					<div
						class="flex items-center gap-4 mb-6"
					>
						<span class="text-4xl">{category.icon}</span>
						<h3 class="text-2xl font-bold text-white">
							{category.name}
						</h3>
					</div>

					<div class="flex flex-wrap gap-4">
						{#each category.skills as item}

							<div class="flex items-center gap-3 px-4 py-3 bg-gradient-to-r from-primary/10 to-secondary/10 border border-primary/30 rounded-xl hover:scale-105 transition-all duration-300">

								<img src={item.img} alt={item.name} class="w-6 h-6 object-contain" />

								<span class="text-sm font-medium text-white">
									{item.name}
								</span>
							</div>
						{/each}
					</div>

				</div>
			{/if}
		{/each}
  </div>
</Section>

<style>
	@keyframes fadeIn {
		from { opacity: 0; transform: translateY(10px); }
		to { opacity: 1; transform: translateY(0); }
	}

	.animate-fadeIn {
		animation: fadeIn 0.3s ease-out;
	}
</style>