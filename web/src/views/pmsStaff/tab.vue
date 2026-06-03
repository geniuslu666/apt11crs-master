<template>
  <div>
    <div class="n-layout-page-header" style="margin: 0">
      <n-card :bordered="false" :header-style="{
                    padding: '25px 20px 20px',
                  }" :content-style="{
                    padding: '0 20px 20px',
                  }">
        <template #header>
          <text style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">员工分销</text>
        </template>
        <div class="statusTab">
          <div :class="defaultTab == 'channelIndex' ? 'active' : ''"><span @click="handleBeforeLeave('channelIndex')">员工管理</span></div>
          <div :class="defaultTab == 'channelWithdraw' ? 'active' : ''"><span @click="handleBeforeLeave('channelWithdraw')">员工提现</span></div>
        </div>
        <IndexList v-if="defaultTab == 'channelIndex'"/>
        <WithdrawList v-if="defaultTab == 'channelWithdraw'"/>
      </n-card>
    </div>
  </div>
</template>

<script lang="ts" setup>
  import { onMounted, ref } from 'vue';
  import IndexList from '@/views/pmsStaff/index.vue';
  import WithdrawList from '@/views/pmsWithdraw/index.vue';

  const defaultTab = ref('channelIndex');

  onMounted(() => {
  });

  function handleBeforeLeave(tabName: string) {
    defaultTab.value = tabName;
  }
</script>

<style lang="less" scoped>
.statusTab{
  display: flex;
  div{
    margin-right: 5px;
    padding: 0 16px;
    height: 30px;
    line-height: 30px;
    text-align: center;
    font-size: 14px;
    color: #4E5969;
    span{
      cursor: pointer;
    }
    &.active{
      background: #F2F3F8;
      border-radius: 30px;
      color: #1664FF;
      font-weight: 500;
    }
  }
}
</style>
