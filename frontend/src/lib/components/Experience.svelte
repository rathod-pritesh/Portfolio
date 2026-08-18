<script>
    import Section from "./Section.svelte";

    export let experience;

</script>

<Section id="experience">
    <div class="space-y-16">

        <div class="text-center space-y-4">
      <h2
        class="text-4xl md:text-5xl font-bold bg-linear-to-r from-primary to-secondary gradient-text underline decoration-secondary"
      >
        Experience
      </h2>
      <div
        class="w-20 h-1 bg-linear-to-r from-primary to-secondary mx-auto rounded-full"
      ></div>
    </div>

        <div class="w-20 h-1 bg-linear-to-r from-primary to-secondary mx-auto rounded-full"></div>

    </div>

    <!-- timeline -->
    <div class="max-w-4xl mx-auto relative">

        <div class="absolute left-5 top-0 bottom-0 w-px bg-gradient-to-b from-primary via-secondary to-transparent"></div>

        <div class="relative pl-14">

            <!-- node -->
            <div class="absolute left-0 top-5  w-10 h-10 rounded-full border border-primary/40 bg-dark backdrop-blur-sm flex items-center justify-center">
                <div class="w-4 h-4 rounded-full bg-primary animate-pulse"></div>
            </div>

            <!-- card -->
            <div
          class="rounded-2xl border border-white/10 bg-white/5 backdrop-blur-md p-8 hover:border-primary/40 transition-all duration-300"
        >

          <div
            class="flex flex-col md:flex-row md:items-center md:justify-between gap-3 mb-6"
          >
            <div>
              <h3 class="text-2xl md:text-3xl font-bold text-white">
                IBM SkillsBuild AICTE Intern
              </h3>

              <p class="text-primary font-medium mt-1">
                Edunet Foundation × IBM SkillsBuild
              </p>
            </div>

            <span
              class="inline-flex items-center rounded-full px-4 py-2 text-sm font-medium bg-primary/10 text-primary border border-primary/20"
            >
              Jun 2026 – Present
            </span>
          </div>

          <p class="text-gray-300 leading-relaxed mb-6">
            Selected for the IBM SkillsBuild AICTE Internship Program, a
            4-week learning experience focused on Artificial Intelligence
            and Cloud technologies. Participating in industry-oriented
            training, mentorship sessions, and hands-on activities designed
            to build practical skills in emerging technologies.
          </p>

          <div class="grid md:grid-cols-2 gap-4">

            <div class="flex gap-3">
              <i class="fa-solid fa-check text-primary mt-1"></i>
              <p class="text-gray-300">
                Learning AI and Cloud fundamentals through structured training.
              </p>
            </div>

            <div class="flex gap-3">
              <i class="fa-solid fa-check text-primary mt-1"></i>
              <p class="text-gray-300">
                Attending expert-led workshops and mentorship sessions.
              </p>
            </div>

            <div class="flex gap-3">
              <i class="fa-solid fa-check text-primary mt-1"></i>
              <p class="text-gray-300">
                Exploring practical use cases and real-world applications.
              </p>
            </div>

            <div class="flex gap-3">
              <i class="fa-solid fa-check text-primary mt-1"></i>
              <p class="text-gray-300">
                Strengthening problem-solving and modern technology skills.
              </p>
            </div>
        </div>

    </div>
</Section>