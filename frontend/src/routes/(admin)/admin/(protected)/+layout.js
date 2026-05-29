import { redirect } from '@sveltejs/kit';
import { API_BASE } from '$lib/config/api';

export async function load({ fetch, url }) {
    // allow login page itself
    if (url.pathname === '/admin') {
        return {};
    }

    const res = await fetch(`${API_BASE}/api/admin/verify`, {
        credentials: 'include'
    });

    if (!res.ok) {
        throw redirect(302, '/admin');
    }

    return {};
}

export const ssr = false;