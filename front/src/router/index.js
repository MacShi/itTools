
import { createRouter, createWebHashHistory } from "vue-router";
import Home from "../components/clientInfo/clientInfo";
import genPwd from "../components/tools/GeneratePasswd";
import steganSting from "../components/tools/SteganographyString";
import formatJson from "../components/tools/FormatJson";
import enCode from "../components/tools/EnCode";
import qrCodeHandle from "../components/tools/QrCodeHandle"

const router = createRouter({
  // baseURL: process.env.VUE_APP_BASE_URL,
  history: createWebHashHistory(process.env.VUE_APP_BASE_URL),
  routes: [
    { path: "/", name: "Home", component: Home },
    { path: "/genPwd", component: genPwd },
    { path: "/hide", component: steganSting },
    { path: "/formatJson", component: formatJson },
    { path: "/enCode", component: enCode },
    { path: "/qrCodeHandle", component: qrCodeHandle },
  ]

})
export default router;
