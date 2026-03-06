<script>
  import { tick } from "svelte";

  let open = false;
  let message = "";
  let messages = [
    {
      from: "bot",
      text: "Hi 👋 Ask me about Pritesh's skills, projects, or education.",
    },
  ];
  let chatContainer;
  let loading = false;
  let inputEl;

  async function getAnswer(userInput) {
    try {
      const res = await fetch("/n8n/webhook/portfolio-reply", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ message: userInput }),
      });

      const text = await res.text();
      return text || "Sorry 😅 I couldn't get a response. Try again.";
    } catch (err) {
      console.error("FULL ERROR:", err);
      return "Error: " + err.message;
    }
  }

  async function sendMessage() {
    if (!message.trim() || loading) return;
    const userText = message;
    messages = [...messages, { from: "user", text: userText }];
    message = "";
    loading = true;
    await tick();
    chatContainer.scrollTop = chatContainer.scrollHeight;

    const reply = await getAnswer(userText);
    messages = [...messages, { from: "bot", text: reply }];
    loading = false;
    await tick();
    chatContainer.scrollTop = chatContainer.scrollHeight;
    inputEl.focus()
  }
</script>

<!-- Toggle button -->
<button
  on:click={() => (open = !open)}
  class="fixed bottom-6 right-6 z-50
         w-14 h-14 rounded-full
         bg-gradient-to-r from-slate-200 to-slate-500
         flex items-center justify-center
         shadow-xl hover:scale-110 transition-transform duration-200"
  aria-label="Open Chat"
>
  <i class="fa-solid fa-robot text-slate-900 text-xl"></i>
</button>

<!-- Chat window -->
{#if open}
  <div
    class="fixed bottom-24 right-6 z-50
           w-80 sm:w-96 h-[420px]
           bg-[#020617] border border-gray-700
           rounded-2xl shadow-2xl flex flex-col overflow-hidden
           animate-slideUp"
  >
    <!-- Header-->
    <div
      class="px-4 py-3 bg-gradient-to-r from-slate-200 to-slate-500 flex justify-between items-center"
    >
      <span class="text-slate-900 font-semibold">Ask About Pritesh</span>
      <button
        on:click={() => (open = false)}
        aria-label="Close"
        class="text-slate-800 hover:text-slate-900 transition-colors"
      >
        <i class="fa-solid fa-xmark text-lg"></i>
      </button>
    </div>

    <!-- Messages -->
    <div
      class="flex-1 p-4 space-y-3 overflow-y-auto text-sm"
      bind:this={chatContainer}
    >
      {#each messages as msg}
        <div
          class="flex {msg.from === 'user' ? 'justify-end' : 'justify-start'}"
        >
          <div
            class="px-3 py-2 rounded-xl max-w-[80%] leading-relaxed
                   {msg.from === 'user'
              ? 'bg-slate-200 text-slate-900'
              : 'bg-gray-800 text-gray-200'}"
          >
            {msg.text}
          </div>
        </div>
      {/each}

      {#if loading}
        <div class="flex justify-start">
          <div
            class="px-3 py-2 rounded-xl bg-gray-800 text-gray-400 text-sm animate-pulse"
          >
            Thinking...
          </div>
        </div>
      {/if}
    </div>

    <!-- Input row -->
    <div class="p-3 border-t border-gray-700 flex gap-2">
      <input
        bind:this={inputEl}
        bind:value={message}
        on:keydown={(e) => e.key === "Enter" && sendMessage()}
        placeholder="Ask something..."
        disabled={loading}
        class="flex-1 bg-gray-900 text-white placeholder-gray-500 rounded-lg px-3 py-2 text-sm
               focus:outline-none focus:ring-1 focus:ring-slate-400 disabled:opacity-50"
      />
      <button
        on:click={sendMessage}
        disabled={loading}
        class="bg-slate-200 text-slate-900 font-semibold px-4 rounded-lg
               hover:bg-slate-300 transition-colors duration-200 disabled:opacity-50"
      >
        Send
      </button>
    </div>
  </div>
{/if}
