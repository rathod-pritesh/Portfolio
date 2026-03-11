<script>
  import { onMount } from "svelte";

  let stats = {
    home: { loaded: false, name: "" },
    about: { loaded: false, title: "" },
    skills: { loaded: false, count: 0 },
    projects: { loaded: false, count: 0 },
  };

  const cards = [
    {
      label: "Home Section",
      href: "/admin/home",
      icon: "fa-house",
      color: "blue",
    },
    {
      label: "About Section",
      href: "/admin/about",
      icon: "fa-user",
      color: "violet",
    },
    {
      label: "Skills Section",
      href: "/admin/skills",
      icon: "fa-code",
      color: "green",
    },
    {
      label: "Projects",
      href: "/admin/projects",
      icon: "fa-folder-open",
      color: "amber",
    },
    {
      label: "Contact Section",
      href: "/admin/contact",
      icon: "fa-envelope",
      color: "rose",
    },
  ];

  const colorMap = {
    blue: "bg-blue-500/10 border-blue-500/20 text-blue-400",
    violet: "bg-violet-500/10 border-violet-500/20 text-violet-400",
    green: "bg-green-500/10 border-green-500/20 text-green-400",
    amber: "bg-amber-500/10 border-amber-500/20 text-amber-400",
    rose: "bg-rose-500/10 border-rose-500/20 text-rose-400",
  };

  onMount(async () => {
    const base = "http://127.0.0.1:8080/api";
    const opts = { credentials: "include" };

    try {
      const [homeRes, aboutRes, skillsRes, projectsRes] = await Promise.all([
        fetch(`${base}/home`, opts),
        fetch(`${base}/about`, opts),
        fetch(`${base}/skills`, opts),
        fetch(`${base}/projects`, opts),
      ]);

      const [home, about, skills, projects] = await Promise.all([
        homeRes.json(),
        aboutRes.json(),
        skillsRes.json(),
        projectsRes.json(),
      ]);

      stats.home = { loaded: true, name: home.name ?? "—" };
      stats.about = { loaded: true, name: about.name ?? "—" };
      stats.skills = { loaded: true, count: skills.categories?.length ?? 0 };
      stats.projects = {
        loaded: true,
        count: Array.isArray(projects) ? projects.length : 0,
      };
    } catch {
      alert("Error while loading page!!!");
    }
  });
</script>

<svelte:head><title>Dashboard | Admin</title></svelte:head>

<div class="space-y-8 max-w-5xl">
  <!-- Welcome -->
  <div>
    <h2 class="text-2xl font-bold text-white">Welcome Back</h2>
    <p class="text-gray-500 text-sm mt-1">Manage your portfolio content</p>
  </div>

  <!-- Quick start -->
  <div class="grid grid-cols-2 lg:grid-cols-4 gap-4">
    <div class="bg-dark border border-gray-800 rounded-xl p-4">
      <p class="text-xs text-gray-500 uppercase tracking-wide mb-1">Name</p>
      <p class="text-primary font-semibold truncate">
        {stats.home.loaded ? stats.home.name : "…"}
      </p>
    </div>
    <div class="bg-dark border border-gray-800 rounded-xl p-4">
      <p class="text-xs text-gray-500 uppercase tracking-wide mb-1">
        About Title
      </p>
      <p class="text-primary font-semibold truncate">
        {stats.about.loaded ? stats.about.title : "…"}
      </p>
    </div>

    <div class="bg-dark border border-gray-800 rounded-xl p-4">
      <p class="text-xs text-gray-500 uppercase tracking-wide mb-1">
        Skill Categories
      </p>
      <p class="text-primary font-semibold">
        {stats.skills.loaded ? stats.skills.count : "…"}
      </p>
    </div>
    <div class="bg-dark border border-gray-800 rounded-xl p-4">
      <p class="text-xs text-gray-500 uppercase tracking-wide mb-1">Projects</p>
      <p class="text-primary font-semibold">
        {stats.projects.loaded ? stats.projects.count : "…"}
      </p>
    </div>
  </div>

  <!-- Section cards -->
  <div>
    <h3
      class="text-sm font-semibold text-gray-400 uppercase tracking-wider mb-4"
    >
      Manage Sections
    </h3>
    <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4">
      {#each cards as card}
        <a
          href={card.href}
          class="group bg-dark border border-gray-800 hover:border-gray-700 rounded-xl p-5
                 flex items-center gap-4 transition-all duration-200
                 hover:shadow-lg hover:shadow-black/20 hover:-translate-y-0.5"
        >
          <div class="w-11 h-11 rounded-xl border flex items-center justify-center shrink-0
                      {colorMap[card.color]}">
            <i class="fa-solid {card.icon}"></i>
          </div>
          <div class="min-w-0">
            <p class="text-sm font-semibold text-white group-hover:text-primary transition-colors">
              {card.label}
            </p>
            <p class="text-xs text-gray-600 mt-0.5">Click to edit</p>
          </div>
          <i class="fa-solid fa-chevron-right text-gray-700 group-hover:text-gray-500 ml-auto text-xs transition-colors"></i>
        </a>
      {/each}
    </div>
  </div>
</div>
