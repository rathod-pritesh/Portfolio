<script>
	import BackgroundTile from '$lib/components/Background_tile.svelte';
	import Home from '$lib/components/Home.svelte';
	import About from '$lib/components/About.svelte';
	import Skills from '$lib/components/Skills.svelte';
	import Projects from '$lib/components/Projects.svelte';
	import Contact from '$lib/components/Contact.svelte';
	import Certification from '$lib/components/Certification.svelte';
	import Experience from '$lib/components/Experience.svelte';

	export let data;
</script>

<svelte:head>
	<title>Pritesh Rathod | {data.home.role}</title>
	<meta name="description" content="Portfolio of Pritesh Rathod" />
</svelte:head>

<BackgroundTile />

<div class="relative z-10">
	<Home home={data.home}/>
	<About about={data.about} education={data.education}/>
	<Experience experience={data.experience} />
	<Projects projects={data.projects}/>
	<Skills skills={data.skills} automations={data.automations} />
	<Certification certifications={data.certifications} />
	<Contact />
</div>
