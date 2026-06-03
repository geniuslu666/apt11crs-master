<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content closable :header-style="{
                    padding: '20px',
                  }" :body-content-style="{
                    padding: '20px',
                  }">
        <template #header>
          <div style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">选择商户</div>
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
import {computed, h, ref} from 'vue';
import {adaModalWidth, adaTableScrollX, getOptionLabel, getOptionTag} from "@/utils/hotgo";
import {BasicTable} from "@/components/Table";
import {options, schemas, loadOptions} from '@/views/thMch/model';
import { List } from '@/api/thMch';
import {BasicForm, useForm} from "@/components/Form";
import {isNullObject} from "@/utils/is";
import {NTag, useMessage} from "naive-ui";

  const emit = defineEmits(['reloadMch']);
  const loading = ref(false);
  const showModal = ref(false);
const message = useMessage();
  const dialogWidth = computed(() => {
    return adaModalWidth(950);
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

const columns = [
  {
    title: '商户名称',
    key: 'name',
    align: 'left',
    width: 120,
  },
  {
    title: '商户分类',
    key: 'thMchCategoryName',
    align: 'left',
    width: 130,
  },
  {
    title: '状态',
    key: 'status',
    align: 'left',
    width: 80,
    render(row) {
      if (isNullObject(row.status)) {
        return ``;
      }
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: getOptionTag(options.value.sys_normal_disable, row.status),
          bordered: false,
        },
        {
          default: () => getOptionLabel(options.value.sys_normal_disable, row.status),
        }
      );
    },
  },
];

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
      message.error('请选择商户');
      return false;
    }
    emit('reloadMch',checkedIds.value);
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
