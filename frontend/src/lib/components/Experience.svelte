<script>
  import Section from "./Section.svelte";

  export let experience;
</script>

<Section id="experience">
  <div class="text-center space-y-4 mb-16">
    <h2
      class="text-4xl md:text-5xl font-bold gradient-text underline decoration-secondary/60 decoration-2 underline-offset-4"
    >
      Experience
    </h2>
    <div class="w-16 h-1 bg-secondary/40 mx-auto rounded-full"></div>
  </div>

  <!-- timeline -->
  <div class="max-w-4xl mx-auto relative">
    <div class="education-line absolute left-5 top-0 bottom-0 w-px"></div>

    <div class="relative pl-14">
      <!-- node -->
      <div
        class="absolute left-0 top-6 w-10 h-10 rounded-full border border-secondary/30 bg-dark backdrop-blur-sm flex items-center justify-center"
      >
        <div class="w-3.5 h-3.5 rounded-full bg-secondary"></div>
      </div>

      <!-- card -->
      <div
        class="rounded-2xl border border-secondary/15 bg-primary/[0.03] p-6 md:p-8 transition-all duration-300 hover:border-secondary/30 hover:-translate-y-1"
      >
        <p class="text-xs tracking-widest uppercase text-secondary/80 font-medium">
          Jun 2026 – Jul 2026
        </p>

        <h3 class="text-xl md:text-2xl font-bold text-primary mt-2">
          IBM SkillsBuild AI & Cloud Intern
        </h3>
        <p class="text-accent font-medium mt-1">
          Edunet Foundation × IBM
        </p>

        <p class="mt-5 text-secondary leading-relaxed">
          Worked hands-on with IBM Watsonx.ai, Watsonx Orchestrate, and IBM
          Cloud services while gaining practical experience in AI and cloud
          technologies.
        </p>

        <ul class="mt-6 space-y-3">
          <li class="flex items-start gap-3">
            <i class="fa-solid fa-check text-secondary/70 mt-1 text-sm"></i>
            <span class="text-secondary">
              Explored Watsonx.ai, Watsonx Orchestrate, and IBM Cloud
              services.
            </span>
          </li>
          <li class="flex items-start gap-3">
            <i class="fa-solid fa-check text-secondary/70 mt-1 text-sm"></i>
            <span class="text-secondary">
              Applied AI and cloud concepts through hands-on projects and
              activities.
            </span>
          </li>
          <li class="flex items-start gap-3">
            <i class="fa-solid fa-check text-secondary/70 mt-1 text-sm"></i>
            <span class="text-secondary">
              Strengthened backend and AI development skills through
              practical work.
            </span>
          </li>
        </ul>

        <div class="mt-8">
          <a
            href="/certificates/Pritesh_Rathod_Edunet_Internship_Certificate.pdf"
            target="_blank"
            rel="noopener noreferrer"
            aria-label="View IBM SkillsBuild AI & Cloud Intern certificate (opens PDF in a new tab)"
            class="inline-flex items-center gap-2 rounded-xl border border-secondary/25 px-5 py-3 text-sm font-medium text-primary
              transition-all duration-200 hover:border-accent/40 hover:text-accent
              focus-visible:outline focus-visible:outline-2 focus-visible:outline-accent focus-visible:outline-offset-2"
          >
            <i class="fa-solid fa-certificate"></i>
            View Certificate
          </a>
        </div>
      </div>
    </div>
  </div>
</Section>