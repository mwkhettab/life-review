type Star = {
  x: number;
  y: number;
  r: number;
  baseAlpha: number;
  speed: number;
  offset: number;
  hue: string;
};

type Layer = {
  count: number;
  minR: number;
  maxR: number;
  alpha: number;
  twinkle: number;
};

export function initStars(): void {
  console.log("Hello");

  const canvas: HTMLCanvasElement = document.createElement("canvas");

  canvas.id = "stars-canvas";

  const canvasStyle: Partial<CSSStyleDeclaration> = {
    position: "fixed",
    inset: "0",
    zIndex: "-1",
    pointerEvents: "none",
  };

  Object.assign(canvas.style, canvasStyle);

  document.body.appendChild(canvas);

  const ctx: CanvasRenderingContext2D | null = canvas.getContext("2d");

  if (ctx === null) {
    throw new Error("2D canvas context not available");
  }

  const LAYERS: ReadonlyArray<Layer> = [
    { count: 320, minR: 0.3, maxR: 0.7, alpha: 0.55, twinkle: 0.012 },
    { count: 120, minR: 0.7, maxR: 1.2, alpha: 0.75, twinkle: 0.018 },
    { count: 40, minR: 1.2, maxR: 2.0, alpha: 0.9, twinkle: 0.025 },
  ];

  let stars: Star[] = [];

  let W: number = 0;
  let H: number = 0;

  const rand = (a: number, b: number): number => {
    return a + Math.random() * (b - a);
  };

  const build = (): void => {
    const builtStars: Star[] = [];

    for (const l of LAYERS) {
      const layerStars: Star[] = Array.from({ length: l.count }, (): Star => {
        const hue: string =
          Math.random() < 0.15
            ? Math.random() < 0.5
              ? "200,220,255"
              : "255,200,200"
            : "255,255,255";

        return {
          x: rand(0, W),
          y: rand(0, H),
          r: rand(l.minR, l.maxR),
          baseAlpha: rand(l.alpha * 0.5, l.alpha),
          speed: rand(l.twinkle * 0.5, l.twinkle * 1.5),
          offset: rand(0, Math.PI * 2),
          hue: hue,
        };
      });

      builtStars.push(...layerStars);
    }

    stars = builtStars;
  };

  const resize = (): void => {
    W = canvas.width = window.innerWidth;
    H = canvas.height = window.innerHeight;

    build();
  };

  let t: number = 0;

  const draw = (): void => {
    ctx.clearRect(0, 0, W, H);

    t += 0.012;

    for (const s of stars) {
      const a: number =
        s.baseAlpha + Math.sin(t * s.speed * 60 + s.offset) * 0.25;

      ctx.beginPath();

      ctx.arc(s.x, s.y, s.r, 0, Math.PI * 2);

      ctx.fillStyle = `rgba(${s.hue},${Math.max(0, a)})`;

      ctx.fill();

      if (s.r > 1.2) {
        const g: CanvasGradient = ctx.createRadialGradient(
          s.x,
          s.y,
          0,
          s.x,
          s.y,
          s.r * 4,
        );

        g.addColorStop(0, `rgba(${s.hue},${a * 0.35})`);

        g.addColorStop(1, `rgba(${s.hue},0)`);

        ctx.beginPath();

        ctx.arc(s.x, s.y, s.r * 4, 0, Math.PI * 2);

        ctx.fillStyle = g;

        ctx.fill();
      }
    }

    requestAnimationFrame(draw);
  };

  const resizeHandler = (): void => {
    resize();
  };

  window.addEventListener("resize", resizeHandler as EventListener);

  resize();

  draw();
}
