<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content closable :header-style="{
                    padding: '20px',
                  }" :body-content-style="{
                    padding: '20px',
                  }">
        <template #header>
          <div style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">选择员工</div>
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
          <BasicForm  ref="employeeSearchFormRef" @register="employeeRegister" @submit="employeeReloadTable" @reset="employeeReloadTable" @keyup.enter="employeeReloadTable">
            <template #statusSlot="{ model, field }">
              <n-input v-model:value="model[field]" />
            </template>
          </BasicForm>
          <BasicTable :openChecked="true" ref="employeeActionRef" :columns="employeeColumns" :request="employeeLoadDataTable" :row-key="(row) => row.id" :scroll-x="employeeScrollX" :resizeHeightOffset="-10000"  :checked-row-keys="employeeCheckedIds" @update:checked-row-keys="handleOnEmployeeCheckedRow"></BasicTable>
        </n-spin>
      </n-drawer-content>
    </n-drawer>
  </div>
</template>

<script lang="ts" setup>
import {computed, ref} from 'vue';
import {adaModalWidth, adaTableScrollX} from "@/utils/hotgo";
import {BasicTable} from "@/components/Table";
import {schemas, columns as employeeColumns, loadOptions} from '@/views/employee/model';
import { List } from '@/api/employee';
import {BasicForm, useForm} from "@/components/Form";
import {useMessage} from "naive-ui";

  const emit = defineEmits(['reloadEmployee']);
  const loading = ref(false);
  const showModal = ref(false);
const message = useMessage();
  const dialogWidth = computed(() => {
    return adaModalWidth(1000);
  });
const employeeCheckedIds = ref<number[]>([]);

const employeeActionRef = ref();
const employeeSearchFormRef = ref<any>({});

  const employeeScrollX = computed(() => {
    return adaTableScrollX(employeeColumns, 0);
  });

const [employeeRegister, {}] = useForm({
  gridProps: { cols: '3' },
  labelWidth: 80,
  schemas,
});

  // 加载表格数据
  const employeeLoadDataTable = async (res) => {
    return await List({ ...employeeSearchFormRef.value?.formModel, ...res });
  };

// 更新选中的行
function handleOnEmployeeCheckedRow(rowKeys) {
  employeeCheckedIds.value = rowKeys;
}

async function openModal(selectedIds) {
  employeeCheckedIds.value = selectedIds
  await loadOptions()
  showModal.value = true;
}

function confirmForm(){
  if(employeeCheckedIds.value.length <= 0){
    message.error('请选择员工');
    return false;
  }
  emit('reloadEmployee',employeeCheckedIds.value);
  showModal.value = false;
}


// 重新加载表格数据
function employeeReloadTable() {
  employeeActionRef.value?.reload();
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
