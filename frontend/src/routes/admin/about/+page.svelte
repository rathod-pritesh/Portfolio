<script>
  import { onMount } from "svelte";

  let form = { _id: "about", title: "", description: "", highlights: [] };
  let loading = true;
  let saving = false;
  let toast = null;

  const API = "http://127.0.0.1:8080/api";

  onMount(async () => {
    try {
      const res = await fetch(`${API}/about`, { credentials: "include" });
      const data = await res.json();
      form = {
        _id: data._id ?? "about",
        title: data.title ?? "",
        description: data.description ?? "",
        highlights: data.highlights ?? [],
      };
    } catch {
      showToast("error", "Failed to load About data");
    } finally {
      loading = false;
    }
  });

  async function save() {
    saving = true;
    try {
      const res = await fetch(`${API}/admin/about`, {
        method: "PUT",
        headers: { "Content-Type": "application/json" },
        credentials: "include",
        body: JSON.stringify(form),
      });
      const data = await res.json();
      if (res.ok) showToast("success", data.message ?? "Saved!");
      else showToast("error", data.error ?? "Save failed");
    } catch {
      showToast("error", "Network error");
    } finally {
      saving = false;
    }
  }

  function addHighlight() {
    form.highlights = [...form.highlights, { label: "", description: "" }];
  }

  function removeHighlight(i) {
    form.highlights = form.highlights.filter((_, idx) => idx !== i);
  }

  function showToast(type, msg) {
    toast = { type, msg };
    setTimeout(() => (toast = null), 3500);
  }
</script>

<svelte:head><title>Edit About | Admin</title></svelte:head>

{#if toast}
  <div
    class="toast {toast.type === 'success' ? 'toast-success' : 'toast-error'}"
  >
    <i
      class="fa-solid {toast.type === 'success'
        ? 'fa-circle-check'
        : 'fa-circle-exclamation'}"
    ></i>
    {toast.msg}
  </div>
{/if}

<div class="max-w-2xl space-y-6">
  <div>
    <h2 class="text-xl font-bold text-white">About Section</h2>
    <p class="text-gray-500 text-sm mt-1">
      Edit your about title, description and highlight cards.
    </p>
  </div>

  {#if loading}
    <div class="flex items-center gap-3 text-gray-500 text-sm">
      <div class="spinner"></div>
       Loading…
    </div>
  {:else}
    <div class="card space-y-5">
      <!-- title -->
      <div class="field-wrap">
        <label class="field-label" for="title">
          <i class="fa-solid fa-heading"></i> Section Title
        </label>
        <input
          id="title"
          bind:value={form.title}
          type="text"
          placeholder="About Me"
          class="field-input"
        />
      </div>

      <!-- desc -->
      <div class="field-wrap">
        <label class="field-label" for="description">
          <i class="fa-solid fa-align-left"></i> Description
        </label>
        <textarea
          id="description"
          bind:value={form.description}
          rows="5"
          placeholder="Write a few paragraphs about yourself…"
          class="field-input"
          style="resize:vertical"
        ></textarea>
      </div>

      <!-- highlights -->
      <div class="field-wrap">
        <div class="flex items-center justify-between mb-2">
          <span class="field-label">
            <i class="fa-solid fa-star"></i> Highlights
          </span>
          <button on:click={addHighlight} class="add-btn">
            <i class="fa-solid fa-plus"></i> Add
          </button>
        </div>

        {#if form.highlights.length === 0}
          <p class="text-xs text-gray-600 italic py-2">
            No highlights yet. Click Add to create one.
          </p>
        {/if}

        <div class="space-y-3">
          {#each form.highlights as hl, i}
            <div class="highlight-card">
              <div class="flex items-center justify-between mb-2">
                <span
                  class="text-xs text-gray-500 font-medium uppercase tracking-wide"
                >
                  Highlight #{i + 1}
                </span>
                <button
                  on:click={() => removeHighlight(i)}
                  class="remove-btn"
                  aria-label="Remove"
                >
                  <i class="fa-solid fa-trash-can"></i>
                </button>
              </div>
              <div class="space-y-2">
                <input
                  bind:value={hl.label}
                  type="text"
                  placeholder="Label (e.g. Experience)"
                  class="field-input text-xs"
                />
                <textarea
                  bind:value={hl.description}
                  rows="2"
                  placeholder="Short description…"
                  class="field-input text-xs"
                  style="resize:vertical"
                ></textarea>
              </div>
            </div>
          {/each}
        </div>
      </div>

      <!-- save -->
      <div class="pt-2">
        <button on:click={save} disabled={saving} class="save-btn">
          {#if saving}
            <div class="spinner"></div>
             Saving…
          {:else}
            <i class="fa-solid fa-floppy-disk"></i> Save Changes
          {/if}
        </button>
      </div>
    </div>
  {/if}
</div>

<style>
  .card {
    background: #0f172a;
    border: 1px solid #1e293b;
    border-radius: 1rem;
    padding: 1.5rem;
  }
  .field-wrap {
    display: flex;
    flex-direction: column;
    gap: 0.375rem;
  }
  .field-label {
    font-size: 0.75rem;
    font-weight: 500;
    color: #94a3b8;
    text-transform: uppercase;
    letter-spacing: 0.05em;
    display: flex;
    align-items: center;
    gap: 0.375rem;
  }
  .field-input {
    width: 100%;
    background: #020617;
    border: 1px solid #1e293b;
    border-radius: 0.5rem;
    padding: 0.625rem 0.875rem;
    font-size: 0.875rem;
    color: #e2e8f0;
    outline: none;
    transition: border-color 0.2s;
    font-family: inherit;
  }
  .field-input:focus {
    border-color: rgba(226, 232, 240, 0.35);
  }
  .field-input::placeholder {
    color: #334155;
  }
  .highlight-card {
    background: #020617;
    border: 1px solid #1e293b;
    border-radius: 0.75rem;
    padding: 0.875rem;
  }
  .add-btn {
    display: inline-flex;
    align-items: center;
    gap: 0.375rem;
    padding: 0.3rem 0.75rem;
    border-radius: 0.4rem;
    font-size: 0.75rem;
    font-weight: 600;
    background: rgba(226, 232, 240, 0.06);
    color: #94a3b8;
    border: 1px solid #1e293b;
    cursor: pointer;
    transition: all 0.2s;
  }
  .add-btn:hover {
    background: rgba(226, 232, 240, 0.1);
    color: #e2e8f0;
  }
  .remove-btn {
    background: none;
    border: none;
    cursor: pointer;
    color: #475569;
    font-size: 0.75rem;
    transition: color 0.2s;
    padding: 0.25rem;
  }
  .remove-btn:hover {
    color: #f87171;
  }
  .save-btn {
    display: inline-flex;
    align-items: center;
    gap: 0.5rem;
    padding: 0.625rem 1.5rem;
    border-radius: 0.5rem;
    font-size: 0.875rem;
    font-weight: 600;
    background: rgba(226, 232, 240, 0.08);
    color: #e2e8f0;
    border: 1px solid rgba(226, 232, 240, 0.2);
    cursor: pointer;
    transition: all 0.2s;
  }
  .save-btn:hover:not(:disabled) {
    background: rgba(226, 232, 240, 0.15);
    border-color: rgba(226, 232, 240, 0.4);
  }
  .save-btn:disabled {
    opacity: 0.5;
    cursor: not-allowed;
  }
  .spinner {
    width: 1rem;
    height: 1rem;
    border: 2px solid rgba(226, 232, 240, 0.2);
    border-top-color: #e2e8f0;
    border-radius: 50%;
    animation: spin 0.7s linear infinite;
    display: inline-block;
  }
  @keyframes spin {
    to {
      transform: rotate(360deg);
    }
  }
  .toast {
    position: fixed;
    top: 1.25rem;
    right: 1.25rem;
    z-index: 50;
    display: flex;
    align-items: center;
    gap: 0.5rem;
    padding: 0.75rem 1.25rem;
    border-radius: 0.75rem;
    font-size: 0.875rem;
    font-weight: 500;
    box-shadow: 0 10px 30px rgba(0, 0, 0, 0.4);
  }
  .toast-success {
    background: rgba(34, 197, 94, 0.15);
    border: 1px solid rgba(34, 197, 94, 0.3);
    color: #4ade80;
  }
  .toast-error {
    background: rgba(239, 68, 68, 0.15);
    border: 1px solid rgba(239, 68, 68, 0.3);
    color: #f87171;
  }
</style>
