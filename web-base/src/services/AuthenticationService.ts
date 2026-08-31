// api/AuthenticationService.ts
import { config } from "../helpers/config";
import { getUsersMe } from "../api";

export class AuthenticationService {
    private static _provider: string = 'google'; 

    public static setProvider(provider: string): void {
        this._provider = provider;
    }

    public static handleCallback(): boolean {
        const url = new URL(window.location.href);
        const token = url.searchParams.get('token');

        if (token) {
            localStorage.setItem('jwt_token', token);
            
            url.searchParams.delete('token');
            url.searchParams.delete('redirect');
            window.history.replaceState({}, document.title, url.toString());
            
            return true;
        }
        return false;
    }

    public static login(provider?: string): void {
        const targetProvider = provider || this._provider;
        window.location.href = `${config.apiBaseUrl}/authentication/${targetProvider}`;
    }

    public static logout(): void {
        localStorage.removeItem('jwt_token');

        window.location.href = `${config.apiBaseUrl}/authentication/logout/${this._provider}`;
    }

    // Silently renews the access token using the httpOnly refresh cookie -
    // credentials:'include' is required for the browser to send it
    // cross-origin. Returns false (without throwing) on any failure, since
    // callers should treat "couldn't refresh" as "session's over", not as
    // an error to surface.
    public static async refreshToken(): Promise<boolean> {
        try {
            const response = await fetch(`${config.apiBaseUrl}/authentication/refresh`, {
                method: 'POST',
                credentials: 'include'
            });
            if (!response.ok) return false;

            const data = await response.json();
            if (!data.token) return false;

            localStorage.setItem('jwt_token', data.token);
            return true;
        } catch {
            return false;
        }
    }
}