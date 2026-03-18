<script>
  import { onMount } from "svelte";

  let form = {
    title: "",
    description: "",
    highlights: [],
  };

  let loading = true;
  let saving = false;
  let success = false;
  let error = "";

  onMount(async () => {
    try {
      const res = await fetch("http://localhost:8080/api/admin/about", {
        credentials: "include",
      });
      if (!res.ok) {
        error = "Failed to load about data.";
        return;
      }

      const data = await res.json();
      form = {
        title: data.title ?? "",
        description: data.description ?? "",
        highlights: data.highlights ?? [],
      };
    } catch {
      error = "Could not reach server.";
    } finally {
      loading = false;
    }
  });

  function addHighlight() {
    form.highlights = [...form.highlights, { label: "", description: "" }];
  }

  function removeHighlight(index) {
    form.highlights = form.highlights.filter((_, i) => i !== index);
  }

  function updateHighlight(index, field, value) {
    form.highlights[index][field] = value;
    form.highlights = [...form.highlights];
  }

  function moveUp(index) {
    if (index === 0) return;
    const arr = [...form.highlights];
    [arr[index - 1], arr[index]] = [arr[index], arr[index - 1]];
    form.highlights = arr;
  }

  function moveDown(index) {
    if (index === form.highlights.length - 1) return;
    const arr = [...form.highlights];
    [arr[index], arr[index + 1]] = [arr[index + 1], arr[index]];
    form.highlights = arr;
  }

  async function handleSave() {
    error = "";
    success = false;

    if (!form.title || !form.description) {
      error = "Title and Description are required.";
      return;
    }

    for (let i = 0; i < form.highlights.length; i++) {
      if (!form.highlights[i].label || !form.highlights[i].description) {
        error = `Highlights #${i + 1}: both Label and Description are required.`;
        return;
      }
    }

    saving = true;
    try {
      const res = await fetch("http://localhost:8080/api/admin/about", {
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
        setTimeout(() => (success = false), 3000);
      }
    } catch {
      error = "Could not reach server.";
    } finally {
      saving = false;
    }
  }
</script>

<svelte:head>
  <title>About — Admin Panel</title>
</svelte:head>

<div class="max-w-5xl mx-auto px-4 sm:px-6">
  <div class="mb-8">
    <div class="flex items-center gap-2 text-white/30 text-sm mb-1">
      <i class="fa-solid fa-shield-halved text-primary"></i>
      <span>Admin</span>
      <i class="fa-solid fa-chevron-right text-[10px]"></i>
      <span class="text-white/60">About</span>
    </div>
    <h1 class="text-2xl font-bold text-white">About Section</h1>
    <p class="text-sm text-white/40 mt-1">
      Edit your about title, description and highlight cards.
    </p>
  </div>

  {#if loading}
    <div class="flex items-center justify-center py-24">
      <span
        class="w-8 h-8 rounded-full border-2 border-white/10 border-t-primary animate-spin"
      ></span>
    </div>
  {:else}
    {#if success}
      <div
        class="flex items-center gap-2 bg-green-500/[0.08] border border-green-500/20 rounded-lg px-4 py-3 mb-6 text-sm text-green-400"
        role="status"
      >
        <i class="fa-solid fa-circle-check shrink-0"></i>
        <span>About section updated successfully.</span>
      </div>
    {/if}

    {#if error}
      <div
        class="flex items-center gap-2 bg-red-500/[0.08] border border-red-500/20 rounded-lg px-4 py-3 mb-6 text-sm text-red-400"
        role="alert"
      >
        <i class="fa-solid fa-circle-exclamation shrink-0"></i>
        <span>{error}</span>
      </div>
    {/if}

    <div
      class="bg-dark border border-white/[0.07] rounded-2xl p-5 sm:p-8 shadow-[0_8px_32px_rgba(0,0,0,0.35)] mb-6"
    >
      <h2
        class="text-sm font-semibold text-white/70 mb-6 flex items-center gap-2"
      >
        <i class="fa-solid fa-pen text-primary text-xs"></i>
        Basic Info
      </h2>

      <div class="space-y-6">
        <div class="flex flex-col gap-1.5">
          <label
            for="title"
            class="text-[11px] font-medium text-white/40 uppercase tracking-widest"
          >
            Title <span class="text-red-400">*</span>
          </label>
          <input
            id="title"
            type="text"
            bind:value={form.title}
            placeholder="e.g. About Me"
            class="w-full bg-darker border border-white/10 rounded-lg px-4 py-3 text-white placeholder:text-white/20 outline-none focus:border-primary/60 focus:ring-2 focus:ring-primary/15 transition-all appearance-none"
          />
        </div>

        <div class="flex flex-col gap-1.5">
          <label
            for="description"
            class="text-[11px] font-medium text-white/40 uppercase tracking-widest"
          >
            Description <span class="text-red-400">*</span>
          </label>
          <textarea
            id="description"
            bind:value={form.description}
            rows="5"
            placeholder="Write a short bio..."
            class="w-full bg-darker border border-white/10 rounded-lg px-4 py-3 text-white placeholder:text-white/20 outline-none focus:border-primary/60 focus:ring-2 focus:ring-primary/15 transition-all appearance-none resize-y"
          ></textarea>
          <p class="text-[10px] text-white/25 text-right">
            {form.description.length} characters
          </p>
        </div>
      </div>
    </div>

    <div
      class="bg-dark border border-white/[0.07] rounded-2xl p-5 sm:p-8 shadow-[0_8px_32px_rgba(0,0,0,0.35)] mb-6"
    >
      <div
        class="flex flex-col sm:flex-row sm:items-center justify-between gap-4 mb-6"
      >
        <h2 class="text-sm font-semibold text-white/70 flex items-center gap-2">
          <i class="fa-solid fa-list-check text-primary text-xs"></i>
          Highlights
          <span class="text-xs font-normal text-white/30 ml-1"
            >({form.highlights.length})</span
          >
        </h2>
        <button
          on:click={addHighlight}
          class="flex items-center justify-center gap-1.5 px-4 py-2 rounded-lg text-xs font-medium text-primary border border-primary/30 bg-primary/5 hover:bg-primary/15 transition-all cursor-pointer"
        >
          <i class="fa-solid fa-plus text-[10px]"></i> Add Highlight
        </button>
      </div>

      {#if form.highlights.length === 0}
        <div
          class="flex flex-col items-center justify-center py-12 border border-dashed border-white/10 rounded-xl text-center"
        >
          <i class="fa-solid fa-rectangle-list text-3xl text-white/10 mb-3"></i>
          <p class="text-sm text-white/30">No highlights yet.</p>
        </div>
      {:else}
        <div class="space-y-4">
          {#each form.highlights as highlight, i}
            <div
              class="border border-white/[0.07] rounded-xl p-4 sm:p-5 bg-darker/30 relative"
            >
              <div class="flex items-center justify-between mb-4">
                <span
                  class="text-[10px] font-bold text-white/20 bg-white/[0.03] px-2 py-1 rounded"
                  >#{i + 1}</span
                >
                <div class="flex items-center gap-2">
                  <button
                    aria-label="move-up"
                    on:click={() => moveUp(i)}
                    disabled={i === 0}
                    class="w-8 h-8 flex items-center justify-center rounded-md border border-white/[0.08] text-white/30 hover:text-white/70 disabled:opacity-20 transition-all"
                    ><i class="fa-solid fa-chevron-up text-xs"></i></button
                  >
                  <button
                    aria-label="move-down"
                    on:click={() => moveDown(i)}
                    disabled={i === form.highlights.length - 1}
                    class="w-8 h-8 flex items-center justify-center rounded-md border border-white/[0.08] text-white/30 hover:text-white/70 disabled:opacity-20 transition-all"
                    ><i class="fa-solid fa-chevron-down text-xs"></i></button
                  >
                  <button
                    aria-label="removehighlight"
                    on:click={() => removeHighlight(i)}
                    class="w-8 h-8 flex items-center justify-center rounded-md border border-white/[0.08] text-white/30 hover:text-red-400 hover:bg-red-400/10 transition-all"
                    ><i class="fa-solid fa-trash-can text-xs"></i></button
                  >
                </div>
              </div>

              <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
                <div class="flex flex-col gap-1.5">
                  <label
                    for="label"
                    class="text-[10px] font-medium text-white/40 uppercase tracking-widest"
                    >Label</label
                  >
                  <input
                    type="text"
                    value={highlight.label}
                    on:input={(e) =>
                      updateHighlight(i, "label", e.currentTarget.value)}
                    placeholder="Experience"
                    class="w-full bg-darker border border-white/10 rounded-lg px-4 py-2.5 text-sm text-white outline-none focus:border-primary/50 transition-all"
                  />
                </div>
                <div class="flex flex-col gap-1.5">
                  <label
                    for="value"
                    class="text-[10px] font-medium text-white/40 uppercase tracking-widest"
                    >Value</label
                  >
                  <input
                    type="text"
                    value={highlight.description}
                    on:input={(e) =>
                      updateHighlight(i, "description", e.currentTarget.value)}
                    placeholder="2+ Years"
                    class="w-full bg-darker border border-white/10 rounded-lg px-4 py-2.5 text-sm text-white outline-none focus:border-primary/50 transition-all"
                  />
                </div>
              </div>
            </div>
          {/each}
        </div>
      {/if}
    </div>

    <div
      class="bg-dark border border-white/[0.07] rounded-2xl px-5 py-4 flex flex-col sm:flex-row items-center justify-between gap-4"
    >
      <p class="text-[11px] text-white/25 italic order-2 sm:order-1">
        <span class="text-red-400">*</span> Required fields
      </p>
      <div class="flex items-center gap-3 w-full sm:w-auto order-1 sm:order-2">
        <button
          on:click={() => window.location.reload()}
          disabled={saving}
          class="flex-1 sm:flex-none px-5 py-2.5 rounded-lg text-sm font-medium text-white/40 border border-white/[0.08] hover:bg-white/[0.04] transition-all"
        >
          Reset
        </button>
        <button
          on:click={handleSave}
          disabled={saving}
          class="flex-1 sm:flex-none flex items-center justify-center gap-2 px-6 py-2.5 rounded-lg text-sm font-semibold text-primary bg-primary/10 border border-primary/30 hover:bg-primary/20 transition-all"
        >
          {#if saving}
            <span
              class="w-3.5 h-3.5 rounded-full border-2 border-primary/30 border-t-primary animate-spin"
            ></span>
            Saving...
          {:else}
            <i class="fa-solid fa-floppy-disk"></i> Save Changes
          {/if}
        </button>
      </div>
    </div>
  {/if}
</div>
