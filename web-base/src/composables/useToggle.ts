import { ref, computed } from 'vue'

export function useToggle(initialState = false) {
  const value = ref(initialState)

  function toggle() {
    value.value = !value.value
  }

  function set(newValue: boolean) {
    value.value = newValue
  }

  return {
    value,
    toggle,
    set,
    isOn: computed(() => value.value),
    isOff: computed(() => !value.value)
  }
}
