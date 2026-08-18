<script>
	import { onMount } from 'svelte';
	import * as THREE from 'three';

	let container;

	const technologies = [
		{ src: '/tech/Python.png',  pos: [-14,  7,   0],  scale: 1.9 },
		{ src: '/tech/LangChain.png',  pos: [ 18,  9,  -22], scale: 1.0 },
		{ src: '/tech/Go.png',      pos: [  6, -4,  -10], scale: 1.5 },
		{ src: '/tech/FastAPI.png', pos: [-20, -9,  -6],  scale: 1.3 },
		{ src: '/tech/MongoDB.png', pos: [  2,  14, -28], scale: 0.95 },
		{ src: '/tech/MySQL.png',   pos: [ 22,  -5, -14], scale: 1.1 },
		{ src: '/tech/n8n.png',     pos: [ 14, -16, -20], scale: 1.0 },
		{ src: '/tech/Svelte.png',  pos: [ -6, -18, -32], scale: 0.9 }
	];

	onMount(() => {
		const scene = new THREE.Scene();
		const camera = new THREE.PerspectiveCamera(50, window.innerWidth / window.innerHeight, 0.1, 1000);
		camera.position.z = 35;

		const renderer = new THREE.WebGLRenderer({ alpha: true, antialias: true });
		renderer.setPixelRatio(Math.min(window.devicePixelRatio, 2));
		renderer.setSize(window.innerWidth, window.innerHeight);
		container.appendChild(renderer.domElement);

		const isLight = () => document.documentElement.getAttribute('data-theme') === 'light';

		const LIGHT = {
			tile:    0xe8ecf2,
			ambient: { color: 0xffffff, intensity: 1.6 },
			key:     { color: 0xffffff, intensity: 3.0 },
			fill:    { color: 0xd0dcff, intensity: 1.0 },
			rim:     { color: 0xffeedd, intensity: 0.5 },
		};

		const DARK = {
			tile:    0x26231e,
			ambient: { color: 0xffffff, intensity: 1.0 },
			key:     { color: 0x7eb8ff, intensity: 2.5 },
			fill:    { color: 0x4466aa, intensity: 1.2 },
			rim:     { color: 0x3355cc, intensity: 0.8 },
		};

		const ambientLight = new THREE.AmbientLight(0xffffff, 1.6);
		scene.add(ambientLight);
		const keyLight = new THREE.DirectionalLight(0xffffff, 3.0);
		keyLight.position.set(8, 12, 18);
		scene.add(keyLight);
		const fillLight = new THREE.DirectionalLight(0xd0dcff, 1.0);
		fillLight.position.set(-12, -6, 10);
		scene.add(fillLight);
		const rimLight = new THREE.DirectionalLight(0xffeedd, 0.5);
		rimLight.position.set(0, -10, -5);
		scene.add(rimLight);

		function applyTheme(tileMaterials) {
			const t = isLight() ? LIGHT : DARK;
			ambientLight.color.setHex(t.ambient.color);
			ambientLight.intensity = t.ambient.intensity;
			keyLight.color.setHex(t.key.color);
			keyLight.intensity = t.key.intensity;
			fillLight.color.setHex(t.fill.color);
			fillLight.intensity = t.fill.intensity;
			rimLight.color.setHex(t.rim.color);
			rimLight.intensity = t.rim.intensity;
			tileMaterials.forEach(mat => {
				mat.color.setHex(t.tile);
				mat.needsUpdate = true;
			});
		}

		function createRoundedRectShape(w, h, r) {
			const shape = new THREE.Shape();
			const x = -w / 2, y = -h / 2;
			shape.moveTo(x + r, y);
			shape.lineTo(x + w - r, y);
			shape.quadraticCurveTo(x + w, y, x + w, y + r);
			shape.lineTo(x + w, y + h - r);
			shape.quadraticCurveTo(x + w, y + h, x + w - r, y + h);
			shape.lineTo(x + r, y + h);
			shape.quadraticCurveTo(x, y + h, x, y + h - r);
			shape.lineTo(x, y + r);
			shape.quadraticCurveTo(x, y, x + r, y);
			return shape;
		}

		// Returns a 0-1 factor: 1 = desktop, shrinks toward 0 on mobile
		function getResponsiveFactor() {
			const w = window.innerWidth;
			if (w >= 1024) return 1.0;       // desktop
			if (w >= 768)  return 0.7;        // tablet
			return 0.45;                       // mobile
		}

		const textureLoader = new THREE.TextureLoader();
		const tiles = [];
		const tileMaterials = [];
		const tileGroups = [];

		technologies.forEach((tech) => {
			const group = new THREE.Group();

			const shape = createRoundedRectShape(4, 5, 0.65);
			const tileGeo = new THREE.ExtrudeGeometry(shape, {
				depth: 0.6,
				bevelEnabled: true,
				bevelSegments: 14,
				bevelSize: 0.14,
				bevelThickness: 0.1
			});

			const tileMat = new THREE.MeshPhysicalMaterial({
				color: 0xe8ecf2,
				roughness: 0.9,
				metalness: 0.02,
				clearcoat: 0.2,
				clearcoatRoughness: 0.6
			});
			tileMaterials.push(tileMat);

			const tile = new THREE.Mesh(tileGeo, tileMat);
			group.add(tile);

			const texture = textureLoader.load(tech.src);
			texture.colorSpace = THREE.SRGBColorSpace;
			const logo = new THREE.Mesh(
				new THREE.PlaneGeometry(2.6, 2.6),
				new THREE.MeshBasicMaterial({ map: texture, transparent: true, alphaTest: 0.05 })
			);
			logo.position.z = 0.72;
			group.add(logo);

			// Apply responsive scaling to position
			const f = getResponsiveFactor();
			group.position.set(tech.pos[0] * f, tech.pos[1] * f, tech.pos[2]);

			// Scale down tiles on small screens too
			const s = (tech.scale || 1) * (window.innerWidth < 768 ? 0.55 : window.innerWidth < 1024 ? 0.75 : 1.0);
			group.scale.set(s, s, s);

			group.rotation.x = (Math.random() - 0.5) * 0.35;
			group.rotation.y = (Math.random() - 0.5) * 0.35;
			scene.add(group);

			tileGroups.push({ group, tech });

			tiles.push({
				mesh:   group,
				phase:  Math.random() * Math.PI * 2,
				speed:  0.22 + Math.random() * 0.13,
				baseY:  tech.pos[1] * f,
				baseRX: (Math.random() - 0.5) * 0.4,
				baseRY: (Math.random() - 0.5) * 0.4
			});
		});

		applyTheme(tileMaterials);

		const observer = new MutationObserver(() => applyTheme(tileMaterials));
		observer.observe(document.documentElement, {
			attributes: true,
			attributeFilter: ['data-theme']
		});

		let mouseX = 0, mouseY = 0;
		const handleMouse = (e) => {
			mouseX = (e.clientX / window.innerWidth  - 0.5) * 2;
			mouseY = (e.clientY / window.innerHeight - 0.5) * 2;
		};
		window.addEventListener('mousemove', handleMouse);

		const handleResize = () => {
			camera.aspect = window.innerWidth / window.innerHeight;
			camera.updateProjectionMatrix();
			renderer.setSize(window.innerWidth, window.innerHeight);

			// Reposition + rescale all tiles on resize
			const f = getResponsiveFactor();
			const tileScale = window.innerWidth < 768 ? 0.55 : window.innerWidth < 1024 ? 0.75 : 1.0;

			tileGroups.forEach(({ group, tech }, i) => {
				group.position.set(tech.pos[0] * f, tech.pos[1] * f, tech.pos[2]);
				const s = (tech.scale || 1) * tileScale;
				group.scale.set(s, s, s);
				// Update baseY for animation
				tiles[i].baseY = tech.pos[1] * f;
			});
		};
		window.addEventListener('resize', handleResize);

		function animate(time) {
			requestAnimationFrame(animate);
			const t = time * 0.001;
			tiles.forEach(({ mesh, phase, speed, baseY, baseRX, baseRY }) => {
				mesh.position.y = baseY + Math.sin(t * speed + phase) * 0.75;
				mesh.rotation.x = baseRX + Math.sin(t * speed       + phase) * 0.07;
				mesh.rotation.y = baseRY + Math.cos(t * speed * 0.9 + phase) * 0.07;
				mesh.rotation.z =          Math.sin(t * speed * 0.5 + phase) * 0.04;
			});
			scene.rotation.y += (mouseX * 0.03  - scene.rotation.y) * 0.03;
			scene.rotation.x += (-mouseY * 0.02 - scene.rotation.x) * 0.03;
			renderer.render(scene, camera);
		}

		animate(0);

		return () => {
			window.removeEventListener('mousemove', handleMouse);
			window.removeEventListener('resize', handleResize);
			observer.disconnect();
			renderer.dispose();
		};
	});
</script>

<div
	bind:this={container}
	class="fixed inset-0 z-0 pointer-events-none overflow-hidden"
></div>