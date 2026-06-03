<template>
  <div>
    <n-spin :show="show" description="请稍候...">
      <n-card
        :bordered="false"
        class="proCard"
        :header-style="{
                    padding: '25px 20px 20px',
                  }"
        :content-style="{
                    padding: '0 20px 20px',
                  }"
      >
        <template #header>
          <text style="font-weight: 500;font-size: 20px;color: #3D3D3D;line-height: 28px;">预约订单</text>
        </template>
        <div class="statusTab">
          <div :class="tab == 0 ? 'active' : ''"><span @click="handleTab(0)">CRS定金模式订单</span></div>
          <div :class="tab == 1 ? 'active' : ''"><span @click="handleTab(1)">CRS全款模式订单</span></div>
          <div :class="tab == 2 ? 'active' : ''"><span @click="handleTab(2)">TORETA订单</span></div>
        </div>
        <CrsOrder ref="crsOrderRef" v-if="tab == 0"/>
        <ToretaOrder ref="toretaOrderRef" v-if="tab == 2"/>
        <CrsAllOrder ref="crsAllOrderRef" v-if="tab == 1"/>
      </n-card>
    </n-spin>
  </div>
</template>
<script lang="ts" setup>
import {ref} from 'vue';
import CrsOrder from "@/views/foodOrder/crs.vue";
import ToretaOrder from "@/views/foodOrder/toreta.vue";
import CrsAllOrder from "@/views/foodOrder/crsall.vue";

const show = ref(false);
const tab = ref(0);
const crsOrderRef = ref();
const toretaOrderRef = ref();
const crsAllOrderRef = ref();

function handleTab(e){
  tab.value = e;
}
</script>
<style lang="less" scoped>
.statusTab{
  display: flex;
  margin-bottom: 4px;
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
