<script>
  import AdminProjectAddForm from "$lib/components/admin/AdminProjectAddForm.svelte";
  import AdminProjectsEditDelete from "$lib/components/admin/AdminProjectsEditDelete.svelte";

  let showAddForm = $state(false);

  function toggleModal() {
    showAddForm = !showAddForm;
  }
</script>

<div class="min-h-screen bg-slate-950 p-8">
  <div class="max-w-6xl mx-auto">
    <div class="flex justify-between items-center mb-8 border-b border-slate-800 pb-6">
      <div>
        <h1 class="text-3xl font-bold text-white">Project Management</h1>
        <p class="text-slate-400 mt-1">Manage and showcase your portfolio work.</p>
      </div>

      <button
        onclick={toggleModal}
        class="px-6 py-2.5 bg-gradient-to-r from-blue-600 to-cyan-500 hover:from-blue-500 hover:to-cyan-400 text-white rounded-lg font-semibold transition-all shadow-lg shadow-blue-500/20 active:scale-95"
      ><i class="fa-solid fa-plus text-sm"></i>Add Project</button>
    </div>

    <AdminProjectsEditDelete />
  </div>
</div>

{#if showAddForm}
  <div class="fixed inset-0 z-50 flex items-center justify-center bg-black/80 backdrop-blur-sm p-4">
  <div class="relative w-full max-w-2xl max-h-[90vh] overflow-y-auto">
    <AdminProjectAddForm onCancel={toggleModal} />
  </div></div>
{/if}

