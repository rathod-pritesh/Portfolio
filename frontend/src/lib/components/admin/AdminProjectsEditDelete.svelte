<script>
  import { onMount } from "svelte";

  let projects = [];
  let loading = true;
  let error = "";

  const fetchProjects = async () => {
    try {
      const res = await fetch("http://localhost:8080/api/projects");
      const data = await res.json();
      projects = data || [];
    } catch (err) {
      error = "Failed to load projects";
      console.error(err);
    } finally {
      loading = false;
    }
  };

  onMount(fetchProjects);

  const handleDelete = async (id) => {
    if (!confirm("Are you sure you want to delete this project?")) return;
    try {
      await fetch(`http://localhost:8080/api/projects/${id}`, {
        method: "DELETE",
      });
      projects = projects.filter((p) => p._id !== id);
    } catch (err) {
      console.error("Delete Failed", err);
    }
  };

  const handleEdit = (project) => {
    console.log("Edit clicked:", project);
  };
</script>

<div
  class="bg-slate-950 rounded-xl border border-slate-800 shadow-2xl overflow-hidden"
>
  <div
    class="px-6 py-4 border-b border-slate-800 flex justify-between items-center bg-slate-900/50"
  >
    <h2 class="text-xl font-bold text-white flex items-center gap-2">
      <i class="fa-solid fa-list-check text-blue-500"></i>
      Project Inventory
    </h2>
    <span
      class="px-3 py-1 bg-slate-800 text-slate-400 text-xs rounded-full border border-slate-700"
    >
      {projects.length} Total Projects
    </span>
  </div>

  {#if loading}
    <div class="p-12 flex flex-col items-center justify-center space-y-4">
      <i class="fa-solid fa-circle-notch fa-spin text-3xl text-blue-500"></i>
      <p class="text-slate-400 animate-pulse text-sm">Retrieving projects...</p>
    </div>
  {:else if projects.length === 0}
    <div class="border-2 border-dashed border-slate-800 rounded-xl p-12 text-center">
    <p class="text-slate-500">
      No projects yet. Click <span class="text-blue-400 font-medium">"Add Project"</span> to create one.
    </p>
  </div>
  {:else}
    <div class="overflow-x-auto">
      <table class="w-full text-left border-collapse">
        <thead>
          <tr
            class="bg-slate-900 text-slate-400 text-xs uppercase tracking-wider font-semibold"
          >
            <th class="px-6 py-4">Project Details</th>
            <th class="px-6 py-4">Technologies</th>
            <th class="px-6 py-4 text-center">Order</th>
            <th class="px-6 py-4 text-center">Links</th>
            <th class="px-6 py-4 text-right">Actions</th>
          </tr>
        </thead>

        <tbody class="divide-y divide-slate-800">
          {#each projects as project (project._id)}
            <tr class="hover:bg-slate-800/40 transition-all group">
              <td class="px-6 py-4 max-w-xs">
                <div class="text-white font-medium truncate">
                  {project.title}
                </div>
                <div class="text-slate-500 text-sm truncate">
                  {project.description}
                </div>
              </td>

              <td class="px-6 py-4">
                <div class="flex flex-wrap gap-1.5">
                  {#if project.technologies}
                    {#each project.technologies as tech}
                      <span
                        class="px-2 py-0.5 bg-blue-500/10 border border-blue-500/20 text-blue-400 text-[10px] rounded uppercase font-bold tracking-tight"
                      >
                        {tech}
                      </span>
                    {/each}
                  {/if}
                </div>
              </td>

              <td class="px-6 py-4 text-center">
                <span
                  class="font-mono text-slate-400 bg-slate-800 px-2 py-1 rounded text-xs"
                >
                  {project.order}
                </span>
              </td>

              <td class="px-6 py-4 text-center">
                {#if project.github}
                  <a
                    href={project.github}
                    target="_blank"
                    class="text-slate-400 hover:text-white transition-colors"
                  >
                    <i class="fa-brands fa-github text-xl"></i>View
                  </a>
                {:else}
                  <span class="text-slate-700 text-xs">N/A</span>
                {/if}
              </td>

              <td class="px-6 py-4 text-right whitespace-nowrap">
                <div class="flex justify-end items-center gap-2">
                  <button
                    on:click={() => handleEdit(project)}
                    class="p-2 bg-slate-800 hover:bg-yellow-500/20 text-slate-400 hover:text-yellow-500 rounded-lg transition-all border border-slate-700"
                    title="Edit Project"
                  >
                    <i class="fa-solid fa-pen-to-square"></i>
                  </button>
                  <button
                    on:click={() => handleDelete(project._id)}
                    class="p-2 bg-slate-800 hover:bg-red-500/20 text-slate-400 hover:text-red-500 rounded-lg transition-all border border-slate-700"
                    title="Delete Project"
                  >
                    <i class="fa-solid fa-trash-can"></i>
                  </button>
                </div>
              </td>
            </tr>
          {/each}
        </tbody>
      </table>
    </div>
  {/if}
</div>

<style>
  .overflow-x-auto::-webkit-scrollbar {
    height: 6px;
  }
  .overflow-x-auto::-webkit-scrollbar-track {
    background: #020617;
  }
  .overflow-x-auto::-webkit-scrollbar-thumb {
    background: #1e293b;
    border-radius: 10px;
  }
  .overflow-x-auto::-webkit-scrollbar-thumb:hover {
    background: #334155;
  }
</style>
