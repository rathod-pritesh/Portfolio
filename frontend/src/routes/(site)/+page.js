import { API_BASE } from "$lib/config/api";
import {
  defaultHome,
  defaultAbout,
  defaultSkills,
  defaultProjects,
  defaultAutomations,
  defaultCertifications,
  defaultEducation
} from "$lib/portfolioData.js";

async function safeFetch(fetchFn, url, fallback) {
  try {
    const controller = new AbortController();
    const timeoutId = setTimeout(() => controller.abort(), 3500);
    const res = await fetchFn(url, { signal: controller.signal });
    clearTimeout(timeoutId);
    if (!res.ok) return fallback;
    const data = await res.json();
    return data ?? fallback;
  } catch {
    return fallback;
  }
}

export async function load({ fetch }) {
  const [home, about, skills, projects, automations, certifications, education] = await Promise.all([
    safeFetch(fetch, `${API_BASE}/api/home`, defaultHome),
    safeFetch(fetch, `${API_BASE}/api/about`, defaultAbout),
    safeFetch(fetch, `${API_BASE}/api/skills`, defaultSkills),
    safeFetch(fetch, `${API_BASE}/api/projects`, defaultProjects),
    safeFetch(fetch, `${API_BASE}/api/automations`, defaultAutomations),
    safeFetch(fetch, `${API_BASE}/api/certifications`, defaultCertifications),
    safeFetch(fetch, `${API_BASE}/api/education`, defaultEducation),
  ]);

  return {
    home,
    about,
    skills,
    projects,
    automations,
    certifications,
    education,
  };
}