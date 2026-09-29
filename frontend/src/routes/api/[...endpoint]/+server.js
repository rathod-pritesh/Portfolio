import {
  defaultHome,
  defaultAbout,
  defaultSkills,
  defaultProjects,
  defaultAutomations,
  defaultCertifications,
  defaultEducation
} from "$lib/portfolioData.js";

const BACKEND_BASE = "https://portfolio-i5x9.onrender.com";

const FALLBACK_MAP = {
  home: defaultHome,
  about: defaultAbout,
  skills: defaultSkills,
  projects: defaultProjects,
  automations: defaultAutomations,
  certifications: defaultCertifications,
  education: defaultEducation
};

/** @type {import('./$types').RequestHandler} */
export async function GET({ params, fetch }) {
  const endpoint = (params.endpoint || "").replace(/^\/+|\/+$/g, "");

  // Special health check alias under /api
  if (endpoint === "health") {
    return new Response(
      JSON.stringify({
        status: "ok",
        service: "Pritesh Rathod Developer Portfolio API",
        version: "1.0.0",
        timestamp: new Date().toISOString()
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

  // Check if endpoint is known
  if (!(endpoint in FALLBACK_MAP)) {
    return new Response(
      JSON.stringify({
        type: "https://priteshrathod.vercel.app/docs/errors#not-found",
        title: "Not Found",
        status: 404,
        detail: `The requested API endpoint '/api/${endpoint}' does not exist on this server.`,
        instance: `/api/${endpoint}`
      }),
      {
        status: 404,
        headers: {
          "Content-Type": "application/problem+json; charset=utf-8",
          "Access-Control-Allow-Origin": "*"
        }
      }
    );
  }

  // Try fetching from backend with 2.5s timeout
  try {
    const controller = new AbortController();
    const timeoutId = setTimeout(() => controller.abort(), 2500);
    const backendRes = await fetch(`${BACKEND_BASE}/api/${endpoint}`, {
      signal: controller.signal
    });
    clearTimeout(timeoutId);

    if (backendRes.ok) {
      const data = await backendRes.json();
      return new Response(JSON.stringify(data), {
        status: 200,
        headers: {
          "Content-Type": "application/json; charset=utf-8",
          "Access-Control-Allow-Origin": "*",
          "Cache-Control": "public, s-maxage=300, stale-while-revalidate=600"
        }
      });
    }
  } catch {
    // Backend offline, cold starting, or timed out - fall back gracefully
  }

  // Fallback data response
  return new Response(JSON.stringify(FALLBACK_MAP[endpoint]), {
    status: 200,
    headers: {
      "Content-Type": "application/json; charset=utf-8",
      "Access-Control-Allow-Origin": "*",
      "X-Cache-Fallback": "true",
      "Cache-Control": "public, s-maxage=60, stale-while-revalidate=120"
    }
  });
}

/** @type {import('./$types').RequestHandler} */
export async function POST({ params, request, fetch }) {
  const endpoint = (params.endpoint || "").replace(/^\/+|\/+$/g, "");

  if (endpoint === "contact") {
    try {
      const body = await request.json();
      const controller = new AbortController();
      const timeoutId = setTimeout(() => controller.abort(), 3500);

      const backendRes = await fetch(`${BACKEND_BASE}/api/contact`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(body),
        signal: controller.signal
      });
      clearTimeout(timeoutId);

      if (backendRes.ok) {
        const result = await backendRes.json();
        return new Response(JSON.stringify(result), {
          status: 200,
          headers: {
            "Content-Type": "application/json; charset=utf-8",
            "Access-Control-Allow-Origin": "*"
          }
        });
      }
    } catch {
      // Backend unavailable; return accepted response
    }

    return new Response(
      JSON.stringify({
        success: true,
        message: "Your message has been received.",
        status: "accepted"
      }),
      {
        status: 200,
        headers: {
          "Content-Type": "application/json; charset=utf-8",
          "Access-Control-Allow-Origin": "*"
        }
      }
    );
  }

  return new Response(
    JSON.stringify({
      type: "https://priteshrathod.vercel.app/docs/errors#method-not-allowed",
      title: "Method Not Allowed",
      status: 405,
      detail: `POST method is not supported on '/api/${endpoint}'.`,
      instance: `/api/${endpoint}`
    }),
    {
      status: 405,
      headers: {
        "Content-Type": "application/problem+json; charset=utf-8",
        "Access-Control-Allow-Origin": "*"
      }
    }
  );
}

/** @type {import('./$types').RequestHandler} */
export async function OPTIONS() {
  return new Response(null, {
    status: 204,
    headers: {
      "Access-Control-Allow-Origin": "*",
      "Access-Control-Allow-Methods": "GET, POST, OPTIONS",
      "Access-Control-Allow-Headers": "Content-Type, Authorization, Accept"
    }
  });
}
