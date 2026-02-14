<script>
	import Section from './Section.svelte';

	let formData = {
		name: '',
		email: '',
		message: ''
	};

	let isSubmitting = false;
	let submitStatus = '';

	function handleSubmit(e) {
		e.preventDefault();
		isSubmitting = true;

		// Simulate form submission
		setTimeout(() => {
			submitStatus = 'success';
			isSubmitting = false;
			formData = { name: '', email: '', message: '' };

			setTimeout(() => {
				submitStatus = '';
			}, 3000);
		}, 1500);
	}

	const socialLinks = [
		{
			name: 'GitHub',
			url: 'https://github.com/rathod-pritesh',
			icon: '/icons/social/GitHub.png'
		},
		{
			name: 'LinkedIn',
			url: 'https://www.linkedin.com/in/rathodpritesh/',
			icon: '/icons/social/LinkedIn.png'
		},
	];
</script>

<Section id="contact">
	<div class="space-y-12">
		<div class="text-center space-y-4">
			<h2 class="text-4xl md:text-5xl font-bold bg-gradient-to-r from-primary to-secondary gradient-text text-primary underline decoration-[#6366f1]">
				Get In Touch
			</h2>
			<div class="w-20 h-1 bg-gradient-to-r from-primary to-secondary mx-auto"></div>
			<p class="text-gray-400 text-lg max-w-2xl mx-auto">
				Have a project in mind or just want to say hello? Feel free to reach out!
			</p>
		</div>

		<div class="grid md:grid-cols-2 gap-12 max-w-6xl mx-auto">
			<!-- Contact Form -->
			<div class="bg-dark/50 border border-gray-800 rounded-xl p-8">
				<form on:submit={handleSubmit} class="space-y-6">
					<div>
						<label for="name" class="block text-sm font-medium text-gray-300 mb-2">
							Name
						</label>
						<input
							type="text"
							id="name"
							bind:value={formData.name}
							required
							class="w-full px-4 py-3 bg-darker border border-gray-700 rounded-lg focus:outline-none focus:border-primary transition-colors text-white"
							placeholder="Your Name"
						/>
					</div>
					<div>
						<label for="email" class="block text-sm font-medium text-gray-300 mb-2">
							Email
						</label>
						<input
							type="email"
							id="email"
							bind:value={formData.email}
							required
							class="w-full px-4 py-3 bg-darker border border-gray-700 rounded-lg focus:outline-none focus:border-primary transition-colors text-white"
							placeholder="your@email.com"
						/>
					</div>
					<div>
						<label for="message" class="block text-sm font-medium text-gray-300 mb-2">
							Message
						</label>
						<textarea
							id="message"
							bind:value={formData.message}
							required
							rows="5"
							class="w-full px-4 py-3 bg-darker border border-gray-700 rounded-lg focus:outline-none focus:border-primary transition-colors text-white resize-none"
							placeholder="Your message here..."
						></textarea>
					</div>
					<button
						type="submit"
						disabled={isSubmitting}
						class="w-full px-8 py-3 bg-gradient-to-r from-primary to-secondary rounded-lg font-semibold hover:shadow-lg hover:shadow-primary/50 transition-all duration-300 hover:scale-105 disabled:opacity-50 border border-white disabled:cursor-not-allowed disabled:hover:scale-100"
					>
						{isSubmitting ? 'Sending...' : 'Send Message'}
					</button>
					{#if submitStatus === 'success'}
						<p class="text-green-400 text-center animate-pulse">Message sent successfully!</p>
					{/if}
				</form>
			</div>

			<!-- Contact Info -->
			<div class="space-y-8">
				<div class="bg-dark/50 border border-gray-800 rounded-xl p-8 space-y-6">
					<h3 class="text-2xl font-bold text-white mb-4">Contact Information</h3>
					<div class="space-y-4">
						<div class="flex items-center gap-4">
							<div class="w-12 h-12 bg-primary/20 rounded-lg flex items-center justify-center">
								<i class="fa-solid fa-envelope"></i>
							</div>
							<div>
								<p class="text-gray-400 text-sm">Email</p>
								<a href="mailto:alex@example.com" class="text-white hover:text-primary transition-colors">rathodpritesh0712@gmail.com</a>
							</div>
						</div>
						<div class="flex items-center gap-4">
							<div class="w-12 h-12 bg-primary/20 rounded-lg flex items-center justify-center">
								<i class="fa-solid fa-location-dot"></i>
							</div>
							<div>
								<p class="text-gray-400 text-sm">Location</p>
								<p class="text-white">Ahmedabad, India</p>
							</div>
						</div>
					</div>
				</div>

				<div class="bg-dark/50 border border-gray-800 rounded-xl p-8">
					<h3 class="text-2xl font-bold text-white mb-6">Follow Me</h3>
					<div class="flex gap-4">
						{#each socialLinks as social}
							<a
								href={social.url}
								target="_blank"
								rel="noopener noreferrer"
								class="w-12 h-12 bg-primary/20 rounded-lg flex items-center justify-center hover:bg-primary transition-all duration-300 hover:scale-110 group"
								aria-label={social.name}
							>
								<img 
									src={social.icon} 
									alt={social.name} 
									class="w-6 h-6 object-contain group-hover:scale-110 transition-transform duration-300">
							</a>
						{/each}
					</div>
				</div>
			</div>
		</div>
	</div>
</Section>
