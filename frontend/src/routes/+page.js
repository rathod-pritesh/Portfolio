export async function load({ fetch }) {
  const [homeRes, aboutRes, skillsRes, projectsRes] = await Promise.all([
    fetch("http://127.0.0.1:8080/api/home"),
    fetch("http://127.0.0.1:8080/api/about"),
    fetch("http://127.0.0.1:8080/api/skills"),
    fetch("http://127.0.0.1:8080/api/projects")
  ]);

  return {
    home: await homeRes.json(),
    about: await aboutRes.json(),
    skills: await skillsRes.json(),
    projects: await projectsRes.json()
  }
}