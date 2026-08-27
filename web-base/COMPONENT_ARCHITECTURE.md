# Web Base - Component Architecture

## Overview

This document outlines the component-based architecture for this template project, designed to maximize code reuse and follow Vue 3 best practices.

## Architecture Principles

All components follow these coding standards from the README:
1. **Error Early** - Validate inputs and handle errors gracefully
2. **Be Explicit** - Clear prop interfaces and explicit intent
3. **No Recursion** - Components should not recursively call themselves
4. **No await in loops** - Async operations are properly handled
5. **Component-Based** - Small, focused, reusable components
6. **Modal Focus Lock** - All modals implement proper focus management
7. **Function Syntax** - Prefer `function` over arrow functions (in setup scripts)

## Directory Structure

```
src/components/
├── index.ts                    # Central export file
├── ui/                         # Base UI components
│   ├── TextInput.vue           # Text input field
│   ├── TextArea.vue            # Multi-line text input
│   ├── PasswordInput.vue        # Password input with toggle
│   ├── Select.vue              # Dropdown select
│   ├── Checkbox.vue            # Single checkbox
│   ├── CheckboxGroup.vue       # Multiple checkboxes
│   ├── Radio.vue               # Single radio button
│   ├── RadioGroup.vue          # Radio group
│   ├── Button.vue              # Primary button component
│   ├── Alert.vue               # Alert/message box
│   └── Modal.vue               # Modal dialog with focus lock
├── common/                     # Semantic components
│   ├── Card.vue                # Flexible card container
│   ├── ProfileHeader.vue       # User profile header
│   ├── TabNavigation.vue       # Tab navigation
│   ├── EmptyState.vue          # Empty/no results state
│   ├── UserAvatar.vue          # User avatar display
│   ├── CountedTextArea.vue     # Text area with char count
│   ├── SearchableSelect.vue    # Select with search/filter
│   ├── RatingInput.vue         # Star rating input
│   ├── ButtonGroup.vue         # Group of toggleable buttons
│   ├── FormCard.vue            # Form wrapper container
│   └── Notification.vue        # Auto-dismissing notification
└── App/                        # Page-level components
    └── Header.vue              # Navigation header

src/composables/
├── index.ts                    # Central export file
├── useForm.ts                  # Form state management
├── useFormValidation.ts        # Form field validation
├── useAsync.ts                 # Async operation handling
├── useModal.ts                 # Modal open/close state
├── useToggle.ts                # Boolean toggle state
├── useClickOutside.ts          # Click outside detection
├── useDarkMode.ts              # Dark mode management
└── useClient.ts                # API client setup
```

## UI Components (Building Blocks)

### TextInput
Single-line text input with validation support.

```vue
<TextInput
  v-model="username"
  label="Username"
  placeholder="Enter username"
  :required="true"
  :error="errors.username"
  @blur="validateUsername"
/>
```

### TextArea
Multi-line text input with optional character limit.

```vue
<TextArea
  v-model="bio"
  label="Bio"
  placeholder="Tell us about yourself"
  :rows="3"
  :error="errors.bio"
/>
```

### Button
Flexible button with multiple variants and loading states.

```vue
<Button
  label="Save Changes"
  variant="primary"        <!-- primary | secondary | outline | danger -->
  :disabled="isLoading"
  :loading-text="'Saving...'"
  @click="handleSave"
/>
```

### Select
Dropdown selection with label and validation.

```vue
<Select
  v-model="selectedType"
  :options="itemTypes"
  label="Item Type"
  placeholder="Choose type"
  :required="true"
/>
```

### Checkbox / CheckboxGroup
Single or multiple choice selections.

```vue
<!-- Single checkbox -->
<Checkbox
  v-model="agreeToTerms"
  label="I agree to terms"
/>

<!-- Multiple checkboxes -->
<CheckboxGroup
  v-model="selectedFeatures"
  :options="featureOptions"
  label="Features"
/>
```

### Radio / RadioGroup
Single selection from multiple options.

```vue
<RadioGroup
  v-model="campAccessType"
  :options="accessOptions"
  name="access"
  label="How to access"
/>
```

### Modal
Dialog with focus lock and keyboard handling.

```vue
<Modal
  :is-open="isModalOpen"
  title="Confirm Action"
  @close="isModalOpen = false"
>
  <p>Are you sure?</p>
  <template #footer>
    <Button label="Cancel" @click="isModalOpen = false" />
    <Button label="Confirm" variant="danger" @click="confirm" />
  </template>
</Modal>
```

### Alert
Contextual alerts for messages, warnings, errors.

```vue
<Alert
  type="success"    <!-- success | error | warning | info -->
  title="Success"
  message="Operation completed"
  :closable="true"
/>
```

### PasswordInput
Text input specifically for passwords with show/hide toggle.

```vue
<PasswordInput
  v-model="password"
  label="Password"
  :required="true"
/>
```

## Common Components (Semantic)

### Card
Flexible container for content with optional slots.

```vue
<Card @click="selectItem">
  <template #icon>📦</template>
  <template #title>Item Name</template>
  <template #description>Item description</template>
  <template #footer>Additional info</template>
</Card>
```

### ProfileHeader
Display user profile information with optional edit button.

```vue
<ProfileHeader
  :profile="userProfile"
  :show-edit-button="true"
  @edit="openEditModal"
/>
```

### TabNavigation
Tab bar for switching between views.

```vue
<TabNavigation
  :tabs="tabs"
  :active-tab="activeTabId"
  @select="activeTabId = $event"
/>
```

### EmptyState
Display when no content is available.

```vue
<EmptyState
  icon="🔍"
  title="No results"
  message="Try adjusting your search filters"
/>
```

### CountedTextArea
TextArea with character counter.

```vue
<CountedTextArea
  v-model="bio"
  label="Bio"
  :max-length="500"
  :show-count="true"
/>
```

### RatingInput
Star rating selector.

```vue
<RatingInput
  v-model="rating"
  label="Rating"
  :max-rating="5"
  :show-rating="true"
/>
```

### ButtonGroup
Group of toggleable buttons.

```vue
<ButtonGroup
  v-model="selectedOptions"
  :options="options"
  label="Choose features"
/>
```

### FormCard
Wrapper component for forms with automatic footer handling.

```vue
<FormCard
  :is-loading="isSaving"
  submit-label="Save"
  @submit="handleSubmit"
  @cancel="closeForm"
>
  <!-- Form fields -->
</FormCard>
```

### Notification
Auto-dismissing notification message.

```vue
<Notification
  icon="✓"
  title="Saved"
  message="Changes saved successfully"
  :duration="3000"
/>
```

## Composables (Logic)

### useForm
Manage form state, values, and loading.

```typescript
const { formData, state, updateField, setLoading, resetForm } = useForm({
  username: '',
  email: ''
})
```

### useFormValidation
Validate form fields with custom rules.

```typescript
const { errors, registerField, validateAll, hasErrors } = useFormValidation()

registerField('email', [
  {
    validate: (v) => /.+@.+\..+/.test(v) || 'Invalid email',
    message: 'Email format is incorrect'
  }
])
```

### useAsync
Handle async operations with loading and error states.

```typescript
const { data, error, isLoading, execute } = useAsync(() =>
  fetchUserData()
)
```

### useModal
Manage modal open/close state.

```typescript
const { isOpen, open, close, toggle } = useModal()
```

### useToggle
Simple boolean state toggling.

```typescript
const { value, toggle, set } = useToggle(false)
```

### useClickOutside
Detect clicks outside an element.

```typescript
const { elementRef } = useClickOutside(() => {
  isMenuOpen.value = false
})

<div ref="elementRef"><!-- content --></div>
```

## Migration Guide: From Old to New

### Before (Manual Form Input)
```vue
<input
  v-model="username"
  type="text"
  class="w-full px-3 py-2 border rounded-lg"
  style="background-color: var(--surface); border-color: var(--border);"
/>
```

### After (Using TextInput Component)
```vue
<TextInput
  v-model="username"
  label="Username"
  placeholder="Enter username"
/>
```

### Before (Manual Button)
```vue
<button
  @click="save"
  :disabled="isSaving"
  class="px-4 py-2 rounded-lg bg-orange-600 text-white hover:bg-orange-700"
>
  {{ isSaving ? 'Saving...' : 'Save' }}
</button>
```

### After (Using Button Component)
```vue
<Button
  label="Save"
  variant="primary"
  :disabled="isSaving"
  loading-text="Saving..."
/>
```

## Best Practices

1. **Compose Components** - Build complex UIs by combining smaller components
2. **Use Composables** - Extract logic into reusable composables
3. **Keep Components Small** - Single responsibility principle
4. **Validate Early** - Use useFormValidation for user inputs
5. **Handle Loading States** - Use Button's loading props for async operations
6. **Proper Error Handling** - Display validation errors near inputs
7. **Test Components Independently** - Each component should work in isolation
8. **Document Slots** - Make it clear what slot content should be
9. **Consistent Styling** - All components use CSS variables (var(--accent), etc)
10. **Accessible** - Components include proper aria labels and keyboard support

## Development Workflow

When adding a new feature:

1. **Identify reusable UI patterns** - Create UI components
2. **Build semantic components** - Combine UI components for specific use cases
3. **Extract logic to composables** - Create composables for repeated logic
4. **Use in views** - Compose views from components and composables
5. **Test thoroughly** - Ensure components work in isolation and together

## Common Patterns

### Form Handling
```vue
<template>
  <FormCard @submit="handleSubmit" :is-loading="isSaving">
    <TextInput v-model="form.username" label="Username" />
    <TextArea v-model="form.bio" label="Bio" />
    <Button label="Save" type="submit" />
  </FormCard>
</template>

<script setup>
const { formData, setLoading } = useForm({ username: '', bio: '' })
const { validateAll, errors } = useFormValidation()

async function handleSubmit() {
  if (!validateAll(formData)) return
  setLoading(true)
  await saveUser(formData)
  setLoading(false)
}
</script>
```

### Modal Dialogs
```vue
<template>
  <Button label="Open" @click="isOpen = true" />
  <Modal :is-open="isOpen" title="Confirm" @close="isOpen = false">
    <p>Are you sure?</p>
    <template #footer>
      <Button label="Cancel" @click="isOpen = false" />
      <Button label="Confirm" variant="danger" @click="confirm" />
    </template>
  </Modal>
</template>

<script setup>
const { isOpen, open, close } = useModal()
</script>
```

## File Size & Performance

- **UI Components**: ~2-4KB each (minified)
- **Common Components**: ~3-5KB each (minified)
- **Composables**: ~1-2KB each (minified)
- **Total Bundle Impact**: ~80KB (all components combined, before gzip)

Tree-shaking is enabled, so unused components are eliminated from production builds.

---

**Last Updated**: February 2026
**Version**: 1.0
