<script>
  import Section from "./Section.svelte";
  export let skills; 

  let activeTab;

  $: if (skills?.categories?.length && !activeTab) {
    activeTab = skills.categories[0].name;
  }
</script>

<Section id="skills">
  <div class="space-y-12">

    <!-- Heading -->
    <div class="text-center space-y-4">
      <h2 class="text-4xl md:text-5xl font-bold bg-gradient-to-r from-primary to-secondary gradient-text underline decoration-[#94a3b8]">
        {skills.title}
      </h2>
      <div class="w-20 h-1 bg-gradient-to-r from-primary to-secondary mx-auto rounded-full"></div>
    </div>

    <!-- Category tabs-->
    <div class="flex flex-wrap justify-center gap-4">
      {#each skills.categories as category}
        <button
          class="px-6 py-2 rounded-full border transition-all duration-300
                 {activeTab === category.name
                   ? 'bg-gradient-to-r from-primary to-secondary text-slate-900 border-primary font-semibold'
                   : 'bg-dark/40 text-gray-300 border-gray-700 hover:border-primary hover:text-white'}"
          on:click={() => (activeTab = category.name)}
        >
          {category.icon} {category.name}
        </button>
      {/each}
    </div>

    <!-- Active tab skill cards -->
    {#each skills.categories as category}
      {#if activeTab === category.name}
        <div class="bg-dark/50 border border-gray-800 rounded-xl p-8 hover:border-primary transition-all duration-300 animate-fadeIn">

          <div class="flex items-center gap-4 mb-6">
            <span class="text-4xl">{category.icon}</span>
            <h3 class="text-2xl font-bold text-white">{category.name}</h3>
          </div>

          <div class="flex flex-wrap gap-4">
            {#each category.skills as item}
              <div class="flex items-center gap-3 px-4 py-3
                          bg-gradient-to-r from-primary/10 to-secondary/10
                          border border-primary/25 rounded-xl
                          hover:scale-105 hover:border-primary/50
                          transition-all duration-300">
                <img src={item.img} alt={item.name} class="w-6 h-6 object-contain" />
                <span class="text-sm font-medium text-gray-200">{item.name}</span>
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
    to   { opacity: 1; transform: translateY(0); }
  }
  .animate-fadeIn { animation: fadeIn 0.3s ease-out; }
</style>