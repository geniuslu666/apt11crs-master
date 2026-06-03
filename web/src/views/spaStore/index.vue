<template>
  <div>
    <div class="n-layout-page-header">
      <n-card :bordered="false" :header-style="{
                    padding: '20px',
                  }">
        <template #header>
          <text style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">门店/地点管理</text>
        </template>
        <!--  这是由系统生成的CURD表格，你可以将此行注释改为表格的描述 -->
      </n-card>
    </div>
    <n-card :bordered="false" class="proCard" :content-style="{
                    padding: '0 20px 20px',
                  }">
      <n-tabs type="card" animated>
        <n-tab-pane name="门店信息" tab="门店信息">
          <Store ref="storeRef"/>
        </n-tab-pane>
        <n-tab-pane name="物业列表" tab="物业列表">
          <Property ref="propertyRef"/>
        </n-tab-pane>
      </n-tabs>
    </n-card>
  </div>
</template>

<script lang="ts" setup>
import { ref } from 'vue';
import Store from "@/views/spaStore/edit.vue";
import Property from "@/views/spaStore/property.vue";

const storeRef = ref();
const propertyRef = ref()
</script>

<style lang="less" scoped></style>

