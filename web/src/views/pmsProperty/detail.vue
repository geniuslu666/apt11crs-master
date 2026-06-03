<template>
  <div class="height100 flex-row">
    <div class="bgfff p-2" style="width: 180px; border-radius: 4px; height: 262px;margin-right: 15px">
      <n-menu
        :options="state.menuOptions"
        :value="state.valuekey"
        :default-expand-all="true"
        @update:value="onmenuOptions"
      />
    </div>
    <div class="flex-item bgfff" style="height: 86vh; overflow-x: hidden;border-radius: 4px">
      <detail v-if="state.valuekey == 'detail'" />
      <plan v-else-if="state.valuekey == 'plan'" />
      <fangx v-else-if="state.valuekey == 'fangx'" />
      <fangj v-else-if="state.valuekey == 'fangj'" />
    </div>
  </div>
</template>

<script lang="ts" setup>
  import { ref, computed, onMounted, reactive } from 'vue';
  import { Edit, View } from '@/api/pmsProperty';
  import detail from './comm/detail.vue';
  import fangx from '../pmsRoomType/index.vue';
  import fangj from '../pmsRoomUnit/index.vue';
  import plan from '../pmsPricePlan/index.vue';
  import {
    BookOutline as BookIcon,
    PersonOutline as PersonIcon,
    WineOutline as WineIcon,
  } from '@vicons/ionicons5';
  import { useMessage } from 'naive-ui';
  import { useRouter } from 'vue-router';

  const router = useRouter();
  import { stat } from 'fs';

  const message = useMessage();
  const state = reactive({
    rowdata: {},
    valuekey: 'detail',
    menuOptions: [
      {
        label: '物业详情',
        key: 'detail',
      },
      {
        label: '价格计划',
        key: 'plan',
      },
      {
        label: '房间管理',
        key: 'gl',
        children: [
          {
            label: '房型',
            key: 'fangx',
          },
          {
            label: '房间',
            key: 'fangj',
          },
        ],
      },
    ],
  });
  const onmenuOptions = (key: string) => {
    state.valuekey = key;
  };

  onMounted(() => {
    if (router.currentRoute.value.query?.rowdata) {
      state.rowdata = JSON.parse(router.currentRoute.value.query.rowdata);
    }
  });
</script>

<style lang="less"></style>
