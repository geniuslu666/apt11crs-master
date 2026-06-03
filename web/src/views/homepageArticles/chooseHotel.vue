<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content closable :header-style="{
                    padding: '20px',
                  }" :body-content-style="{
                    padding: '20px',
                  }">
        <template #header>
          <div style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">选择民宿</div>
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
import { List } from '@/api/pmsProperty';
import {BasicForm, useForm} from "@/components/Form";
import {NTag, useMessage} from "naive-ui";
import { FormSchema } from '@/components/Form';
import {getlang} from "@/utils/smjcomm";
import {useUserStore} from "@/store/modules/user";

const emit = defineEmits(['reloadHotel']);
const userStore = useUserStore();
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
    label: '物业名称',
    componentProps: {
      placeholder: '物业名称',
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
    title: '物业名称',
    key: 'name',
    align: 'left',
    width: 300,
    render: function (row){
      var labelH = ""

      if(row.deletedAt != null){
        labelH = h(
          NTag,
          {
            style: {
              marginRight: '6px',
            },
            type: 'error',
            bordered: false,
            size: 'small'
          },
          {
            default: () => "已删除",
          }
        )
      }
      var propertyName = "--"
      if(row.nameLanguage){
        propertyName = getlang(row.nameLanguage, userStore.language).content
      }

      return h(
        'div',
        null,
        [
          h(
            'span',
            {
              style: {
                display: "inline-block",
                marginRight: '6px'
              }
            },
            {
              default: () => propertyName,
            }
          ),
          labelH
        ]

      )
    }
  },
  {
    title: '房型数',
    key: 'roomTypeNum',
    align: 'left',
    width: 100,
  },
  {
    title: '房间数',
    key: 'roomUnitNum',
    align: 'left',
    width: 100,
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
  // 筛选条件 已启用
  res.propertyStatus = 1;
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
    message.error('请选择物业 ');
    return false;
  }
  console.log('reloadHotel_before',checkedIds.value)
  emit('reloadHotel',checkedIds.value);
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
