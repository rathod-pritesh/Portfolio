<script>
  import { onMount } from "svelte";
  let canvas;

  onMount(() => {
    const ctx = canvas.getContext("2d");
    let raf;
    const mouse = { x: -9999, y: -9999 };

    const N = 50,
      LINK = 140,
      MR = 200,
      SPD = 0.35,
      DR = 2;

    let W, H, pts;

    const resize = () => {
      W = canvas.width = innerWidth;
      H = canvas.height = innerHeight;
    };

    const init = () => {
      resize();
      pts = Array.from({ length: N }, () => ({
        x: Math.random() * W,
        y: Math.random() * H,
        vx: (Math.random() - 0.5) * SPD,
        vy: (Math.random() - 0.5) * SPD,
      }));
    };

    const isDark = () =>
      document.documentElement.getAttribute("data-theme") === "dark" ||
      document.documentElement.classList.contains("dark");

    const loop = () => {
      ctx.clearRect(0, 0, W, H);
      const dark = isDark();
      const rgb = dark ? "255,255,255" : "15,15,15";

      for (const p of pts) {
        const dx = p.x - mouse.x,
          dy = p.y - mouse.y;
        const d = Math.hypot(dx, dy);
        if (d < MR && d > 0) {
          const f = ((MR - d) / MR) ** 2 * 1.2;
          p.vx += (dx / d) * f;
          p.vy += (dy / d) * f;
        }
        p.vx *= 0.96;
        p.vy *= 0.96;
        p.x = (p.x + p.vx + W) % W;
        p.y = (p.y + p.vy + H) % H;
      }

      for (let i = 0; i < pts.length; i++) {
        for (let j = i + 1; j < pts.length; j++) {
          const d = Math.hypot(pts[i].x - pts[j].x, pts[i].y - pts[j].y);
          if (d < LINK) {
            const a = ((1 - d / LINK) * (dark ? 0.38 : 0.22)).toFixed(3);
            ctx.beginPath();
            ctx.strokeStyle = `rgba(${rgb},${a})`;
            ctx.lineWidth = 0.8;
            ctx.moveTo(pts[i].x, pts[i].y);
            ctx.lineTo(pts[j].x, pts[j].y);
            ctx.stroke();
          }
        }
      }

      ctx.fillStyle = dark ? "rgba(255,255,255,0.72)" : "rgba(15,15,15,0.55)";
      for (const p of pts) {
        ctx.beginPath();
        ctx.arc(p.x, p.y, DR, 0, Math.PI * 2);
        ctx.fill();
      }

      raf = requestAnimationFrame(loop);
    };

    init();
    loop();

    const mm = (e) => {
      mouse.x = e.clientX;
      mouse.y = e.clientY;
    };
    const ml = () => {
      mouse.x = mouse.y = -9999;
    };
    addEventListener("mousemove", mm);
    addEventListener("mouseleave", ml);
    addEventListener("resize", resize);

    return () => {
      cancelAnimationFrame(raf);
      removeEventListener("mousemove", mm);
      removeEventListener("mouseleave", ml);
      removeEventListener("resize", resize);
    };
  });
</script>

<canvas
  bind:this={canvas}
  style="position:fixed;inset:0;z-index:-10;pointer-events:none;width:100vw;height:100vh;"
></canvas>
