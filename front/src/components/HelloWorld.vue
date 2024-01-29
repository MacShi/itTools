<template>
  <a-calendar v-model:value="value" @panelChange="onPanelChange" />
</template>
<script>
import { defineComponent, ref } from 'vue';
export default defineComponent({
  setup() {
    const value = ref();
    const onPanelChange = (value, mode) => {
      console.log(value, mode);
    };
    return {
      value,
      onPanelChange,
    };
  },
});
</script>