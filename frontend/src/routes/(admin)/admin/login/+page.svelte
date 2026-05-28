<script>
  import { goto } from "$app/navigation";
  import { onMount } from "svelte";
  import { API_BASE } from "$lib/config/api";

  let email = "";
  let password = "";
  let error = "";
  let loading = false;
  let showPassword = false;
  let fadeIn = false;

  onMount(async () => {
    try {
      const res = await fetch(`${API_BASE}/api/admin/verify`, {
        credentials: "include",
      });
      if (res.ok) goto("/admin/home");
    } catch {
      console.warn("Backend not reachable.");
    }
  });

  async function handleLogin() {
    error = "";

    if (!email || !password) {
      error = "Email and password are required.";
      return;
    }

    loading = true;
    try {
      const res = await fetch(`${API_BASE}/api/admin/login`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ email, password }),
        credentials: "include", //cookie store in browser
      });

      const data = await res.json();

      if (!res.ok) {
        error = data.error ?? "Login failed";
      } else {
        loading = false;
        fadeIn = true;
        await new Promise((r) => setTimeout(r, 900));
        goto("/admin/home");
      }
    } catch {
      error = "Could not reach server. Is backend running?";
    } finally {
      loading = false;
    }
  }

  function handleKeydown(e) {
    if (e.key === "Enter") handleLogin();
  }
</script>

<svelte:head>
  <title>Admin Login | Portfolio</title>
</svelte:head>

<div
  class="relative min-h-screen bg-[#0a0a0f] flex items-center justify-center p-4 sm:p-6 overflow-hidden select-none"
>
  <div
    class="absolute w-[300px] sm:w-[500px] h-[300px] sm:h-[500px] rounded-full bg-primary/10 blur-[80px] sm:blur-[120px] top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 pointer-events-none z-0"
  ></div>

  <div class="relative z-10 w-full max-w-[420px] transition-all duration-300">
    <div
      class="bg-dark/60 backdrop-blur-xl border border-white/[0.06] rounded-2xl p-6 sm:p-8 shadow-[0_32px_64px_-16px_rgba(0,0,0,0.7)] hover:border-white/[0.09] transition-all duration-300"
    >
      <div class="text-center mb-8">
        <div
          class="inline-flex items-center justify-center w-12 h-12 sm:w-14 sm:h-14 rounded-xl bg-primary/10 border border-primary/20 text-primary text-lg sm:text-xl mb-4 shadow-sm shadow-primary/5"
        >
          <i class="fa-solid fa-shield-halved"></i>
        </div>
        <h1 class="text-2xl sm:text-3xl font-bold tracking-tight text-white">
          Admin Login
        </h1>
        <p class="text-xs sm:text-sm text-white/40 mt-1.5 font-medium">
          Portfolio Management Panel
        </p>
      </div>

      {#if error}
        <div
          class="flex items-start gap-3 bg-red-500/10 border border-red-500/20 rounded-xl px-4 py-3.5 mb-6 text-sm text-red-400 animate-in fade-in slide-in-from-top-2 duration-200"
          role="alert"
        >
          <i
            class="fa-solid fa-circle-exclamation shrink-0 mt-0.5 text-red-400/80"
          ></i>
          <span class="leading-snug">{error}</span>
        </div>
      {/if}

      <div class="flex flex-col gap-5">
        <div class="flex flex-col gap-2">
          <label
            for="email"
            class="text-[11px] font-semibold text-white/40 uppercase tracking-wider"
          >
            Email Address
          </label>
          <div class="relative group">
            <i
              class="fa-solid fa-envelope absolute left-4 top-1/2 -translate-y-1/2 text-sm text-white/30 group-focus-within:text-primary transition-colors pointer-events-none"
            ></i>
            <input
              id="email"
              type="email"
              bind:value={email}
              on:keydown={handleKeydown}
              placeholder="admin@example.com"
              autocomplete="email"
              class="w-full bg-[#0d0d14]/80 border border-white/10 rounded-xl pl-11 pr-4 py-3 text-sm sm:text-base text-white placeholder:text-white/20 outline-none focus:border-primary/60 focus:ring-4 focus:ring-primary/10 transition-all appearance-none"
            />
          </div>
        </div>

        <div class="flex flex-col gap-2">
          <label
            for="password"
            class="text-[11px] font-semibold text-white/40 uppercase tracking-wider"
          >
            Password
          </label>
          <div class="relative group">
            <i
              class="fa-solid fa-lock absolute left-4 top-1/2 -translate-y-1/2 text-sm text-white/30 group-focus-within:text-primary transition-colors pointer-events-none"
            ></i>
            <input
              id="password"
              type={showPassword ? "text" : "password"}
              bind:value={password}
              on:keydown={handleKeydown}
              placeholder="••••••••"
              autocomplete="current-password"
              class="w-full bg-[#0d0d14]/80 border border-white/10 rounded-xl pl-11 pr-12 py-3 text-sm sm:text-base text-white placeholder:text-white/20 outline-none focus:border-primary/60 focus:ring-4 focus:ring-primary/10 transition-all appearance-none"
            />
            <button
              type="button"
              on:click={() => (showPassword = !showPassword)}
              aria-label="Toggle password visibility"
              class="absolute right-3 top-1/2 -translate-y-1/2 w-9 h-9 flex items-center justify-center rounded-lg text-sm text-white/30 hover:text-white/60 hover:bg-white/5 active:scale-95 transition-all cursor-pointer border-none"
            >
              <i class="fa-solid {showPassword ? 'fa-eye-slash' : 'fa-eye'}"
              ></i>
            </button>
          </div>
        </div>

        <button
          on:click={handleLogin}
          disabled={loading}
          class="w-full flex items-center justify-center gap-2.5 mt-2 py-3 px-4 bg-primary/10 hover:bg-primary/15 active:scale-[0.98] border border-primary/25 hover:border-primary/45 rounded-xl text-primary text-[15px] font-semibold cursor-pointer shadow-sm hover:shadow-md hover:shadow-primary/5 disabled:opacity-40 disabled:pointer-events-none transition-all duration-200"
        >
          {#if loading}
            <span
              class="w-4 h-4 rounded-full border-2 border-primary/30 border-t-primary animate-spin"
            ></span>
            <span>Signing in…</span>
          {:else}
            <i class="fa-solid fa-arrow-right-to-bracket text-sm"></i>
            <span>Sign In</span>
          {/if}
        </button>
      </div>

      <p
        class="text-center text-xs text-white/20 mt-6 font-medium tracking-wide"
      >
        Session expires after 24 hours
      </p>
    </div>

    <div class="text-center mt-5">
      <a
        href="/"
        class="inline-flex items-center gap-2 text-xs text-white/30 hover:text-white/60 font-medium transition-colors no-underline group"
      >
        <i
          class="fa-solid fa-arrow-left group-hover:-translate-x-0.5 transition-transform"
        ></i>
        Back to Portfolio
      </a>
    </div>
  </div>
</div>

{#if fadeIn}
  <div class="login-overlay">
    <div class="flex flex-col items-center gap-4">
      <span
        class="w-9 h-9 rounded-full border-2 border-white/10 border-t-white animate-spin"
      ></span>
      <p class="text-white/60 text-sm font-medium tracking-wider">
        Loading dashboard...
      </p>
    </div>
  </div>
{/if}

<style>
  .login-overlay {
    position: fixed;
    inset: 0;
    z-index: 9999;
    display: flex;
    align-items: center;
    justify-content: center;
    background: rgba(10, 10, 15, 0);
    animation: fadeToBlack 0.9s cubic-bezier(0.16, 1, 0.3, 1) forwards;
  }
  @keyframes fadeToBlack {
    from {
      background: rgba(10, 10, 15, 0);
      backdrop-blur: 0px;
    }
    to {
      background: rgba(10, 10, 15, 1);
      backdrop-blur: 12px;
    }
  }
</style>
