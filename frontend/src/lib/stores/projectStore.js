import { writable } from "svelte/store";

const BASE = "http://localhost:8080/api/admin";

export const projects = writable([]);
export const loading = writable(true);
export const storeError = writable("");

const authFetch = (url, options = {}) => 
  fetch(url, { credentials: "include", ...options });

export async function fetchProjects() {
  loading.set(true);
  storeError.set("");
  try {
    const res = await authFetch(`${BASE}/projects`);
    if (!res.ok) throw new Error(`HTTP ${res.status}`);
    const data = await res.json();
    projects.set(data || []);
  } catch (err) {
    storeError.set("Failed to load projects");
    console.error(err);
  } finally {
    loading.set(false);
  }
}

// add new project
export async function addProject(payload) {
  const res = await authFetch(`${BASE}/projects`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(payload),
  });
  if (!res.ok) {
    const err = await res.json();
    throw new Error(err.error || "Failed to add project");
  }
  await fetchProjects();
}

// update existing project
export async function updateProject(id, payload) {
  const res = await authFetch(`${BASE}/projects/${id}`, {
    method: "PUT",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(payload),
  });
  if (!res.ok) {
    const err = await res.json();
    throw new Error(err.error || "Failed to update project");
  }
  await fetchProjects();
}

// delete project by id
export async function deleteProject(id) {
  const res = await authFetch(`${BASE}/projects/${id}`, { method: "DELETE" });
  if (!res.ok) {
    const err = await res.json();
    throw new Error(err.error || "Failed to delete project");
  }
  // remove locally
  projects.update((list) => list.filter((p) => p._id !== id));
}
