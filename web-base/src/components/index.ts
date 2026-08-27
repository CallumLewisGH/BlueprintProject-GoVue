/**
 * UI Components - Base form and interactive elements
 * These are small, focused, reusable components following Vue best practices
 */

// Form Inputs
export { default as TextInput } from './ui/TextInput.vue'
export { default as TextArea } from './ui/TextArea.vue'
export { default as PasswordInput } from './ui/PasswordInput.vue'
export { default as Select } from './ui/Select.vue'
export { default as Checkbox } from './ui/Checkbox.vue'
export { default as CheckboxGroup } from './ui/CheckboxGroup.vue'
export { default as Radio } from './ui/Radio.vue'
export { default as RadioGroup } from './ui/RadioGroup.vue'

// Buttons and Actions
export { default as Button } from './ui/Button.vue'

// Feedback
export { default as Alert } from './ui/Alert.vue'

// Modals
export { default as Modal } from './ui/Modal.vue'
export { default as ConfirmDeleteModal } from './ui/ConfirmDeleteModal.vue'

/**
 * Common Components - Semantic components combining UI elements
 * These provide common patterns and functionality
 */

// Layout
export { default as ProfileHeader } from './common/ProfileHeader.vue'
export { default as TabNavigation } from './common/TabNavigation.vue'
export { default as Card } from './common/Card.vue'
export { default as EmptyState } from './common/EmptyState.vue'
export { default as UserAvatar } from './common/UserAvatar.vue'
export { default as ImagePicker } from './common/ImagePicker.vue'

// Forms
export { default as CountedTextArea } from './common/CountedTextArea.vue'
export { default as SearchableSelect } from './common/SearchableSelect.vue'
export { default as RatingInput } from './common/RatingInput.vue'
export { default as ButtonGroup } from './common/ButtonGroup.vue'
export { default as FormCard } from './common/FormCard.vue'

// Notifications
export { default as Notification } from './common/Notification.vue'
