import { api_get_projects } from "$lib/api.js";

export async function load({ fetch }) {
  try {
    let projects = await api_get_projects(fetch);
    return {
      projects: null === projects ? [] : projects,
      error: "",
    };
  } catch (err) {
    return {
      error: err.toString(),
    };
  }
}
