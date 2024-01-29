<template>
  <a-tabs v-model:activeKey="activeKey">
    <a-tab-pane key="1" tab="隐藏信息">
       <a-form
      ref="formRef1"
      :model="formState1"
      :rules="rules"
      :label-col="labelCol"
      :wrapper-col="wrapperCol"
    >
      <a-form-item label="载体字符串：" name="carrierStr">
        <a-input v-model:value="formState1.carrierStr" />
      </a-form-item>
      <a-form-item label="隐藏字符串：" name="hideStr">
        <a-input v-model:value="formState1.hideStr" />
      </a-form-item>
      <a-form-item :wrapper-col="{ span: 14, offset: 4 }">
        <a-button type="primary" @click="onSubmit1">确认</a-button>
        <a-button style="margin-left: 10px" @click="resetForm1">重置</a-button>
      </a-form-item>
        <a-descriptions title="隐藏信息" column="1">
    <a-descriptions-item >{{hideresultStr1}}</a-descriptions-item>
  </a-descriptions>
    </a-form>
    </a-tab-pane>
    <a-tab-pane key="2" tab="提取隐藏信息" force-render>
      <a-form
      ref="formRef2"
      :model="formState2"
      :rules="rules"
      :label-col="labelCol"
      :wrapper-col="wrapperCol"
    >
      <a-form-item label="隐藏字符串：" name="hideStr">
        <a-input v-model:value="formState2.hideStr" />
      </a-form-item>
      <a-form-item :wrapper-col="{ span: 14, offset: 4 }">
        <a-button type="primary" @click="onSubmit2">确认</a-button>
        <a-button style="margin-left: 10px" @click="resetForm2">重置</a-button>
      </a-form-item>
        <a-descriptions title="隐藏后的字符串" column="1">
    <a-descriptions-item >{{hideresultStr2}}</a-descriptions-item>
  </a-descriptions>
    </a-form>
      
       </a-tab-pane>
  </a-tabs>
  
</template>
<script>
import { defineComponent, ref, toRaw, reactive } from "vue";
import http from "../../http"
export default defineComponent({
  setup() {
    const formRef1 = ref();
    const formRef2 = ref();
    const hideresultStr1 = ref("NUll");
    const hideresultStr2 = ref("NUll");
    const formState1 = reactive({
      carrierStr: "",
      hideStr:""
    });
    const formState2 = reactive({
      carrierStr: "",
      hideStr:""
    });
    const rules = {
      carrierStr: {
        required: true,
        message: '请输入载体字符串',
      },
      hideStr: {
        required: true,
        message: '请输入需要隐藏的字符串',
      },
    };
    const onSubmit1 = () => {
      formRef1.value
        .validate()
        .then(() => {        
          var dataStr={
            carrier_str:formState1.carrierStr,
            hide_str:formState1.hideStr,
            optiontype:"1",
            hide_result_str:"",
          }
          http.post("/apis/hide",dataStr).then(function(response){
            hideresultStr1.value = response["data"]["data"]["hide_result_str"]
            console.log("fa",response["data"]["data"]["hide_result_str"])
          })
          console.log('values', formState1, toRaw(formState1),dataStr);
        })
        .catch(error => {
          console.log('error', error);
        });
    };
    const resetForm1 = () => {
      formRef1.value.resetFields();
    };
    const onSubmit2 = () => {
      formRef2.value
        .validate()
        .then(() => {        
          var dataStr={
            carrier_str:"",
            hide_str:"",
            optiontype:"2",
            hide_result_str:formState2.hideStr,
          }
          http.post("/apis/hide",dataStr).then(function(response){
            hideresultStr2.value = response["data"]["data"]["hide_str"]
            console.log("fa",response)
          })
          console.log('values', formState2, toRaw(formState2),dataStr);
        })
        .catch(error => {
          console.log('error', error);
        });
    };
    const resetForm2 = () => {
      formRef2.value.resetFields();
    };
    return {
      formRef1,
      formState1,
      formRef2,
      formState2,
      rules,
      hideresultStr1,
      hideresultStr2,
      resetForm1,
      onSubmit1,
      resetForm2,
      onSubmit2,
      activeKey: ref("1"),
      labelCol: {
        span: 4,
      },
      wrapperCol: {
        span: 14,
      },
      
    };
  },
});
</script>