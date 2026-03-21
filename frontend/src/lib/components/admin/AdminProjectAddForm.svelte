<script>
  import { addProject, updateProject } from "$lib/stores/projectStore";
  import { or } from "ajv/dist/compile/codegen";

  let { onCancel, project=null } = $props();

  const isEdit = $derived(!!project);
  const projectId = $derived(project?._id ?? "");

  let title = $state("");
  $effect(() => { title = project?.title ?? "" })
  let description = $state("");
  $effect(() => { description = project?.description ?? "" })
  let image = $state("");
  $effect(() => { image = project?.image ?? "" })
  let github = $state("");
  $effect(() => { github = project?.github ?? "" })
  let technologies = $state("");
  $effect(() => { technologies = (project?.technologies ?? []).join(", ") })
  let gradient = $state("");
  $effect(() => { gradient = project?.gradient ?? "" })
  let order = $state("");
  $effect(() => { order = project?.order ?? 0 })
  let isFeatured = $state("");
  $effect(() => { isFeatured =  project?.isFeatured ?? false })

  let submitting = $state(false);
  let formError = $state("");

  async function handleSubmit(e) {
    e.preventDefault();
    submitting = true;
    formError = "";

    const payload = {
      title,
      description,
      image,
      github,
      // convert comma-separated string back to a trimmed array, drop empty entries
      technologies: technologies.split(",").map((t) => t.trim()).filter(Boolean),
      gradient,
      order: Number(order),
      isFeatured,
    };

    try {
      if (isEdit) {
        await updateProject(project._id, payload);
      } else {
        await addProject(payload);
      }
      onCancel();
    } catch (err) {
      formError = err.message;
    } finally {
      submitting = false;
    }
  }
</script>
<div class="bg-slate-900 border border-slate-800 rounded-xl shadow-2xl p-8">
  <h2 class="text-2xl font-bold mb-6 text-white border-b border-slate-800 pb-4">
    <i class="fa-solid fa-diagram-project mr-2 text-blue-500"></i>
    {isEdit ? "Edit Project" : "Add New Project"}
  </h2>
 
  <form onsubmit={handleSubmit} class="space-y-5">
    <div class="grid grid-cols-1 md:grid-cols-4 gap-4">
      <div class="md:col-span-3">
        <label for="title" class="block text-sm font-medium mb-2 text-slate-400">Title*</label>
        <input
          type="text"
          id="title"
          bind:value={title}
          placeholder="Enter project title"
          required
          class="w-full bg-slate-800 border border-slate-700 rounded-lg px-4 py-2.5 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent transition-all"
        />
      </div>
 
      <div>
        <label for="order" class="block text-sm font-medium mb-2 text-slate-400">Order*</label>
        <input
          type="number"
          id="order"
          bind:value={order}
          placeholder="0"
          required
          class="w-full bg-slate-800 border border-slate-700 rounded-lg px-4 py-2.5 focus:outline-none focus:ring-2 focus:ring-blue-500 transition-all"
        />
      </div>
    </div>
 
    <div>
      <label for="description" class="block text-sm font-medium mb-2 text-slate-400">Description*</label>
      <textarea
        id="description"
        bind:value={description}
        rows="3"
        placeholder="Describe the project..."
        required
        class="w-full bg-slate-800 border border-slate-700 rounded-lg px-4 py-2.5 focus:outline-none focus:ring-2 focus:ring-blue-500 transition-all"
      ></textarea>
    </div>
 
    <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
      <div>
        <label for="image" class="block text-sm font-medium mb-2 text-slate-400">Image URL</label>
        <input
          type="text"
          id="image"
          bind:value={image}
          placeholder="https://..."
          class="w-full bg-slate-800 border border-slate-700 rounded-lg px-4 py-2.5 focus:outline-none focus:ring-2 focus:ring-blue-500 transition-all"
        />
      </div>
      <div>
        <label for="github" class="block text-sm font-medium mb-2 text-slate-400">GitHub URL</label>
        <input
          type="text"
          id="github"
          bind:value={github}
          placeholder="https://github.com/..."
          class="w-full bg-slate-800 border border-slate-700 rounded-lg px-4 py-2.5 focus:outline-none focus:ring-2 focus:ring-blue-500 transition-all"
        />
      </div>
    </div>
 
      <div>
        <label for="gradient" class="block text-sm font-medium mb-2 text-slate-400">Gradient</label>
        <input
          type="text"
          id="gradient"
          bind:value={gradient}
          placeholder="from-blue-500 to-purple-600"
          class="w-full bg-slate-800 border border-slate-700 rounded-lg px-4 py-2.5 focus:outline-none focus:ring-2 focus:ring-blue-500 transition-all"
        />
      </div>
 
    <div>
      <label for="technologies" class="block text-sm font-medium mb-2 text-slate-400">Technologies (comma separated)</label>
      <input
        type="text"
        id="technologies"
        bind:value={technologies}
        placeholder="Go, SvelteKit, Tailwind"
        class="w-full bg-slate-800 border border-slate-700 rounded-lg px-4 py-2.5 focus:outline-none focus:ring-2 focus:ring-blue-500 transition-all"
      />
    </div>
 
    <div class="flex items-center space-x-3 p-4 bg-slate-800/50 rounded-lg border border-slate-700">
      <input
        type="checkbox"
        id="featured"
        bind:checked={isFeatured}
        class="w-5 h-5 rounded border-slate-700 bg-slate-900 text-blue-500 focus:ring-blue-500 focus:ring-offset-slate-900"
      />
      <label for="featured" class="text-sm font-medium text-slate-300">Mark as Featured Project</label>
    </div>
 
    <!-- Error message -->
    {#if formError}
      <div class="flex items-center gap-2 p-3 bg-red-500/10 border border-red-500/30 rounded-lg text-red-400 text-sm">
        <i class="fa-solid fa-circle-exclamation"></i>
        {formError}
      </div>
    {/if}
 
    <div class="flex items-start gap-2 text-sm text-slate-400">
      <span class="text-blue-400">ℹ</span>
      <p>Fields marked with <span class="text-red-500 font-medium">*</span> are required.</p>
    </div>
 
    <div class="flex justify-end space-x-4 pt-4">
      <button
        type="button"
        onclick={onCancel}
        class="px-6 py-2.5 rounded-lg text-sm font-semibold text-slate-400 hover:text-white hover:bg-slate-800 transition-all"
        disabled={submitting}
      >
        <i class="fa-solid fa-xmark text-xl"></i>
        Cancel
      </button>
      <button
        type="submit"
        disabled={submitting}
        class="px-6 py-2.5 bg-blue-600 hover:bg-blue-500 disabled:opacity-50 disabled:cursor-not-allowed text-white rounded-lg text-sm font-semibold shadow-lg shadow-blue-900/20 transition-all flex items-center gap-2"
      >
        {#if submitting}
          <i class="fa-solid fa-circle-notch fa-spin"></i>
          {isEdit ? "Saving..." : "Creating..."}
        {:else}
          <i class="fa-solid fa-paper-plane"></i>
          {isEdit ? "Save Changes" : "Create Project"}
        {/if}
      </button>
    </div>
  </form>
</div>