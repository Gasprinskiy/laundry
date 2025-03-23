<script lang="ts" setup>
import { computed, ref, shallowRef } from 'vue';
import { required } from '@vuelidate/validators';
import useVuelidate from '@vuelidate/core';
import type { Validation } from '@vuelidate/core';
import type { FormValidationStatus } from 'naive-ui/es/form/src/interface';
import { NSelect, NButton, NIcon, NInputNumber, NTag, useMessage } from 'naive-ui';
import type { SelectMixedOption } from 'naive-ui/es/select/src/interface';
import { Add } from '@vicons/ionicons5';

import type { ServiceItem, Service, ServiceSubService } from '@/packages/api/types/services';
import { fetchServiceSubService, fetchServiceItems, fetchSubServiceItems } from '@/packages/api/methods/service';
import { UnitType } from '@/packages/api/types/unit-type';
import { UnitTypePointer, UnitTypeTitle } from '@/packages/api/constatns/unit-type';

import type { AddServiceModalEmits, AddedItem, AddServiceModalProps } from './types';
import type { ItemType } from '@/packages/api/types/items';

const props = defineProps<AddServiceModalProps>();
const emit = defineEmits<AddServiceModalEmits>();

const message = useMessage();

const subServices = shallowRef<ServiceSubService[] | null>(null);
const serviceItems = shallowRef<ServiceItem[] | null>(null);
const addedItems = ref<AddedItem[]>([]);

const chosenServiceId = shallowRef<number | null>(null);
const chosenItemTypeId = shallowRef<number | null>(null);
const chosenSubServiceId = shallowRef<number | null>(null);
const chosenItemId = shallowRef<number | null>(null);
const chosenItemUnit = shallowRef<number | null>(null);

const addedItemsLen = computed<number | null>(() => addedItems.value.length > 0 ? addedItems.value.length : null);

const mainFormValidationRules = computed(() => {
  const rules = {
    chosenServiceId: { required },
    chosenItemTypeId: { required },
    chosenSubServiceId: {},
    addedItemsLen: {},
  };

  if (subServices.value) {
    rules.chosenSubServiceId = { required };
  }

  if (serviceItems.value) {
    rules.addedItemsLen = { required };
  }

  return rules;
});
const itemValidatorRules = {
  chosenItemId: { required },
  chosenItemUnit: { required },
};

const itemValidator = useVuelidate(itemValidatorRules, { chosenItemId, chosenItemUnit });
const mainValidator = useVuelidate(mainFormValidationRules, {
  chosenServiceId,
  chosenItemTypeId,
  chosenSubServiceId,
  addedItemsLen,
});

const chosenItemIdStatus = computed<FormValidationStatus | undefined>(() => valudationStatus(itemValidator.value, 'chosenItemId'));
const chosenItemUnitStatus = computed<FormValidationStatus | undefined>(() => valudationStatus(itemValidator.value, 'chosenItemUnit'));
const chosenServiceIdStatus = computed<FormValidationStatus | undefined>(() => valudationStatus(mainValidator.value, 'chosenServiceId'));
const chosenItemTypeIdStatus = computed<FormValidationStatus | undefined>(() => valudationStatus(mainValidator.value, 'chosenItemTypeId'));
const chosenSubServiceIdStatus = computed<FormValidationStatus | undefined>(() => valudationStatus(mainValidator.value, 'chosenSubServiceId'));
const addedItemsLenStatus = computed<FormValidationStatus | undefined>(() => valudationStatus(mainValidator.value, 'addedItemsLen'));

const serviceOptions = computed<SelectMixedOption[]>(() => props.services.map(service => ({
  label: service.name,
  value: service.id,
})));
const itemOptions = computed<SelectMixedOption[]>(() => props.itemType.map(item => ({
  label: item.name,
  value: item.id,
})));
const subServicesOptions = computed<SelectMixedOption[]>(() => {
  if (subServices.value) {
    return subServices.value.map(item => ({
      label: item.name,
      value: item.id,
    }));
  }

  return [];
});
const serviceItemOptions = computed<SelectMixedOption[]>(() => {
  if (serviceItems.value) {
    return serviceItems.value.map(item => ({
      label: item.item_name,
      value: item.id,
    }));
  }

  return [];
});

const serivceOptionsMap = computed<Map<number, Service>>(() => {
  const initialValue = new Map<number, Service>();
  return props.services.reduce((acc, cur) => {
    acc.set(cur.id, cur);
    return acc;
  }, initialValue);
});
const itemTypesMap = computed<Map<number, ItemType>>(() => {
  const initialValue = new Map<number, ItemType>();
  return props.itemType.reduce((acc, cur) => {
    acc.set(cur.id, cur);
    return acc;
  }, initialValue);
});
const subSerivceOptionsMap = computed<Map<number, ServiceSubService>>(() => {
  const initialValue = new Map<number, ServiceSubService>();
  if (subServices.value) {
    return subServices.value.reduce((acc, cur) => {
      acc.set(cur.id, cur);
      return acc;
    }, initialValue);
  }

  return initialValue;
});
const itemOptionsMap = computed<Map<number, ServiceItem>>(() => {
  const initialValue = new Map<number, ServiceItem>();
  if (serviceItems.value) {
    return serviceItems.value.reduce((acc, cur) => {
      acc.set(cur.id, cur);
      return acc;
    }, initialValue);
  }

  return initialValue;
});

const chosenServiceUnit = computed<UnitType>(() => serivceOptionsMap.value.get(chosenServiceId.value || 0)?.unit_id || UnitType.UnitPCS);
const itemUnitPrecision = computed<number>(() => chosenServiceUnit.value === UnitType.UnitPCS ? 0 : 2);

async function onUpdateServiceOption(id: number): Promise<void> {
  itemValidator.value.$reset();

  subServices.value = null;
  chosenItemId.value = null;
  chosenItemUnit.value = null;
  addedItems.value = [];

  const chosenService = serivceOptionsMap.value.get(id);
  if (!chosenService) {
    return;
  }

  if (chosenService.has_sub_service) {
    subServices.value = await fetchServiceSubService(id);
    return;
  }

  await fetchItems(id);
}

async function fetchItems(id: number, isSub: boolean = false): Promise<void> {
  if (isSub) {
    serviceItems.value = await fetchSubServiceItems(id);
    return;
  }

  serviceItems.value = await fetchServiceItems(id);
}

async function addItem() {
  const valid = await itemValidator.value.$validate();
  if (!valid) {
    return;
  }

  const itemToAdd: ServiceItem | undefined = itemOptionsMap.value.get(chosenItemId.value || 0);
  if (!itemToAdd) {
    return;
  }

  const hasIndex = addedItems.value.findIndex(value => value.item.id === itemToAdd.id);
  if (hasIndex >= 0) {
    addedItems.value[hasIndex].quantity += chosenItemUnit.value!;
  } else {
    addedItems.value = [...addedItems.value, {
      item: itemToAdd,
      quantity: chosenItemUnit.value!,
    }];
  }

  itemValidator.value.$reset();
  chosenItemId.value = null;
  chosenItemUnit.value = null;
}

function filterAddedItem(id: number): void {
  addedItems.value = addedItems.value.filter(value => value.item.id !== id);
}

async function addService(): Promise<void> {
  const valid = await mainValidator.value.$validate();
  if (!valid) {
    if (addedItemsLenStatus.value === 'error') {
      message.error('Добавьте вещи', {
        duration: 3000,
      });
    }

    return;
  }

  const chosenService = serivceOptionsMap.value.get(chosenServiceId.value!);
  const itemType = itemTypesMap.value.get(chosenItemTypeId.value!);
  const chosenSubService = subSerivceOptionsMap.value.get(chosenSubServiceId.value || 0);
  if (!chosenService || !itemType) {
    return;
  }

  emit('onAdd', {
    chosenService,
    chosenSubService,
    addedItems: addedItems.value,
    itemType,
  });
}

function valudationStatus<T extends Validation>(validator: T, key: keyof T): FormValidationStatus | undefined {
  if (validator[key].$invalid && validator.$dirty) {
    return 'error';
  }
  return undefined;
}
</script>

<template>
  <div class="chose-service-modal">
    <h3 class="chose-service-modal__title">
      Добавить услугу
    </h3>

    <div class="chose-service-modal__body">
      <div class="chose-service-modal__body-item chose-service-modal__body-item-service">
        <h4>Выберите услугу</h4>
        <div class="chose-service-modal__body-item-service-selects">
          <NSelect
            v-model:value="chosenServiceId"
            class="chose-service-modal__body-item-service-selects__service"
            placeholder="Услуга"
            :options="serviceOptions"
            :status="chosenServiceIdStatus"
            @update:value="onUpdateServiceOption"
          />
          <NSelect
            v-model:value="chosenItemTypeId"
            class="chose-service-modal__body-item-service-selects__item-type"
            placeholder="Тип одежды"
            :options="itemOptions"
            :status="chosenItemTypeIdStatus"
          />
        </div>
      </div>

      <div
        v-if="subServices"
        class="chose-service-modal__body-item chose-service-modal__body-item-service"
      >
        <h4>Выберите под услугу</h4>
        <div class="chose-service-modal__body-item-service-selects">
          <NSelect
            v-model:value="chosenSubServiceId"
            class="chose-service-modal__body-item-service-selects__service"
            placeholder="Услуга"
            :options="subServicesOptions"
            :status="chosenSubServiceIdStatus"
            @update:value="fetchItems($event, true)"
          />
        </div>
      </div>

      <div
        v-if="serviceItems"
        class="chose-service-modal__body-item"
      >
        <h4>Добавьте вещи</h4>
        <div class="chose-service-modal__body-item-service-selects">
          <NSelect
            v-model:value="chosenItemId"
            class="chose-service-modal__body-item-service-selects__service"
            placeholder="Вещь"
            :options="serviceItemOptions"
            :status="chosenItemIdStatus"
          />

          <NInputNumber
            v-model:value="chosenItemUnit"
            :show-button="false"
            :precision="itemUnitPrecision"
            :placeholder="UnitTypeTitle[chosenServiceUnit]"
            :status="chosenItemUnitStatus"
          />

          <NButton
            tertiary
            type="primary"
            @click="addItem"
          >
            <template #icon>
              <NIcon>
                <Add />
              </NIcon>
            </template>
          </NButton>
        </div>

        <div
          v-if="addedItems.length > 0"
          class="chose-service-modal__body-item"
        >
          <h4>Добавленные вещи</h4>
          <div class="chose-service-modal__added-items-list">
            <NTag
              v-for="value in addedItems"
              :key="value.item.id"
              class="chose-service-modal__added-items-list__value"
              type="primary"
              closable
              @close="filterAddedItem(value.item.id)"
            >
              <span>{{ value.item.item_name }}</span>
              <span>{{ `${value.quantity}${UnitTypePointer[chosenServiceUnit]}` }}</span>
            </NTag>
          </div>
        </div>
      </div>
    </div>

    <div class="chose-service-modal__footer">
      <NButton
        type="primary"
        @click="addService"
      >
        <template #icon>
          <NIcon>
            <Add />
          </NIcon>
        </template>
        Добавить
      </NButton>
    </div>
  </div>
</template>

<style lang="scss" src="./style.scss" />
