<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content closable :header-style="{
                    padding: '20px',
                  }" :body-content-style="{
                    padding: '20px',
                  }">
        <template #header>
          <div style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">选择按摩服务</div>
        </template>
        <template #footer>
          <n-button @click="closeForm" style="width: 70px;height: 35px;margin-right: 10px">
            取消
          </n-button>
          <n-button type="info" @click="confirmForm" style="width: 70px;height: 35px;">
            确认
          </n-button>
        </template>
        <n-spin :show="loading" description="请稍候...">
          <BasicForm  ref="searchFormRef" @register="register" @submit="reloadTable" @reset="reloadTable" @keyup.enter="reloadTable">
            <template #statusSlot="{ model, field }">
              <n-input v-model:value="model[field]" />
            </template>
          </BasicForm>
          <BasicTable :openChecked="true" ref="actionRef" :columns="columns" :request="loadDataTable" :row-key="(row) => row.id" :scroll-x="scrollX" :resizeHeightOffset="-10000"  :checked-row-keys="checkedIds" @update:checked-row-keys="handleOnCheckedRow"></BasicTable>
        </n-spin>
      </n-drawer-content>
    </n-drawer>
  </div>
</template>

<script lang="ts" setup>
import {computed, ref, h} from 'vue';
import {adaModalWidth, adaTableScrollX} from "@/utils/hotgo";
import {BasicTable} from "@/components/Table";
import { List } from '@/api/spaService';
import {BasicForm, useForm} from "@/components/Form";
import {useMessage} from "naive-ui";
import { FormSchema } from '@/components/Form';

const emit = defineEmits(['reloadSpaService']);
const loading = ref(false);
const showModal = ref(false);
const message = useMessage();
const dialogWidth = computed(() => {
  return adaModalWidth(1000);
});
const checkedIds = ref<number[]>([]);

const actionRef = ref();
const searchFormRef = ref<any>({});

// 表格搜索表单
const schemas = ref<FormSchema[]>([
  {
    field: 'name',
    component: 'NInput',
    label: '服务名称',
    componentProps: {
      placeholder: '服务名称',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
]);

// 表格列
const columns = [
  {
    title: 'ID',
    key: 'Id',
    align: 'left',
    width: 80,
    render(row){
      return row.id
    }
  },
  {
    title: '服务名称',
    key: 'name',
    align: 'left',
    width: 200,
    render(row){
      return row.name
    }
  },
  {
    title: '所属服务商',
    key: 'ispId',
    align: 'left',
    width: 120,
    render(row){
      if(row.ispId > 0){
        return row.spaIspDetail.name
      }else{
        return '--'
      }
    }
  },
];

const scrollX = computed(() => {
  return adaTableScrollX(columns, 0);
});

const [register, {}] = useForm({
  gridProps: { cols: '3' },
  labelWidth: 80,
  schemas,
});

// 加载表格数据
const loadDataTable = async (res) => {
  // 筛选条件 营业中、可预约
  res.openStatus = "OPENING";
  res.canOrder = 1;
  return await List({ ...searchFormRef.value?.formModel, ...res });
};

// 更新选中的行
function handleOnCheckedRow(rowKeys) {
  checkedIds.value = rowKeys;
}

async function openModal(selectedIds) {
  checkedIds.value = selectedIds
  showModal.value = true;
}

function confirmForm(){
  if(checkedIds.value.length <= 0){
    message.error('请选择服务 ');
    return false;
  }
  console.log('reloadSpaService_before',checkedIds.value)
  emit('reloadSpaService',checkedIds.value);
  showModal.value = false;
}


// 重新加载表格数据
function reloadTable() {
  actionRef.value?.reload();
}

function closeForm() {
  showModal.value = false;
  loading.value = false;
}

defineExpose({
  openModal,
});
</script>

<style lang="less"></style>
