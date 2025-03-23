<script lang="ts" setup>
import { RouteLinks } from './constants';
import { DryCleaningTwotone } from '@vicons/material';
</script>

<template>
  <div class="app-page-header">
    <div class="app-page-header__logo">
      <DryCleaningTwotone />
      <div class="app-page-header__logo-text">
        Laundry
      </div>
    </div>

    <div class="app-page-header__links">
      <RouterLink
        v-for="link in RouteLinks"
        :key="link.path"
        :to="link.path"
        class="app-page-header__link"
      >
        {{ link.label }}
      </RouterLink>
    </div>
  </div>
</template>

<style lang="scss" src="./style.scss" />
