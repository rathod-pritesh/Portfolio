<script>
  import Section from "./Section.svelte";
  export let home; 
</script>

<Section id="home">
  <div class="min-h-[calc(100vh-4rem)] flex items-center">
    <div class="grid md:grid-cols-2 gap-12 items-center w-full">

      <!-- Left -->
      <div class="space-y-6 order-2 md:order-1">
        <div class="space-y-2">
          <h1 class="text-5xl md:text-7xl font-bold bg-gradient-to-r from-white via-slate-200 to-slate-400 gradient-text">
            {home.name}
          </h1>
          <h2 class="text-2xl md:text-3xl text-gray-300 font-light">
            {home.role}
          </h2>
        </div>

        <p class="text-lg md:text-xl text-gray-400 leading-relaxed">
          {home.tagline}
        </p>

        <!-- CTA Buttons -->
        <div class="flex flex-wrap gap-6 pt-4 items-center justify-center">

          <a href="#projects"
            class="group relative flex h-12 min-w-[160px] items-center justify-center overflow-hidden rounded-lg p-[1.5px] transition-all hover:scale-105 active:scale-95"
          >
            <div class="animate-rotate absolute inset-0 h-full w-full bg-[conic-gradient(from_0deg,transparent_20deg,rgba(255,255,255,0.8)_120deg,transparent_120deg)]"></div>
            <div class="relative flex h-full w-full items-center justify-center rounded-[7px] bg-slate-950 px-6 py-2.5 font-medium text-white transition-all group-hover:bg-slate-900">
              View Projects
            </div>
          </a>

          <a href="#contact"
            class="group relative flex h-12 min-w-[160px] items-center justify-center overflow-hidden rounded-lg p-[1.5px] transition-all hover:scale-105 active:scale-95"
          >
            <div class="animate-rotate absolute inset-0 h-full w-full bg-[conic-gradient(from_0deg,transparent_20deg,rgba(255,255,255,0.8)_120deg,transparent_120deg)]"></div>
            <div class="relative flex h-full w-full items-center justify-center rounded-[7px] bg-slate-950 px-6 py-2.5 font-medium text-white transition-all group-hover:bg-slate-900">
              Contact Me
            </div>
          </a>

          <a href={home.resumeUrl} target="_blank" rel="noopener noreferrer"
            class="group relative flex h-12 min-w-[180px] items-center justify-center overflow-hidden rounded-lg p-[1.5px] transition-all hover:scale-105 active:scale-95"
          >
            <div class="animate-rotate absolute inset-0 h-full w-full bg-[conic-gradient(from_0deg,transparent_20deg,rgba(255,255,255,0.8)_120deg,transparent_120deg)]"></div>
            <div class="relative flex h-full w-full items-center justify-center gap-2 rounded-[7px] bg-slate-950 px-6 py-2.5 font-medium text-white transition-all group-hover:bg-slate-900">
              <i class="fa-solid fa-download"></i>
              Download Resume
            </div>
          </a>

        </div>
      </div>

      <!-- Right -->
      <div class="order-1 md:order-2 flex justify-center">
        <div class="relative w-72 h-72 md:w-96 md:h-96">
          <div class="absolute inset-0 bg-gradient-to-r from-white to-slate-400 rounded-full blur-3xl opacity-20 animate-pulse"></div>
          <div class="relative w-full h-full rounded-full overflow-hidden border-4 border-primary shadow-2xl shadow-primary/30">
            <img src={home.profileImage} alt={home.name} class="w-full h-full object-cover" />
          </div>
        </div>
      </div>

    </div>
  </div>
</Section>