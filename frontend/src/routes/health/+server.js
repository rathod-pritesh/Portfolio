/** @type {import('./$types').RequestHandler} */
export async function GET() {
  return new Response(
    JSON.stringify({
      status: "ok",
      service: "Pritesh Rathod Developer Portfolio API",
      version: "1.0.0",
      timestamp: new Date().toISOString(),
      links: {
        documentation: "https://priteshrathod.vercel.app/docs",
        openapi: "https://priteshrathod.vercel.app/openapi.json",
        projects: "https://priteshrathod.vercel.app/api/projects",
        skills: "https://priteshrathod.vercel.app/api/skills"
      }
    }),
    {
      status: 200,
      headers: {
        "Content-Type": "application/json; charset=utf-8",
        "Access-Control-Allow-Origin": "*",
        "Cache-Control": "public, max-age=60"
      }
    }
  );
}

export async function OPTIONS() {
  return new Response(null, {
    status: 204,
    headers: {
      "Access-Control-Allow-Origin": "*",
      "Access-Control-Allow-Methods": "GET, OPTIONS",
      "Access-Control-Allow-Headers": "Content-Type, Authorization, Accept"
    }
  });
}
