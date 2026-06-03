<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content closable :header-style="{
                    padding: '20px',
                  }" :body-content-style="{
                    padding: '20px',
                  }">
        <template #header>
          <div style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">选择优惠券</div>
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
import { List } from '@/api/pmsCouponType';
import {BasicForm, useForm} from "@/components/Form";
import {useMessage} from "naive-ui";
import { FormSchema } from '@/components/Form';

const emit = defineEmits(['reloadCoupon']);
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
      }else if(record.scene == 5){
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
            default: () => '储物柜',
          }
        );
      }
      return '--'
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
  res.status = 1;
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
    message.error('请选择优惠券');
    return false;
  }
  console.log('reloadCoupon_before',checkedIds.value)
  emit('reloadCoupon',checkedIds.value);
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
