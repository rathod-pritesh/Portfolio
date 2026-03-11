<script>
	import { onMount } from 'svelte';

	let canvas;
	let ctx;
	let particles = [];
	let animationFrameId;

	function getColors() {
		const isLight = document.body.classList.contains('light');
		return {
			particle: isLight ? `rgba(30, 77, 123, 0.4)` : `rgba(226, 232, 240, 0.5)`,
			line: isLight ? `15, 45, 82` : `148, 163, 184`
		}
	}

	class Particle {
		constructor(width, height) {
			this.x = Math.random() * width;
			this.y = Math.random() * height;
			this.vx = (Math.random() - 0.5) * 0.5;
			this.vy = (Math.random() - 0.5) * 0.5;
			this.radius = Math.random() * 2 + 1;
		}

		update(width, height) {
			this.x += this.vx;
			this.y += this.vy;
			if (this.x < 0 || this.x > width)  this.vx = -this.vx;
			if (this.y < 0 || this.y > height)  this.vy = -this.vy;
		}

		draw(ctx) {
			ctx.beginPath();
			ctx.arc(this.x, this.y, this.radius, 0, Math.PI * 2);
			ctx.fillStyle = getColors().particle;
			ctx.fill();
		}
	}

	function initParticles() {
		if (!canvas) return;
		const width  = window.innerWidth;
		const height = window.innerHeight;
		canvas.width  = width;
		canvas.height = height;
		particles = [];
		const count = Math.floor((width * height) / 10000);
		for (let i = 0; i < count; i++) particles.push(new Particle(width, height));
	}

	function drawLines() {
		const maxDist = 150;
		const {line} = getColors();
		for (let i = 0; i < particles.length; i++) {
			for (let j = i + 1; j < particles.length; j++) {
				const dx   = particles[i].x - particles[j].x;
				const dy   = particles[i].y - particles[j].y;
				const dist = Math.sqrt(dx * dx + dy * dy);
				if (dist < maxDist) {
					const opacity = (1 - dist / maxDist) * 0.25;
					ctx.beginPath();
					ctx.strokeStyle = `rgba(${line}, ${opacity})`;
					ctx.lineWidth = 1;
					ctx.moveTo(particles[i].x, particles[i].y);
					ctx.lineTo(particles[j].x, particles[j].y);
					ctx.stroke();
				}
			}
		}
	}

	function animate() {
		if (!ctx) return;
		ctx.clearRect(0, 0, canvas.width, canvas.height);
		particles.forEach(p => { p.update(canvas.width, canvas.height); p.draw(ctx); });
		drawLines();
		animationFrameId = requestAnimationFrame(animate);
	}

	onMount(() => {
		if (canvas) {
			ctx = canvas.getContext('2d');
			initParticles();
			animate();
		}
		const onResize = () => initParticles();
		window.addEventListener('resize', onResize);
		return () => {
			window.removeEventListener('resize', onResize);
			if (animationFrameId) cancelAnimationFrame(animationFrameId);
		};
	});
</script>

<div class="fixed inset-0 z-0 pointer-events-none">
	<canvas bind:this={canvas} class="w-full h-full"></canvas>
</div>