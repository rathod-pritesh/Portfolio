<script>
	import { onMount } from 'svelte';
	import { browser } from '$app/environment';
	import { replaceState } from '$app/navigation';

	export let id = '';

	let isVisible = false;

	onMount(() => {
		if (!browser) return;

		const observer = new IntersectionObserver(
			(entries) => {
				entries.forEach((entry) => {
					if (entry.isIntersecting) {
						isVisible = true;

						if (window.location.hash !== `#${id}`) {
							replaceState(`#${id}`)
						}
					}
				});
			},
			{ 
				rootMargin: '-80px 0px -70% 0px',
				threshold: 0 
			}
		);

		const element = document.getElementById(id);
		if (element) observer.observe(element);

		return () => observer.disconnect();
	});
</script>

<section {id} class="section-padding container-max relative z-10">
	<div class="transition-all duration-1000 {isVisible ? 'opacity-100 translate-y-0' : 'opacity-0 translate-y-10'}">
		<slot />
	</div>
</section>
