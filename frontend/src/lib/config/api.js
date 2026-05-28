const dev = import.meta.env.DEV;

export const API_BASE = dev
    ? "http://localhost:8080"
    : "https://portfolio-i5x9.onrender.com";