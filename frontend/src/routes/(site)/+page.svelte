<script>
	import ParticlesBackground from '$lib/components/ParticlesBackground.svelte';
	import Home from '$lib/components/Home.svelte';
	import About from '$lib/components/About.svelte';
	import Skills from '$lib/components/Skills.svelte';
	import Projects from '$lib/components/Projects.svelte';
	import Contact from '$lib/components/Contact.svelte';
	import Certification from '$lib/components/Certification.svelte';

	export let data;
</script>

<svelte:head>
	<title>Pritesh Rathod | Aspiring AI/ML Engineer</title>
	<meta name="description" content="Portfolio of Pritesh Rathod - Aspiring AI/ML Engineer" />
</svelte:head>

<ParticlesBackground />

<div class="relative z-10">
	<Home home={data.home}/>
	<About about={data.about} education={data.education}/>
	<Skills skills={data.skills} automations={data.automations} />
	<Certification certifications={data.certifications} />
	<Projects projects={data.projects}/>
	<Contact />
</div>
