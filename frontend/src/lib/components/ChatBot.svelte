<script>
  import data from "$lib/data/chatbot-data.json";
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

  function getAnswer(userInput) {
    const input = userInput.toLowerCase();

    for (const item of data) {
      for (const q of item.questions) {
        if (input.includes(q)) {
          return item.answer;
        }
      }
    }
    return "Sorry 😅 I don’t have an answer for that yet. Try asking about skills or projects.";
  }

  async function sendMessage() {
    if (!message.trim()) return;

    const userText = message;

    messages = [...messages, { from: "user", text: userText }];
    message = "";

    await tick();
    chatContainer.scrollTop = chatContainer.scrollHeight;

    setTimeout(async () => {
      const reply = getAnswer(userText);
      messages = [...messages, { from: "bot", text: reply }];

      await tick()
      chatContainer.scrollTop = chatContainer.scrollHeight;
    }, 400);
  }
</script>

<!-- Chat button -->
<button
  on:click={() => (open = !open)}
  class="fixed bottom-6 right-6 z-50
         w-14 h-14 rounded-full
         bg-gradient-to-r from-[#6366f1] to-[#06b6d4]
         flex items-center justify-center
         shadow-xl hover:scale-110 transition"
  aria-label="Open Chat"
>
  <i class="fa-solid fa-robot text-white text-xl"></i>
</button>

<!-- Chat window -->
{#if open}
  <div
    class="fixed bottom-24 right-6 z-50 w-80 sm:w-96 h-[420px] bg-[#0f172a] border border-gray-800
           rounded-2xl shadow-2xl
           flex flex-col overflow-hidden
           animate-slideUp"
  >
    <!-- Header -->
    <div
      class="px-4 py-3 bg-gradient-to-r from-[#6366f1] to-[#06b6d4]
                text-white font-semibold flex justify-between items-center"
    >
      <span>Ask Pritesh</span>
      <button on:click={() => (open = false)} aria-label="Close">
        <i class="fa-solid fa-xmark"></i>
      </button>
    </div>

    <!-- messages -->
    <div class="flex-1 p-4 space-y-3 overflow-y-auto text-sm" bind:this={chatContainer}>
      {#each messages as msg}
        <div
          class={`flex ${msg.from === "user" ? "justify-end" : "justify-start"}`}
        >
          <div
            class={`px-3 py-2 rounded-xl max-w-[80%] ${
              msg.from === "user"
                ? "bg-[#6366f1] text-white"
                : "bg-gray-800 text-gray-200"
            }`}
          >
            {msg.text}
          </div>
        </div>
      {/each}
    </div>

    <!-- Input -->
    <div class="p-3 border-t border-gray-800 flex gap-2">
      <input
        bind:value={message}
        on:keydown={(e) => e.key === "Enter" && sendMessage()}
        placeholder="Ask something..."
        class="flex-1 bg-gray-900 text-white
               rounded-lg px-3 py-2 text-sm
               focus:outline-none focus:ring-1 focus:ring-[#6366f1]"
      />
      <button
        on:click={sendMessage}
        class="bg-[#6366f1] text-white
               px-4 rounded-lg hover:opacity-90"
      >
        Send
      </button>
    </div>
  </div>
{/if}
