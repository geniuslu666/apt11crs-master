<template>
  <div>
    <n-spin :show="show" description="请稍候...">
      <n-card :bordered="false" :header-style="{
                    padding: '20px',
                  }">
        <template #header>
          <text style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 28px;">核销终端-配置</text>
        </template>
        <div class="statusTab">
          <div :class="tab == brandItem.id ? 'active' : ''" v-for="brandItem in brandListArr" :key="brandItem.brandModel"><span @click="handleTab(brandItem.id)">{{ brandItem.brandModel }}</span></div>
        </div>
        <BrandConfig :brandId="tab"/>
      </n-card>

    </n-spin>
  </div>
</template>

<script lang="ts" setup>
  import { ref, onMounted } from 'vue';
  import { BrandList} from '@/api/terminal';
  import BrandConfig from "@/views/terminal/brand_config.vue";

  const show = ref(false);
  const tab = ref(0);

  const brandListArr = ref([]);

  onMounted(() => {
    load();
  });

  function handleTab(e){
    tab.value = e;
  }

  function load() {
    show.value = true;
    new Promise((_resolve, _reject) => {
      BrandList({ pagination: false })
        .then((res) => {
          brandListArr.value = res.list;
          tab.value = res.list[0].id
          show.value = false;
        })
    });
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
