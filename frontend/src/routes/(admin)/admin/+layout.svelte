<script>
  import "$lib/app.css";
  import { goto } from "$app/navigation";
  import { page } from "$app/stores";
  import { onMount } from "svelte";

  let checking = true;

  onMount(async () => {
    if ($page.url.pathname === "/admin/login") {
      checking = false;
      return;
    }

    try {
      const res = await fetch("http://localhost:8080/api/admin/verify", {
        credentials: "include",
      });
      if (!res.ok) {
        goto("/admin/login");
        return;
      }
    } catch {
      goto("/admin/login");
      return;
    }
    checking = false;
  });

  async function handleLogout() {
    await fetch("http://localhost:8080/api/admin/logout", {
      method: "POST",
      credentials: "include",
    });
    goto("/admin/login");
  }

  const navLinks = [
    { href: "/admin/home", label: "Home", icon: "fa-house" },
    { href: "/admin/about", label: "About", icon: "fa-user" },
    { href: "/admin/skills", label: "Skills", icon: "fa-layer-group" },
    { href: "/admin/projects", label: "Projects", icon: "fa-diagram-project" },
  ];
</script>

{#if checking && $page.url.pathname !== "/admin/login"}
  <div class="min-h-screen bg-darker flex items-center justify-center">
    <span
      class="w-8 h-8 rounded-full border-2 border-white/8 border-t-primary animate-spin"
    ></span>
  </div>
{:else if $page.url.pathname === "/admin/login"}
  <slot />
{:else}
  <div class="min-h-screen bg-darker flex flex-col">
    <header
      class="sticky top-0 z-50 border-b border-white/6 bg-[rgba(10,10,15,0.85)] backdrop-blur-md"
    >
      <div
        class="max-w-[1200px] mx-auto px-4 h-16 flex items-center justify-between gap-4"
      >
        <a
          href="/admin/home"
          class="flex items-center gap-2 text-[15px] font-semibold text-primary no-underline shrink-0"
        >
          <i class="fa-solid fa-shield-halved"></i>
          <span>Admin</span>
        </a>

        <nav class="flex items-center gap-1 overflow-x-auto no-scrollbar py-2">
          {#each navLinks as link}
            <a
              href={link.href}
              class="flex items-center gap-1.5 px-3 py-1.5 rounded-lg text-[13px] font-medium no-underline whitespace-nowrap transition-colors duration-150
          {$page.url.pathname === link.href
                ? 'text-primary bg-primary/10'
                : 'text-white/45 hover:text-white/85 hover:bg-white/5'}"
            >
              <i class="fa-solid {link.icon}"></i>
              <span>{link.label}</span>
            </a>
          {/each}
        </nav>

        <div class="flex items-center gap-1.5 shrink-0">
          <a
            href="/"
            target="_blank"
            title="View portfolio"
            class="flex items-center justify-center gap-2 px-3 h-9 rounded-lg border border-white/8 text-white/40 text-[13px] hover:text-white/85 hover:border-white/15 hover:bg-white/5 no-underline transition-colors"
          >
            <i class="fa-solid fa-arrow-up-right-from-square"></i>
            <span>Portfolio</span>
          </a>

          <button
            on:click={handleLogout}
            title="Logout"
            class="flex items-center justify-center gap-2 px-3 h-9 rounded-lg border border-white/8 bg-transparent text-white/40 text-[13px] hover:text-red-400 hover:border-red-400/30 hover:bg-red-400/8 cursor-pointer transition-colors"
          >
            <i class="fa-solid fa-right-from-bracket"></i>
            <span>Logout</span>
          </button>
        </div>
      </div>
    </header>

    <main class="flex-1 p-4 md:p-8 w-full max-w-[1200px] mx-auto">
      <slot />
    </main>
  </div>
{/if}

<style>
  /* Hides scrollbar for the nav on mobile but keeps it scrollable */
  .no-scrollbar::-webkit-scrollbar {
    display: none;
  }
  .no-scrollbar {
    -ms-overflow-style: none;
    scrollbar-width: none;
  }
</style>
