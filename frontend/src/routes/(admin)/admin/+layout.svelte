<script>
  import "$lib/app.css";
  import { goto } from "$app/navigation";
  import { page } from "$app/stores";
  import { onMount } from "svelte";
  import ParticlesBackground from "$lib/components/ParticlesBackground.svelte";
  import ToastMessage from "$lib/components/admin/ToastMessage.svelte";
  import { showToast } from "$lib/stores/toastStore";

  let checking    = true;
  let loggingOut  = false;
  let showConfirm = false;
  let fadeOut     = false;

  onMount(async () => {
    if ($page.url.pathname === "/admin/login") {
      checking = false;
      return;
    }
    try {
      const res = await fetch("http://localhost:8080/api/admin/verify", {
        credentials: "include",
      });
      if (!res.ok) { checking = false; goto("/admin/login"); return; }
    } catch {
      checking = false; goto("/admin/login"); return;
    }
    checking = false;
  });

  $: if ($page.url.pathname === "/admin/login") {
    loggingOut  = false;
    fadeOut     = false;
    showConfirm = false;
  }

  function requestLogout() { showConfirm = true; }
  function cancelLogout()  { showConfirm = false; }

  async function confirmLogout() {
    showConfirm = false;
    loggingOut  = true;
    await fetch("http://localhost:8080/api/admin/logout", { method: "POST", credentials: "include" });
    showToast("Logged out successfully");
    await new Promise((r) => setTimeout(r, 400));
    fadeOut = true;
    await new Promise((r) => setTimeout(r, 700));
    goto("/admin/login");
  }

  const navLinks = [
    { href: "/admin/home",     label: "Home",     icon: "fa-house" },
    { href: "/admin/about",    label: "About",    icon: "fa-user" },
    { href: "/admin/skills",   label: "Skills",   icon: "fa-layer-group" },
    { href: "/admin/projects", label: "Projects", icon: "fa-diagram-project" },
  ];
</script>

{#if checking && $page.url.pathname !== "/admin/login"}
  <div class="min-h-screen bg-darker flex items-center justify-center">
    <span class="w-8 h-8 rounded-full border-2 border-white/8 border-t-primary animate-spin"></span>
  </div>

{:else if $page.url.pathname === "/admin/login"}
  <slot />

{:else}
  <ParticlesBackground />

  {#if fadeOut}
    <div class="logout-overlay">
      <div class="flex flex-col items-center gap-4">
        <span class="w-10 h-10 rounded-full border-2 border-white/20 border-t-white animate-spin"></span>
        <p class="text-white/70 text-sm font-medium tracking-wide">Signing out...</p>
      </div>
    </div>
  {/if}

  <div class="min-h-screen bg-darker flex flex-col">
    <header class="sticky top-0 z-50 border-b border-white/6 bg-[rgba(2,6, 23, 0.92)] backdrop-blur-md">

      <div class="max-w-[1200px] mx-auto px-4 h-14 flex items-center justify-between gap-3">

        <a href="/admin/home"
          class="flex items-center gap-2 text-[15px] font-semibold text-primary no-underline shrink-0">
          <i class="fa-solid fa-shield-halved"></i>
          <span>Admin</span>
        </a>

        <!-- sm screen -->
        <div class="flex items-center gap-1.5 shrink-0">
          <a href="/" target="_blank" title="View portfolio"
            class="flex items-center justify-center gap-1.5 px-2.5 h-8 rounded-lg border border-white/8
                   text-white/40 text-[13px] hover:text-white/85 hover:border-white/15 hover:bg-white/5
                   no-underline transition-colors">
            <i class="fa-solid fa-arrow-up-right-from-square"></i>
            <span class="hidden sm:inline">Portfolio</span>
          </a>

          <button
            on:click={requestLogout}
            disabled={loggingOut}
            title="Logout"
            class="flex items-center justify-center gap-1.5 px-2.5 h-8 rounded-lg border border-white/8
                   bg-transparent text-white/40 text-[13px] cursor-pointer transition-all duration-200
                   hover:text-red-400 hover:border-red-400/30 hover:bg-red-400/8
                   disabled:opacity-60 disabled:cursor-not-allowed"
          >
            {#if loggingOut}
              <span class="w-3 h-3 rounded-full border-2 border-white/20 border-t-red-400 animate-spin"></span>
              <span class="hidden sm:inline">Signing out...</span>
            {:else}
              <i class="fa-solid fa-right-from-bracket"></i>
              <span class="hidden sm:inline">Logout</span>
            {/if}
          </button>
        </div>
      </div>

      <div class="border-t border-white/5">
        <nav class="max-w-[1200px] mx-auto px-3 flex items-center gap-0.5 overflow-x-auto no-scrollbar py-1.5 justify-center">
          {#each navLinks as link}
            <a href={link.href}
              class="flex items-center gap-1.5 px-3 py-1.5 rounded-lg text-[13px] font-medium
                     no-underline whitespace-nowrap transition-colors duration-150 shrink-0
                     {$page.url.pathname === link.href
                       ? 'text-primary bg-primary/10'
                       : 'text-white/45 hover:text-white/85 hover:bg-white/5'}">
              <i class="fa-solid {link.icon}"></i>
              <span>{link.label}</span>
            </a>
          {/each}
        </nav>
      </div>

    </header>

    <main class="flex-1 p-4 md:p-8 w-full max-w-[1200px] mx-auto">
      <slot />
    </main>
  </div>

  <ToastMessage />

  {#if showConfirm}
    <div class="fixed inset-0 z-[9998] flex items-center justify-center p-4">
      <button type="button" on:click={cancelLogout} tabindex="-1"
        class="absolute inset-0 w-full h-full bg-black/60 backdrop-blur-sm cursor-default"
        aria-label="Cancel logout">
      </button>
      <div class="confirm-dialog relative bg-[#0f1117] border border-white/10 rounded-2xl
                  w-full max-w-sm p-6 shadow-2xl flex flex-col gap-5">
        <div class="flex items-start gap-4">
          <div class="w-10 h-10 rounded-xl bg-red-500/10 border border-red-500/20
                      flex items-center justify-center shrink-0 text-red-400">
            <i class="fa-solid fa-right-from-bracket"></i>
          </div>
          <div>
            <h3 class="text-white font-bold text-base mb-1">Sign out?</h3>
            <p class="text-white/40 text-sm">You'll need to log in again to access the admin panel.</p>
          </div>
        </div>
        <div class="flex gap-3">
          <button type="button" on:click={confirmLogout}
            class="flex-1 py-2.5 rounded-xl bg-red-500 hover:bg-red-600
                   text-white text-sm font-bold transition-colors duration-150">
            <i class="fa-solid fa-right-from-bracket mr-2"></i>Sign out
          </button>
          <button type="button" on:click={cancelLogout}
            class="px-5 py-2.5 rounded-xl border border-white/10 text-white/40
                   text-sm hover:border-white/20 hover:text-white/70 transition-colors duration-150">
            Cancel
          </button>
        </div>
      </div>
    </div>
  {/if}
{/if}

<style>
  .no-scrollbar::-webkit-scrollbar { display: none; }
  .no-scrollbar { -ms-overflow-style: none; scrollbar-width: none; }

  .logout-overlay {
    position: fixed; inset: 0; z-index: 9999;
    display: flex; align-items: center; justify-content: center;
    background: rgba(10, 10, 15, 0);
    animation: fadeToBlack 0.7s ease-out forwards;
  }
  @keyframes fadeToBlack {
    from { background: rgba(10, 10, 15, 0); }
    to   { background: rgba(10, 10, 15, 1); }
  }

  .confirm-dialog { animation: popIn 0.2s cubic-bezier(0.34, 1.56, 0.64, 1); }
  @keyframes popIn {
    from { opacity: 0; transform: scale(0.92); }
    to   { opacity: 1; transform: scale(1); }
  }
</style>