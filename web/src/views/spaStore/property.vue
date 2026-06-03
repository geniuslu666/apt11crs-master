<template>
  <div>
    <BasicTable
      ref="actionRef"
      :columns="columns"
      :request="loadDataTable"
      :row-key="(row) => row.id"
      :scroll-x="scrollX"
      :resizeHeightOffset="-10000"
      :checked-row-keys="checkedIds"
      :pagination="false"
      @update:checked-row-keys="handleOnCheckedRow"
    >
      <template #tableTitle>
      </template>
    </BasicTable>
    <View ref="viewRef" />
  </div>
</template>

<script lang="ts" setup>
  import { useDesignSettingStore } from '@/store/modules/designSetting';

  import { ref, computed, onMounted } from 'vue';
  import { BasicTable } from '@/components/Table';
  import { List } from '@/api/pmsProperty';
  import { columns } from '@/views/spaStore/property_model';
  import { adaTableScrollX } from '@/utils/hotgo';

  const designStore = useDesignSettingStore();
  const actionRef = ref();
  const viewRef = ref();
  const checkedIds = ref([]);
  const scrollX = computed(() => {
    return adaTableScrollX(columns, 0);
  });

  // 加载表格数据
  const loadDataTable = async () => {
    const aa = await List({ page: 1, pageSize: 200 });
    return aa;
  };

  // 更新选中的行
  function handleOnCheckedRow(rowKeys) {
    checkedIds.value = rowKeys;
  }
  onMounted(() => {
    // loadDataTable();
  });
</script>

<style lang="less" scoped>
  .meunpms {
    box-shadow: 2px 0 8px 0 rgba(29, 35, 41, 0.05);
    font-size: 14px;
  }
  .pmsitem {
    padding: 12px 12px 12px 22px;
  }
  .pmsitem:hover {
    background-color: hsla(0, 0%, 87%, 0.341);
    border-radius: 5px;
  }
</style>
