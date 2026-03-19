<script>
  const BASE = 'http://localhost:8080/api/admin/certifications';

  let certs         = [];
  let loading       = true;
  let toast         = null;    // { msg, type }
  let certModal     = null;    // { mode, id, name, company, link, issueDate, order }
  let deleteConfirm = null;    // { label, onConfirm }

  //  fetch 
  async function fetchCerts() {
    loading = true;
    try {
      const res = await fetch(BASE, { credentials: 'include' });
      certs = await res.json();
    } catch {
      showToast('Failed to load certifications', 'error');
    } finally {
      loading = false;
    }
  }
  fetchCerts();

  //  toast 
  function showToast(msg, type = 'success') {
    toast = { msg, type };
    setTimeout(() => (toast = null), 3000);
  }

  //  helpers 
  function formatDate(dateStr) {
    if (!dateStr) return '';
    return new Date(dateStr).toLocaleDateString('en-US', { month: 'short', year: 'numeric' });
  }
  function toInputDate(dateStr) {
    if (!dateStr) return '';
    return new Date(dateStr).toISOString().slice(0, 10);
  }

  //  modals 
  function openAddCert() {
    certModal = { mode: 'add', id: null, name: '', company: '', link: '', issueDate: '', order: 0 };
  }
  function openEditCert(cert) {
    certModal = {
      mode: 'edit',
      id: cert.id || cert._id,
      name: cert.name,
      company: cert.company || '',
      link: cert.link || '',
      issueDate: toInputDate(cert.issueDate),
      order: cert.order || 0,
    };
  }

  async function submitCert() {
    const { mode, id, name, company, link, issueDate, order } = certModal;
    if (!name.trim()) return;
    const payload = {
      name, company, link, order: Number(order),
      issueDate: issueDate ? new Date(issueDate).toISOString() : '',
    };
    const url    = mode === 'add' ? BASE : `${BASE}/${id}`;
    const method = mode === 'add' ? 'POST' : 'PUT';
    const res = await fetch(url, {
      method, credentials: 'include',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(payload),
    });
    if (res.ok) {
      certModal = null;
      showToast(mode === 'add' ? 'Certification added' : 'Certification updated');
      await fetchCerts();
    } else showToast('Failed to save certification', 'error');
  }

  function confirmDeleteCert(cert) {
    deleteConfirm = {
      label: `Delete "${cert.name}"?`,
      onConfirm: async () => {
        const id  = cert.id || cert._id;
        const res = await fetch(`${BASE}/${id}`, { method: 'DELETE', credentials: 'include' });
        if (res.ok) { showToast('Certification deleted'); await fetchCerts(); }
        else showToast('Failed to delete', 'error');
        deleteConfirm = null;
      },
    };
  }

  function handleKeydown(e) {
    if (e.key !== 'Escape') return;
    if (deleteConfirm) { deleteConfirm = null; return; }
    if (certModal)     { certModal     = null; return; }
  }
</script>

<svelte:window on:keydown={handleKeydown} />

<!--  loading  -->
{#if loading}
  <div class="flex items-center justify-center py-20">
    <div class="w-8 h-8 border-2 border-primary border-t-transparent rounded-full animate-spin"></div>
  </div>

{:else}

<!--  toast  -->
{#if toast}
  <div role="status" aria-live="polite"
    class="fixed top-5 right-5 z-50 px-5 py-3 rounded-xl text-sm font-medium shadow-xl animate-slideIn
           {toast.type === 'error' ? 'bg-red-500/90 text-white' : 'bg-emerald-500/90 text-white'}">
    <i class="fa-solid {toast.type === 'error' ? 'fa-circle-xmark' : 'fa-circle-check'} mr-2"></i>{toast.msg}
  </div>
{/if}

<div class="space-y-6">

  <!-- header row -->
  <div class="flex items-center justify-between">
    <p class="text-gray-500 text-sm">
      {certs.length} certification{certs.length !== 1 ? 's' : ''}
    </p>
    <button type="button" on:click={openAddCert}
      class="flex items-center gap-2 px-4 py-2 rounded-xl bg-primary/10 border border-primary/30
             text-primary text-sm font-semibold hover:bg-primary/20 transition">
      <i class="fa-solid fa-plus"></i> Add Certification
    </button>
  </div>

  <!-- empty state -->
  {#if certs.length === 0}
    <div class="text-center py-16 border border-dashed border-gray-700 rounded-2xl text-gray-500">
      <i class="fa-solid fa-graduation-cap text-4xl mb-3 block"></i>
      <p>No certifications yet. Add your first one!</p>
    </div>

  {:else}
    <div class="grid grid-cols-1 sm:grid-cols-2 xl:grid-cols-3 gap-4">
      {#each certs as cert}
        <div class="bg-dark/40 border border-gray-800 rounded-2xl p-5
                    hover:border-gray-700 transition-all duration-200 flex flex-col gap-3">

          <!-- badge + title -->
          <div class="flex items-start gap-3">
            <div class="w-10 h-10 rounded-xl bg-primary/10 border border-primary/20
                        flex items-center justify-center shrink-0 text-primary"
                 aria-hidden="true">
              <i class="fa-solid fa-certificate"></i>
            </div>
            <div class="flex-1 min-w-0">
              <p class="text-white font-semibold text-sm leading-snug line-clamp-2">{cert.name}</p>
              <p class="text-gray-400 text-xs mt-0.5">{cert.company || '—'}</p>
            </div>
          </div>

          <!-- meta -->
          <div class="flex items-center justify-between text-xs text-gray-500">
            <span>
              {#if cert.issueDate}
                <i class="fa-regular fa-calendar mr-1"></i>Issued: {formatDate(cert.issueDate)}
              {:else}
                No date set
              {/if}
            </span>
            {#if cert.order}
              <span class="px-2 py-0.5 rounded-full bg-gray-800 text-gray-400">
                <i class="fa-solid fa-hashtag text-[10px]"></i>{cert.order}
              </span>
            {/if}
          </div>

          {#if cert.link}
            <a href={cert.link} target="_blank" rel="noopener noreferrer"
               class="text-xs text-primary hover:text-secondary transition truncate">
              <i class="fa-solid fa-arrow-up-right-from-square mr-1"></i>{cert.link}
            </a>
          {/if}

          <!-- actions -->
          <div class="flex gap-2 pt-1 border-t border-gray-800/60 mt-auto">
            <button type="button" on:click={() => openEditCert(cert)}
              class="flex-1 py-1.5 rounded-lg border border-gray-700 text-gray-400 text-xs
                     font-medium hover:border-blue-500/50 hover:text-blue-400 transition
                     flex items-center justify-center gap-1.5">
              <i class="fa-solid fa-pen text-[11px]"></i> Edit
            </button>
            <button type="button" on:click={() => confirmDeleteCert(cert)}
              class="flex-1 py-1.5 rounded-lg border border-gray-700 text-gray-400 text-xs
                     font-medium hover:border-red-500/50 hover:text-red-400 transition
                     flex items-center justify-center gap-1.5">
              <i class="fa-solid fa-trash text-[11px]"></i> Delete
            </button>
          </div>

        </div>
      {/each}
    </div>
  {/if}

</div>
{/if}

<!--  cert  -->
{#if certModal}
  <div role="dialog" aria-modal="true"
       aria-label="{certModal.mode === 'add' ? 'Add' : 'Edit'} Certification"
       class="fixed inset-0 z-40 flex items-center justify-center p-4">

    <button type="button" on:click={() => (certModal = null)} tabindex="-1"
      class="absolute inset-0 w-full h-full bg-black/60 backdrop-blur-sm cursor-default"
      aria-label="Close modal">
    </button>

    <div class="relative bg-[#0f1117] border border-gray-700 rounded-2xl w-full max-w-lg p-6 space-y-5 shadow-2xl">
      <div class="flex items-center justify-between">
        <h3 class="text-white font-bold text-lg">
          {certModal.mode === 'add' ? 'Add Certification' : 'Edit Certification'}
        </h3>
        <button type="button" on:click={() => (certModal = null)} aria-label="Close"
          class="w-8 h-8 flex items-center justify-center rounded-lg text-gray-500
                 hover:text-white hover:bg-white/5 transition">
          <i class="fa-solid fa-xmark"></i>
        </button>
      </div>

      <div class="space-y-4">
        <div>
          <label for="cert-name" class="block text-xs text-gray-400 mb-1.5 font-medium">
            Name <span class="text-red-400" aria-hidden="true">*</span>
          </label>
          <input id="cert-name" bind:value={certModal.name}
            placeholder="e.g. AWS Solutions Architect"
            class="w-full bg-dark/60 border border-gray-700 focus:border-primary rounded-xl
                   px-4 py-2.5 text-white text-sm outline-none transition-colors" />
        </div>

        <div class="grid grid-cols-2 gap-4">
          <div>
            <label for="cert-company" class="block text-xs text-gray-400 mb-1.5 font-medium">
              Issuing Company
            </label>
            <input id="cert-company" bind:value={certModal.company} placeholder="e.g. Amazon"
              class="w-full bg-dark/60 border border-gray-700 focus:border-primary rounded-xl
                     px-4 py-2.5 text-white text-sm outline-none transition-colors" />
          </div>
          <div>
            <label for="cert-order" class="block text-xs text-gray-400 mb-1.5 font-medium">Order</label>
            <input id="cert-order" type="number" bind:value={certModal.order} min="0"
              class="w-full bg-dark/60 border border-gray-700 focus:border-primary rounded-xl
                     px-4 py-2.5 text-white text-sm outline-none transition-colors" />
          </div>
        </div>

        <div>
          <label for="cert-link" class="block text-xs text-gray-400 mb-1.5 font-medium">
            Credential Link
          </label>
          <input id="cert-link" bind:value={certModal.link} placeholder="https://..."
            class="w-full bg-dark/60 border border-gray-700 focus:border-primary rounded-xl
                   px-4 py-2.5 text-white text-sm outline-none transition-colors" />
        </div>

        <div>
          <label for="cert-date" class="block text-xs text-gray-400 mb-1.5 font-medium">
            Issue Date
          </label>
          <input id="cert-date" type="date" bind:value={certModal.issueDate}
            class="w-full bg-dark/60 border border-gray-700 focus:border-primary rounded-xl
                   px-4 py-2.5 text-white text-sm outline-none transition-colors" />
        </div>
      </div>

      <div class="flex gap-3 pt-1">
        <button type="button" on:click={submitCert}
          class="flex-1 py-2.5 rounded-xl bg-primary text-slate-900 text-sm font-bold
                 hover:brightness-110 transition">
          {certModal.mode === 'add' ? 'Add Certification' : 'Save Changes'}
        </button>
        <button type="button" on:click={() => (certModal = null)}
          class="px-5 py-2.5 rounded-xl border border-gray-700 text-gray-400 text-sm
                 hover:border-gray-500 transition">
          Cancel
        </button>
      </div>
    </div>
  </div>
{/if}

<!-- delete -->
{#if deleteConfirm}
  <div role="alertdialog" aria-modal="true" aria-label="Confirm deletion"
       class="fixed inset-0 z-50 flex items-center justify-center p-4">

    <button type="button" on:click={() => (deleteConfirm = null)} tabindex="-1"
      class="absolute inset-0 w-full h-full bg-black/70 backdrop-blur-sm cursor-default"
      aria-label="Close modal">
    </button>

    <div class="relative bg-[#0f1117] border border-red-900/50 rounded-2xl w-full max-w-sm p-6 space-y-5 shadow-2xl">
      <div class="flex items-start gap-4">
        <div class="w-10 h-10 rounded-xl bg-red-500/10 border border-red-500/20
                    flex items-center justify-center shrink-0 text-red-400" aria-hidden="true">
          <i class="fa-solid fa-triangle-exclamation"></i>
        </div>
        <div>
          <h3 class="text-white font-bold text-base mb-1">Confirm Delete</h3>
          <p class="text-gray-400 text-sm">{deleteConfirm.label}</p>
        </div>
      </div>
      <div class="flex gap-3">
        <button type="button" on:click={deleteConfirm.onConfirm}
          class="flex-1 py-2.5 rounded-xl bg-red-500 text-white text-sm font-bold
                 hover:bg-red-600 transition">
          <i class="fa-solid fa-trash mr-2"></i>Delete
        </button>
        <button type="button" on:click={() => (deleteConfirm = null)}
          class="px-5 py-2.5 rounded-xl border border-gray-700 text-gray-400 text-sm
                 hover:border-gray-500 transition">
          Cancel
        </button>
      </div>
    </div>
  </div>
{/if}

<style>
  @keyframes slideIn {
    from { opacity: 0; transform: translateX(1rem); }
    to   { opacity: 1; transform: translateX(0); }
  }
  .animate-slideIn { animation: slideIn 0.25s ease-out; }
</style>