import { client } from '@/api/client.gen'
import router from '@/router';

let isInitialized = false;

export function useClient() {
    if (isInitialized) return;

    client.setConfig({
        baseUrl: import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080',
        throwOnError: false
    });

    client.interceptors.request.use((request) => {
        const token = localStorage.getItem('jwt_token');
        if (token) {
            request.headers.set('Authorization', `Bearer ${token}`);
        }
        return request;
    });

    client.interceptors.response.use((response) => {
        if (response.status === 401) {
            console.warn('Unauthorized! Redirecting to login...');
            
            localStorage.removeItem('jwt_token');

            router.push({ 
                name: 'Login', 
                query: { redirect: router.currentRoute.value.fullPath } 
            });
        }
        return response;
    });

    isInitialized = true;
}