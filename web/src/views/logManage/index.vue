<template>
  <div>
    <n-spin :show="show" description="请稍候...">
      <n-card :bordered="false" :header-style="{
                    padding: '20px',
                  }">
        <template #header>
          <text style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">订单日志</text>
        </template>
        <n-tabs type="card" animated>
          <n-tab-pane name="酒店订单" tab="酒店订单">
            <HotelOrderLog ref="hotelOrderLogRef"/>
          </n-tab-pane>
          <n-tab-pane name="餐厅订单" tab="餐厅订单">
            <FoodOrderLog ref="foodOrderLogRef"/>
          </n-tab-pane>
          <n-tab-pane name="接送机订单" tab="接送机订单">
            <CarOrderLog ref="carOrderLogRef"/>
          </n-tab-pane>
          <n-tab-pane name="按摩订单" tab="按摩订单">
            <SpaOrderLog ref="spaOrderLogRef"/>
          </n-tab-pane>
        </n-tabs>
      </n-card>
    </n-spin>
  </div>
</template>

<script lang="ts" setup>
  import { ref, onMounted } from 'vue';
  import { useMessage } from 'naive-ui';
  import {useProjectSettingStore} from "@/store/modules/projectSetting";
  import CarOrderLog from "@/views/logManage/car_order_log.vue";
  import FoodOrderLog from "@/views/logManage/food_order_log.vue";
  import SpaOrderLog from "@/views/logManage/spa_order_log.vue";
  import HotelOrderLog from "@/views/logManage/hotel_order_log.vue";

  const show = ref(false);
  const message = useMessage();
  const formBtnLoading = ref(false);
  const settingStore = useProjectSettingStore();
  const carOrderLogRef = ref();
  const foodOrderLogRef = ref();
  const spaOrderLogRef = ref();
  const hotelOrderLogRef = ref();


  onMounted(() => {
    load();
  });

  function load() {

  }
</script>
