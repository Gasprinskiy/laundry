<script lang="ts" setup>
import { NButton, NIcon, NEmpty, NCard, NScrollbar, NTag } from 'naive-ui';
import { Add, CalculatorOutline, RemoveSharp } from '@vicons/ionicons5';
import { defineAsyncComponent, onBeforeMount, shallowRef } from 'vue';

import type { ServicesCommonResponse } from '@/packages/api/types/services';
import { fetchAllAbleServices } from '@/packages/api/methods/service';
import { useModal } from '@/composables/use-modal';
import type { AddServiceModalEmits, AddServiceModalProps, AddServiceParams } from '@/components/add-service-modal/types';
import type { ConvertEmitType } from '@/tools/types';
import { UnitTypePointer } from '@/packages/api/constatns/unit-type';

const { showModal } = useModal();

const shit = {
  chosenService: { id: 1, name: 'Химчистка', unit_id: 2, has_sub_service: false },
  addedItems: [{ item: { id: 103, item_id: 1, item_name: 'Пальто' }, quantity: 1 }],
  itemType: { id: 2, name: 'Взрослая', modifier_id: null },
};

const serivceCommonItems = shallowRef<ServicesCommonResponse | null>(null);
const addedServices = shallowRef<AddServiceParams[]>([shit]);

async function fetchCommonServiceItems(): Promise<void> {
  const response = await fetchAllAbleServices();
  serivceCommonItems.value = response;
}

function addService(item: AddServiceParams) {
  // const has = addedServices.value.findIndex((value) => {
  //   const valueServiceId = value.chosenService.id;
  //   const subServiceId = value.chosenSubService?.id;
  //   const itemTypeId = value.itemType.id;

  //   const serviceIdEqual = valueServiceId === item.chosenService.id;
  //   const itemTypeIdEqual = itemTypeId === item.itemType.id;

  //   let subServiceIdEqual = false;
  //   if (subServiceId && item.chosenSubService?.id) {
  //     subServiceIdEqual = subServiceId === item.chosenSubService.id;
  //   }

  //   return serviceIdEqual && subServiceIdEqual && itemTypeIdEqual;
  // });

  // if (has >= 0) {
  //   const hasItem = addedServices.value.
  // }

  addedServices.value = [...addedServices.value, item];
}

function filterAddedServices(itemIndex: number) {
  addedServices.value = addedServices.value.filter((_, index) => index !== itemIndex);
}

function openAddServiceModal() {
  if (!serivceCommonItems.value) {
    return;
  }

  const component = defineAsyncComponent(() => import('@/components/add-service-modal/AddServiceModal.vue'));

  const props: AddServiceModalProps = {
    services: serivceCommonItems.value.services,
    itemType: serivceCommonItems.value.item_types,
  };

  const emits: ConvertEmitType<AddServiceModalEmits> = {
    onAdd: addService,
  };

  showModal({
    component,
    props,
    emits,
  });
}

onBeforeMount(fetchCommonServiceItems);
</script>

<template>
  <div class="home-view">
    <div class="home-view__right">
      <div class="home-view__head">
        <h2>Оформление заказа</h2>
        <NButton type="primary" @click="openAddServiceModal">
          <template #icon>
            <NIcon>
              <Add />
            </NIcon>
          </template>
          Добавить услугу
        </NButton>
      </div>

      <div class="home-view__body">
        <NEmpty v-if="addedServices.length <= 0">
          Нет добавленных услуг
        </NEmpty>

        <NScrollbar
          v-else
          class="home-view__body_scrollbar"
        >
          <NCard
            v-for="(service, index) in addedServices"
            :key="index"
            :title="service.chosenService.name"
          >
            <template #header-extra>
              <NButton
                quaternary
                type="error"
                @click="filterAddedServices(index)"
              >
                <template #icon>
                  <NIcon>
                    <RemoveSharp />
                  </NIcon>
                </template>
                Удалить
              </NButton>
            </template>
            <div class="home-view__service-card">
              <div class="home-view__service-card__right">
                <div
                  v-if="service.chosenSubService"
                  class="home-view__service-card__item"
                >
                  <div class="home-view__service-card__item-title">
                    Подуслуга:
                  </div>
                  <NTag type="primary">
                    {{ service.chosenSubService.name }}
                  </NTag>
                </div>

                <div class="home-view__service-card__item">
                  <div class="home-view__service-card__item-title">
                    Тип одежды:
                  </div>
                  <NTag type="primary">
                    {{ service.itemType.name }}
                  </NTag>
                </div>

                <div class="home-view__service-card__item">
                  <div class="home-view__service-card__item-title">
                    Вещи:
                  </div>
                  <NTag
                    v-for="value in service.addedItems"
                    :key="value.item.id"
                    class="home-view__service-card__item-tag"
                    type="primary"
                  >
                    <span>{{ value.item.item_name }}</span>
                    <span>{{ `${value.quantity}${UnitTypePointer[service.chosenService.unit_id]}` }}</span>
                  </NTag>
                </div>
              </div>
            </div>
          </NCard>
        </NScrollbar>
      </div>

      <div class="home-view__right-footer">
        <NButton type="primary">
          <template #icon>
            <NIcon>
              <CalculatorOutline />
            </NIcon>
          </template>
          Расчитать стоимость
        </NButton>
      </div>
    </div>
  </div>
</template>

<style lang="scss" src="./style.scss" />
