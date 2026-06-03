<template>
  <div >
    <div  style="padding-bottom: 20px;color: #fff; width: 120px; font-size: 20px; ">
      <div class="pb-2">账号密码登录</div>
      <div style="height: 3px;background-color: #fff; width: 100%"></div>
    </div>
    <Form @updateActiveModule="updateActiveModule" mode="account" />
  </div>
</template>

<script lang="ts" setup>
  import Form from './form.vue';

  const emit = defineEmits(['updateActiveModule']);

  function updateActiveModule(key: string) {
    emit('updateActiveModule', key);
  }
</script>

<style lang="less" scoped></style>
