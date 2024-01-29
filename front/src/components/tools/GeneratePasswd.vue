<template>
  <div>
    <a-checkbox
      v-model:checked="checkAll"
      :indeterminate="indeterminate"
      @change="onCheckAllChange"
    >
      所有
    </a-checkbox>
  </div>
  <a-divider />
  <a-row type="flex">
    <a-col :flex="2">
      字符类型：
      <a-checkbox-group v-model:value="checkedList" :options="plainOptions" />
    </a-col>
    <a-col :flex="5"></a-col>
  </a-row>
  <br />
  <a-row>
    <a-col align="bottom">
      密码长度：
      <a-input-number
        id="inputNumber"
        v-model:value="pwdLength"
        :min="8"
        :max="32"
      />
    </a-col>
  </a-row>
  <br />
  <a-row>
    <a-col>
      <a-button type="submit" @click="genPwd">生成</a-button>
    </a-col>
  </a-row>
  <a-divider />
  <a-descriptions title="随机密码" column="1">
    <a-descriptions-item>{{ randomStr }}</a-descriptions-item>
  </a-descriptions>
</template>
<script>
import { message } from "ant-design-vue";
import { defineComponent, reactive, toRefs, watch, ref } from "vue";
import http from "../../http"
const plainOptions = ["小写字母", "大写字母", "数字", "特殊字符"];
export default defineComponent({
  setup() {
    const randomStr = ref("Null");
    const pwdLength = ref(8);
    const state = reactive({
      indeterminate: true,
      checkAll: false,
      checkedList: ["小写字母", "大写字母", "数字"],
    });
    const onCheckAllChange = (e) => {
      Object.assign(state, {
        checkedList: e.target.checked ? plainOptions : [],
        indeterminate: false,
      });
    };
    watch(
      () => state.checkedList,
      (val) => {
        state.indeterminate = !!val.length && val.length < plainOptions.length;
        state.checkAll = val.length === plainOptions.length;
      }
    );
    const genPwd = () => {
      if (state.checkedList.length < 1) {
        message.error("至少选择1个字符类型");
        console.log("fads");
      } else {
        var sum = 0;
        //'小写字母', '大写字母', '数字','特殊字符'
        if (state.checkedList.indexOf("小写字母") > -1) {
          sum = sum + 8;
        }
        if (state.checkedList.indexOf("大写字母") > -1) {
          sum = sum + 4;
        }
        if (state.checkedList.indexOf("数字") > -1) {
          sum = sum + 2;
        }
        if (state.checkedList.indexOf("特殊字符") > -1) {
          sum = sum + 1;
        }
        // var a =state.checkedList.indexOf("大写字母")
        console.log(pwdLength.value, state.checkedList, sum);
        var dataJson = {
          pwd_length: pwdLength.value,
          pwd_type: sum,
        };
        http.post("/apis/getPwd",dataJson).then(function(response){
          randomStr.value = response["data"]["data"]
          console.log("afdsa",response["data"]["data"])
        })
      }
    };
    return {
      ...toRefs(state),
      plainOptions,
      onCheckAllChange,
      pwdLength,
      genPwd,
      randomStr,
    };
  },
});
</script>