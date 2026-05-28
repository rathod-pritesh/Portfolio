<script>
  import Section from "./Section.svelte";
  export let about;
  export let education = [];
</script>

<Section id="about">
  <div class="space-y-12">
    <!-- heading -->
    <div class="text-center space-y-4">
      <h2
        class="text-4xl md:text-5xl font-bold bg-linear-to-r from-primary to-secondary gradient-text underline decoration-secondary"
      >
        {about.title}
      </h2>
      <div
        class="w-20 h-1 bg-linear-to-r from-primary to-secondary mx-auto rounded-full"
      ></div>
    </div>

    <!-- desc -->
    <div class="max-w-3xl mx-auto space-y-5">
      {#each about.description
        .split("\n\n")
        .filter((p) => p.trim()) as paragraph}
        <p class="text-lg md:text-xl text-gray-300 leading-relaxed text-center">
          {paragraph}
        </p>
      {/each}
    </div>

    <!-- education -->
    <div class="max-w-5xl mx-auto w-full pt-8">
      <div class="text-center mb-16">
        <h3 class="text-3xl md:text-4xl font-bold gradient-text inline-block">
          Education
        </h3>

        <div class="w-20 h-px bg-white/20 mx-auto mt-4"></div>
      </div>

      <div class="relative">
        <!-- center line -->
        <div
          class="education-line absolute left-1/2 top-0 bottom-0 w-px hidden md:block"
        ></div>

        <div class="space-y-20">
          {#each education as edu, i}
            <div
              class={`relative flex flex-col md:flex-row items-center ${i % 2 === 0 ? "md:flex-row" : "md:flex-row-reverse"}`}
            >
              <!-- content -->
              <div class="w-full md:w-1/2 px-0 md:px-10">
                <div
                  class={`space-y-3 ${i % 2 === 0 ? "md:text-right" : "md:text-left"} text-center md:text-inherit`}
                >
                  <!-- year -->
                  <p
                    class="text-xs tracking-[0.35em] uppercase text-primary font-semibold"
                  >
                    {edu.yearStart} — {edu.yearEnd}
                  </p>

                  <!-- degree -->
                  <h4
                    class="text-2xl md:text-3xl font-bold text-white leading-tight"
                  >
                    {edu.degree}
                  </h4>

                  <!-- institute -->
                  <p class="text-base text-white/55">
                    {edu.institution}
                  </p>

                  <!-- focus -->
                  {#if edu.focus?.trim()}
                    <p
                      class="text-sm md:text-base text-white/40 leading-relaxed max-w-md
    {i % 2 === 0 ? 'md:ml-auto' : ''}"
                    >
                      {edu.focus}
                    </p>
                  {/if}
                </div>
              </div>

              <!-- center node -->
              <div
                class="hidden md:flex absolute left-1/2 -translate-x-1/2 items-center justify-center"
              >
                <div
                  class="w-5 h-5 rounded-full border border-primary/40 bg-dark flex items-center justify-center backdrop-blur-sm"
                >
                  <div class="w-2 h-2 rounded-full bg-primary"></div>
                </div>
              </div>

              <!-- empty side -->
              <div class="hidden md:block w-1/2"></div>
            </div>
          {/each}
        </div>
      </div>
    </div>
  </div>
</Section>
