<script lang="ts" setup>
import { getTodayOrders } from '@/packages/api/methods/orders';
import type { Order } from '@/packages/api/types/orders';
import { NNumberAnimation, NTable, NEmpty } from 'naive-ui';
import { computed, onBeforeMount, shallowRef } from 'vue';
import type { TodayOrdersCommonSum } from './types';

const ordersList = shallowRef<Order[] | null>(null);

const totalReduced = computed<TodayOrdersCommonSum | null>(() => {
  if (!ordersList.value) {
    return null;
  }

  const reduced = ordersList.value.reduce((acc, cur) => {
    acc.final += cur.final;
    acc.total += cur.total;

    return acc;
  }, {
    total: 0,
    final: 0,
  });

  return {
    discounts: reduced.total - reduced.final,
    final: reduced.final,
  };
});

async function getOrders() {
  ordersList.value = await getTodayOrders();
}

onBeforeMount(getOrders);
</script>

<template>
  <div class="today-orders-view">
    <div class="today-orders-view__head">
      <h2>Заказы за сегодня</h2>

      <div
        v-if="totalReduced"
        class="today-orders-view__reduced"
      >
        <div class="today-orders-view__reduced-item">
          <span class="text">Скидки: </span>
          <span class="value">
            <NNumberAnimation
              :from="0"
              :to="totalReduced.discounts"
              :precision="1"
            />
          </span>
        </div>

        <div class="today-orders-view__reduced-item">
          <span class="text">Общая выручка: </span>
          <span class="value">
            <NNumberAnimation
              :from="0"
              :to="totalReduced.final"
              :precision="1"
            />
          </span>
        </div>
      </div>
    </div>

    <div class="today-orders-view__body">
      <NEmpty v-if="!ordersList">
        Заказы не найдены
      </NEmpty>

      <div
        v-else
        class="today-orders-view__order-list"
      >
        <NTable>
          <thead>
            <tr>
              <th>№</th>
              <th>Имя пользователя</th>
              <th>Номер телефона</th>
              <th>Общая сумма</th>
              <th>Итоговая сумма</th>
            </tr>
          </thead>

          <tbody>
            <tr v-for="order in ordersList" :key="order.id">
              <td>{{ order.id }}</td>
              <td>{{ order.user_name }}</td>
              <td>{{ order.phone_number }}</td>
              <td>{{ order.total }}</td>
              <td>{{ order.final }}</td>
            </tr>
          </tbody>
        </NTable>
      </div>
    </div>
  </div>
</template>

<style lang="scss" src="./style.scss" />
