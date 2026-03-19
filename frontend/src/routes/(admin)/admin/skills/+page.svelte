<script>
  import AdminCertification from "$lib/components/admin/AdminCertification.svelte";
  import AdminSkills from "$lib/components/admin/AdminSkills.svelte";

  let activeTab = "skills";

  const tabs = [
    { id: "skills", label: "Tech Skills", icon: "fa-screwdriver-wrench" },
    {
      id: "certifications",
      label: "Certifications",
      icon: "fa-graduation-cap",
    },
  ];
</script>

<svelte:head>
  <title>Admin: Skills & Certifications</title>
</svelte:head>

<div class="min-h-screen bg-[#080b12] text-white px-4 py-10 md:px-10">
  <!-- page header -->
  <div class="mb-8">
    <h1 class="text-2xl md:text-3xl font-extrabold text-white tracking-tight">
      <i class="fa-solid fa-layer-group mr-3 text-primary"></i>Skills &amp;
      Certifications
    </h1>
    <p class="text-gray-500 text-sm mt-1 ml-9">
      Manage tech skill categories, individual skills, and portfolio
      certifications.
    </p>
  </div>

  <!-- tabs -->
  <div class="flex gap-1 mb-8 border-b border-gray-800" role="tablist">
    {#each tabs as tab}
      <button
        type="button"
        role="tab"
        aria-selected={activeTab === tab.id}
        aria-controls="tabpanel-{tab.id}"
        on:click={() => (activeTab = tab.id)}
        class="relative px-5 py-2.5 text-sm font-semibold transition-colors duration-200 {activeTab ===
        tab.id
          ? 'text-white'
          : 'text-gray-500 hover:text-gray-300'}"
      >
        <i class="fa-solid {tab.icon} mr-2"></i>{tab.label}

        <!-- active underline -->
        {#if activeTab === tab.id}
          <span
            class="absolute bottom-0 left-0 right-0 h-0.5
                       bg-gradient-to-r from-primary to-secondary rounded-full"
            aria-hidden="true"
          ></span>
        {/if}
      </button>
    {/each}
  </div>

  <!-- tab panels -->
  <div class="max-w-5xl">
    <div
      class="tabpanel-skills"
      role="tabpanel"
      hidden={activeTab !== "skills"}
    >
      <AdminSkills />
    </div>
    <div
      class="tabpanel-certifications"
      role="tabpanel"
      hidden={activeTab !== "certifications"}
    >
      <AdminCertification />
    </div>
  </div>
</div>
