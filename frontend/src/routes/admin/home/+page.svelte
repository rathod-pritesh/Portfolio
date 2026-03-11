<script>
  import { onMount } from "svelte";

  let form = {
    _id: "home",
    name: "",
    role: "",
    tagline: "",
    resumeUrl: "",
    profileImage: "",
  };
  let loading = true;
  let saving = false;
  let toast = null;

  const API = "http://127.0.0.1:8080/api";

  onMount(async () => {
    try {
      const res = await fetch(`${API}/home`, { credentials: "include" });
      const data = await res.json();
      form = { ...form, ...data };
    } catch {
      showToast("error", "Failed to load Home data");
    } finally {
      loading = false;
    }
  });

  async function save() {
    saving = true;
    try {
      const res = await fetch(`${API}/admin/home`, {
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

  function showToast(type, msg) {
    toast = { type, msg };
    setTimeout(() => {
      toast = null;
    }, 3500);
  }
</script>

<svelte:head><title>Edit Home | Admin</title></svelte:head>

{#if toast}
  <div
    class="fixed top-5 right-5 z-50 flex items-center gap-2 px-5 py-3 rounded-xl text-sm font-medium shadow-xl
    {toast.type === 'success' ? 'toast-success' : 'toast-error'}"
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
    <h2 class="text-xl font-bold text-white">Home Section</h2>
    <p class="text-gray-500 text-sm mt-1">
      Edit your hero / landing section content
    </p>
  </div>

  {#if loading}
    <div class="flex items-center gap-3 text-gray-500 text-sm">
      <div class="spinner"></div>
      Loading...
    </div>
  {:else}
    <div class="card space-y-5">
      <div class="field-wrap">
        <label for="name" class="field-label">
          <i class="fa-solid fa-signature"></i> Name
        </label>
        <input
          id="name"
          bind:value={form.name}
          type="text"
          placeholder="Pritesh Rathod"
          class="field-input"
        />
      </div>

      <div class="field-wrap">
        <label for="role" class="field-label">
          <i class="fa-solid fa-briefcase"></i> Role / Title
        </label>
        <input
          id="role"
          bind:value={form.role}
          type="text"
          placeholder="Full-Stack Developer"
          class="field-input"
        />
      </div>

      <div class="field-wrap">
        <label for="tagline" class="field-label">
          <i class="fa-solid fa-quote-left"></i> Tagline
        </label>
        <textarea
          id="tagline"
          bind:value={form.tagline}
          rows="3"
          placeholder="A short tagline about yourself…"
          class="field-input"
          style="resize:vertical"
        ></textarea>
      </div>

      <div class="field-wrap">
        <label for="resumeUrl" class="field-label">
          <i class="fa-solid fa-file-pdf"></i> Resume URL
        </label>
        <input
          id="resumeUrl"
          type="url"
          bind:value={form.resumeUrl}
          placeholder="https://.../resume.pdf"
          class="field-input"
        />
      </div>

      <div class="field-wrap">
        <label class="field-label" for="profileImage">
          <i class="fa-solid fa-image"></i> Profile Image URL
        </label>
        <input
          id="profileImage"
          bind:value={form.profileImage}
          type="url"
          placeholder="https://…/photo.jpg"
          class="field-input"
        />
      </div>

      {#if form.profileImage}
        <div class="flex items-center gap-3 pt-1">
          <img
            src={form.profileImage}
            alt="Profile preview"
            class="w-14 h-14 rounded-full object-cover border border-gray-700"
          />
          <p class="text-xs text-gray-500">Image preview</p>
        </div>
      {/if}

      <div class="pt-2">
        <button on:click={save} disabled={saving} class="save-btn">
          {#if saving}
            <div class="spinner"></div>
            Saving...
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
