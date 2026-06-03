<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content closable :header-style="{
                    padding: '20px',
                  }" :body-content-style="{
                    padding: '25px 20px',
                  }">
        <template #header>
          <div style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">选择会员</div>
        </template>
        <n-spin :show="loading" description="请稍候...">
          <BasicForm  ref="searchFormRef" @register="register" @submit="reloadTable" @reset="reloadTable" @keyup.enter="reloadTable">
            <template #statusSlot="{ model, field }">
              <n-input v-model:value="model[field]" />
            </template>
          </BasicForm>
          <BasicTable  ref="actionRef" :openChecked="false" :columns="columns" :request="loadDataTable" :actionColumn="actionColumn" :scroll-x="scrollX" :resizeHeightOffset="-10000"></BasicTable>
        </n-spin>
      </n-drawer-content>
    </n-drawer>
  </div>
</template>

<script lang="ts" setup>
import {ref, computed, reactive, h} from 'vue';
import {Bind} from '@/api/spaIsp';
import { useMessage } from 'naive-ui';
import {adaModalWidth, adaTableScrollX} from '@/utils/hotgo';
import {List} from "@/api/pmsMember/index";
import {BasicForm, FormSchema, useForm} from "@/components/Form";
import {BasicTable, TableAction} from "@/components/Table";

const emit = defineEmits(['reloadTable']);
const message = useMessage();
const loading = ref(false);
const showModal = ref(false);
const dialogWidth = computed(() => {
  return adaModalWidth(1100);
});
const actionRef = ref();
const searchFormRef = ref<any>({});
const initIspId = ref(0);
const chooseMemberId = ref(0);


// 表格搜索表单
const schemas = ref<FormSchema[]>([
  {
    field: 'memberNo',
    component: 'NInput',
    label: '会员号',
    componentProps: {
      placeholder: '会员号',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'phone',
    component: 'NInput',
    label: '手机号',
    componentProps: {
      placeholder: '手机号',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'fullName',
    component: 'NInput',
    label: '全称',
    componentProps: {
      placeholder: '会员全称',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
]);

// 表格列
const columns = [
  {
    title: '会员号',
    key: 'memberNo',
    align: 'left',
    width: 150
  },
  {
    title: '姓名',
    key: 'fullName',
    align: 'left',
    width: 150,
  },
  {
    title: '手机',
    key: 'phone',
    align: 'left',
    width: 200,
    render(row){
      return row.phoneArea + row.phone
    }
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
          label: '绑定',
          onClick: handleChoose.bind(null, record),
        },
      ],
    });
  },
});

const scrollX = computed(() => {
  return adaTableScrollX(columns, actionColumn.width);
});

const [register, {}] = useForm({
  gridProps: { cols: '1 s:1 m:2 l:3 xl:4 2xl:4' },
  labelWidth: 100,
  schemas,
});

// 加载表格数据
const loadDataTable = async (res) => {
  res.status = 1;
  return await List({ ...searchFormRef.value?.formModel, ...res });
};


// 重新加载表格数据
function reloadTable() {
  actionRef.value?.reload();
}

function openModal(id,memberId) {
  initIspId.value = id;
  chooseMemberId.value = memberId;
  showModal.value = true;
}

function handleChoose(record: Recordable) {
  loading.value = true
  Bind({
    id: initIspId.value,
    memberId: record.id
  }).then((_res) => {
    message.success('操作成功');
    setTimeout(() => {
      loading.value = false;
      showModal.value = false;
      emit('reloadTable');
    });
  });
}

defineExpose({
  openModal,
});
</script>

<style lang="less"></style>


