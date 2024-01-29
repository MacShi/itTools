<template>
  <div class="testBox">
    <a-textarea v-model:value="userInput" placeholder="输入 JSON 字符串" />
    <a-divider />
    <a-button type="primary" @click="formatJson">格式化</a-button>
    <!-- <div class="box" v-if="formattedJson">{{ formattedJson }}</div> -->
    <div class="box" v-if="formattedJson">
      <json-viewer :value="formattedJson"></json-viewer>
    </div>
  </div>
</template>

<script>
import { message } from 'ant-design-vue';
import JsonViewer from 'vue-json-viewer';

export default {
  components: {
    JsonViewer,
  },
  data() {
    return {
      userInput: '',
      formattedJson: null,
    };
  },
  methods: {
    formatJson() {
      try {
        this.formattedJson = JSON.stringify(JSON.parse(this.userInput), null, 2);
      } catch (error) {
        console.error('Invalid JSON format:', error.message);
        this.formattedJson = null;
        message.error(error.message)
      }
    },
  },
};
</script>
