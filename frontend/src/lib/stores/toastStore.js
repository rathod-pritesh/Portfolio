import { writable } from "svelte/store";

export const toast = writable(null);

let timer = null;

export function showToast(msg, type = "success") {
  if (timer) clearTimeout(timer);
  toast.set({ msg, type });
  timer = setTimeout(() => {
    toast.set(null);
    timer = null;
  }, 3000);
}