<script>
	import Section from './Section.svelte';

	// Form state
	let formData     = { name: '', email: '', message: '' };
	let isSubmitting = false;
	let submitStatus = '';

	function handleSubmit(e) {
		e.preventDefault();
		isSubmitting = true;
		setTimeout(() => {
			submitStatus = 'success';
			isSubmitting = false;
			formData = { name: '', email: '', message: '' };
			setTimeout(() => { submitStatus = ''; }, 3000);
		}, 1500);
	}

	const socialLinks = [
		{ name: 'GitHub',   url: 'https://github.com/rathod-pritesh',              icon: '/icons/social/GitHub.png'   },
		{ name: 'LinkedIn', url: 'https://www.linkedin.com/in/rathodpritesh/',     icon: '/icons/social/LinkedIn.png' },
	];
</script>

<Section id="contact">
	<div class="space-y-12">

		<!-- Heading -->
		<div class="text-center space-y-4">
			<h2 class="text-4xl md:text-5xl font-bold bg-gradient-to-r from-primary to-secondary gradient-text underline decoration-[#94a3b8]">
				Get In Touch
			</h2>
			<div class="w-20 h-1 bg-gradient-to-r from-primary to-secondary mx-auto rounded-full"></div>
			<p class="text-gray-400 text-lg max-w-2xl mx-auto">
				Have a project in mind or just want to say hello? Feel free to reach out!
			</p>
		</div>

		<div class="grid md:grid-cols-2 gap-12 max-w-6xl mx-auto">

			<!-- ── Contact Form ── -->
			<div class="bg-dark/50 border border-gray-800 rounded-xl p-8">
				<form on:submit={handleSubmit} class="space-y-6">

					<div>
						<label for="name" class="block text-sm font-medium text-gray-300 mb-2">Name</label>
						<input
							type="text" id="name"
							bind:value={formData.name}
							required
							placeholder="Your Name"
							class="w-full px-4 py-3 bg-darker border border-gray-700 rounded-lg
							       text-white placeholder-gray-500
							       focus:outline-none focus:border-primary transition-colors"
						/>
					</div>

					<div>
						<label for="email" class="block text-sm font-medium text-gray-300 mb-2">Email</label>
						<input
							type="email" id="email"
							bind:value={formData.email}
							required
							placeholder="your@email.com"
							class="w-full px-4 py-3 bg-darker border border-gray-700 rounded-lg
							       text-white placeholder-gray-500
							       focus:outline-none focus:border-primary transition-colors"
						/>
					</div>

					<div>
						<label for="message" class="block text-sm font-medium text-gray-300 mb-2">Message</label>
						<textarea
							id="message"
							bind:value={formData.message}
							required rows="5"
							placeholder="Your message here..."
							class="w-full px-4 py-3 bg-darker border border-gray-700 rounded-lg
							       text-white placeholder-gray-500 resize-none
							       focus:outline-none focus:border-primary transition-colors"
						></textarea>
					</div>

					<button
						type="submit"
						disabled={isSubmitting}
						class="w-full px-8 py-3 bg-gradient-to-r from-primary to-secondary
						       text-slate-900 font-semibold rounded-lg
						       hover:shadow-lg hover:shadow-primary/25 hover:scale-105
						       transition-all duration-300
						       disabled:opacity-50 disabled:cursor-not-allowed disabled:hover:scale-100"
					>
						{isSubmitting ? 'Sending...' : 'Send Message'}
					</button>

					{#if submitStatus === 'success'}
						<p class="text-green-400 text-center animate-pulse">✓ Message sent successfully!</p>
					{/if}

				</form>
			</div>

			<div class="space-y-8">

				<div class="bg-dark/50 border border-gray-800 rounded-xl p-8 space-y-6">
					<h3 class="text-2xl font-bold text-white">Contact Information</h3>
					<div class="space-y-5">

						<div class="flex items-center gap-4">
							<div class="w-12 h-12 bg-primary/10 border border-primary/20 rounded-lg flex items-center justify-center text-primary shrink-0">
								<i class="fa-solid fa-envelope"></i>
							</div>
							<div>
								<p class="text-gray-500 text-xs uppercase tracking-wider mb-0.5">Email</p>
								<a href="mailto:rathodpritesh0712@gmail.com" class="text-gray-200 hover:text-primary transition-colors text-sm">
									rathodpritesh0712@gmail.com
								</a>
							</div>
						</div>

						<div class="flex items-center gap-4">
							<div class="w-12 h-12 bg-primary/10 border border-primary/20 rounded-lg flex items-center justify-center text-primary shrink-0">
								<i class="fa-solid fa-location-dot"></i>
							</div>
							<div>
								<p class="text-gray-500 text-xs uppercase tracking-wider mb-0.5">Location</p>
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
									class="w-6 h-6 object-contain
									       transition-all duration-300
									       group-hover:brightness-0 group-hover:invert"
								/>
							</a>
						{/each}
					</div>
				</div>

			</div>
		</div>

	</div>
</Section>