<script>
  import "../app.css";
  import { onMount } from "svelte";
  import { browser } from "$app/environment";
  import ChatBot from "$lib/components/ChatBot.svelte";

  let showScrollTop = false;
  let activeSection = "home";
  let isMenuOpen = false;
  let isDark = true;

  const navItems = [
    { label: "Home",     href: "#home",     id: "home"     },
    { label: "About",    href: "#about",    id: "about"    },
    { label: "Skills",   href: "#skills",   id: "skills"   },
    { label: "Projects", href: "#projects", id: "projects" },
    { label: "Contact",  href: "#contact",  id: "contact"  },
  ];

  function toggleTheme() {
    isDark = !isDark;
    document.body.classList.toggle('light', !isDark);
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
    window.addEventListener("scroll", handleScroll);
    return () => window.removeEventListener("scroll", handleScroll);
  });

  $: if (browser) {
    document.body.style.overflow = isMenuOpen ? "hidden" : "auto";
  }

  function scrollToTop() { window.scrollTo({ top: 0, behavior: "smooth" }); }
  function closeMenu()    { isMenuOpen = false; }
</script>

<div class="min-h-screen relative">

  <!-- Navbar -->
  <nav class="fixed top-0 left-0 right-0 z-50 bg-darker/80 backdrop-blur-md border-b border-gray-800">
    <div class="container-max">
      <div class="flex items-center justify-between py-4 px-6">

        <!-- Logo -->
        <a
          href="#home"
          class="text-2xl font-bold bg-gradient-to-r from-primary to-secondary gradient-text"
          aria-label="home"
        >
          Pritesh Rathod
        </a>

        <!-- Desktop links -->
        <ul class="hidden md:flex items-center gap-8">
          {#each navItems as item}
            <li>
              <a
                href={item.href}
                class="relative text-gray-400 hover:text-white transition-colors duration-300 group"
                class:text-primary={activeSection === item.id}
              >
                {item.label}
                <!-- Active underline -->
                <span
                  class="absolute left-0 -bottom-1 h-0.5 w-full bg-primary origin-left transform transition-transform duration-300
                         {activeSection === item.id ? 'scale-x-100' : 'scale-x-0 group-hover:scale-x-100'}"
                ></span>
              </a>
            </li>
          {/each}
        </ul>

        <button 
          on:click={toggleTheme}
          class="flex items-center justify-center w-9 h-9 rounded-full border border-gray-600 hover:border-gray-400 transition-colors"
          aria-label="Toggle theme"
        >

          {#if isDark}
            <i class="fa-solid fa-circle-half-stroke"></i>
          {:else}
            <i class="fa-solid fa-circle-half-stroke"></i>
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
        <div class="fixed top-16 left-0 right-0 z-50 md:hidden bg-darker border-t border-gray-800 px-6 py-4 animate-slideDown">
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
             bg-gradient-to-r from-primary to-secondary text-slate-900
             p-3 rounded-full shadow-lg hover:shadow-xl
             transition-all duration-300 hover:scale-110"
      aria-label="Scroll to top"
    >
      <i class="fa-solid fa-arrow-up font-bold"></i>
    </button>
  {/if}

</div>