<script>
  import { goto } from "$app/navigation";
  import { onMount } from "svelte";

  let email = "";
  let password = "";
  let error = "";
  let loading = false;
  let showPassword = false;

  // If already logged in redirect to dashboard
  onMount(async () => {
    try {
      const res = await fetch("http://localhost:8080/api/admin/verify", {
        credentials: "include", // used localhost bcz in verification cookie localhost 127.0.0.1 different not works browser treat both as different
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
      const res = await fetch("http://localhost:8080/api/admin/login", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ email, password }),
        credentials: "include" //cookie store in browser
      });

      const data = await res.json();

      if (!res.ok) {
        error = data.error ?? "Login failed";
      } else {
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

<div class="relative min-h-screen bg-darker flex items-center justify-center p-4 overflow-hidden">
 
  <div class="absolute w-96 h-96 rounded-full bg-primary/5 blur-3xl
              top-1/3 left-1/2 -translate-x-1/2 -translate-y-1/2 pointer-events-none">
  </div>
 
  <!-- Card wrapper -->
  <div class="relative z-10 w-full max-w-105">
 
    <!-- Card -->
    <div class="bg-dark border border-white/[0.07] rounded-2xl p-6 sm:p-9
                shadow-[0_24px_64px_rgba(0,0,0,0.55)]">
 
      <!-- Header -->
      <div class="text-center mb-7">
        <div class="inline-flex items-center justify-center w-14 h-14
                    rounded-xl bg-primary/10 border border-primary/25
                    text-primary text-xl mb-3.5">
          <i class="fa-solid fa-shield-halved"></i>
        </div>
        <h1 class="text-2xl font-bold gradient-text">Admin Login</h1>
        <p class="text-sm text-white/35 mt-1">Portfolio Management Panel</p>
      </div>
 
      <!-- Error banner -->
      {#if error}
        <div class="flex items-start gap-2 bg-red-500/8 border border-red-500/20
                    rounded-lg px-4 py-3 mb-5 text-sm text-red-400 leading-relaxed"
             role="alert">
          <i class="fa-solid fa-circle-exclamation shrink-0 mt-0.5"></i>
          <span>{error}</span>
        </div>
      {/if}
 
      <!-- Fields -->
      <div class="flex flex-col gap-4">
 
        <!-- Email -->
        <div class="flex flex-col gap-1.5">
          <label for="email"
                 class="text-[11px] font-medium text-white/40 uppercase tracking-widest">
            Email
          </label>
          <div class="relative">
            <i class="fa-solid fa-envelope absolute left-3.5 top-1/2 -translate-y-1/2
                      text-sm text-white/25 pointer-events-none"></i>
            <input
              id="email"
              type="email"
              bind:value={email}
              on:keydown={handleKeydown}
              placeholder="admin@example.com"
              autocomplete="email"
              class="w-full bg-darker border border-white/10 rounded-lg
                     pl-10 pr-4 py-3 text-base text-white placeholder:text-white/20
                     outline-none focus:border-primary/60 focus:ring-2 focus:ring-primary/15
                     transition-colors appearance-none"
            />
          </div>
        </div>
 
        <!-- Password -->
        <div class="flex flex-col gap-1.5">
          <label for="password"
                 class="text-[11px] font-medium text-white/40 uppercase tracking-widest">
            Password
          </label>
          <div class="relative">
            <i class="fa-solid fa-lock absolute left-3.5 top-1/2 -translate-y-1/2
                      text-sm text-white/25 pointer-events-none"></i>
            <input
              id="password"
              type={showPassword ? "text" : "password"}
              bind:value={password}
              on:keydown={handleKeydown}
              placeholder="••••••••"
              autocomplete="current-password"
              class="w-full bg-darker border border-white/10 rounded-lg
                     pl-10 pr-12 py-3 text-base text-white placeholder:text-white/20
                     outline-none focus:border-primary/60 focus:ring-2 focus:ring-primary/15
                     transition-colors appearance-none"
            />
            <button
              type="button"
              on:click={() => (showPassword = !showPassword)}
              aria-label="Toggle password visibility"
              class="absolute right-3 top-1/2 -translate-y-1/2
                     w-8 h-8 flex items-center justify-center
                     text-sm text-white/30 hover:text-white/70
                     transition-colors cursor-pointer bg-transparent border-none"
            >
              <i class="fa-solid {showPassword ? 'fa-eye-slash' : 'fa-eye'}"></i>
            </button>
          </div>
        </div>
 
        <!-- Submit -->
        <button
          on:click={handleLogin}
          disabled={loading}
          class="w-full flex items-center justify-center gap-2 mt-1 py-3 px-4
                 bg-primary/10 hover:bg-primary/20
                 border border-primary/30 hover:border-primary/55
                 rounded-lg text-primary text-[15px] font-semibold cursor-pointer
                 hover:shadow-lg hover:shadow-primary/10
                 disabled:opacity-50 disabled:cursor-not-allowed
                 transition-all duration-200"
        >
          {#if loading}
            <span class="w-4 h-4 rounded-full border-2 border-primary/30
                         border-t-primary animate-spin"></span>
            Signing in…
          {:else}
            <i class="fa-solid fa-arrow-right-to-bracket"></i>
            Sign In
          {/if}
        </button>
 
      </div>
 
      <!-- Session note -->
      <p class="text-center text-xs text-white/20 mt-5">
        Session expires after 24 hours
      </p>
 
    </div>
 
    <!-- Back link -->
    <div class="text-center mt-4">
      <a href="/"
         class="inline-flex items-center gap-1.5 text-xs text-white/25
                hover:text-white/55 transition-colors no-underline">
        <i class="fa-solid fa-arrow-left"></i>
        Back to Portfolio
      </a>
    </div>
 
  </div>
</div>