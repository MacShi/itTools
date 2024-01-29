<template>
  <a-descriptions title="远程客户端信息" :column="1">
    <a-descriptions-item label="IP地址">{{ IpProxy }}</a-descriptions-item>
    <a-descriptions-item label="IP地址(WebRTC)">{{
      RealIp
    }}</a-descriptions-item>
    <a-descriptions-item label="User Agent">{{ UA }}</a-descriptions-item>
  </a-descriptions>
</template>
<script>
import { ref } from "vue";
import http from "../../http";

export default {
  setup() {
    const UA = ref("Null");
    const IpProxy = ref("Null");
    const RealIp = ref("Null");

    function getClientInfo() {
      http.get("/apis/clientInfo").then(function (response) {
        IpProxy.value = response.data["data"]["remote_ip"];
        UA.value = response.data["data"]["user_agent"];
      });
    }
    function getIp() {
      function handleCandidate(candidate) {
        if (candidate.indexOf("srflx") != -1) {
          var regex =
            /([0-9]{1,3}(\.[0-9]{1,3}){3}|[a-f0-9]{1,4}(:[a-f0-9]{1,4}){7})/;
          var ip_addr = regex.exec(candidate)[0];
          RealIp.value = ip_addr;
        }
      }

      if (window.RTCPeerConnection) {
        const config = {
          iceServers: [
            {
              urls: ["stun:stun.l.google.com:19302", "stun:stun.voippro.com"], // stun.voippro.com  stun.voipraider.com  这里使用谷歌，线上部署直接替换
              // urls: "stun:stun.l.google.com:19302"
            },
          ],
        };

        // 构建
        let pc = new RTCPeerConnection(config);
        pc.onicecandidate = function (event) {
          if (event.candidate) handleCandidate(event.candidate.candidate);
        };
        pc.createDataChannel("");
        pc.createOffer(
          function (result) {
            pc.setLocalDescription(result);
          },
          function () {}
        );
      } else {
        RealIp.value = "不支持 RTCPeerConnection";
      }
    }

    getClientInfo();
    getIp();
    return {
      UA,
      IpProxy,
      getClientInfo,
      RealIp,
    };
  },
};
</script>
