<script>
  import Section from "./Section.svelte";
  export let about;
  export let education = [];
</script>

<Section id="about">
  <div class="space-y-12">

    <!-- heading -->
    <div class="text-center space-y-4">
      <h2 class="text-4xl md:text-5xl font-bold bg-linear-to-r from-primary to-secondary gradient-text underline decoration-secondary">
        {about.title}
      </h2>
      <div class="w-20 h-1 bg-linear-to-r from-primary to-secondary mx-auto rounded-full"></div>
    </div>

    <!-- desc -->
    <div class="max-w-3xl mx-auto">
      <p class="text-lg md:text-xl text-gray-300 leading-relaxed text-center">
        {about.description}
      </p>
    </div>

    <!-- education -->
    <div class="max-w-2xl mx-auto w-full pt-4">
      <h3 class="text-2xl font-bold text-white mb-8 text-center underline decoration-primary">Education</h3>

      <div class="flex flex-col">
        {#each education as edu, i}

          <!-- left column : Dot + line -->
          <div class="flex gap-6">

            <div class="flex flex-col items-center">
              <!-- dot -->
              <div class="w-7 h-7 rounded-full border-2 flex items-center justify-center shrink-0
                transition-all duration-300 hover:scale-125 hover:shadow-lg z-10 cursor-default
                {edu.status === 'In Progress'
                  ? 'border-primary bg-primary/20 hover:bg-primary/40 hover:shadow-primary/40'
                  : 'border-gray-600 bg-transparent hover:border-primary hover:shadow-primary/30'}">
                <div class="w-2.5 h-2.5 rounded-full transition-all duration-300
                  {edu.status === 'In Progress'
                    ? 'bg-primary animate-pulse'
                    : 'bg-gray-500 hover:bg-primary'}">
                </div>
              </div>
              <!-- line -->
              {#if i < education.length - 1}
                <div class="w-px bg-gray-700 grow mt-1"></div>
              {/if}
            </div>

            <!-- right: gap for cotnent -->
            <div class="flex-1 {i < education.length - 1 ? 'pb-8' : 'pb-0'}">
              <h4 class="text-base font-bold text-white leading-snug mt-0.5">
                {edu.degree}
              </h4>
              <p class="text-sm text-gray-400 mt-1 leading-relaxed">
                {#if edu.institution}{edu.institution} | {/if} {edu.yearStart} – {edu.yearEnd}{#if edu.status !== 'In Progress' && edu.score} | {edu.score}{/if}
              </p>
              {#if edu.status === 'In Progress'}
                <span class="inline-block mt-1.5 text-xs text-primary font-medium">In Progress</span>
              {/if}
            </div>

          </div>

        {/each}
      </div>

    </div>

  </div>
</Section>