<script>
  import Section from "./Section.svelte";
  import CertificationCard from "./CertificationCard.svelte";

  export let certifications = [];
</script>

<Section id="certifications">
  <div class="space-y-10">
    <div class="text-center space-y-4">
      <h2 class="text-4xl md:text-5xl font-bold bg-linear-to-r from-primary to-secondary gradient-text underline decoration-secondary">
        Certifications
      </h2>
      <div class="w-20 h-1 bg-linear-to-r from-primary to-secondary mx-auto rounded-full"></div>
    </div>

    <!-- Cards -->
    {#if certifications.length === 0}
      <p class="text-center text-gray-500 py-12">No certifications yet.</p>
    {:else}
      <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4">
        {#each certifications as cert}
          <CertificationCard {cert} />
        {/each}
      </div>
    {/if}
  </div>
</Section>