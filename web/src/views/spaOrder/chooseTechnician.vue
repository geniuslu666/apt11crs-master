<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content closable :header-style="{
                    padding: '20px',
                  }" :body-content-style="{
                    padding: '20px',
                  }">
        <template #header>
          <div style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">选择技师</div>
        </template>
        <template #footer>
          <n-button @click="closeForm" style="width: 70px;height: 35px;">
            取消
          </n-button>
          <n-button type="primary" :loading="formBtnLoading" @click="confirmForm" style="width: 70px;height: 35px;margin-left: 10px">
            确认
          </n-button>
        </template>
        <n-spin :show="loading" description="请稍候...">
          <BasicTable  ref="actionRef" :openChecked="true" :columns="columns" :request="loadDataTable" :row-key="(row) => row.id+'-'+row.nickname+'-'+row.phoneArea+'-'+row.phone" :scroll-x="scrollX" :resizeHeightOffset="-10000" :checked-row-keys="checkedInfos" @update:checked-row-keys="handleOnCheckedRow"></BasicTable>
        </n-spin>
      </n-drawer-content>
    </n-drawer>
  </div>
</template>

<script lang="ts" setup>
import {computed, h, ref} from 'vue';
import {adaModalWidth, adaTableScrollX} from "@/utils/hotgo";
import {BasicTable} from "@/components/Table";
import {NTag, useMessage} from "naive-ui";
import {isNullObject} from "@/utils/is";
import {OrderTechnician} from "@/api/spaOrder";

const message = useMessage();
  const emit = defineEmits(['reloadTechnician']);
const loading = ref(false);
const showModal = ref(false);
const dialogWidth = computed(() => {
  return adaModalWidth(500);
});
const actionRef = ref();
const searchFormRef = ref<any>({});

const formBtnLoading = ref(false);
const spaOrderId = ref(0);
const checkedInfos = ref([]);
const orderSelectNum = ref(0);

// 表格列
const columns = [
  {
    type: 'selection',
    disabled(row) {
      return row.workStatusEnum == 1 || row.workStatusEnum == 2
    }
  },
  {
    title: '技师信息',
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
      return row.phoneArea + row.phone
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
            default: () => '可预约',
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
            default: () => '工作中',
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

  const scrollX = computed(() => {
    return adaTableScrollX(columns, 0);
  });

    // 加载表格数据
    const loadDataTable = async (res) => {
      res.id = spaOrderId.value;
      return await OrderTechnician({ ...searchFormRef.value?.formModel, ...res });
    };

  function openModal(orderId, selectNum, technicianInfos) {
    spaOrderId.value = orderId;
    orderSelectNum.value = selectNum;
    if(technicianInfos){
      checkedInfos.value = technicianInfos.split(',')
    }else{
      checkedInfos.value = []
    }
    showModal.value = true;
  }

// 更新选中的行
function handleOnCheckedRow(rowKeys) {
  checkedInfos.value = rowKeys;
}

function confirmForm() {
  formBtnLoading.value = true;
  if(orderSelectNum.value !=  checkedInfos.value.length){
    message.error('请选择'+orderSelectNum.value+'名技师');
    formBtnLoading.value = false;
    return false;
  }
  emit('reloadTechnician',checkedInfos.value.join(','));
  formBtnLoading.value = false;
  closeForm();
}


function closeForm() {
  showModal.value = false;
  loading.value = false;
  spaOrderId.value = 0;
  orderSelectNum.value = 0;
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
