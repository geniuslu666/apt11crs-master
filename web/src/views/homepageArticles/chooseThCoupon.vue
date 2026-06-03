<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content closable :header-style="{
                    padding: '20px',
                  }" :body-content-style="{
                    padding: '20px',
                  }">
        <template #header>
          <div style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">选择礼品券</div>
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
import {computed, ref} from 'vue';
import {adaModalWidth, adaTableScrollX} from "@/utils/hotgo";
import {BasicTable} from "@/components/Table";
import {schemas, loadOptions, columns} from '@/views/thCoupon/model';
import { List } from '@/api/thCoupon';
import {BasicForm, useForm} from "@/components/Form";
import {useMessage} from "naive-ui";

const emit = defineEmits(['reloadThCoupon']);
const loading = ref(false);
const showModal = ref(false);
const message = useMessage();
const dialogWidth = computed(() => {
  return adaModalWidth(1000);
});
const checkedIds = ref<number[]>([]);

const actionRef = ref();
const searchFormRef = ref<any>({});

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
  return await List({ ...searchFormRef.value?.formModel, ...res });
};

// 更新选中的行
function handleOnCheckedRow(rowKeys) {
  checkedIds.value = rowKeys;
}

async function openModal(selectedIds) {
  checkedIds.value = selectedIds
  await loadOptions()
  showModal.value = true;
}

function confirmForm(){
  if(checkedIds.value.length <= 0){
    message.error('请选择礼品券');
    return false;
  }
  emit('reloadThCoupon',checkedIds.value);
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
