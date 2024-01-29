<template>
  <div>
    <a-layout>
      <a-layout-sider
        breakpoint="lg"
        collapsed-width="0"
        @collapse="onCollapse"
        @breakpoint="onBreakpoint"
        :style="{ height: '300 px' }"
      >
        <div />
        <a-menu
          v-model:selectedKeys="selectedKeys"
          theme="dark"
          mode="inline"
          @click="handleClick"
        >
          <a-menu-item key="1">
            <user-outlined />
            <router-link to="/" >
              <span class="nav-text"> 客户端信息</span>
            </router-link>
          </a-menu-item>
          <a-menu-item key="2">
            <TagOutlined />
            <router-link to="/genPwd">
              <span class="nav-text"> 密码生成</span>
            </router-link>
          </a-menu-item>
          <a-menu-item key="3">
            <TagOutlined />
            <router-link to="/hide" />
            <span> 文字隐写</span>
          </a-menu-item>
           <a-menu-item key="4">
            <TagOutlined />
            <router-link to="/formatJson" />
            <span> json格式化</span>
          </a-menu-item>
          <a-menu-item key="5">
            <TagOutlined />
            <router-link to="/enCode" />
            <span> 编码</span>
          </a-menu-item>
           <a-menu-item key="6">
            <TagOutlined />
            <router-link to="/qrCodeHandle" />
            <span> 二维码</span>
          </a-menu-item>
        </a-menu>
      </a-layout-sider>
      <a-layout>
        <a-layout-header
          :style="{
            background: '#fff',
            padding: 0,
            margin: '16px 16px 16px 16px',
          }"
        />
        <a-layout-content :style="{ margin: '16px 16px 16px 16px' }">
          <div
            :style="{ padding: '24px', background: '#fff', minHeight: '560px' }"
          >
            <router-view></router-view>
          </div>
        </a-layout-content>
        <a-layout-footer style="text-align: center">
          Tools ©2023 Created by it2it
        </a-layout-footer>
      </a-layout>
    </a-layout>
  </div>
</template>
<script>
import { UserOutlined, TagOutlined } from "@ant-design/icons-vue";
import { defineComponent, ref } from "vue";
export default defineComponent({
  components: {
    UserOutlined,
    TagOutlined,
  },
  setup() {
    const onCollapse = (collapsed, type) => {
      console.log(collapsed, type);
    };
    const onBreakpoint = (broken) => {
      console.log(broken);
    };
    return {
      selectedKeys: ref(["1"]),
      onCollapse,
      onBreakpoint,
    };
  },
  methods: {
    handleClick(selectedKeys) {
      // if (selectedKeys.key == "4") {
      //   this.$router.push({
      //     path: "/Home",
      //   });
      // }
      console.log(selectedKeys.key);
      // console.log("fadsf",devServer["host"])
    },
  },
});
</script>
<style>
#components-layout-demo-responsive .logo {
  height: 32px;
  background: rgba(255, 255, 255, 0.2);
  margin: 16px;
}

.site-layout-sub-header-background {
  background: #fff;
}

.site-layout-background {
  background: #fff;
}

[data-theme="dark"] .site-layout-sub-header-background {
  background: #141414;
}
</style>