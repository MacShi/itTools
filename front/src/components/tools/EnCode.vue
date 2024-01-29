<template>
  <a-form :model="formState" :label-col="labelCol" :wrapper-col="wrapperCol">
    <a-form-item label="类型">
      <a-radio-group v-model:value="formState.type">
        <a-radio value="1">Base64</a-radio>
        <a-radio value="2">UrlCode</a-radio>
      </a-radio-group>
    </a-form-item>
    <a-form-item label="待编码或待解码">
      <a-textarea v-model:value="formState.desc1" />
    </a-form-item>
    <a-form-item :wrapper-col="{ span: 14, offset: 4 }">
      <a-button type="primary" @click="enCode">编码</a-button>
      <a-button style="margin-left: 10px" @click="dencode">解码</a-button>
    </a-form-item>
    <a-form-item label="编码或解码结果">
      <a-textarea v-model:value="formState.desc2" />
    </a-form-item>
  </a-form>
</template>
<script>
import { defineComponent, reactive, toRaw } from 'vue';
import { message } from 'ant-design-vue';
export default defineComponent({
  setup() {
    const formState = reactive({
      type: "1",
      desc1: '',
      desc2: '',
    });
    const enCode = () => {
      if (!formState.desc1.trim()) {
        // 如果 desc1 为空，弹出错误提示
        message.error('请输入待编码或待解码的内容');
        // return;
      }

      if (formState.type==1) {
        formState.desc2 = btoa(formState.desc1);
      } else if (formState.type==2) {
        formState.desc2 = encodeURIComponent(formState.desc1);
      }

      console.log('submit!', toRaw(formState));
    };
    const dencode =() =>{
      if (!formState.desc1.trim()) {
        // 如果 desc1 为空，弹出错误提示
        message.error('请输入待编码或待解码的内容');
        // return;
      }

      if (formState.type==1) {
        formState.desc2 = atob(formState.desc1);
      } else if (formState.type==2) {
        formState.desc2 = decodeURIComponent(formState.desc1);
      }

    };
    // const isBase64 = (str) => {
    //   try {
    //     return btoa(atob(str)) === str;
    //   } catch (err) {
    //     return false;
    //   }
    // };

    // const isUrlEncoded = (str) => {
    //   try {
    //     return decodeURIComponent(str) !== str;
    //   } catch (err) {
    //     return false;
    //   }
    // };
    return {
      labelCol: {
        style: {
          width: '150px',
        },
      },
      wrapperCol: {
        span: 14,
      },
      formState,
      enCode,
      dencode,
    };
  },
});
</script>