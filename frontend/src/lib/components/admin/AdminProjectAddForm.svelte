<script>
  let { onCancel } = $props();
</script>

<div
  class="min-h-screen bg-slate-950 text-slate-100 p-8 flex justfiy-cen
items-center"
>
  <div class="bg-slate-900 border border-slate-800 rounded-xl shadow-2xl p-8">
    <h2
      class="text-2xl font-bold mb-6 text-white border-b border-slate-800 pb-4"
    >
      <i class="fa-solid fa-diagram-project mr-2 text-blue-500"></i>
      Add New Project
    </h2>

    <form class="space-y-5">
      <div class="grid grid-cols-1 md:grid-cols-4 gap-4">
        <div class="md:col-span-3">
          <label
            for="title"
            class="block text-sm font-medium mb-2 text-slate-400">Title*</label
          >
          <input
            type="text"
            name="title"
            placeholder="Enter project title"
            required
            class="w-full bg-slate-800 border border-slate-700 rounded-lg px-4 py-2.5 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent transition-all"
          />
        </div>

        <div>
          <label
            for="order"
            class="block text-sm font-medium mb-2 text-slate-400">Order*</label
          >
          <input
            type="number"
            name="order"
            placeholder="0"
            required
            class="w-full bg-slate-800 border border-slate-700 rounded-lg px-4 py-2.5 focus:outline-none focus:ring-2 focus:ring-blue-500 transition-all"
          />
        </div>
      </div>

      <div>
        <label
          for="description"
          class="block text-sm font-medium mb-2 text-slate-400"
          >Description*</label
        >
        <textarea
          name="description"
          rows="3"
          placeholder="Describe the project..."
          required
          class="w-full bg-slate-800 border border-slate-700 rounded-lg px-4 py-2.5 focus:outline-none focus:ring-2 focus:ring-blue-500 transition-all"
        ></textarea>
      </div>

      <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
        <div>
          <label
            for="imageUrl"
            class="block text-sm font-medium mb-2 text-slate-400"
            >Image URL*</label
          >
          <input
            type="text"
            name="image"
            placeholder="https://..."
            required
            class="w-full bg-slate-800 border border-slate-700 rounded-lg px-4 py-2.5 focus:outline-none focus:ring-2 focus:ring-blue-500 transition-all"
          />
        </div>
        <div>
          <label
            for="github"
            class="block text-sm font-medium mb-2 text-slate-400"
            >GitHub URL*</label
          >
          <input
            type="text"
            name="github"
            placeholder="https://github.com/..."
            required
            class="w-full bg-slate-800 border border-slate-700 rounded-lg px-4 py-2.5 focus:outline-none focus:ring-2 focus:ring-blue-500 transition-all"
          />
        </div>
      </div>

      <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
        <div>
          <label
            for="technologies"
            class="block text-sm font-medium mb-2 text-slate-400"
            >Technologies* (Comma separated)</label
          >
          <input
            type="text"
            name="technologies"
            placeholder="Go, SvelteKit, Tailwind"
            required
            class="w-full bg-slate-800 border border-slate-700 rounded-lg px-4 py-2.5 focus:outline-none focus:ring-2 focus:ring-blue-500 transition-all"
          />
        </div>
        <div>
          <label
            for="gradient"
            class="block text-sm font-medium mb-2 text-slate-400"
            >Gradient</label
          >
          <input
            type="text"
            name="gradient"
            placeholder="from-blue-500 to-purple-600"
            class="w-full bg-slate-800 border border-slate-700 rounded-lg px-4 py-2.5 focus:outline-none focus:ring-2 focus:ring-blue-500 transition-all"
          />
        </div>
      </div>

      <div
        class="flex items-center space-x-3 p-4 bg-slate-800/50 rounded-lg border border-slate-700"
      >
        <input
          type="checkbox"
          name="isFeatured"
          id="featured"
          class="w-5 h-5 rounded border-slate-700 bg-slate-900 text-blue-500 focus:ring-blue-500 focus:ring-offset-slate-900"
        />
        <label for="featured" class="text-sm font-medium text-slate-300"
          >Mark as Featured Project</label
        >
      </div>

      <div class="flex items-start gap-2 text-sm text-slate-400 mb-4">
        <span class="text-blue-400">ℹ</span>
        <p>
          Fields marked with <span class="text-red-500 font-medium">*</span> are
          required. Please complete them before submitting.
        </p>
      </div>

      <div class="flex justify-end space-x-4 pt-4">
        <button
          type="button"
          onclick={onCancel}
          class="px-6 py-2.5 rounded-lg text-sm font-semibold text-slate-400 hover:text-white hover:bg-slate-800 transition-all"
          ><i class="fa-solid fa-xmark text-xl"></i>
          Cancel
        </button>
        <button
          type="submit"
          class="px-6 py-2.5 bg-blue-600 hover:bg-blue-500 text-white rounded-lg text-sm font-semibold shadow-lg shadow-blue-900/20 transition-all"
          ><i class="fa-solid fa-paper-plane"></i>
          Create Project
        </button>
      </div>
    </form>
  </div>
</div>
