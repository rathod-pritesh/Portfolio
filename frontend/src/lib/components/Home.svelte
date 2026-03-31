<script>
  import Section from "./Section.svelte";
  export let home;
</script>

<Section id="home">
  <div class="min-h-[calc(100vh-4rem)] flex items-center">
    <div class="grid md:grid-cols-2 gap-12 md:gap-16 items-center w-full">

      <div class="order-1 md:order-2 flex justify-center items-center">

        <div class="relative flex items-center justify-center w-65 h-80 sm:w-75 sm:h-92 md:w-90 md:h-110">

          <div class="absolute top-0 right-0 w-28 h-28 rounded-full bg-slate-400/10 blur-3xl animate-[orb-float_6s_ease-in-out_infinite]"></div>
          <div class="absolute bottom-2 left-0 w-20 h-20 rounded-full bg-slate-500/15 blur-2xl animate-[orb-float_6s_ease-in-out_-2s_infinite]"></div>
          <div class="absolute top-1/2 -right-4 w-14 h-14 rounded-full bg-slate-300/10 blur-xl animate-[orb-float_6s_ease-in-out_-4s_infinite]"></div>

          <div class="group relative w-55 h-70 sm:w-65 sm:h-82 md:w-75 md:h-95 -rotate-3 hover:rotate-0 hover:scale-105 transition-transform duration-500 ease-out cursor-pointer">

            <div class="absolute -inset-px rounded-2xl animate-[spin-border_4s_linear_infinite] group-hover:animate-[spin-border_1.5s_linear_infinite] z-0"
              style="background: conic-gradient(from 0deg, transparent 0deg, rgba(255,255,255,0.6) 60deg, rgba(148,163,184,0.8) 120deg, transparent 180deg)">
            </div>

            <div class="absolute -top-1.5 -left-1.5 w-4 h-4 border-t-2 border-l-2 border-slate-300/70 rounded-tl group-hover:border-white group-hover:w-5 group-hover:h-5 transition-all duration-300 z-10"></div>
            <div class="absolute -top-1.5 -right-1.5 w-4 h-4 border-t-2 border-r-2 border-slate-300/70 rounded-tr group-hover:border-white group-hover:w-5 group-hover:h-5 transition-all duration-300 z-10"></div>
            <div class="absolute -bottom-1.5 -left-1.5 w-4 h-4 border-b-2 border-l-2 border-slate-300/70 rounded-bl group-hover:border-white group-hover:w-5 group-hover:h-5 transition-all duration-300 z-10"></div>
            <div class="absolute -bottom-1.5 -right-1.5 w-4 h-4 border-b-2 border-r-2 border-slate-300/70 rounded-br group-hover:border-white group-hover:w-5 group-hover:h-5 transition-all duration-300 z-10"></div>

            <div class="relative w-full h-full rounded-2xl overflow-hidden bg-slate-950 shadow-2xl group-hover:shadow-[0_35px_80px_rgba(0,0,0,0.7),0_0_0_1px_rgba(255,255,255,0.1)] transition-shadow duration-500 z-1">

              <img
                src={home.profileImage}
                alt={home.name}
                class="w-full h-full object-cover object-top brightness-92 contrast-105 group-hover:brightness-100 group-hover:scale-106 transition-all duration-500 ease-out"
              />

              <div class="absolute inset-0 bg-gradient-to-r from-transparent via-white/5 to-transparent -translate-x-full group-hover:translate-x-full transition-transform duration-700 pointer-events-none"></div>

              <div class="absolute bottom-0 left-0 right-0 pt-7 pb-3 px-3.5 bg-gradient-to-t from-slate-950/85 to-transparent flex items-center gap-2 translate-y-1 opacity-80 group-hover:opacity-100 group-hover:translate-y-0 transition-all duration-300">
                <span class="size-2 rounded-full bg-slate-400 shrink-0 animate-[pulse-dot_2s_ease-in-out_infinite]"></span>
                <span class="text-xs font-medium text-slate-300 tracking-widest uppercase">{home.name}</span>
              </div>

            </div>
          </div>

          <!-- Side accent lines (hidden on mobile) -->
          <div class="hidden sm:block absolute left-3 top-1/2 -translate-y-1/2 w-px h-3/5 bg-gradient-to-b from-transparent via-slate-500/50 to-transparent"></div>
          <div class="hidden sm:block absolute right-3 top-1/2 -translate-y-1/2 w-px h-3/5 bg-gradient-to-b from-transparent via-slate-500/50 to-transparent"></div>

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

        <div class="flex flex-col sm:flex-row flex-wrap gap-4 pt-4 items-center md:justify-start justify-center">

          <a href="#projects"
            class="group relative flex h-12 w-full sm:w-auto min-w-40 items-center justify-center overflow-hidden rounded-lg p-px transition-all hover:scale-105 active:scale-95"
          >
            <div class="animate-rotate absolute inset-0 h-full w-full bg-[conic-gradient(from_0deg,transparent_20deg,rgba(255,255,255,0.8)_120deg,transparent_120deg)]"></div>
            <div class="relative flex h-full w-full items-center justify-center rounded-lg bg-slate-950 px-6 py-2.5 font-medium text-white transition-all group-hover:bg-slate-900">
              View Projects
            </div>
          </a>

          <a href="#contact"
            class="group relative flex h-12 w-full sm:w-auto min-w-40 items-center justify-center overflow-hidden rounded-lg p-px transition-all hover:scale-105 active:scale-95"
          >
            <div class="animate-rotate absolute inset-0 h-full w-full bg-[conic-gradient(from_0deg,transparent_20deg,rgba(255,255,255,0.8)_120deg,transparent_120deg)]"></div>
            <div class="relative flex h-full w-full items-center justify-center rounded-lg bg-slate-950 px-6 py-2.5 font-medium text-white transition-all group-hover:bg-slate-900">
              Contact Me
            </div>
          </a>

          <a href={home.resumeUrl} target="_blank" rel="noopener noreferrer"
            class="group relative flex h-12 w-full sm:w-auto min-w-45 items-center justify-center overflow-hidden rounded-lg p-px transition-all hover:scale-105 active:scale-95"
          >
            <div class="animate-rotate absolute inset-0 h-full w-full bg-[conic-gradient(from_0deg,transparent_20deg,rgba(255,255,255,0.8)_120deg,transparent_120deg)]"></div>
            <div class="relative flex h-full w-full items-center justify-center gap-2 rounded-lg bg-slate-950 px-6 py-2.5 font-medium text-white transition-all group-hover:bg-slate-900">
              <i class="fa-solid fa-download"></i>
              Download Resume
            </div>
          </a>

        </div>
      </div>

    </div>
  </div>
</Section>

<style>
  @keyframes orb-float {
    0%, 100% { transform: translateY(0) scale(1); }
    50%       { transform: translateY(-14px) scale(1.05); }
  }
  @keyframes spin-border {
    to { transform: rotate(360deg); }
  }
  @keyframes pulse-dot {
    0%, 100% { opacity: 1;   transform: scale(1); }
    50%       { opacity: 0.4; transform: scale(0.75); }
  }
</style>