/**
 * Compares two objects and returns a partial object containing only the fields
 * that differ. Normalizes null/undefined values to empty strings for comparison.
 * * @template T - The type of the objects being compared.
 * @param {T} original - The baseline object (usually from the API).
 * @param {T} modified - The current state containing potential changes.
 * @returns {Partial<T>} An object containing only the modified key-value pairs.
 */
export function getDirtyFields<T extends object>(original: T, modified: T): Partial<T> {
  const diff: Partial<T> = {};

  for (const key in modified) {
    const originalValue = original[key];
    const modifiedValue = modified[key];

    const normalizedOriginal = originalValue ?? '';
    const normalizedModified = modifiedValue ?? '';

    if (normalizedOriginal !== normalizedModified) {
      diff[key] = modifiedValue;
    }
  }

  return diff;
}

/**
 * Maps properties from a source object to a target request object.
 * Automatically handles casing differences by mapping target keys (e.g., 'Bio') 
 * to lowercase source keys (e.g., 'bio').
 * * @template T - The type of the target request object.
 * @template S - The type of the source data object.
 * @param {T} target - The reactive request object to be populated.
 * @param {S} source - The source data (usually an API response).
 */
export function syncRequest<T extends object, S extends object>(target: T, source: S) {
  for (const key in target) {
    const sourceKey = (key.charAt(0).toLowerCase() + key.slice(1)) as keyof S;
    const value = source[sourceKey];

    if (value !== undefined) {
      target[key] = (value ?? '') as any;
    }
  }
}

/**
 * Determines if there are any differences between the original and modified objects.
 * Useful for disabling save buttons or triggering unsaved changes warnings.
 * * @template T - The type of the objects being compared.
 * @param {T} original - The baseline object.
 * @param {T} modified - The object to check for changes.
 * @returns {boolean} True if at least one field has changed.
 */
export function hasChanged<T extends object>(original: T, modified: T): boolean {
  const dirty = getDirtyFields(original, modified);
  return Object.keys(dirty).length > 0;
}