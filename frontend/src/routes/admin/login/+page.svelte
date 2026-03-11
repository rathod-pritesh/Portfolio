<script>
  import { goto } from "$app/navigation";
  import { includes } from "lodash";
  import { onMount } from "svelte";

  let email = "";
  let password = "";
  let error = "";
  let loading = false;
  let showPassword = false;

  // If already logged in redirect to dashboard
  onMount(async () => {
    try {
      const res = await fetch("http://127.0.0.1:8080/api/admin/veridy", {
        credentials: "include",
      });
      if (res.ok) goto("/admin/dashboard");
    } catch {
      alert("An error occured!!");
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
      const res = await fetch("http://127.0.0.1:8080/api/admin/login", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ email, password }),
        credentials: "include",
      });

      const data = await res.json();

      if (!res.ok) {
        error = data.error ?? "Login failed";
      } else {
        goto("/admin/dashboard");
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

<div class="min-h-screen bg-darker flex items-center justify-center px-4">
  <!-- Background -->
  <div class="absolute inset-0 overflow-hidden pointer-events-none">
    <div
      class="absolute top-1/3 left-1/2 -translate-x-1/2 -translate-y-1/2
                w-96 h-96 bg-primary/5 rounded-full blur-3xl"
    ></div>
  </div>

  <div class="relative w-full max-w-md">
    <!-- Card -->
    <div class="bg-dark border border-gray-800 rounded-2xl p-8 shadow-2xl">
      <!-- Header -->
      <div class="text-center mb-8">
        <div
          class="inline-flex items-center justify-center w-14 h-14 rounded-2xl
                    bg-primary/10 border border-primary/20 mb-4"
        >
          <i class="fa-solid fa-shield-halved text-primary text-xl"></i>
        </div>
        <h1 class="text-2xl font-bold gradient-text">Admin Login</h1>
        <p class="text-gray-500 text-sm mt-1">Portfolio Management Panel</p>
      </div>

      <!-- Error banner -->
      {#if error}
        <div
          class="flex items-center gap-2 bg-red-500/10 border border-red-500/20 rounded-lg px-4 py-3 mb-5 text-sm text-red-400"
        >
          <i class="fa-solid fa-circle-exclamation"></i>
          {error}
        </div>
      {/if}

      <!-- Form -->
      <div class="space-y-5">
        <!-- Email -->
        <div class="space-y-1.5">
          <label
            class="text-xs font-medium text-gray-400 uppercase tracking-wider"
            for="email"
          >
            Email
          </label>
          <div class="relative">
            <span
              class="absolute left-3.5 top-1/2 -translate-y-1/2 text-gray-500 text-sm"
            >
              <i class="fa-solid fa-envelope"></i>
            </span>
            <input
              id="email"
              type="email"
              bind:value={email}
              on:keydown={handleKeydown}
              placeholder="admin@example.com"
              class="w-full bg-darker border border-gray-700 rounded-lg pl-10 pr-4 py-3
                     text-sm text-white placeholder-gray-600
                     focus:outline-none focus:border-primary/60 focus:ring-1 focus:ring-primary/30
                     transition-colors"
              autocomplete="email"
            />
          </div>
        </div>

        <!-- Password -->
        <div class="space-y-1.5">
          <label
            class="text-xs font-medium text-gray-400 uppercase tracking-wider"
            for="password"
          >
            Password
          </label>
          <div class="relative">
            <span
              class="absolute left-3.5 top-1/2 -translate-y-1/2 text-gray-500 text-sm"
            >
              <i class="fa-solid fa-lock"></i>
            </span>
            <input
              id="password"
              type={showPassword ? "text" : "password"}
              bind:value={password}
              on:keydown={handleKeydown}
              placeholder="••••••••"
              class="w-full bg-darker border border-gray-700 rounded-lg pl-10 pr-12 py-3
                     text-sm text-white placeholder-gray-600
                     focus:outline-none focus:border-primary/60 focus:ring-1 focus:ring-primary/30
                     transition-colors"
              autocomplete="current-password"
            />
            <button
              type="button"
              class="absolute right-3.5 top-1/2 -translate-y-1/2 text-gray-500 hover:text-gray-300 text-sm"
              on:click={() => (showPassword = !showPassword)}
              aria-label="Toggle password"
            >
              <i class="fa-solid {showPassword ? 'fa-eye-slash' : 'fa-eye'}"
              ></i>
            </button>
          </div>
        </div>

        <!-- Submit -->
        <button
          on:click={handleLogin}
          disabled={loading}
          class="w-full flex items-center justify-center gap-2
                 bg-primary/10 hover:bg-primary/20 text-primary
                 border border-primary/30 hover:border-primary/60
                 rounded-lg py-3 text-sm font-semibold
                 transition-all duration-200 disabled:opacity-50 disabled:cursor-not-allowed
                 hover:shadow-lg hover:shadow-primary/10 mt-2"
        >
          {#if loading}
            <div
              class="w-4 h-4 border-2 border-primary/30 border-t-primary rounded-full animate-spin"
            ></div>
            Signing in…
          {:else}
            <i class="fa-solid fa-arrow-right-to-bracket"></i>
            Sign In
          {/if}
        </button>
      </div>

      <!-- Footer note -->
      <p class="text-center text-xs text-gray-600 mt-6">
        Session expires after 24 hours
      </p>
    </div>

    <!-- Back link -->
    <div class="text-center mt-4">
      <a
        href="/"
        class="text-xs text-gray-600 hover:text-gray-400 transition-colors"
      >
        <i class="fa-solid fa-arrow-left mr-1"></i>Back to Portfolio
      </a>
    </div>
  </div>
</div>
