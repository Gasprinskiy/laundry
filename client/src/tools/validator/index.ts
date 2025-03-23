import type { Validation } from '@vuelidate/core';
import type { FormValidationStatus } from 'naive-ui/es/form/src/interface';

export function valudationStatus<T extends Validation>(validator: T, key: keyof T): FormValidationStatus | undefined {
  if (validator[key].$invalid && validator.$dirty) {
    return 'error';
  }
  return undefined;
}
