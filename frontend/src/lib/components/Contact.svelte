<script>
  import Section from "./Section.svelte";
  import { API_BASE } from "$lib/config/api.js";

  // Form state
  let formData = { name: "", email: "", message: "" };
  let isSubmitting = false;
  let submitStatus = "";

  async function handleSubmit(e) {
    e.preventDefault();
    isSubmitting = true;

    try {
      const response = await fetch(`${API_BASE}/api/contact`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(formData),
      });

      if (response.ok) {
        submitStatus = "success";
        formData = { name: "", email: "", message: "" };
      } else {
        submitStatus = "error";
      }
    } catch (error) {
      submitStatus = "error";
    }

    isSubmitting = false;

    setTimeout(() => {
      submitStatus = "";
    }, 3000);
  }

  const socialLinks = [
    {
      name: "GitHub",
      url: "https://github.com/rathod-pritesh",
      icon: "/icons/social/GitHub.png",
    },
    {
      name: "LinkedIn",
      url: "https://www.linkedin.com/in/rathodpritesh/",
      icon: "/icons/social/LinkedIn.png",
    },
  ];
</script>

<Section id="contact">
  <div class="space-y-12">
    <!-- Heading -->
    <div class="text-center space-y-4">
      <h2
        class="text-4xl md:text-5xl font-bold bg-linear-to-r from-primary to-secondary gradient-text underline decoration-secondary"
      >
        Get In Touch
      </h2>
      <div
        class="w-20 h-1 bg-linear-to-r from-primary to-secondary mx-auto rounded-full"
      ></div>
      <p class="text-gray-400 text-lg max-w-2xl mx-auto">
        Have a project in mind or just want to say hello? Feel free to reach
        out!
      </p>
    </div>

    <div class="grid md:grid-cols-2 gap-12 max-w-6xl mx-auto">
      <!-- ── Contact Form ── -->
      <div class="bg-dark/50 border border-gray-800 rounded-xl p-8">
        <form on:submit={handleSubmit} class="space-y-6">
          <div>
            <label
              for="name"
              class="block text-sm font-medium text-gray-300 mb-2">Name</label
            >
            <input
              type="text"
              id="name"
              bind:value={formData.name}
              required
              placeholder="Your Name"
              class="contact-input w-full px-4 py-3 bg-darker border border-gray-700 rounded-lg placeholder-gray-500
							       focus:outline-none focus:border-primary transition-colors"
            />
          </div>

          <div>
            <label
              for="email"
              class="block text-sm font-medium text-gray-300 mb-2">Email</label
            >
            <input
              type="email"
              id="email"
              bind:value={formData.email}
              required
              placeholder="your@email.com"
              class="contact-input w-full px-4 py-3 bg-darker border border-gray-700 rounded-lg placeholder-gray-500
							       focus:outline-none focus:border-primary transition-colors"
            />
          </div>

          <div>
            <label
              for="message"
              class="block text-sm font-medium text-gray-300 mb-2"
              >Message</label
            >
            <textarea
              id="message"
              bind:value={formData.message}
              required
              rows="5"
              placeholder="Your message here..."
              class="contact-input w-full px-4 py-3 bg-darker border border-gray-700 rounded-lg placeholder-gray-500
							       focus:outline-none focus:border-primary transition-colors"
            ></textarea>
          </div>

          <button
            type="submit"
            disabled={isSubmitting}
            class="w-full px-8 py-3 bg-linear-to-r from-primary to-secondary
						       gradient-btn-text font-semibold rounded-lg
						       hover:shadow-lg hover:shadow-primary/25 hover:scale-105
						       transition-all duration-300
						       disabled:opacity-50 disabled:cursor-not-allowed disabled:hover:scale-100"
          >
            {isSubmitting ? "Sending..." : "Send Message"}
          </button>

          {#if submitStatus === "success"}
            <p class="text-green-400 text-center animate-pulse">
              ✓ Message sent successfully!
            </p>
          {/if}

          {#if submitStatus === "error"}
            <p class="text-red-400 text-center">
              Something went wrong. Please try again.
            </p>
          {/if}
        </form>
      </div>

      <div class="space-y-8">
        <div class="bg-dark/50 border border-gray-800 rounded-xl p-8 space-y-6">
          <h3 class="text-2xl font-bold text-white">Contact Information</h3>
          <div class="space-y-5">
            <div class="flex items-center gap-4">
              <div
                class="w-12 h-12 bg-primary/10 border border-primary/20 rounded-lg flex items-center justify-center text-primary shrink-0"
              >
                <i class="fa-solid fa-envelope"></i>
              </div>
              <div>
                <p
                  class="text-gray-500 text-xs uppercase tracking-wider mb-0.5"
                >
                  Email
                </p>
                <a
                  href="mailto:rathodpritesh0712@gmail.com"
                  class="text-gray-200 hover:text-primary transition-colors text-sm"
                >
                  rathodpritesh0712@gmail.com
                </a>
              </div>
            </div>

            <div class="flex items-center gap-4">
              <div
                class="w-12 h-12 bg-primary/10 border border-primary/20 rounded-lg flex items-center justify-center text-primary shrink-0"
              >
                <i class="fa-solid fa-location-dot"></i>
              </div>
              <div>
                <p
                  class="text-gray-500 text-xs uppercase tracking-wider mb-0.5"
                >
                  Location
                </p>
                <p class="text-gray-200 text-sm">Ahmedabad, India</p>
              </div>
            </div>
          </div>
        </div>

        <!-- Social links -->
        <div class="bg-dark/50 border border-gray-800 rounded-xl p-8">
          <h3 class="text-2xl font-bold text-white mb-6">Follow Me</h3>
          <div class="flex gap-4">
            {#each socialLinks as social}
              <a
                href={social.url}
                target="_blank"
                rel="noopener noreferrer"
                aria-label={social.name}
                class="w-12 h-12 bg-primary/10 border border-primary/20 rounded-lg
								       flex items-center justify-center
								       hover:bg-slate-700 hover:border-slate-500 hover:scale-110
								       transition-all duration-300 group"
              >
                <img
                  src={social.icon}
                  alt={social.name}
                  class={`w-6 h-6 object-contain
          transition-all duration-300
          group-hover:brightness-0 group-hover:invert
          ${social.name === 'GitHub' ? 'github-icon' : ''}`}
                />
              </a>
            {/each}
          </div>
        </div>
      </div>
    </div>
  </div>
</Section>

<style>
  .contact-input {
    color: white !important;
  }

  :global([data-theme="light"]) .contact-input {
    color: black !important;
  }

  .contact-input::placeholder {
    color: #6b7280;
  }

  :global([data-theme="light"]) .contact-input::placeholder {
    color: #9ca3af;
  }

   :global([data-theme="light"]) .github-icon {
    filter: brightness(0);
  }

  :global([data-theme="light"]) .group:hover .github-icon {
    filter: brightness(0) invert(1);
  }
</style>