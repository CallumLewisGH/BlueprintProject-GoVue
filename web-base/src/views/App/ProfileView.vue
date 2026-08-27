<template>
  <div v-if="userProfile" class="max-w-2xl mx-auto p-6 min-h-screen">
    <div class="relative mb-8">
      <button
        @click="fileInput?.click()"
        :disabled="isUploadingPicture"
        class="absolute -top-2 -right-2 p-2 rounded-full shadow-lg hover:opacity-90 transition-opacity z-10 disabled:opacity-60"
        style="background-color: var(--accent); color: var(--header-text);"
      >
        <ArrowPathIcon v-if="isUploadingPicture" class="w-5 h-5 animate-spin" />
        <PencilSquareIcon v-else class="w-5 h-5" />
      </button>
      <input
        type="file"
        ref="fileInput"
        class="hidden"
        accept="image/*"
        @change="handleFileChange"
      />
      <ProfileHeader :profile="userProfile" />
    </div>

    <div class="p-6 rounded-xl border"
         :style="{ backgroundColor: 'var(--card-bg)', borderColor: 'var(--border)' }">
      <div class="flex items-center justify-between mb-6">
        <h3 class="font-bold text-lg" :style="{ color: 'var(--text-color)' }">Edit Profile</h3>
        <Button
          label="Delete Account"
          variant="danger"
          :disabled="deleting"
          @click="handleUserDelete"
        />
      </div>

      <form @submit.prevent="handleSaveProfile" class="space-y-4">
        <TextInput
          :model-value="editRequest.username || ''"
          label="Username"
          required
          @update:model-value="editRequest.username = $event"
        />

        <TextInput
          type="email"
          :model-value="editRequest.email || ''"
          label="Email"
          required
          @update:model-value="editRequest.email = $event"
        />

        <CountedTextArea
          :model-value="editRequest.bio || ''"
          label="Bio"
          :max-length="500"
          :show-count="true"
          :rows="3"
          @update:model-value="editRequest.bio = $event"
        />

        <div class="pt-4 flex gap-3">
          <Button
            label="Save Changes"
            type="submit"
            variant="primary"
            :disabled="isSaving || !hasChanges"
          />
          <Button
            label="Cancel"
            variant="secondary"
            @click="resetEdit"
          />
        </div>
      </form>
    </div>
  </div>

  <div v-else class="flex justify-center items-center min-h-[400px]">
    <div class="animate-spin rounded-full h-12 w-12 border-b-2" :style="{ borderColor: 'var(--accent)' }"></div>
  </div>

  <Modal :is-open="showErrorModal" title="Update Failed" @close="showErrorModal = false">
    <p class="text-sm whitespace-pre-wrap" :style="{ color: 'var(--text-color)' }">
      {{ errorMessage }}
    </p>
    <template #footer>
      <Button label="Close" variant="primary" @click="showErrorModal = false" />
    </template>
  </Modal>

  <ConfirmDeleteModal
    :is-open="showDeleteModal"
    title="Delete Account"
    message="Are you sure you want to delete your account? This action cannot be undone."
    :is-deleting="deleting"
    @confirm="confirmDeleteUser"
    @close="showDeleteModal = false"
  />
</template>

<script setup lang="ts">
import { ref, reactive, onMounted, computed } from 'vue'
import { useRouter } from 'vue-router'
import type { UpdateUserRequest, UserPrivateProfile } from '@/api'
import { getDirtyFields, hasChanged, syncRequest } from '@/helpers/diff'
import { useImageUpload } from '@/composables/useImageUpload'
import { Modal, Button } from '@/components'
import ConfirmDeleteModal from '@/components/ui/ConfirmDeleteModal.vue'
import ProfileHeader from '@/components/common/ProfileHeader.vue'
import TextInput from '@/components/ui/TextInput.vue'
import CountedTextArea from '@/components/common/CountedTextArea.vue'
import { PencilSquareIcon, ArrowPathIcon } from '@heroicons/vue/24/outline'
import { UsersService } from '@/services/UsersService'
import { AuthenticationService } from '@/services/AuthenticationService'

const router = useRouter()

const deleting = ref(false)
const showDeleteModal = ref(false)
const userProfile = ref<UserPrivateProfile | null>(null)

const isSaving = ref(false)
const fileInput = ref<HTMLInputElement | null>(null)
const showErrorModal = ref(false)
const errorMessage = ref('')

const { isUploading: isUploadingPicture, error: uploadError, uploadImage } = useImageUpload()

const editRequest = reactive<UpdateUserRequest>({
  username: undefined,
  bio: undefined,
  email: undefined,
  profilePicture: undefined
})

const hasChanges = computed(() => {
  if (!userProfile.value) return false
  return hasChanged(userProfile.value, editRequest)
})

onMounted(async () => {
  const result = await UsersService.getCurrentUser()
  if (result.data) {
    userProfile.value = result.data
    syncRequest(editRequest, result.data)
  }
})

async function handleSaveProfile() {
  if (!userProfile.value || !hasChanges.value) return
  isSaving.value = true
  const payload = getDirtyFields(userProfile.value, editRequest)
  const result = await UsersService.updateCurrentUser(payload)

  if (result.error || result.data === null) {
    errorMessage.value = result?.error?.errors?.[0]?.message || 'Unknown error'
    showErrorModal.value = true
    isSaving.value = false
    return
  }

  userProfile.value = result.data
  syncRequest(editRequest, result.data)
  isSaving.value = false
}

async function handleFileChange(event: Event) {
  const target = event.target as HTMLInputElement
  const file = target.files?.[0]
  target.value = ''
  if (!file) return

  const publicUrl = await uploadImage(file, 'profile')
  if (!publicUrl) {
    errorMessage.value = uploadError.value || 'Upload failed.'
    showErrorModal.value = true
    return
  }

  editRequest.profilePicture = publicUrl
  await handleSaveProfile()
}

function resetEdit() { if (userProfile.value) syncRequest(editRequest, userProfile.value) }

async function handleUserDelete() { showDeleteModal.value = true }

async function confirmDeleteUser() {
  deleting.value = true
  try {
    await UsersService.deleteCurrentUser()
    AuthenticationService.logout()
    router.push('/')
  } finally {
    deleting.value = false
    showDeleteModal.value = false
  }
}
</script>
