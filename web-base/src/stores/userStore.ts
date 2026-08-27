import { ref } from 'vue'
import { defineStore } from 'pinia'
import { AuthenticationService } from '@/services/AuthenticationService'
import type { UserPrivateProfile } from '@/api'
import { UsersService } from '@/services/UsersService'

export const useUserStore = defineStore('userStore', () => {
  const user = ref<UserPrivateProfile | null>(null)
  const loading = ref(false)
  const error = ref<string | null>(null)

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