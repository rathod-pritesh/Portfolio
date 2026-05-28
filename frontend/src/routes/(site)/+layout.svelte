<script>
  import "$lib/app.css";
  import { onMount } from "svelte";
  import { browser } from "$app/environment";
  import ChatBot from "$lib/components/ChatBot.svelte";

  let showScrollTop = false;
  let activeSection = "home";
  let isMenuOpen = false;
  let isDark = false;
  let skillsDropdownOpen = false;

  function toggleSkillsDropdown(e) {
    e.stopPropagation();
    skillsDropdownOpen = !skillsDropdownOpen;
  }

  function closeSkillsDropdown() {
    skillsDropdownOpen = false;
  }

  const navItems = [
    { label: "Home", href: "#home", id: "home" },
    { label: "About", href: "#about", id: "about" },
    { label: "Skills", href: "#skills", id: "skills" },
    { label: "Projects", href: "#projects", id: "projects" },
    { label: "Contact", href: "#contact", id: "contact" },
  ];

  function toggleTheme() {
    isDark = !isDark;
    document.documentElement.dataset.theme = isDark ? "dark" : "light";
  }

  onMount(() => {
    const handleScroll = () => {
      showScrollTop = window.scrollY > 300;
      for (const item of navItems) {
        const el = document.getElementById(item.id);
        if (el) {
          const rect = el.getBoundingClientRect();
          if (rect.top <= 120 && rect.bottom >= 120) {
            activeSection = item.id;
            break;
          }
        }
      }
    };

    document.documentElement.dataset.theme = "light";
    window.addEventListener("scroll", handleScroll);
    return () => window.removeEventListener("scroll", handleScroll);
  });

  $: if (browser) {
    document.body.style.overflow = isMenuOpen ? "hidden" : "auto";
  }

  function scrollToTop() {
    window.scrollTo({ top: 0, behavior: "smooth" });
  }
  function closeMenu() {
    isMenuOpen = false;
  }
</script>

<div class="min-h-screen relative">
  <!-- Navbar -->
  <nav
    class="fixed top-0 left-0 right-0 z-50 bg-black border-b border-gray-800"
  >
    <div class="container-max">
      <div class="flex items-center justify-between py-4 px-6">
        <!-- Logo -->
        <a
          href="#home"
          class="text-2xl font-bold bg-linear-to-r from-primary to-secondary gradient-text"
          aria-label="home"
        >
          Pritesh Rathod
        </a>

        <!-- Desktop links -->
        <ul class="hidden md:flex items-center gap-8">
          {#each navItems as item}
            <li class="relative">
              {#if item.id === "skills"}
                <!-- skills with dropdown -->
                <div class="relative">
                  <button
                    on:click={toggleSkillsDropdown}
                    class="relative flex items-center gap-1 text-gray-400 hover:text-white transition-colors duration-300 group {activeSection ===
                      'skills' || activeSection === 'certifications'
                      ? 'text-primary'
                      : ''}"
                    >{item.label}
                    <i
                      class="fa-solid fa-chevron-down text-[10px] transition-transform duration-200
                              {skillsDropdownOpen ? 'rotate-180' : ''}"
                    ></i>

                    <span
                      class="absolute left-0 -bottom-1 h-0.5 w-full bg-primary origin-left transform transition-transform duration-300
                             {activeSection === 'skills' ||
                      activeSection === 'certifications'
                        ? 'scale-x-100'
                        : 'scale-x-0 group-hover:scale-x-100'}"
                    ></span>
                  </button>

                  {#if skillsDropdownOpen}
                    <div
                      class="nav-dropdown absolute top-full left-1/2 -translate-x-1/2 mt-3 z-50 w-44 rounded-xl bg-gray-900 border border-white/10 shadow-xl animate-fadeIn"
                    >
                      <div class="py-1.5">
                        <a
                          href="#skills"
                          on:click={closeSkillsDropdown}
                          class="flex items-center gap-2 px-4 py-2 text-sm text-gray-300 hover:bg-white/5 hover:text-white transition-colors duration-150"
                        >
                          <i class="fa-solid fa-code text-xs text-primary"></i>
                          Tech Skills
                        </a>
                        <a
                          href="#certifications"
                          on:click={closeSkillsDropdown}
                          class="flex items-center gap-2 px-4 py-2 text-sm text-gray-300
                                 hover:bg-white/5 hover:text-white transition-colors duration-150"
                        >
                          <i
                            class="fa-solid fa-certificate text-xs text-primary"
                          ></i>
                          Certifications
                        </a>
                      </div>
                    </div>
                  {/if}
                </div>
              {:else}
                <a
                  href={item.href}
                  class="relative text-gray-400 hover:text-white transition-colors duration-300 group"
                  class:text-primary={activeSection === item.id}
                >
                  {item.label}
                  <span
                    class="absolute left-0 -bottom-1 h-0.5 w-full bg-primary origin-left transform transition-transform duration-300
                           {activeSection === item.id
                      ? 'scale-x-100'
                      : 'scale-x-0 group-hover:scale-x-100'}"
                  ></span>
                </a>
              {/if}
            </li>
          {/each}
        </ul>

        <button
          on:click={toggleTheme}
          class="flex items-center justify-center w-9 h-9 rounded-full border border-gray-600 text-gray-300 hover:border-primary hover:text-primary transition-colors"
          aria-label="Toggle theme"
        >
          {#if isDark}
            <!-- Sun Icon -->
            <svg
              xmlns="http://www.w3.org/2000/svg"
              class="w-5 h-5"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor"
              stroke-width="2"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                d="M12 3v2m0 14v2m9-9h-2M5 12H3m15.364 6.364-1.414-1.414M7.05 7.05 5.636 5.636m12.728 0-1.414 1.414M7.05 16.95l-1.414 1.414M12 8a4 4 0 100 8 4 4 0 000-8z"
              />
            </svg>
          {:else}
            <!-- Moon Icon -->
            <svg
              xmlns="http://www.w3.org/2000/svg"
              class="w-5 h-5"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor"
              stroke-width="2"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                d="M21 12.79A9 9 0 1111.21 3c0 .34.02.67.05 1A7 7 0 0020 13c.33.03.66.05 1 .05z"
              />
            </svg>
          {/if}
        </button>

        <!-- Mobile hamburger -->
        <button
          class="md:hidden text-white text-2xl transition-transform duration-200"
          on:click={() => (isMenuOpen = !isMenuOpen)}
          aria-label={isMenuOpen ? "Close Menu" : "Open Menu"}
        >
          <i class={isMenuOpen ? "fa-solid fa-xmark" : "fa-solid fa-bars"}></i>
        </button>
      </div>

      {#if isMenuOpen}
        <!-- Backdrop -->
        <button
          type="button"
          class="fixed inset-0 bg-black/40 z-40 md:hidden cursor-default"
          on:click={closeMenu}
          aria-label="Close menu"
        ></button>

        <!-- Mobile menu -->
        <div
          class="mobile-menu fixed top-16 left-0 right-0 z-50 md:hidden bg-black border-t border-gray-800 px-6 py-4 animate-slideDown"
        >
          <ul class="space-y-4">
            {#each navItems as item}
              <li>
                <a
                  href={item.href}
                  class="block text-gray-400 hover:text-white transition-colors py-1"
                  class:text-primary={activeSection === item.id}
                  on:click={closeMenu}
                >
                  {item.label}
                </a>
                {#if item.id === "skills"}
                  <div
                    class="pl-4 mt-1 space-y-1 border-l border-gray-700 ml-1"
                  >
                    <a
                      href="#skills"
                      on:click={closeMenu}
                      class="flex items-center gap-2 text-gray-500 hover:text-white transition-colors py-1 text-sm"
                      class:text-primary={activeSection === "skills"}
                      ><i class="fa-solid fa-code text-xs"></i> Tech Skills</a
                    >
                    <a
                      href="#certifications"
                      on:click={closeMenu}
                      class="flex items-center gap-2 text-gray-500 hover:text-white transition-colors py-1 text-sm"
                      class:text-primary={activeSection === "certifications"}
                    >
                      <i class="fa-solid fa-certificate text-xs"></i> Certifications
                    </a>
                  </div>
                {/if}
              </li>
            {/each}
          </ul>
        </div>
      {/if}
    </div>
  </nav>

  <ChatBot />

  <!-- Main content -->
  <main class="pt-16">
    <slot />
  </main>

  <!-- Scroll-to-top -->
  {#if showScrollTop}
    <button
      on:click={scrollToTop}
      class="fixed bottom-24 right-6 z-40
             bg-linear-to-r from-primary to-secondary gradient-btn-text
             p-3 rounded-full shadow-lg hover:shadow-xl
             transition-all duration-300 hover:scale-110"
      aria-label="Scroll to top"
    >
      <i class="fa-solid fa-arrow-up font-bold"></i>
    </button>
  {/if}
</div>

<style>
  @keyframes fadeIn {
    from {
      opacity: 0;
      transform: translateY(-6px);
    }
    to {
      opacity: 1;
      transform: translateY(0);
    }
  }
  .animate-fadeIn {
    animation: fadeIn 0.15s ease-out;
  }
</style>