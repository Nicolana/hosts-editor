<script lang="ts" setup>
import { ref } from "vue";

const link = ref("https://base64.us/");
</script>

<template>
  <iframe class="base-page" :src="link" frameborder="0"></iframe>
</template>

<style scoped>
.base-page {
  width: 600px;
  height: 500px;
}
</style>
