<script>
  import { onMount } from "svelte";

  let form = {
    name: "",
    role: "",
    tagline: "",
    resumeUrl: "",
    profileImage: "",
  };

  let loading = true;
  let saving = false;
  let success = false;
  let error = "";

  // Load current home data
  onMount(async () => {
    try{
      const res = await fetch("http://localhost:8080/api/admin/home", {
        credentials: "include",
      });
      if (!res.ok) {
        error = "Failed to laod home data.";
        return;
      }

      const data = await res.json();
      form = {
        name: data.name ?? "",
        role: data.role ?? "",
        tagline: data.tagline ?? "",
        resumeUrl: data.resumeUrl ?? "",
        profileImage: data.profileImage ?? "",
      };
    } catch {
      error = "Could not reach server.";
    } finally {
      loading = false;
    }
  });

  async function handleSave() {
    error = "";
    success = false;

    if (!form.name || !form.role || !form.tagline) {
      error = "Name, Role and Tagline are required.";
      return;
    }

    saving = true;

    try {
      const res = await fetch("http://localhost:8080/api/admin/home", {
        method: "PUT",
        headers: { "Content-Type": "application/json" },
        credentials: "include",
        body: JSON.stringify(form),
      });

      const data = await res.json();

      if (!res.ok) {
        error = data.error ?? "Failed to save.";
      } else {
        success = true;
        setTimeout(() => {
          success = false;
        }, 3000);
      }
    } catch {
      error = "Could not reach server.";
    } finally {
      saving = false;
    }
  }

  function handleReset() {
    error = "";
    success = false;
  }
</script>

<svelte:head>
  <title>Home - Admin Panel</title>
</svelte:head>

<!-- Page header -->
<div class="mb-8">
  <div class="flex items-center gap-2 text-white/30 text-sm mb-1">
    <i class="fa-solid fa-shield-halved text-primary"></i>
    <span>Admin</span>
    <i class="fa-solid fa-chevron-right text-xs"></i>
    <span class="text-white/60">Home</span>
  </div>
  <h1 class="text-2xl font-bold text-white">Home Section</h1>
  <p class="text-sm text-white/40 mt-1">
    Edit the content shown in your portfolio's hero / home section.
  </p>
</div>

{#if loading}
  <div class="flex items-center justify-center py-24">
    <span class="w-8 h-8 rounded-full border-2 border-white/10 border-t-primary animate-spin"></span>
  </div>

{:else}
  <!-- success banner -->
  {#if success}
    <div class="flex items-center gap-2 bg-green-500/[0.08] border border-green-500/20
                rounded-lg px-4 py-3 mb-6 text-sm text-green-400"
         role="status">
      <i class="fa-solid fa-circle-check shrink-0"></i>
      <span>Home section updated successfully.</span>
    </div>
  {/if}

  <!-- error banner -->
  {#if error}
    <div class="flex items-center gap-2 bg-red-500/[0.08] border border-red-500/20
                rounded-lg px-4 py-3 mb-6 text-sm text-red-400"
         role="alert">
      <i class="fa-solid fa-circle-exclamation shrink-0"></i>
      <span>{error}</span>
    </div>
  {/if}

  <!-- form -->
  <div class="bg-dark border border-white/[0.07] rounded-2xl p-6 sm:p-8
              shadow-[0_8px_32px_rgba(0,0,0,0.35)]">
    <div class="flex flex-col gap-6">

      <!-- name + role -->
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">

        <!-- name -->
        <div class="flex flex-col gap-1.5">
          <label for="name" class="text-[11px] font-medium text-white/40 uppercase tracking-widest">Name <span class="text-red-400">*</span></label>
          <input 
            id="name" 
            type="text" 
            bind:value={form.name} 
            placeholder="e.g Pritesh Rathod"
            class="w-full bg-darker border border-white/10 rounded-lg px-4 py-3 text-base text-white placeholder:text-white/20 outline-none focus:border-primary/60 focus:ring-2 focus:ring-primary/15 transition-colors appearance-none"
          />
        </div>

        <!-- role -->
        <div class="flex flex-col gap-1.5">
          <label for="role" class="text-[11px] font-medium text-white/40 uppercase tracking-widest">Role <span class="text-red-400">*</span></label>
          <input 
            id="role" 
            type="text" 
            bind:value={form.role} 
            placeholder="e.g Full Stack Developer"
            class="w-full bg-darker border border-white/10 rounded-lg px-4 py-3 text-base text-white placeholder:text-white/20 outline-none focus:border-primary/60 focus:ring-2 focus:ring-primary/15 transition-colors appearance-none"
          />
        </div>

      </div>

      <!-- tagline -->
      <div class="flex flex-col gap-1.5">
        <label for="tagline"
               class="text-[11px] font-medium text-white/40 uppercase tracking-widest">
          Tagline <span class="text-red-400">*</span>
        </label>
        <textarea
          id="tagline"
          bind:value={form.tagline}
          rows="3"
          placeholder="e.g. I build things for the web..."
          class="w-full bg-darker border border-white/10 rounded-lg px-4 py-3
                 text-base text-white placeholder:text-white/20 outline-none
                 focus:border-primary/60 focus:ring-2 focus:ring-primary/15
                 transition-colors appearance-none resize-none"
        ></textarea>
      </div>

      <!-- Resume URL -->
      <div class="flex flex-col gap-1.5">
        <label for="resumeUrl"
               class="text-[11px] font-medium text-white/40 uppercase tracking-widest">
          Resume URL
        </label>
        <div class="relative">
          <i class="fa-solid fa-link absolute left-3.5 top-1/2 -translate-y-1/2
                    text-sm text-white/25 pointer-events-none"></i>
          <input type="url" id="resumeUrl" bind:value={form.resumeUrl} placeholder="https://drive.google.com/..."
          class="w-full bg-darker border border-white/10 rounded-lg pl-10 pr-4 py-3
                   text-base text-white placeholder:text-white/20 outline-none
                   focus:border-primary/60 focus:ring-2 focus:ring-primary/15
                   transition-colors appearance-none"
          />
        </div>
        <!-- preview resume -->
        {#if form.resumeUrl}
          <a href={form.resumeUrl} target="_blank" rel="noopener noreferrer" class="inline-flex items-center gap-1.5 text-xs text-primary/70
                    hover:text-primary transition-colors mt-0.5 w-fit">
                    <i class="fa-solid fa-arrow-up-right-from-square text-[10px]"></i>
            Open link</a>
        {/if}
      </div>

      <!-- profileImage URL -->
      <div class="flex flex-col gap-1.5">
        <label for="profileImage" class="text-[11px] font-medium text-white/40 uppercase tracking-widest">
        Profile Image URL</label>
        <div class="flex items-start gap-4">

          <!-- input -->
          <div class="relative flex-1">
            <i class="fa-solid fa-image absolute left-3.5 top-1/2 -translate-y-1/2
                      text-sm text-white/25 pointer-events-none"></i>
            <input 
              id="profileImage"
              type="url"
              bind:value={form.profileImage}
              placeholder="https://..."
              class="w-full bg-darker border border-white/10 rounded-lg pl-10 pr-4 py-3
                     text-base text-white placeholder:text-white/20 outline-none
                     focus:border-primary/60 focus:ring-2 focus:ring-primary/15
                     transition-colors appearance-none"
            />
          </div>

          <!-- live img preview -->
          {#if form.profileImage}
            <div class="shrink-0 w-14 h-14 rounded-xl border border-white/10 overflow-hidden bg-darker">
              <img 
                src={form.profileImage}
                alt="Profile preview"
                class="w-full h-full object-cover"
                on:error={(e) => e.currentTarget.style.display = 'none'}
              />
            </div>
            {:else}
              <div class="shrink-0 w-14 h-14 rounded-xl border border-white/10
                        bg-darker flex items-center justify-center text-white/15">
              <i class="fa-solid fa-user text-xl"></i>
            </div>
          {/if}
        </div>
      </div>
    </div>

    <!-- Divider -->
    <div class="border-t border-white/[0.06] mt-8 pt-6 flex items-center justify-between gap-4 flex-wrap">
      <p class="text-xs text-white/25">
        <span class="text-red-400">*</span> Required fields
      </p>

      <div class="flex items-center gap-3">
        <button
          on:click={handleReset}
          disabled={saving}
          class="px-5 py-2.5 rounded-lg text-sm font-medium text-white/40
                 border border-white/[0.08] hover:text-white/70 hover:border-white/15
                 hover:bg-white/[0.04] disabled:opacity-40 disabled:cursor-not-allowed
                 transition-all duration-150 cursor-pointer bg-transparent"
        >
          Reset
        </button>

        <button
          on:click={handleSave}
          disabled={saving}
          class="flex items-center gap-2 px-6 py-2.5 rounded-lg text-sm font-semibold
                 text-primary bg-primary/10 border border-primary/30
                 hover:bg-primary/20 hover:border-primary/55 hover:shadow-lg hover:shadow-primary/10
                 disabled:opacity-50 disabled:cursor-not-allowed
                 transition-all duration-150 cursor-pointer"
        >
          {#if saving}
            <span class="w-3.5 h-3.5 rounded-full border-2 border-primary/30 border-t-primary animate-spin"></span>Saving...
          {:else}
            <i class="fa-solid fa-floppy-disk"></i>
            Save Changes
          {/if}

        </button>
      </div>  

    </div>
  </div>
{/if}