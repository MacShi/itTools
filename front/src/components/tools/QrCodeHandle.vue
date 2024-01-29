<template>
  <a-tabs default-active-key="generate">
    <a-tab-pane key="generate" tab="生成二维码">
      <div>
        <a-textarea v-model:value="textToEncode" placeholder="输入要生成的文本"  />
        <a-button type="primary" @click="generateQRCode" style="margin-top: 10px">生成二维码</a-button>
         <a-divider />
        <div v-if="generatedQRCode">
          <img :src="generatedQRCode" alt="QR Code" />
        </div>
      </div>
    </a-tab-pane>
    
    <a-tab-pane key="decode" tab="解析二维码">
      <div>
        <a-upload :before-upload="beforeUpload" :show-upload-list="false">
          <a-button icon="upload">上传二维码图片</a-button>
        </a-upload>
        <a-divider />
        <div v-if="decodedText">
          <h3>解析结果:</h3>
          <p>{{ decodedText }}</p>
        </div>
      </div>
    </a-tab-pane>
  </a-tabs>
</template>

<script>
import QRCode from 'qrcode';
import http from "../../http"
import { message } from 'ant-design-vue';

export default {
  data() {
    // const decodedText="";
    return {
      textToEncode: '',
      generatedQRCode: '',
      decodedText: '',
    };
  },
  methods: {
    generateQRCode() {
      if (this.textToEncode) {
        QRCode.toDataURL(this.textToEncode)
          .then((url) => {
            this.generatedQRCode = url;
          })
          .catch((err) => {
            console.error('Error generating QR code:', err);
          });
      }
    },
     beforeUpload(file) {
      const formData = new FormData();
      formData.append('file', file);
       http.post("/apis/upload",formData,{
          headers: {
            'Content-Type': 'multipart/form-data',
          },
        }).then((response)=>{
            console.log(0);
            console.log(response["data"]);
            if (response["data"]["code"]==0){
              console.log(response["data"]["data"]);
              this.decodedText=response["data"]["data"];
              message.success("解析成功");
            }else{
              message.error("解析失败")
            }
          })

    },

  },
};
</script>
