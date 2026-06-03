<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content closable :header-style="{
                    padding: '20px',
                  }" :body-content-style="{
                    padding: '20px',
                  }">
        <template #header>
          <div style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">选择司机</div>
        </template>
        <n-spin :show="loading" description="请稍候...">
          <BasicTable  ref="actionRef" :openChecked="false" :columns="columns" :request="loadDataTable" :actionColumn="actionColumn" :scroll-x="scrollX" :resizeHeightOffset="-10000"></BasicTable>
        </n-spin>
      </n-drawer-content>
    </n-drawer>
  </div>
</template>

<script lang="ts" setup>
import {computed, h, reactive, ref} from 'vue';
import {adaModalWidth, adaTableScrollX} from "@/utils/hotgo";
import {BasicTable, TableAction} from "@/components/Table";
import {OrderDriver} from '@/api/carOrder';
import {NTag} from "naive-ui";
import {isNullObject} from "@/utils/is";

  const emit = defineEmits(['reloadDriver']);
  const loading = ref(false);
  const showModal = ref(false);
  const dialogWidth = computed(() => {
    return adaModalWidth(500);
  });
const actionRef = ref();
const searchFormRef = ref<any>({});
const carOrderId = ref(0);
const chooseDriverId = ref(0);

// 表格列
const columns = [
  {
    title: '司机信息',
    key: 'nickname',
    align: 'left',
    width: 100
  },
  {
    title: '手机',
    key: 'phone',
    align: 'left',
    width: 130,
    render(row){
      return row.phoneArea + ' ' + row.phone
    }
  },
  {
    title: '工作状态',
    key: 'workStatus',
    align: 'left',
    width: 80,
    render(row) {
      if (isNullObject(row.workStatus)) {
        return ``;
      }
      if(row.workStatusEnum == 3){
        // 可预订
        return h(
          NTag,
          {
            style: {
              marginRight: '6px',
            },
            type: 'success',
            bordered: false,
          },
          {
            default: () => '未出车',
          }
        );
      }else if(row.workStatusEnum == 2){
        // 工作中
        return h(
          NTag,
          {
            style: {
              marginRight: '6px',
            },
            type: 'warning',
            bordered: false,
          },
          {
            default: () => '出车中',
          }
        );
      }else{
        // 休息
        return h(
          NTag,
          {
            style: {
              marginRight: '6px',
            },
            type: 'default',
            bordered: false,
          },
          {
            default: () => '休假中',
          }
        );
      }
    },
  },
];

  const actionColumn = reactive({
    width: 90,
    title: '操作',
    key: 'action',
    fixed: 'right',
    render(record) {
      return h(TableAction as any, {
        style: 'button',
        actions: [
          {
            label: chooseDriverId.value == record.id ? '已派单' : '派单',
            onClick: handleChoose.bind(null, record),
            disabled: (record.dispatchMode == 1 && chooseDriverId.value == record.id) || (record.dispatchMode == 2 && (record.workStatusEnum == 1 || record.workStatusEnum == 2 || chooseDriverId.value == record.id))
          },
        ],
      });
    },
  });

  const scrollX = computed(() => {
    return adaTableScrollX(columns, actionColumn.width);
  });

    // 加载表格数据
    const loadDataTable = async (res) => {
      res.id = carOrderId.value
      return await OrderDriver({ ...searchFormRef.value?.formModel, ...res });
    };

  function openModal(orderId, driverId) {
    carOrderId.value = orderId;
    chooseDriverId.value = driverId;
    showModal.value = true;
  }

function handleChoose(record: Recordable) {
  emit('reloadDriver',record);
  showModal.value = false;
}

  // 重新加载表格数据
  function reloadTable() {
    actionRef.value?.reload();
  }

  defineExpose({
    openModal,
  });
</script>

<style lang="less"></style>
