<script>
  import AutomationCard from "./AutomationCard.svelte";

  export let automations = [];
</script>

<section id="automations" class="min-h-100">
  <div class="max-w-6xl mx-auto">

    <div class="text-center space-y-4">
			<h2 class="text-4xl md:text-5xl font-bold bg-linear-to-r from-primary to-secondary gradient-text underline decoration-secondary">
				Automations
			</h2>
      
			<div class="w-20 h-1 bg-linear-to-r from-primary to-secondary mx-auto rounded-full"></div>

      <p class="text-gray-400 text-center mb-12">
      Workflows I've built with automation tools
    </p>
		</div>

    <div class="space-y-8">

      {#each automations as automation}
        <AutomationCard {automation}/>
      {/each}

    </div>

  </div>
</section>