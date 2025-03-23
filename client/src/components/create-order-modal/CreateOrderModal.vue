<script lang="ts" setup>
import { computed, shallowReactive, toRaw } from 'vue';
import { required, minLength, numeric } from '@vuelidate/validators';
import useVuelidate from '@vuelidate/core';
import { NButton, NIcon, NInput } from 'naive-ui';

import { valudationStatus } from '@/tools/validator';
import type { CreateOrderParam } from '@/packages/api/types/orders';

import type { CreateOrderModalEmits } from './types';
import { CreateOutline } from '@vicons/ionicons5';

const emit = defineEmits<CreateOrderModalEmits>();

const userParams = shallowReactive<CreateOrderParam>({
  user_name: '',
  user_phone_number: '',
});

const rules = {
  user_name: { required, minLength: minLength(2) },
  user_phone_number: { required, numeric, minLength: minLength(7) },
};

const validator = useVuelidate(rules, userParams);

const userNameStatus = computed(() => valudationStatus(validator.value, 'user_name'));
const userPhoneStatus = computed(() => valudationStatus(validator.value, 'user_phone_number'));

async function createOrder() {
  const valid = await validator.value.$validate();
  if (!valid) {
    return;
  }

  emit('onCreate', Object.assign({}, toRaw(userParams)));

  validator.value.$reset();
  userParams.user_name = '';
  userParams.user_phone_number = '';
}
</script>

<template>
  <div class="create-order-modal">
    <h3 class="create-order-modal__title">
      Создать заказ
    </h3>

    <div class="create-order-modal__body">
      <NInput
        v-model:value="userParams.user_name"
        placeholder="Имя пользователя"
        :status="userNameStatus"
      />
      <NInput
        v-model:value="userParams.user_phone_number"
        placeholder="Номер телефона"
        :status="userPhoneStatus"
      />
    </div>

    <div class="create-order-modal__footer">
      <NButton type="primary" @click="createOrder">
        <template #icon>
          <NIcon>
            <CreateOutline />
          </NIcon>
        </template>
        Создать
      </NButton>
    </div>
  </div>
</template>

<style lang="scss" src="./style.scss" />
