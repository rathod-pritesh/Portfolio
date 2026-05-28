<script>
  import { tick } from "svelte";
  import { knowledgeBase } from "$lib/data/knowledgeBase";
  import Fuse from "fuse.js";

  let open = false;
  let message = "";
  let loading = false;

  let inputEl;
  let chatContainer;

  let messages = [
    {
      from: "bot",
      text: "Hi, I'm Pritesh's portfolio assistant. Ask me about projects, skills, or experience."
    }
  ];

  // to ignore during keyword matching
  const STOPWORDS = new Set([
    "pritesh", "rathod", "what", "which", "who", "how", "is", "are",
    "the", "his", "your", "you", "me", "my", "do", "did", "does",
    "have", "has", "been", "can", "about", "tell", "give", "and",
    "or", "of", "in", "a", "an", "for", "to", "any", "not",
  ]);

  const flatKB = knowledgeBase.map((item) => ({
    ...item,
    tagsFlat: Array.isArray(item.tags) ? item.tags.join(" ") : "",
  }));

  const fuse = new Fuse(flatKB, {
    keys: [
      {
        name: "tagsFlat",
        weight: 0.6
      },
      {
        name: "question",
        weight: 0.3
      },
      {
        name: "answer",
        weight: 0.1
      }
    ],
    includeScore: true,
    threshold: 0.55,
    ignoreLocation: true,
    minMatchCharLength: 2,
    distance: 200
  });

  function preprocessText(text) {
    return text
      .toLowerCase()
      .replace(/[.,\/#!$%\^&\*;:?{}=\-_`~()]/g, "")
      .trim();
  }

  // checks if any word in the user query directly matches any tag.
  function keywordMatch(cleanedInput) {
    const words = cleanedInput
      .split(/\s+/)
      .filter((w) => w.length > 2 && !STOPWORDS.has(w));

    if (words.length === 0) return null;

    let bestItem = null;
    let bestCount = 0;

    for (const item of knowledgeBase) {
      const tags = Array.isArray(item.tags) ? item.tags : [];
      let score = 0;

      for (const word of words) {
        for (const tag of tags) {
          if (tag === word) {
            score += 2;
          } else if (tag.includes(word) || word.includes(tag)) {
            score += 1;
          }
        }
      }

      if (score > bestCount) {
        bestCount = score;
        bestItem = item;
      }
    }

    return bestCount >= 2 ? bestItem : null;
  }

  function findBestMatch(userInput) {
    const cleanedInput = preprocessText(userInput);

    const keywordResult = keywordMatch(cleanedInput);
    if (keywordResult) {
      return keywordResult.answer;
    }

    const results = fuse.search(cleanedInput);

    if (results.length > 0) {
      if (results.length > 0 && results[0].score < 0.6) {
        return results[0].item.answer;
      }
    }

    return "I'm not sure about that one. Try asking about Pritesh's skills, projects, education, certifications, or GitHub!";
  }

  async function sendMessage() {
    if (!message.trim()) return;

    const userMessage = message;

    messages = [
      ...messages,
      {
        from: "user",
        text: userMessage
      }
    ];

    message = "";
    loading = true;

    await tick();

    scrollToBottom();

    setTimeout(async () => {
      const botReply = findBestMatch(userMessage);

      messages = [
        ...messages,
        {
          from: "bot",
          text: botReply
        }
      ];

      loading = false;

      await tick();

      scrollToBottom();
    }, 500);
  }

  function scrollToBottom() {
    if (chatContainer) {
      chatContainer.scrollTop = chatContainer.scrollHeight;
    }
  }
</script>

<!-- Toggle button -->
<button
  on:click={() => (open = !open)}
  class="fixed bottom-6 right-6 z-50
         w-14 h-14 rounded-full
         bg-linear-to-r from-slate-200 to-slate-500
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
           w-80 sm:w-96 h-105
           bg-darker border border-gray-700
           rounded-2xl shadow-2xl flex flex-col overflow-hidden
           animate-slideUp"
  >
    <!-- Header-->
    <div
      class="px-4 py-3 bg-linear-to-r from-slate-200 to-slate-500 flex justify-between items-center"
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
          class="flex" 
          class:justify-end={msg.from === 'user'}
          class:justify-start={msg.from !== 'user'}
        >
          <div
            class="px-3 py-2 rounded-xl max-w-[80%] leading-relaxed"
            class:bg-slate-200={msg.from === 'user'}
            class:text-slate-900={msg.from === 'user'}
            class:bg-gray-800={msg.from !== 'user'}
            class:text-gray-200={msg.from !== 'user'}
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
