import { ref } from 'vue'
import { UploadsService } from '@/services/UploadsService'

const ALLOWED_TYPES = ['image/jpeg', 'image/png', 'image/webp']

// Two-step upload: ask the backend for a signed URL scoped to this user and
// purpose, then PUT the file directly to storage. The backend never sees
// the file bytes. `purpose` must match a case the backend's upload command
// recognizes (currently just "profile" - add more there before using them here).
export function useImageUpload() {
  const isUploading = ref(false)
  const error = ref<string | null>(null)

  async function uploadImage(file: File, purpose: string): Promise<string | null> {
    error.value = null

    if (!ALLOWED_TYPES.includes(file.type)) {
      error.value = 'Please choose a JPEG, PNG, or WebP image.'
      return null
    }

    isUploading.value = true
    try {
      const signResult = await UploadsService.createUpload({ purpose, contentType: file.type })
      if (!signResult.data) {
        error.value = signResult.error?.detail || 'Could not start the upload.'
        return null
      }

      const putResponse = await fetch(signResult.data.uploadUrl, {
        method: 'PUT',
        headers: { 'Content-Type': file.type },
        body: file
      })

      if (!putResponse.ok) {
        error.value = 'The upload failed partway through. Please try again.'
        return null
      }

      return signResult.data.publicUrl
    } finally {
      isUploading.value = false
    }
  }

  return { isUploading, error, uploadImage }
}
