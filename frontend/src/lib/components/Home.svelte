<script>
  import Section from "./Section.svelte";
  export let home; 
</script>

<Section id="home">
  <div class="min-h-[calc(100vh-4rem)] flex items-center">
    <div class="grid md:grid-cols-2 gap-8 md:gap-12 items-center w-full">

      <div class="order-1 md:order-2 flex justify-center">
        <div class="relative w-48 h-48 sm:w-60 sm:h-60 md:w-96 md:h-96">
          <div class="absolute inset-0 bg-linear-to-r from-white to-slate-400 rounded-full blur-3xl opacity-20 animate-pulse"></div>
          <div class="relative w-full h-full rounded-full overflow-hidden border-4 border-primary shadow-2xl shadow-primary/30">
            <img src={home.profileImage} alt={home.name} class="w-full h-full object-cover" />
          </div>
        </div>
      </div>

      <div class="space-y-5 order-2 md:order-1 text-center md:text-left">
        <div class="space-y-2">
          <h1 class="text-4xl sm:text-5xl md:text-7xl font-bold bg-linear-to-r from-white via-slate-200 to-slate-400 gradient-text">
            {home.name}
          </h1>
          <h2 class="text-xl sm:text-2xl md:text-3xl text-gray-300 font-light">
            {home.role}
          </h2>
        </div>

        <p class="text-base md:text-xl text-gray-400 leading-relaxed">
          {home.tagline}
        </p>

        <!-- CTA Buttons -->
        <div class="flex flex-col sm:flex-row flex-wrap gap-4 pt-4 items-center md:justify-start justify-center">

          <a href="#projects"
            class="group relative flex h-12 w-full sm:w-auto min-w-40 items-center justify-center overflow-hidden rounded-lg p-[1.5px] transition-all hover:scale-105 active:scale-95"
          >
            <div class="animate-rotate absolute inset-0 h-full w-full bg-[conic-gradient(from_0deg,transparent_20deg,rgba(255,255,255,0.8)_120deg,transparent_120deg)]"></div>
            <div class="relative flex h-full w-full items-center justify-center rounded-[7px] bg-slate-950 px-6 py-2.5 font-medium text-white transition-all group-hover:bg-slate-900">
              View Projects
            </div>
          </a>

          <a href="#contact"
            class="group relative flex h-12 w-full sm:w-auto min-w-40 items-center justify-center overflow-hidden rounded-lg p-[1.5px] transition-all hover:scale-105 active:scale-95"
          >
            <div class="animate-rotate absolute inset-0 h-full w-full bg-[conic-gradient(from_0deg,transparent_20deg,rgba(255,255,255,0.8)_120deg,transparent_120deg)]"></div>
            <div class="relative flex h-full w-full items-center justify-center rounded-[7px] bg-slate-950 px-6 py-2.5 font-medium text-white transition-all group-hover:bg-slate-900">
              Contact Me
            </div>
          </a>

          <a href={home.resumeUrl} target="_blank" rel="noopener noreferrer"
            class="group relative flex h-12 w-full sm:w-auto min-w-45 items-center justify-center overflow-hidden rounded-lg p-[1.5px] transition-all hover:scale-105 active:scale-95"
          >
            <div class="animate-rotate absolute inset-0 h-full w-full bg-[conic-gradient(from_0deg,transparent_20deg,rgba(255,255,255,0.8)_120deg,transparent_120deg)]"></div>
            <div class="relative flex h-full w-full items-center justify-center gap-2 rounded-[7px] bg-slate-950 px-6 py-2.5 font-medium text-white transition-all group-hover:bg-slate-900">
              <i class="fa-solid fa-download"></i>
              Download Resume
            </div>
          </a>

        </div>
      </div>

    </div>
  </div>
</Section>