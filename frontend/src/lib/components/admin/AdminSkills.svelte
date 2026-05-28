<script>
  import { showToast } from "$lib/stores/toastStore";
  const BASE = 'http://localhost:8080/api/admin/skills';

  let skills      = null;
  let loading     = true;
  let toast       = null;       // { msg, type: 'success'|'error' }
  let editingTitle = false;
  let titleDraft   = '';
  let catModal     = null;      // { mode:'add'|'edit', ci, name, icon }
  let skillModal   = null;      // { mode:'add'|'edit', ci, si, name, img }
  let deleteConfirm = null;     // { label, onConfirm }

  //  fetch 
  async function fetchSkills() {
    loading = true;
    try {
      const res = await fetch(BASE, { credentials: 'include' });
      skills = await res.json();
      titleDraft = skills.title;
    } catch {
      showToast('Failed to load skills', 'error');
    } finally {
      loading = false;
    }
  }
  fetchSkills();

  //  title 
  async function saveTitle() {
    const res = await fetch(`${BASE}/title`, {
      method: 'PUT', credentials: 'include',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ title: titleDraft }),
    });
    if (res.ok) { skills.title = titleDraft; editingTitle = false; showToast('Title updated'); }
    else showToast('Failed to update title', 'error');
  }

  //  categories 
  function openAddCat()    { catModal = { mode: 'add', ci: null, name: '', icon: '' }; }
  function openEditCat(ci) {
    const c = skills.categories[ci];
    catModal = { mode: 'edit', ci, name: c.name, icon: c.icon };
  }

  async function submitCat() {
    const { mode, ci, name, icon } = catModal;
    if (!name.trim()) return;
    const url    = mode === 'add' ? `${BASE}/category` : `${BASE}/category/${ci}`;
    const method = mode === 'add' ? 'POST' : 'PUT';
    const res = await fetch(url, {
      method, credentials: 'include',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ name, icon }),
    });
    if (res.ok) {
      const d = await res.json(); skills = d.data; catModal = null;
      showToast(mode === 'add' ? 'Category added' : 'Category updated');
    } else showToast('Failed to save category', 'error');
  }

  function confirmDeleteCat(ci) {
    deleteConfirm = {
      label: `Delete category "${skills.categories[ci].name}" and all its skills?`,
      onConfirm: async () => {
        const res = await fetch(`${BASE}/category/${ci}`, { method: 'DELETE', credentials: 'include' });
        if (res.ok) { const d = await res.json(); skills = d.data; showToast('Category deleted'); }
        else showToast('Failed to delete category', 'error');
        deleteConfirm = null;
      },
    };
  }

  //  skills 
  function openAddSkill(ci)      { skillModal = { mode: 'add', ci, si: null, name: '', img: '' }; }
  function openEditSkill(ci, si) {
    const s = skills.categories[ci].skills[si];
    skillModal = { mode: 'edit', ci, si, name: s.name, img: s.img };
  }

  async function submitSkill() {
    const { mode, ci, si, name, img } = skillModal;
    if (!name.trim()) return;
    const url    = mode === 'add' ? `${BASE}/category/${ci}/skill` : `${BASE}/category/${ci}/skill/${si}`;
    const method = mode === 'add' ? 'POST' : 'PUT';
    const res = await fetch(url, {
      method, credentials: 'include',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ name, img }),
    });
    if (res.ok) {
      const d = await res.json(); skills = d.data; skillModal = null;
      showToast(mode === 'add' ? 'Skill added' : 'Skill updated');
    } else showToast('Failed to save skill', 'error');
  }

  function confirmDeleteSkill(ci, si) {
    deleteConfirm = {
      label: `Delete skill "${skills.categories[ci].skills[si].name}"?`,
      onConfirm: async () => {
        const res = await fetch(`${BASE}/category/${ci}/skill/${si}`, { method: 'DELETE', credentials: 'include' });
        if (res.ok) { const d = await res.json(); skills = d.data; showToast('Skill deleted'); }
        else showToast('Failed to delete skill', 'error');
        deleteConfirm = null;
      },
    };
  }

  function handleKeydown(e) {
    if (e.key !== 'Escape') return;
    if (deleteConfirm) { deleteConfirm = null; return; }
    if (skillModal)    { skillModal    = null; return; }
    if (catModal)      { catModal      = null; return; }
  }
</script>

<svelte:window on:keydown={handleKeydown} />

<!--  loading  -->
{#if loading}
  <div class="flex items-center justify-center py-20">
    <div class="w-8 h-8 border-2 border-primary border-t-transparent rounded-full animate-spin"></div>
  </div>

{:else if skills}

<div class="space-y-8">

  <!--  section title  -->
  <div class="bg-dark/40 border border-gray-800 rounded-2xl p-6">
    <p class="text-xs font-semibold text-gray-500 uppercase tracking-widest mb-3">Section Title</p>

    {#if editingTitle}
      <div class="flex gap-3 items-center">
        <input bind:value={titleDraft} placeholder="Skills section title"
          class="flex-1 bg-dark/60 border border-gray-700 focus:border-primary rounded-xl
                 px-4 py-2.5 text-white text-sm outline-none transition-colors" />
        <button type="button" on:click={saveTitle}
          class="px-5 py-2.5 rounded-xl bg-primary text-slate-900 text-sm font-semibold hover:brightness-110 transition">
          Save
        </button>
        <button type="button" on:click={() => { editingTitle = false; titleDraft = skills.title; }}
          class="px-4 py-2.5 rounded-xl border border-gray-700 text-gray-400 text-sm hover:border-gray-500 transition">
          Cancel
        </button>
      </div>
    {:else}
      <div class="flex items-center justify-between">
        <span class="text-white font-bold text-xl">{skills.title}</span>
        <button type="button" on:click={() => (editingTitle = true)}
          class="flex items-center gap-2 px-4 py-2 rounded-xl border border-gray-700 text-gray-400
                 text-sm hover:border-primary hover:text-primary transition">
          <i class="fa-solid fa-pen-to-square"></i> Edit
        </button>
      </div>
    {/if}
  </div>

  <!--  categories  -->
  <div class="space-y-4">
    <div class="flex items-center justify-between">
      <h3 class="text-white font-bold text-lg">Categories</h3>
      <button type="button" on:click={openAddCat}
        class="flex items-center gap-2 px-4 py-2 rounded-xl bg-primary/10 border border-primary/30
               text-primary text-sm font-semibold hover:bg-primary/20 transition">
        <i class="fa-solid fa-plus"></i> Add Category
      </button>
    </div>

    {#if skills.categories.length === 0}
      <div class="text-center py-12 text-gray-500 border border-dashed border-gray-700 rounded-xl">
        No categories yet. Add one to get started.
      </div>
    {/if}

    {#each skills.categories as cat, ci}
      <div class="bg-dark/40 border border-gray-800 rounded-2xl overflow-hidden">

        <!-- category header -->
        <div class="flex items-center justify-between px-6 py-4 border-b border-gray-800/60">
          <div class="flex items-center gap-3">
            <span class="text-2xl" aria-hidden="true">{cat.icon || '📁'}</span>
            <div>
              <p class="text-white font-semibold">{cat.name}</p>
              <p class="text-xs text-gray-500">
                {cat.skills?.length ?? 0} skill{cat.skills?.length !== 1 ? 's' : ''}
              </p>
            </div>
          </div>

          <div class="flex gap-2">
            <button type="button" on:click={() => openAddSkill(ci)}
              class="px-3 py-1.5 rounded-lg bg-primary/10 text-primary text-xs font-semibold
                     border border-primary/20 hover:bg-primary/20 transition">
              <i class="fa-solid fa-plus mr-1"></i>Skill
            </button>
            <button type="button" on:click={() => openEditCat(ci)}
              aria-label="Edit category {cat.name}"
              class="w-8 h-8 rounded-lg border border-gray-700 text-gray-400
                     flex items-center justify-center
                     hover:border-blue-500/60 hover:text-blue-400 transition">
              <i class="fa-solid fa-pen text-xs"></i>
            </button>
            <button type="button" on:click={() => confirmDeleteCat(ci)}
              aria-label="Delete category {cat.name}"
              class="w-8 h-8 rounded-lg border border-gray-700 text-gray-400
                     flex items-center justify-center
                     hover:border-red-500/60 hover:text-red-400 transition">
              <i class="fa-solid fa-trash text-xs"></i>
            </button>
          </div>
        </div>

        <!-- Skills Grid -->
        <div class="p-5">
          {#if !cat.skills || cat.skills.length === 0}
            <p class="text-sm text-gray-600 italic">No skills yet.</p>
          {:else}
            <div class="flex flex-wrap gap-2">
              {#each cat.skills as skill, si}
                <div class="group relative flex items-center gap-2 px-3 py-2 rounded-xl
                            bg-primary/5 border border-primary/20 hover:border-primary/50 transition">
                  {#if skill.img}
                    <img src={skill.img} alt="" aria-hidden="true" class="w-5 h-5 object-contain" />
                  {/if}
                  <span class="text-sm text-gray-200">{skill.name}</span>

                  <span class="hidden group-hover:inline-flex items-center gap-1 ml-1">
                    <button type="button" on:click={() => openEditSkill(ci, si)}
                      aria-label="Edit skill {skill.name}"
                      class="p-0.5 text-blue-400 hover:text-blue-300 transition">
                      <i class="fa-solid fa-pen text-[10px]"></i>
                    </button>
                    <button type="button" on:click={() => confirmDeleteSkill(ci, si)}
                      aria-label="Delete skill {skill.name}"
                      class="p-0.5 text-red-400 hover:text-red-300 transition">
                      <i class="fa-solid fa-xmark text-[10px]"></i>
                    </button>
                  </span>
                </div>
              {/each}
            </div>
          {/if}
        </div>

      </div>
    {/each}
  </div>

</div>
{/if}

<!-- ══ category  -->
{#if catModal}
  <div role="dialog" aria-modal="true"
       aria-label="{catModal.mode === 'add' ? 'Add' : 'Edit'} Category"
       class="fixed inset-0 z-40 flex items-center justify-center p-4">

    <!-- Backdrop: a full-size button so click + keyboard both work -->
    <button type="button" on:click={() => (catModal = null)} tabindex="-1"
      class="absolute inset-0 w-full h-full bg-black/60 backdrop-blur-sm cursor-default"
      aria-label="Close modal">
    </button>

    <div class="relative bg-[#0f1117] border border-gray-700 rounded-2xl w-full max-w-md p-6 space-y-5 shadow-2xl">
      <div class="flex items-center justify-between">
        <h3 class="text-white font-bold text-lg">
          {catModal.mode === 'add' ? 'Add Category' : 'Edit Category'}
        </h3>
        <button type="button" on:click={() => (catModal = null)} aria-label="Close"
          class="w-8 h-8 flex items-center justify-center rounded-lg text-gray-500
                 hover:text-white hover:bg-white/5 transition">
          <i class="fa-solid fa-xmark"></i>
        </button>
      </div>

      <div class="space-y-4">
        <div>
          <label for="cat-name" class="block text-xs text-gray-400 mb-1.5 font-medium">
            Category Name <span class="text-red-400" aria-hidden="true">*</span>
          </label>
          <input id="cat-name" bind:value={catModal.name} placeholder="e.g. Frontend"
            class="w-full bg-dark/60 border border-gray-700 focus:border-primary rounded-xl
                   px-4 py-2.5 text-white text-sm outline-none transition-colors" />
        </div>
        <div>
          <label for="cat-icon" class="block text-xs text-gray-400 mb-1.5 font-medium">Icon (emoji)</label>
          <input id="cat-icon" bind:value={catModal.icon} placeholder="e.g. ⚛️"
            class="w-full bg-dark/60 border border-gray-700 focus:border-primary rounded-xl
                   px-4 py-2.5 text-white text-sm outline-none transition-colors" />
        </div>
      </div>

      <div class="flex gap-3 pt-1">
        <button type="button" on:click={submitCat}
          class="flex-1 py-2.5 rounded-xl bg-primary text-slate-900 text-sm font-bold
                 hover:brightness-110 transition">
          {catModal.mode === 'add' ? 'Add Category' : 'Save Changes'}
        </button>
        <button type="button" on:click={() => (catModal = null)}
          class="px-5 py-2.5 rounded-xl border border-gray-700 text-gray-400 text-sm
                 hover:border-gray-500 transition">
          Cancel
        </button>
      </div>
    </div>
  </div>
{/if}

<!-- ══ skill  -->
{#if skillModal}
  <div role="dialog" aria-modal="true"
       aria-label="{skillModal.mode === 'add' ? 'Add' : 'Edit'} Skill"
       class="fixed inset-0 z-40 flex items-center justify-center p-4">

    <button type="button" on:click={() => (skillModal = null)} tabindex="-1"
      class="absolute inset-0 w-full h-full bg-black/60 backdrop-blur-sm cursor-default"
      aria-label="Close modal">
    </button>

    <div class="relative bg-[#0f1117] border border-gray-700 rounded-2xl w-full max-w-md p-6 space-y-5 shadow-2xl">
      <div class="flex items-center justify-between">
        <h3 class="text-white font-bold text-lg">
          {skillModal.mode === 'add' ? 'Add Skill' : 'Edit Skill'}
        </h3>
        <button type="button" on:click={() => (skillModal = null)} aria-label="Close"
          class="w-8 h-8 flex items-center justify-center rounded-lg text-gray-500
                 hover:text-white hover:bg-white/5 transition">
          <i class="fa-solid fa-xmark"></i>
        </button>
      </div>

      <div class="space-y-4">
        <div>
          <label for="skill-name" class="block text-xs text-gray-400 mb-1.5 font-medium">
            Skill Name <span class="text-red-400" aria-hidden="true">*</span>
          </label>
          <input id="skill-name" bind:value={skillModal.name} placeholder="e.g. React"
            class="w-full bg-dark/60 border border-gray-700 focus:border-primary rounded-xl
                   px-4 py-2.5 text-white text-sm outline-none transition-colors" />
        </div>
        <div>
          <label for="skill-img" class="block text-xs text-gray-400 mb-1.5 font-medium">Image URL</label>
          <input id="skill-img" bind:value={skillModal.img} placeholder="https://..."
            class="w-full bg-dark/60 border border-gray-700 focus:border-primary rounded-xl
                   px-4 py-2.5 text-white text-sm outline-none transition-colors" />
        </div>

        {#if skillModal.img}
          <div class="flex items-center gap-3 p-3 bg-dark/40 rounded-xl border border-gray-800">
            <img src={skillModal.img} alt="Icon preview" class="w-8 h-8 object-contain" />
            <span class="text-sm text-gray-300">{skillModal.name || 'Preview'}</span>
          </div>
        {/if}
      </div>

      <div class="flex gap-3 pt-1">
        <button type="button" on:click={submitSkill}
          class="flex-1 py-2.5 rounded-xl bg-primary text-slate-900 text-sm font-bold
                 hover:brightness-110 transition">
          {skillModal.mode === 'add' ? 'Add Skill' : 'Save Changes'}
        </button>
        <button type="button" on:click={() => (skillModal = null)}
          class="px-5 py-2.5 rounded-xl border border-gray-700 text-gray-400 text-sm
                 hover:border-gray-500 transition">
          Cancel
        </button>
      </div>
    </div>
  </div>
{/if}

<!-- delete  -->
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