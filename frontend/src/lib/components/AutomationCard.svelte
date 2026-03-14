<script>
  import { browser } from "$app/environment";
  import { onMount } from "svelte";

  export let automation;

  let expanded = false;
  let ready = false;
  let n8nLoading = false;

  function toggle() {
    expanded = !expanded;
    if (expanded) {
      n8nLoading = true;
      setTimeout(() => (n8nLoading = false), 1400);
    }
  }

  onMount(() => {
    if (!browser) return;

    const check = () => {
      if (customElements.get("n8n-demo")) {
        ready = true;
      } else {
        setTimeout(check, 100);
      }
    };

    check();
  });

  function toWorkflow(auto) {
    const nodes = (auto.nodes ?? []).map((n, i) => ({
      name: n.label,
      type: nodeType(n.id),
      position: [200 + i * 220, 300],
      parameters: {},
      typeVersion: 1,
    }));

    const connections = {};

    for (const e of auto.edges ?? []) {
      const src = auto.nodes.find((n) => n.id === e.from);
      const dst = auto.nodes.find((n) => n.id === e.to);

      if (!src || !dst) continue;

      if (!connections[src.label]) {
        connections[src.label] = { main: [[]] };
      }

      connections[src.label].main[0].push({
        node: dst.label,
        type: "main",
        index: 0,
      });
    }

    return JSON.stringify({ nodes, connections });
  }

  function nodeType(id = "") {
    const lower = id.toLowerCase();
    if (lower.includes("webhook")) return "n8n-nodes-base.webhook";
    if (lower.includes("http")) return "n8n-nodes-base.httpRequest";
    if (lower.includes("respond")) return "n8n-nodes-base.respondToWebhook";
    if (lower.includes("gmail")) return "n8n-nodes-base.gmail";
    if (lower.includes("schedule")) return "n8n-nodes-base.scheduleTrigger";
    if (lower.includes("rss")) return "n8n-nodes-base.rssFeedRead";
    if (lower.includes("ai") || lower.includes("openai"))
      return "@n8n/n8n-nodes-langchain.openAi";
    if (lower.includes("slack")) return "n8n-nodes-base.slack";
    if (lower.includes("edit") || lower.includes("set") || lower.includes("fields")) return "n8n-nodes-base.set";
    if (lower.includes("sheet") || lower.includes("google"))
      return "n8n-nodes-base.googleSheets";

    return "n8n-nodes-base.noOp";
  }

  $: skeletonCount = Math.max(3, (automation.nodes ?? []).length);
</script>

<div
  class="automation-card border border-gray-700 rounded-2xl overflow-hidden bg-gray-900/70 backdrop-blur"
>
  <div
    class="flex flex-col md:flex-row md:items-center md:justify-between gap-4 p-6 border-b border-gray-800"
  >
    <div>
      <div
        class="inline-flex items-center gap-1.5 bg-purple-500/10 border border-purple-500/30 rounded-full px-3 py-1 mb-2"
      >
        <span class="text-purple-400 font-bold text-sm">#</span>
        <span class="text-purple-300 font-bold text-sm">{automation.flow}</span>
      </div>

      <h3
        class="text-xl md:text-2xl font-bold bg-linear-to-r from-white to-gray-300 bg-clip-text text-transparent leading-snug"
      >
        {automation.title}
      </h3>

      <p
        class="text-sm md:text-base text-gray-400 mt-2 max-w-xl leading-relaxed"
      >
        {automation.description}
      </p>
    </div>

    <button
      on:click={toggle}
      class="toggle-btn shrink-0 self-start sm:self-center px-4 py-2 text-sm rounded-lg bg-gray-800 hover:bg-gray-700 transition text-white whitespace-nowrap"
    >
      {expanded ? "Hide Flow ▲" : "Show Flow ▼"}
    </button>
  </div>

  {#if expanded}
    <div class="bg-[#0a0a0f] p-3 sm:p-4 overflow-x-auto">
      <div class="relative h-65 sm:h-75" style="min-width: min(100%, 500px);">

        {#if n8nLoading || !ready}
          <div class="skeleton-canvas absolute inset-0 flex items-center px-6 gap-0 overflow-hidden">
            {#each Array(skeletonCount) as _, i}
              <div class="skeleton-node shrink-0" style="animation-delay: {i * 120}ms;">
                <div class="node-icon shimmer"></div>
                <div class="node-label shimmer"></div>
              </div>

              <!-- Arrow -->
              {#if i < skeletonCount - 1}
                <div class="connector" style="animation-delay: { i * 120 + 80 }ms;">
                  <div class="connector-line shimmer"></div>
                  <i class="fa-solid fa-chevron-right arrow-head"></i>
                </div>
              {/if}
            {/each}
          </div>
        {/if}

        <!-- Actual n8n flow -->
        {#if browser && ready}
          <n8n-demo workflow={toWorkflow(automation)} tidyup="true" style="width: 100%;height: 100%;display: block;opacity: {n8nLoading ? 0 : 1};transition: opacity 0.4s ease;"></n8n-demo>
        {/if}

      </div>
    </div>
  {/if}
</div>


<style>
  .skeleton-canvas {
    background: #0a0a0f;
    border-radius: 4px;
  }

  .skeleton-node {
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 8px;
    opacity: 0;
    animation: nodeAppear 0.35s ease forwards;
  }
 
  .node-icon {
    width: 56px;
    height: 56px;
    border-radius: 12px;
  }
 
  .node-label {
    width: 64px;
    height: 10px;
    border-radius: 6px;
  }
 
  .connector {
    display: flex;
    align-items: center;
    flex: 1;
    min-width: 40px;
    max-width: 80px;
    opacity: 0;
    animation: nodeAppear 0.3s ease forwards;
  }
 
  .connector-line {
    flex: 1;
    height: 2px;
    border-radius: 2px;
  }
 
  .arrow-head {
    width: 14px;
    height: 14px;
    flex-shrink: 0;
    margin-left: -2px;
    opacity: 0.6;
  }
 
  .shimmer {
    background: linear-gradient(
      90deg,
      #1e1e2e 25%,
      #2d2b55 50%,
      #1e1e2e 75%
    );
    background-size: 200% 100%;
    animation: shimmer 1.4s infinite linear;
  }
 
  @keyframes shimmer {
    0%   { background-position: 200% 0; }
    100% { background-position: -200% 0; }
  }
 
  @keyframes nodeAppear {
    from { opacity: 0; transform: translateY(8px); }
    to   { opacity: 1; transform: translateY(0); }
  }
</style>