<script lang="ts" setup>
import { NButton, NIcon } from 'naive-ui';
import { Close } from '@vicons/ionicons5';

import { useModal } from '@/composables/use-modal';

const { isVisible, state, closeModal } = useModal();
</script>

<template>
  <div
    v-if="isVisible"
    class="modal-provier"
  >
    <div class="modal-provier__wrapper">
      <NButton
        class="modal-provier__close-btn"
        type="error"
        quaternary
        circle
        @click="closeModal"
      >
        <template #icon>
          <NIcon size="30">
            <Close />
          </NIcon>
        </template>
      </NButton>
      <component
        :is="state.component"
        v-bind="state.props"
        v-on="state.emits"
      />
    </div>
  </div>
</template>

<style lang="scss" src="./style.scss" />
