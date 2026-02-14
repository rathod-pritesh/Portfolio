<script>
	import ParticlesBackground from '$lib/components/ParticlesBackground.svelte';
	import Home from '$lib/components/Home.svelte';
	import About from '$lib/components/About.svelte';
	import Skills from '$lib/components/Skills.svelte';
	import Projects from '$lib/components/Projects.svelte';
	import Contact from '$lib/components/Contact.svelte';

	export let data;
</script>

<svelte:head>
	<title>Pritesh Rathod | Aspiring Full-Stack Developer</title>
	<meta name="description" content="Portfolio of Pritesh Rathod - Aspiring Full-Stack Developer" />
</svelte:head>

<ParticlesBackground />

<div class="relative z-10">
	<Home home={data.home}/>
	<About about={data.about}/>
	<Skills skills={data.skills}/>
	<Projects projects={data.projects}/>
	<Contact />
</div>
