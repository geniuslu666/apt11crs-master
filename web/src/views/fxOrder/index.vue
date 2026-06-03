<template>
  <div>
    <div class="n-layout-page-header" style="margin: 0">
      <n-card :bordered="false" :header-style="{
                    padding: '25px 20px 20px',
                  }" :content-style="{
                    padding: '0 20px 20px',
                  }">
        <template #header>
          <text style="font-weight: 500;font-size: 20px;color: #3D3D3D;line-height: 28px;">分销订单</text>
        </template>
        <div class="statusTab">
          <div :class="defaultTab == 'hotel' ? 'active' : ''"><span @click="handleBeforeLeave('hotel')">酒店订单</span></div>
        </div>
        <FoodOrderLog ref="foodOrderLogRef"/>
      </n-card>
    </div>
  </div>
</template>

<script lang="ts" setup>
  import { ref } from 'vue';
  import FoodOrderLog from "@/views/fxOrder/food_order_log.vue";

  const defaultTab = ref('hotel');

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
