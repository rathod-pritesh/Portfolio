<script>
	import { onMount } from 'svelte';

	let canvas;

	// Color theme definitions
	const DARK_THEME = {
		isDark: true,
		particleColors: [
			'rgba(122, 139, 111, ', // #7a8b6f sage accent
			'rgba(163, 176, 150, ', // #a3b096 light sage
			'rgba(236, 238, 234, ', // #eceeea light primary
			'rgba(150, 161, 144, ', // #96a190 sage grey secondary
			'rgba(122, 139, 111, '  // #7a8b6f accent
		],
		lineColor: [150, 161, 144], // #96a190 muted sage grey
		mouseLineColor: [122, 139, 111], // #7a8b6f accent
		glowColor1: 'rgba(122, 139, 111, 0.05)',
		glowColor2: 'rgba(163, 176, 150, 0.035)',
		baseAlpha: 0.75,
		maxLineDistance: 130
	};

	const LIGHT_THEME = {
		isDark: false,
		particleColors: [
			'rgba(95, 115, 85, ',   // #5f7355 sage accent
			'rgba(64, 90, 55, ',    // #405a37 deep olive
			'rgba(24, 28, 22, ',    // #181c16 dark primary
			'rgba(95, 107, 88, ',   // #5f6b58 muted olive secondary
			'rgba(95, 115, 85, '    // #5f7355 accent
		],
		lineColor: [95, 107, 88],    // #5f6b58 muted sage grey
		mouseLineColor: [95, 115, 85], // #5f7355 accent
		glowColor1: 'rgba(95, 115, 85, 0.045)',
		glowColor2: 'rgba(64, 90, 55, 0.03)',
		baseAlpha: 0.65,
		maxLineDistance: 125
	};

	function createParticle(w, h, theme, initial = true) {
		const baseSize = 1.2 + Math.random() * 1.8;
		return {
			x: initial ? Math.random() * w : (Math.random() > 0.5 ? 0 : w),
			y: initial ? Math.random() * h : Math.random() * h,
			baseSize,
			size: baseSize,
			vx: (Math.random() - 0.5) * 0.42,
			vy: (Math.random() - 0.5) * 0.42,
			colorIndex: Math.floor(Math.random() * theme.particleColors.length),
			alpha: 0.2 + Math.random() * 0.6,
			baseAlpha: 0.2 + Math.random() * 0.6,
			pulseSpeed: 0.015 + Math.random() * 0.02,
			pulseAngle: Math.random() * Math.PI * 2
		};
	}

	onMount(() => {
		if (!canvas) return;

		const ctx = canvas.getContext('2d');
		let animationFrameId;
		let width = 0;
		let height = 0;
		let dpr = 1;

		// Mouse state
		const mouse = {
			x: -1000,
			y: -1000,
			targetX: -1000,
			targetY: -1000,
			radius: 140,
			active: false
		};

		function getCurrentTheme() {
			const isLight = document.documentElement.getAttribute('data-theme') === 'light';
			return isLight ? LIGHT_THEME : DARK_THEME;
		}

		let theme = getCurrentTheme();
		let particles = [];

		function initParticles() {
			const area = width * height;
			const count = Math.min(110, Math.max(38, Math.floor(area / 16000)));
			particles = [];
			for (let i = 0; i < count; i++) {
				particles.push(createParticle(width, height, theme, true));
			}
		}

		function resize() {
			width = window.innerWidth;
			height = window.innerHeight;
			dpr = Math.min(window.devicePixelRatio || 1, 2);

			canvas.width = width * dpr;
			canvas.height = height * dpr;
			canvas.style.width = `${width}px`;
			canvas.style.height = `${height}px`;

			ctx.setTransform(1, 0, 0, 1, 0, 0);
			ctx.scale(dpr, dpr);

			initParticles();
		}

		function updateTheme() {
			theme = getCurrentTheme();
			for (let p of particles) {
				p.colorIndex = Math.floor(Math.random() * theme.particleColors.length);
			}
		}

		const themeObserver = new MutationObserver(() => {
			updateTheme();
		});

		themeObserver.observe(document.documentElement, {
			attributes: true,
			attributeFilter: ['data-theme', 'class']
		});

		const handleMouseMove = (e) => {
			mouse.targetX = e.clientX;
			mouse.targetY = e.clientY;
			mouse.active = true;
		};

		const handleMouseLeave = () => {
			mouse.active = false;
		};

		const handleTouchMove = (e) => {
			if (e.touches && e.touches.length > 0) {
				mouse.targetX = e.touches[0].clientX;
				mouse.targetY = e.touches[0].clientY;
				mouse.active = true;
			}
		};

		const handleTouchEnd = () => {
			mouse.active = false;
		};

		window.addEventListener('mousemove', handleMouseMove, { passive: true });
		document.addEventListener('mouseleave', handleMouseLeave);
		window.addEventListener('touchmove', handleTouchMove, { passive: true });
		window.addEventListener('touchend', handleTouchEnd);
		window.addEventListener('resize', resize, { passive: true });

		let isRunning = true;
		const handleVisibilityChange = () => {
			if (document.hidden) {
				isRunning = false;
				if (animationFrameId) cancelAnimationFrame(animationFrameId);
			} else {
				isRunning = true;
				loop();
			}
		};
		document.addEventListener('visibilitychange', handleVisibilityChange);

		resize();

		function loop() {
			if (!isRunning) return;

			// Smooth cursor interpolation
			if (mouse.active) {
				mouse.x += (mouse.targetX - mouse.x) * 0.12;
				mouse.y += (mouse.targetY - mouse.y) * 0.12;
			} else {
				mouse.x += (-1000 - mouse.x) * 0.1;
				mouse.y += (-1000 - mouse.y) * 0.1;
			}

			ctx.clearRect(0, 0, width, height);

			// Draw soft ambient spotlight following cursor
			if (mouse.active && mouse.x > -100 && mouse.x < width + 100 && mouse.y > -100 && mouse.y < height + 100) {
				const spotlight = ctx.createRadialGradient(mouse.x, mouse.y, 0, mouse.x, mouse.y, 350);
				spotlight.addColorStop(0, theme.glowColor1);
				spotlight.addColorStop(0.6, theme.glowColor2);
				spotlight.addColorStop(1, 'transparent');
				ctx.fillStyle = spotlight;
				ctx.fillRect(0, 0, width, height);
			}

			// Update & render particles
			const pLen = particles.length;
			for (let i = 0; i < pLen; i++) {
				const p = particles[i];

				// Organic pulse
				p.pulseAngle += p.pulseSpeed;
				p.alpha = p.baseAlpha + Math.sin(p.pulseAngle) * 0.18;

				// Gentle drift
				p.x += p.vx;
				p.y += p.vy;

				// Boundary wrap
				if (p.x < -20) p.x = width + 20;
				else if (p.x > width + 20) p.x = -20;
				if (p.y < -20) p.y = height + 20;
				else if (p.y > height + 20) p.y = -20;

				// Mouse deflection
				if (mouse.active) {
					const dx = mouse.x - p.x;
					const dy = mouse.y - p.y;
					const distance = Math.hypot(dx, dy);

					if (distance < mouse.radius && distance > 0) {
						const force = (1 - distance / mouse.radius) * 1.5;
						const angle = Math.atan2(dy, dx);
						p.x -= Math.cos(angle) * force;
						p.y -= Math.sin(angle) * force;
					}
				}

				// Draw particle
				const col = theme.particleColors[p.colorIndex];
				const a = Math.max(0.05, Math.min(1, p.alpha * theme.baseAlpha));

				ctx.beginPath();
				ctx.arc(p.x, p.y, p.size, 0, Math.PI * 2);
				ctx.fillStyle = `${col}${a})`;
				ctx.fill();

				if (p.baseSize > 2.0) {
					ctx.beginPath();
					ctx.arc(p.x, p.y, p.size * 2.2, 0, Math.PI * 2);
					ctx.fillStyle = `${col}${a * 0.2})`;
					ctx.fill();
				}
			}

			// Connect nearby particles with subtle lines
			const maxDist = theme.maxLineDistance;
			const maxDistSq = maxDist * maxDist;
			const [r, g, b] = theme.lineColor;

			ctx.lineWidth = 0.75;
			for (let i = 0; i < pLen; i++) {
				const pi = particles[i];
				for (let j = i + 1; j < pLen; j++) {
					const pj = particles[j];
					const dx = pi.x - pj.x;
					const dy = pi.y - pj.y;
					const distSq = dx * dx + dy * dy;

					if (distSq < maxDistSq) {
						const dist = Math.sqrt(distSq);
						const alpha = (1 - dist / maxDist) * 0.18 * (theme.isDark ? 1 : 0.85);
						ctx.strokeStyle = `rgba(${r}, ${g}, ${b}, ${alpha})`;
						ctx.beginPath();
						ctx.moveTo(pi.x, pi.y);
						ctx.lineTo(pj.x, pj.y);
						ctx.stroke();
					}
				}

				// Connect particles to mouse cursor within proximity
				if (mouse.active) {
					const mdx = pi.x - mouse.x;
					const mdy = pi.y - mouse.y;
					const mDistSq = mdx * mdx + mdy * mdy;
					const mRadiusSq = mouse.radius * mouse.radius;

					if (mDistSq < mRadiusSq) {
						const mDist = Math.sqrt(mDistSq);
						const mAlpha = (1 - mDist / mouse.radius) * 0.32;
						const [mr, mg, mb] = theme.mouseLineColor;
						ctx.strokeStyle = `rgba(${mr}, ${mg}, ${mb}, ${mAlpha})`;
						ctx.beginPath();
						ctx.moveTo(pi.x, pi.y);
						ctx.lineTo(mouse.x, mouse.y);
						ctx.stroke();
					}
				}
			}

			animationFrameId = requestAnimationFrame(loop);
		}

		loop();

		return () => {
			isRunning = false;
			if (animationFrameId) cancelAnimationFrame(animationFrameId);
			window.removeEventListener('mousemove', handleMouseMove);
			document.removeEventListener('mouseleave', handleMouseLeave);
			window.removeEventListener('touchmove', handleTouchMove);
			window.removeEventListener('touchend', handleTouchEnd);
			window.removeEventListener('resize', resize);
			document.removeEventListener('visibilitychange', handleVisibilityChange);
			themeObserver.disconnect();
		};
	});
</script>

<canvas
	bind:this={canvas}
	class="fixed inset-0 z-0 pointer-events-none overflow-hidden select-none"
	aria-hidden="true"
></canvas>