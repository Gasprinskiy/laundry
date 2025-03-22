<script setup lang="ts">
import { $api } from "@/packages/api/worker"
import { onMounted } from "vue";

async function suck() {
  const response = await $api('/services')
  console.log('response: ', response);
  
}

onMounted(() => {
  console.log("import.meta: ",import.meta.env);
})
</script>

<template>
  <button @click="suck">SUCK SOME</button>
</template>
