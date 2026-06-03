<template>
  <n-space :vertical="true">
    <br/>
    <n-divider> <span class="cfff">演示角色登录</span></n-divider>
    <n-space justify="center">
      <n-button
        v-for="item in accounts"
        :key="item.username"
        type="primary"
        class="primary_button"
        @click="login(item.username, item.password)"
      >
        {{ item.label }}
      </n-button>
    </n-space>
  </n-space>
</template>

<script lang="ts" setup>
  interface Emits {
    (e: 'login', param: { username: string; password: string }): void;
  }

  const emit = defineEmits<Emits>();

  const accounts = [
    {
      label: '超级管理员',
      username: 'admin',
      password: '123456',
    },
    {
      label: '管理员',
      username: 'test',
      password: '123456',
    },
    // {
    //   label: '代理商',
    //   username: 'ameng',
    //   password: '123456',
    // },
  ];

  function login(username: string, password: string) {
    emit('login', { username, password });
  }
</script>

<style scoped>
  .n-divider:not(.n-divider--vertical) {
    margin: 0 12px 0;
  }
  .primary_button {
    border-radius: 6px;
  }
</style>
