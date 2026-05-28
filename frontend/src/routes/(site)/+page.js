import { API_BASE } from "$lib/config/api";

export async function load({ fetch }) {
  const [homeRes, aboutRes, skillsRes, projectsRes, automationRes, certificationRes, educationRes] = await Promise.all([
    fetch(`${API_BASE}/api/home`),
    fetch(`${API_BASE}/api/about`),
    fetch(`${API_BASE}/api/skills`),
    fetch(`${API_BASE}/api/projects`),
    fetch(`${API_BASE}/api/automations`),
    fetch(`${API_BASE}/api/certifications`),
    fetch(`${API_BASE}/api/education`),
  ]);

  return {
    home: await homeRes.json(),
    about: await aboutRes.json(),
    skills: await skillsRes.json(),
    projects: await projectsRes.json(),
    automations: await automationRes.json(),
    certifications: await certificationRes.json(),
    education: await educationRes.json(),
  }
}