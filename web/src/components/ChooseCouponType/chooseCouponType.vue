<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content closable :header-style="{
                    padding: '20px',
                  }" :body-content-style="{
                    padding: '25px 20px',
                  }">
        <template #header>
          <div style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">选择优惠券</div>
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
import {computed, h, reactive, ref} from 'vue';
import {adaModalWidth, adaTableScrollX} from "@/utils/hotgo";
import {BasicTable, TableAction} from "@/components/Table";
import {BasicForm, useForm} from "@/components/Form";
import { FormSchema } from '@/components/Form';
import {List} from '@/api/pmsCouponType';

  const emit = defineEmits(['reloadCouponList']);
  const loading = ref(false);
  const showModal = ref(false);
  const dialogWidth = computed(() => {
    return adaModalWidth(1100);
  });
const actionRef = ref();
const searchFormRef = ref<any>({});

// 表格搜索表单
const schemas = ref<FormSchema[]>([
  {
    field: 'type',
    component: 'NSelect',
    label: '优惠券类型',
    componentProps: {
      placeholder: '优惠券类型',
      options: [
        {
          labelField: '满减',
          valueField: 'reward',
        },
        {
          labelField: '折扣',
          valueField: 'discount',
        }
      ],
      labelField: 'labelField',
      valueField: 'valueField',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'couponName',
    component: 'NInput',
    label: '优惠券名称',
    componentProps: {
      placeholder: '优惠券名称',
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
    title: '发放数量',
    key: 'count',
    align: 'left',
    width: 100,
    render(record) {
      if(record.count==0){
        return "无限制";
      }
      return record.count
    }
  },
  {
    title: '已领取数量',
    key: 'leadCount',
    align: 'left',
    width: 100,
  },
  {
    title: '优惠券名称',
    key: 'couponName',
    align: 'left',
    width: 200,
    render(row){
      return row.couponName
    }
  },
  {
    title: '优惠券类型',
    key: 'type',
    align: 'left',
    width: 90,
    render(record) {
      if(record.type == 'reward'){
        return h(
          'div',
          {
            style: {
              color: '#17A158',
              width: '38px',
              height: '22px',
              lineHeight: '22px',
              textAlign: 'center',
              fontSize: '14px',
              borderRadius: '2px',
              fontWeight: '400',
              border: '1px solid #17A158'
            },
            bordered: false,
          },
          {
            default: () => '满减',
          }
        );
      }
      if(record.type == 'discount'){
        return h(
          'div',
          {
            style: {
              color: '#F48720',
              width: '38px',
              height: '22px',
              lineHeight: '22px',
              textAlign: 'center',
              fontSize: '14px',
              borderRadius: '2px',
              fontWeight: '400',
              border: '1px solid #F48720'
            },
            bordered: false,
          },
          {
            default: () => '折扣',
          }
        );
      }
      return '--'
    }
  },
  {
    title: '优惠券金额/折扣',
    key: 'couponName',
    align: 'left',
    width: 130,
    render(record) {
      if(record.type=='reward'){
        return record.money + "JPY";
      }else if(record.type == 'discount'){
        return record.discount + "折";
      }
    }
  },
  {
    title: '满多少元使用',
    key: 'atLeast',
    align: 'left',
    width: 120,
  },
  {
    title: '适用场景',
    key: 'scene',
    align: 'left',
    width: 110,
    render(record) {
      if(record.scene == 1){
        return h(
          'div',
          {
            style: {
              color: '#26A763',
              padding: '0 5px',
              height: '22px',
              background: '#E3F4EB',
              lineHeight: '22px',
              textAlign: 'center',
              fontSize: '14px',
              borderRadius: '2px',
              fontWeight: '400',
            },
            bordered: false,
          },
          {
            default: () => '民宿',
          }
        );
      }else if(record.scene == 2){
        return h(
          'div',
          {
            style: {
              color: '#3F9EFF',
              padding: '0 5px',
              height: '22px',
              background: '#ECF5FF',
              lineHeight: '22px',
              textAlign: 'center',
              fontSize: '14px',
              borderRadius: '2px',
              fontWeight: '400',
            },
            bordered: false,
          },
          {
            default: () => '订餐',
          }
        );
      }else if(record.scene == 3){
        return h(
          'div',
          {
            style: {
              color: '#EFA020',
              padding: '0 5px',
              height: '22px',
              background: '#FFECCE',
              lineHeight: '22px',
              textAlign: 'center',
              fontSize: '14px',
              borderRadius: '2px',
              fontWeight: '400',
            },
            bordered: false,
          },
          {
            default: () => '按摩',
          }
        );
      }else if(record.scene == 4){
        return h(
          'div',
          {
            style: {
              color: '#F56C6C',
              padding: '0 5px',
              height: '22px',
              background: '#FEF0F0',
              lineHeight: '22px',
              textAlign: 'center',
              fontSize: '14px',
              borderRadius: '2px',
              fontWeight: '400',
            },
            bordered: false,
          },
          {
            default: () => '接送机/包车',
          }
        );
      }
      return '--'
    }
  },
  {
    title: '有效期',
    key: 'time',
    align: 'left',
    width: 150,
    render(row){
      if(row.validityType == 1){
        return row.endUseTime
      }else if(row.validityType == 2){
        return '领取后'+row.fixedTerm+'天有效'
      }else{
        return '长期有效'
      }
    }
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
            label: '已领完',
            onClick: handleChoose.bind(null, record),
            disabled: true,
            type: 'default',
            ifShow: () => {
              return record.count > 0 && record.count == record.leadCount;
            },
          },
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
              return (record.count == 0 || (record.count > 0 && record.count > record.leadCount)) && couponTypeIdsArr1.value.indexOf(record.id) === -1;
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
    gridProps: { cols: 3 },
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
