<template>
  <div>
    <n-modal
      v-model:show="showModal"
      :mask-closable="false"
      :show-icon="false"
      preset="dialog"
      transform-origin="center"
      title="选择餐厅"
      :style="{
        width: dialogWidth,
      }"
    >
      <n-spin :show="loading" description="请稍候...">
        <BasicForm  ref="searchFormRef" @register="register" @submit="reloadTable" @reset="reloadTable" @keyup.enter="reloadTable">
          <template #statusSlot="{ model, field }">
            <n-input v-model:value="model[field]" />
          </template>
        </BasicForm>
        <BasicTable  ref="actionRef" :openChecked="false" :columns="columns" :request="loadDataTable" :actionColumn="actionColumn" :scroll-x="scrollX" :resizeHeightOffset="-10000"></BasicTable>
      </n-spin>
    </n-modal>
  </div>
</template>

<script lang="ts" setup>
import {computed, h, onMounted, reactive, ref} from 'vue';
import {adaModalWidth, adaTableScrollX, getOptionLabel, getOptionTag} from "@/utils/hotgo";
import {BasicTable, TableAction} from "@/components/Table";
import {BasicForm, useForm} from "@/components/Form";
import { FormSchema } from '@/components/Form';
import {List} from '@/api/foodRestaurant';
import {isNullObject} from "@/utils/is";
import {NTag} from "naive-ui";
import {loadOptions, options} from "@/views/foodRestaurant/model";

  const emit = defineEmits(['reloadCouponList']);
  const loading = ref(false);
  const showModal = ref(false);
  const dialogWidth = computed(() => {
    return adaModalWidth(1040);
  });
const actionRef = ref();
const searchFormRef = ref<any>({});

// 表格搜索表单
const schemas = ref<FormSchema[]>([
  {
    field: 'name',
    component: 'NInput',
    label: '餐厅名称',
    componentProps: {
      placeholder: '餐厅名称',
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
    title: '餐厅名称',
    key: 'name',
    align: 'left',
    width: 200,
    render(row){
      return row.name
    }
  },
  {
    title: '营业类型',
    key: 'cooperateTypeId',
    align: 'left',
    width: 80,
    render(row) {
      if (isNullObject(row.cooperateTypeId)) {
        return `--`;
      }
      return row.cooperateTypeDetail.typeName
    },
  },
  {
    title: '地址信息',
    key: 'address',
    align: 'left',
    width: 240,
    ellipsis: false,
    render(row) {
      let full_address = '--';
      if(row.areaPid > 0){
        full_address = row.pAreaDetail.areaName
      }
      if(row.areaId > 0){
        full_address = full_address + '-' + row.areaDetail.areaName
      }
      if(row.detailAddress){
        full_address = full_address + '-' + row.detailAddress
      }
      return h(
        'div',
        null,
        [
          h(
            'div',
            {},
            {
              default: () => '联系电话：' + row.phone ? row.phone : '--',
            }
          ),
          h(
            'div',
            {},
            {
              default: () => '地址：' + full_address,
            }
          )
        ]
      )
    },
  },
  {
    title: '餐厅状态',
    key: 'status',
    align: 'left',
    width: 100,
    render(row) {
      if (isNullObject(row.openStatus)) {
        return ``;
      }
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: getOptionTag(options.value.open_status, row.openStatus),
          bordered: false,
        },
        {
          default: () => getOptionLabel(options.value.open_status, row.openStatus),
        }
      );
    },
  },
];

  const actionColumn = reactive({
    width: 95,
    title: '操作',
    key: 'action',
    fixed: 'right',
    render(record) {
      return h(TableAction as any, {
        style: 'button',
        actions: [
          {
            label: '已选择',
            onClick: handleChoose.bind(null, record),
            disabled: true,
            type: 'default',
            ifShow: () => {
              return couponTypeIdsArr1.value.indexOf(record.id) !== -1;
            },
          },
          {
            label: '选择',
            onClick: handleChoose.bind(null, record),
            ifShow: () => {
              return couponTypeIdsArr1.value.indexOf(record.id) === -1;
            },
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

  const couponTypeIdsArr1 = ref([]);

  function openModal(openCouponTypeIdsArr) {
    couponTypeIdsArr1.value = [];
    showModal.value = true;
    couponTypeIdsArr1.value = openCouponTypeIdsArr
  }

function handleChoose(record: Recordable) {
  emit('reloadCouponList',record);
  // showModal.value = false;
}

  // 重新加载表格数据
  function reloadTable() {
    actionRef.value?.reload();
  }

  defineExpose({
    openModal,
  });
onMounted(() => {
  loadOptions();
});
</script>

<style lang="less"></style>
