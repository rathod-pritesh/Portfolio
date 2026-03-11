<script>
  import { goto } from "$app/navigation";
  import { page } from "$app/stores";
  import { onMount } from "svelte";
  import { browser } from "$app/environment";

  let isOpen = false;
  let adminEmail = '';
  let loading = true;

  const navItems = [
    { label: 'Dashboard', href: '/admin/dashboard', icon: 'fa-gauge-high' },
    { label: 'Home', href: '/admin/home', icon: 'fa-house' },
    { label: 'About', href: '/admin/about', icon: 'fa-user' },
    { label: 'Skills', href: '/admin/skills', icon: 'fa-code' },
    { label: 'Projects', href: '/admin/projects', icon: 'fa-folder-open' },
    { label: 'Contact', href: '/admin/contact', icon: 'fa-envelope' },
  ];

  onMount(async () => {
    try {
      const res = await fetch('http://127.0.0.1:8080/api/admin/verify', {
        credentials: 'include'
      });
      if (!res.ok) {
        goto('/admin/login');
        return;
      }
      const data = await res.json();
      adminEmail = data.email ?? '';
    } catch {
      goto('/admin/login')
    } finally {
      loading = false;
    }
  });

  async function logout() {
    await fetch('http://127.0.0.1:8080/api/logout', {
      method: 'POST',
      credentials: 'include'
    });
    goto('/admin/login');
  }

  $: if (browser) {
    document.body.style.overflow = isOpen ? 'hidden' : '';
  }

  $: currentPath = $page.url.pathname;
</script>

{#if loading}
  <div class="fixed inset-0 bg-darker flex items-center justify-center z-50">
    <div class="flex flex-col items-center gap-4">
      <div class="w-10 h-10 border-2 border-primary/30 border-t-primary rounded-full animate-spin"></div>
      <p class="text-secondary text-sm tracking-widest uppercase">Verifying session…</p>
    </div>
  </div>
{:else}
  <div class="min-h-screen bg-darker flex font-sans">

  <!-- ── Overlay (mobile)  -->
  {#if isOpen}
    <button
      type="button"
      class="fixed inset-0 bg-black/60 z-30 md:hidden"
      on:click={() => (isOpen = false)}
      aria-label="Close sidebar"
    ></button>
  {/if}

  <!-- ── Sidebar  -->
  <aside
    class="fixed top-0 left-0 h-full w-64 z-40 flex flex-col
           bg-dark border-r border-gray-800
           transition-transform duration-300 ease-in-out
           {isOpen ? 'translate-x-0' : '-translate-x-full'}
           md:translate-x-0 md:static md:flex"
  >
    <!-- Logo area -->
    <div class="flex items-center justify-between px-6 py-5 border-b border-gray-800">
      <a href="/admin/dashboard" class="flex items-center gap-2">
        <span class="text-xl font-bold gradient-text tracking-tight">Admin Panel</span>
      </a>
      <!-- Close button only on mobile -->
      <button
        class="md:hidden text-gray-400 hover:text-white text-xl"
        on:click={() => (isOpen = false)}
        aria-label="Close sidebar"
      >
        <i class="fa-solid fa-xmark"></i>
      </button>
    </div>

    <!-- Nav links -->
    <nav class="flex-1 overflow-y-auto py-4 px-3">
      <ul class="space-y-1">
        {#each navItems as item}
          <li>
            <a
              href={item.href}
              class="flex items-center gap-3 px-4 py-3 rounded-lg text-sm font-medium transition-all duration-200
                     {currentPath === item.href
                       ? 'bg-primary/10 text-primary border border-primary/20'
                       : 'text-gray-400 hover:text-white hover:bg-white/5'}"
              on:click={() => (isOpen = false)}
            >
              <i class="fa-solid {item.icon} w-4 text-center"></i>
              {item.label}
            </a>
          </li>
        {/each}
      </ul>
    </nav>

    <!-- logout -->
    <div class="border-t border-gray-800 p-4 space-y-3">
      <div class="flex items-center gap-3 px-2">
        <div class="w-8 h-8 rounded-full bg-primary/20 border border-primary/30 flex items-center justify-center">
          <i class="fa-solid fa-user text-primary text-xs"></i>
        </div>
        <span class="text-xs text-gray-400 truncate">{adminEmail}</span>
      </div>
      <button
        on:click={logout}
        class="w-full flex items-center gap-3 px-4 py-2.5 rounded-lg text-sm
               text-gray-400 hover:text-red-400 hover:bg-red-400/10
               border border-transparent hover:border-red-400/20
               transition-all duration-200"
      >
        <i class="fa-solid fa-right-from-bracket w-4 text-center"></i>
        Logout
      </button>
    </div>
  </aside>

  <!-- ── Main content  -->
  <div class="flex-1 flex flex-col min-w-0">

    <!-- Top bar -->
    <header class="sticky top-0 z-20 bg-dark/80 backdrop-blur-md border-b border-gray-800">
      <div class="flex items-center justify-between px-4 md:px-6 py-4">

        <!-- Hamburger -->
        <button
          class="md:hidden text-gray-400 hover:text-white text-xl transition-colors"
          on:click={() => (isOpen = !isOpen)}
          aria-label="Toggle sidebar"
        >
          <i class="fa-solid {isOpen ? 'fa-xmark' : 'fa-bars'}"></i>
        </button>

        <!-- Page title derived from path -->
        <h1 class="text-sm font-semibold text-secondary uppercase tracking-widest">
          {navItems.find(n => n.href === currentPath)?.label ?? 'Admin'}
        </h1>

        <!-- View portfolio link -->
        <a
          href="/"
          target="_blank"
          rel="noopener"
          class="flex items-center gap-2 text-xs text-gray-500 hover:text-primary transition-colors"
        >
          <i class="fa-solid fa-arrow-up-right-from-square"></i>
          <span class="hidden sm:inline">View Portfolio</span>
        </a>
      </div>
    </header>

    <!-- Slot -->
    <main class="flex-1 p-4 md:p-8 overflow-y-auto">
      <slot />
    </main>
  </div>
  </div>
{/if}

<style>
  @import url('https://fonts.googleapis.com/css2?family=Space+Mono:wght@400;700&display=swap');

  :global(body) {
    background-color: #020617;
    color: #e2e8f0;
  }

  .gradient-text {
    background: linear-gradient(to right, #e2e8f0, #94a3b8);
    -webkit-background-clip: text;
    -webkit-text-fill-color: transparent;
    background-clip: text;
  }
</style>