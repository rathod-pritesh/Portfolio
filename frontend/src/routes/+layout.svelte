<script>
  import "../app.css";
  import { onMount } from "svelte";
  import { browser } from "$app/environment";
  import ChatBot from "$lib/components/ChatBot.svelte";

  let showScrollTop = false;
  let activeSection = "home";
  let isMenuOpen = false;

  const navItems = [
    { label: "Home", href: "#home", id: "home" },
    { label: "About", href: "#about", id: "about" },
    { label: "Skills", href: "#skills", id: "skills" },
    { label: "Projects", href: "#projects", id: "projects" },
    { label: "Contact", href: "#contact", id: "contact" },
  ];

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
    class="fixed top-0 left-0 right-0 z-50 bg-darker/80 backdrop-blur-md border-b border-gray-800"
  >
    <div class="container-max">
      <div class="flex items-center justify-between py-4 px-6">
        <a
          href="#home"
          class="text-2xl font-bold bg-gradient-to-r from-primary to-secondary gradient-text text-transparent"
          aria-label="home"
        >
        </a>
        <ul class="hidden md:flex items-center gap-8">
          {#each navItems as item}
            <li>
              <a
                href={item.href}
                class="relative text-gray-300 hover:text-white transition-colors duration-300 group"
                class:text-primary={activeSection === item.id}
              >
                {item.label}
                {#if activeSection === item.id}
                  <span
                    class="absolute left-0 -bottom-1 h-0.5 w-full bg-primary origin-left transform transition-transform duration-300 scale-x-0 group-hover:scale-x-100"
                    class:scale-x-100={activeSection === item.id}
                  ></span>
                {/if}
              </a>
            </li>
          {/each}
        </ul>

        <button
          class="md:hidden text-white text-2xl transition-transform duration-200"
          on:click={() => (isMenuOpen = !isMenuOpen)}
          aria-label={isMenuOpen ? "Close Menu" : "Open Menu"}
        >
          {#if isMenuOpen}
            <i class="fa-solid fa-xmark"></i>
          {:else}
            <i class="fa-solid fa-bars"></i>
          {/if}
        </button>
      </div>

      {#if isMenuOpen}
        <!-- Overlay -->
        <button
          type="button"
          class="fixed inset-0 bg-black/40 z-40 md:hidden cursor-default"
          on:click={closeMenu}
          aria-label="Close menu"
        ></button>

        <!-- Mobile Menu -->
        <div
          class="fixed top-16 left-0 right-0 z-50 md:hidden
			   bg-darker border-t border-gray-800
			   px-6 py-4 animate-slideDown"
        >
          <ul class="space-y-4">
            {#each navItems as item}
              <li>
                <a
                  href={item.href}
                  class="block text-gray-300 hover:text-white"
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
  <!-- Main Content -->
  <main class="pt-16">
    <slot />
  </main>

  <!-- Scroll to Top Button -->
  {#if showScrollTop}
    <button
      on:click={scrollToTop}
      class="fixed bottom-24 right-6 z-40
       bg-gradient-to-r from-primary to-secondary
       p-3 rounded-full shadow-lg
       hover:shadow-xl transition-all duration-300
       hover:scale-110"
      aria-label="Scroll to top"
    >
      <i class="fa-solid fa-arrow-up"></i>
    </button>
  {/if}
</div>
