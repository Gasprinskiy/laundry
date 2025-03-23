<script lang="ts" setup>
import type { SelectMixedOption } from 'naive-ui/es/select/src/interface';
import { NButton, NIcon, NEmpty, NCard, NScrollbar, NTag, NSelect, NTable, NNumberAnimation, useMessage } from 'naive-ui';
import { Add, CalculatorOutline, CreateOutline, RemoveSharp } from '@vicons/ionicons5';
import { computed, defineAsyncComponent, onBeforeMount, shallowRef } from 'vue';

import type { ServicesCommonResponse } from '@/packages/api/types/services';
import { fetchAllAbleServices } from '@/packages/api/methods/service';
import { useModal } from '@/composables/use-modal';
import type { AddServiceModalEmits, AddServiceModalProps, AddServiceParams } from '@/components/add-service-modal/types';
import type { ConvertEmitType } from '@/tools/types';
import { UnitTypePointer } from '@/packages/api/constatns/unit-type';
import type { CalculateOrderService, CalculateOrderItem, CalculateOrderResponse, CreateOrderParam } from '@/packages/api/types/orders';
import { calculateOrder, createOrder } from '@/packages/api/methods/orders';
import { ItemTypeTitle } from '@/packages/api/constatns/item-type';
import type { CreateOrderModalEmits } from '@/components/create-order-modal/types';

const { showModal, closeModal } = useModal();
const message = useMessage();

const serivceCommonItems = shallowRef<ServicesCommonResponse | null>(null);
const addedServices = shallowRef<AddServiceParams[]>([]);
const caclResult = shallowRef<CalculateOrderResponse | null>(null);
const fulfillmentId = shallowRef<number>(2);
const openAccordion = shallowRef<number | null>(null);

const fullfilmentOptions = computed<SelectMixedOption[]>(() => {
  if (serivceCommonItems.value) {
    return serivceCommonItems.value.fulfillment_types.map(service => ({
      label: service.name,
      value: service.id,
    }));
  }

  return [];
});

async function fetchCommonServiceItems(): Promise<void> {
  serivceCommonItems.value = await fetchAllAbleServices();
}

function addService(item: AddServiceParams) {
  // TO DO make merge service items
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
  closeModal();
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

async function calcOrder() {
  if (!serivceCommonItems.value) {
    return;
  }

  const services: CalculateOrderService[] = addedServices.value.map((service) => {
    const items: CalculateOrderItem[] = service.addedItems.map(value => ({
      id: value.item.id,
      quantity: value.quantity,
    }));

    return {
      service_id: service.chosenService.id,
      service_name: service.chosenService.name,
      subservice_id: service.chosenSubService?.id,
      subservice_name: service.chosenSubService?.name,
      items,
      unit_id: service.chosenService.unit_id,
      item_type_id: service.itemType.id,
    };
  });

  const fulfillment = serivceCommonItems.value.fulfillment_types.find(item => item.id === fulfillmentId.value);
  if (!fulfillment) {
    return;
  }

  caclResult.value = await calculateOrder({
    fulfillment,
    services,
  });

  fulfillmentId.value = 2;
  addedServices.value = [];
}

async function onOrderCreate(param: CreateOrderParam) {
  if (!caclResult.value) {
    return;
  }

  try {
    const created = await createOrder(caclResult.value.temporary_id, param);

    message.info(`Заказ №${created.id} успешно создан`, {
      duration: 3000,
    });
    caclResult.value = null;
  } catch (e) {
    message.error('Не удалось создать заказ', {
      duration: 3000,
    });
    console.error(e);
  } finally {
    closeModal();
  }
}

function openCreateOrderModal() {
  const component = defineAsyncComponent(() => import('@/components/create-order-modal/CreateOrderModal.vue'));

  const emits: ConvertEmitType<CreateOrderModalEmits> = {
    onCreate: onOrderCreate,
  };

  showModal({
    component,
    emits,
  });
}

function toggleAccordion(index: number) {
  openAccordion.value = openAccordion.value === index ? null : index;
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
        <NSelect
          v-model:value="fulfillmentId"
          placeholder="Скорость выполнения"
          :options="fullfilmentOptions"
        />
        <NButton
          type="primary"
          :disabled="addedServices.length <= 0"
          @click="calcOrder"
        >
          <template #icon>
            <NIcon>
              <CalculatorOutline />
            </NIcon>
          </template>
          Расчитать стоимость
        </NButton>
      </div>
    </div>

    <div
      class="home-view__divider"
      vertical
    />

    <div class="home-view__left">
      <h2>Итоговый результат</h2>

      <div class="home-view__result">
        <NEmpty v-if="!caclResult">
          Итоговый результат пуст
        </NEmpty>

        <NScrollbar
          v-else
          class="home-view__result_scrollbar"
        >
          <div class="home-view__result_data">
            <div class="home-view__result_tables">
              <h3>Услуги</h3>
              <NTable>
                <thead>
                  <tr>
                    <th>Услуга</th>
                    <th>Подуслуга</th>
                    <th>Тип одежды</th>
                    <th>Общая сумма</th>
                    <th>Итоговая сумма</th>
                    <th>Подробнее</th>
                  </tr>
                </thead>

                <tbody>
                  <template
                    v-for="(service, index) in caclResult.order_services"
                    :key="index"
                  >
                    <tr>
                      <td>{{ service.service_name }}</td>
                      <td>{{ service.subservice_name || '-' }}</td>
                      <td>{{ ItemTypeTitle[service.items_type_id] }}</td>
                      <td>{{ service.total }}</td>
                      <td>{{ service.final }}</td>
                      <td>
                        <NButton size="small" @click="toggleAccordion(index)">
                          {{ openAccordion === index ? 'Скрыть' : 'Показать' }}
                        </NButton>
                      </td>
                    </tr>

                    <tr v-if="openAccordion === index" style="width: 100%;">
                      <td
                        style="width: 100%;"
                        colspan="10"
                      >
                        <div>
                          <h4>Список вещей</h4>
                          <NTable>
                            <thead>
                              <tr>
                                <th>Название вещи</th>
                                <th>Количество</th>
                                <th>Цена за единицу</th>
                                <th>Общая цена</th>
                              </tr>
                            </thead>

                            <tbody>
                              <tr v-for="item in service.items" :key="item.id">
                                <td>{{ item.item_name }}</td>
                                <td>{{ item.quantity }}{{ UnitTypePointer[service.unit_id] }}</td>
                                <td>{{ item.price_for_one }}</td>
                                <td>{{ item.price_for_all }}</td>
                              </tr>
                            </tbody>
                          </NTable>
                        </div>

                        <div v-if="service.discounts">
                          <h4>Скидки</h4>
                          <NTable>
                            <thead>
                              <tr>
                                <th>Описание</th>
                                <th>Процент</th>
                              </tr>
                            </thead>

                            <tbody>
                              <tr v-for="discount in service.discounts" :key="discount.modifier_id">
                                <td>{{ discount.description }}</td>
                                <td>{{ discount.percent }}%</td>
                              </tr>
                            </tbody>
                          </NTable>
                        </div>

                        <div v-if="service.markups">
                          <h4>Наценки</h4>
                          <NTable>
                            <thead>
                              <tr>
                                <th>Описание</th>
                                <th>Процент</th>
                              </tr>
                            </thead>

                            <tbody>
                              <tr v-for="markup in service.markups" :key="markup.modifier_id">
                                <td>{{ markup.description }}</td>
                                <td>{{ markup.percent }}%</td>
                              </tr>
                            </tbody>
                          </NTable>
                        </div>
                      </td>
                    </tr>
                  </template>
                </tbody>
              </NTable>
            </div>

            <div>
              <div
                v-if="caclResult.discounts"
                class="home-view__result_tables"
              >
                <h3>Общие скидки</h3>
                <NTable>
                  <thead>
                    <tr>
                      <th>Описание</th>
                      <th>Процент</th>
                    </tr>
                  </thead>

                  <tbody>
                    <tr v-for="discount in caclResult.discounts" :key="discount.modifier_id">
                      <td>{{ discount.description }}</td>
                      <td>{{ discount.percent }}</td>
                    </tr>
                  </tbody>
                </NTable>
              </div>

              <div
                v-if="caclResult.markups"
                class="home-view__result_tables"
              >
                <h3>Общие наценки</h3>
                <NTable>
                  <thead>
                    <tr>
                      <th>Описание</th>
                      <th>Процент</th>
                    </tr>
                  </thead>

                  <tbody>
                    <tr v-for="markup in caclResult.markups" :key="markup.modifier_id">
                      <td>{{ markup.description }}</td>
                      <td>{{ markup.percent }}</td>
                    </tr>
                  </tbody>
                </NTable>
              </div>
            </div>

            <div class="home-view__result_tables">
              <h3>Общие сведения</h3>

              <NTable>
                <tbody>
                  <tr>
                    <td>Скорость выполнения</td>
                    <td class="home-view__result-common-td">
                      <NTag type="primary">
                        {{ caclResult.fulfillment.name }}
                      </NTag>
                    </td>
                  </tr>
                  <tr>
                    <td>Общая сумма</td>
                    <td class="home-view__result-common-td number-animation">
                      <NNumberAnimation
                        :from="0"
                        :to="caclResult.total"
                        :precision="1"
                      />
                    </td>
                  </tr>
                  <tr>
                    <td>Итоговая сумма</td>
                    <td class="home-view__result-common-td number-animation">
                      <NNumberAnimation
                        :from="0"
                        :to="caclResult.final"
                        :precision="1"
                      />
                    </td>
                  </tr>
                </tbody>
              </NTable>
            </div>
          </div>
        </NScrollbar>
      </div>

      <div class="home-view__left-footer">
        <NButton
          type="primary"
          :disabled="caclResult === null"
          @click="openCreateOrderModal"
        >
          <template #icon>
            <NIcon>
              <CreateOutline />
            </NIcon>
          </template>
          Создать заказ
        </NButton>
      </div>
    </div>
  </div>
</template>

<style lang="scss" src="./style.scss" />
