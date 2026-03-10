<script>
  import { browser } from "$app/environment";
  import { onMount } from "svelte";

  export let automation;

  let expanded = false;
  let ready = false;

  function toggle() {
    expanded = !expanded;
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
</script>

<div
  class="border border-gray-700 rounded-2xl overflow-hidden bg-gray-900/70 backdrop-blur"
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
        class="text-xl md:text-2xl font-bold bg-gradient-to-r from-white to-gray-300 bg-clip-text text-transparent leading-snug"
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
      class="self-start md:self-auto px-4 py-2 text-sm rounded-lg bg-gray-800 hover:bg-gray-700 transition text-white"
    >
      {expanded ? "Hide Flow ▲" : "Show Flow ▼"}
    </button>
  </div>

  {#if expanded}
    <div class="bg-[#0a0a0f] p-4 overflow-hidden">
      <div class="min-w-[500px] h-[280px]">
        {#if browser && ready}
          <n8n-demo
            workflow={toWorkflow(automation)}
            tidyup="true"
            style="width:100%;height:100%;display:block;"
          ></n8n-demo>
        {:else}
          <div class="flex items-center justify-center h-full gap-2">
            <div
              class="w-2 h-2 bg-purple-500 rounded-full animate-bounce"
            ></div>
            <div
              class="w-2 h-2 bg-purple-500 rounded-full animate-bounce delay-150"
            ></div>
            <div
              class="w-2 h-2 bg-purple-500 rounded-full animate-bounce delay-300"
            ></div>
          </div>
        {/if}
      </div>
    </div>
  {/if}
</div>
