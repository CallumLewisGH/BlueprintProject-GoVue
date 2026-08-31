import { ref } from 'vue'
import { defineStore } from 'pinia'
import { AuthenticationService } from '@/services/AuthenticationService'
import type { UserPrivateProfile } from '@/api'
import { UsersService } from '@/services/UsersService'

// Refresh comfortably before the 30-minute access token expires, so an
// occasional slow network moment still leaves room to retry before the
// token actually goes stale.
const REFRESH_INTERVAL_MS = 20 * 60 * 1000

export const useUserStore = defineStore('userStore', () => {
  const user = ref<UserPrivateProfile | null>(null)
  const loading = ref(false)
  const error = ref<string | null>(null)
  let refreshTimer: ReturnType<typeof setInterval> | null = null

  function startRefreshTimer() {
    stopRefreshTimer()
    refreshTimer = setInterval(async () => {
      const ok = await AuthenticationService.refreshToken()
      if (!ok) {
        // The refresh session is gone (expired, logged out elsewhere, etc.) -
        // just drop local state. The next route navigation's auth check will
        // naturally redirect to login.
        stopRefreshTimer()
        localStorage.removeItem('jwt_token')
        user.value = null
      }
    }, REFRESH_INTERVAL_MS)
  }

  function stopRefreshTimer() {
    if (refreshTimer) {
      clearInterval(refreshTimer)
      refreshTimer = null
    }
  }

  async function checkAuth() {
    loading.value = true;

    AuthenticationService.handleCallback();

    try {
      if (localStorage.getItem('jwt_token')) {
        const result = await UsersService.getCurrentUser();
        if (!result.data) {
          throw new Error('No user data returned');
        }

        user.value = result.data;
        startRefreshTimer();
      }
      error.value = null;
      return true;
    } catch (err: any) {
      if (localStorage.getItem('jwt_token')) {
          console.error("Token verification failed", err);
          AuthenticationService.logout();
      }
      user.value = null;
      return false;
    } finally {
      loading.value = false;
    }
  }

  async function login(provider?: string) {
    if (provider) AuthenticationService.setProvider(provider)
    AuthenticationService.login()
  }

  async function logout() {
    stopRefreshTimer()
    AuthenticationService.logout()
    user.value = null
  }

  return {
    user,
    loading,
    error,
    login,
    logout,
    checkAuth
  }
})