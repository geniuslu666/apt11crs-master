<template>
  <div>
    <div class="n-layout-page-header">
      <n-card :bordered="false" :header-style="{
                    padding: '20px',
                  }">
        <template #header>
          <text style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">餐厅管理</text>
        </template>
        <!--  这是由系统生成的CURD表格，你可以将此行注释改为表格的描述 -->
      </n-card>
    </div>
    <n-card :bordered="false" class="proCard" :content-style="{
                    padding: '0 20px 20px',
                  }">
      <n-tabs type="line" animated v-model:value="tabValue" @update:value="handleUpdateValue">
        <n-tab-pane name="ALL" tab="全部">
        </n-tab-pane>
        <n-tab-pane :name="item.value" :tab="item.label" v-for="item in options.open_status">
        </n-tab-pane>
      </n-tabs>
      <BasicForm  ref="searchFormRef" @register="register" @submit="reloadTable" @reset="reloadTable" @keyup.enter="reloadTable">
        <template #statusSlot="{ model, field }">
          <n-input v-model:value="model[field]" />
        </template>
      </BasicForm>
      <BasicTable  ref="actionRef" :openChecked="false" :columns="columns" :request="loadDataTable" :row-key="(row) => row.id" :actionColumn="actionColumn" :scroll-x="scrollX" :resizeHeightOffset="-10000"  :checked-row-keys="checkedIds" @update:checked-row-keys="handleOnCheckedRow"
                   @update:sorter="handleUpdateSorter">
        <template #tableTitle>
          <n-button type="primary"  @click="addTable" class="min-left-space" v-if="hasPermission(['/foodRestaurant/edit'])">
            <template #icon>
              <n-icon>
                <PlusOutlined />
              </n-icon>
            </template>
            添加
          </n-button>
<!--          <n-button type="error" @click="handleBatchDelete" class="min-left-space" v-if="hasPermission(['/foodRestaurant/delete'])">-->
<!--            <template #icon>-->
<!--              <n-icon>-->
<!--                <DeleteOutlined />-->
<!--              </n-icon>-->
<!--            </template>-->
<!--            批量删除-->
<!--          </n-button>-->
        </template>
      </BasicTable>
    </n-card>
    <View ref="viewRef" />
    <Status ref="statusRef" @reloadTable="reloadTable"/>
  </div>
</template>

<script lang="ts" setup>
import { h, reactive, ref, computed, onMounted } from 'vue';
import { useDialog, useMessage } from 'naive-ui';
import { BasicTable, TableAction } from '@/components/Table';
import { BasicForm, useForm } from '@/components/Form/index';
import { usePermission } from '@/hooks/web/usePermission';
import {List, Delete, ResetVerifyCode} from '@/api/foodRestaurant';
import { PlusOutlined, DeleteOutlined } from '@vicons/antd';
import { columns, schemas, options, loadOptions } from './model';
import { adaTableScrollX } from '@/utils/hotgo';
import View from './view.vue';
import Status from './status_desc.vue';
import {useRouter} from "vue-router";
import {useSorter} from "@/hooks/common";

const tabValue = ref('ALL')
const router = useRouter();
const dialog = useDialog();
const message = useMessage();
const { hasPermission } = usePermission();
const actionRef = ref();
const searchFormRef = ref<any>({});
const viewRef = ref();
const checkedIds = ref([]);
const statusRef = ref();

const { updateSorter: handleUpdateSorter, sortStatesRef: sortStatesRef } = useSorter(reloadTable);

const actionColumn = reactive({
  width: 340,
  title: '操作',
  key: 'action',
  fixed: 'right',
  render(record) {
    return h(TableAction as any, {
      style: 'button',
      actions: [
        {
          label: '餐厅设置',
          onClick: handleEdit.bind(null, record),
          auth: ['/foodRestaurant/edit'],
        },

        {
          label: '停业',
          onClick: handleStatus.bind(null, record, 'CLOSE'),
          type: 'warning',
          ifShow: () => {
            return record.openStatus === "OPENING";
          },
          auth: ['/foodRestaurant/status'],
        },
        {
          label: '开业',
          onClick: handleStatus.bind(null, record, 'OPENING'),
          type: 'success',
          ifShow: () => {
            return record.openStatus === "CLOSE";
          },
          auth: ['/foodRestaurant/status'],
        },
        {
          label: '套餐管理',
          onClick: handleGoods.bind(null, record),
          auth: ['/foodGoods/list'],
        },
        // {
        //   label: '重置核销码',
        //   onClick: handleResetVerifyCode.bind(null, record),
        //   type: 'success',
        //   auth: ['/foodRestaurant/resetVerifyCode'],
        // },
      ],
    });
  },
});

const scrollX = computed(() => {
  return adaTableScrollX(columns, actionColumn.width);
});

const [register, {}] = useForm({
  gridProps: { cols: '1 s:1 m:2 l:3 xl:4 2xl:4' },
  labelWidth: 80,
  schemas,
});

// 加载表格数据
const loadDataTable = async (res) => {
  if(tabValue.value != 'ALL'){
    res.openStatus = tabValue.value
  }

  console.log("sortStatesRef", sortStatesRef)
  if(sortStatesRef.value.length > 0){
    var sort = "";
    sortStatesRef.value.forEach(element => {
      if(element.columnKey == 'sort'){
        sort = element.order
      }
    });
    return await List({ ...searchFormRef.value?.formModel, ...res, sort: sort });
  }else{
    return await List({ ...searchFormRef.value?.formModel, ...res });
  }


};

// 更新选中的行
function handleOnCheckedRow(rowKeys) {
  checkedIds.value = rowKeys;
}

// 重新加载表格数据
function reloadTable() {
  actionRef.value?.reload();
}

function handleUpdateValue(e){
  tabValue.value = e;
  reloadTable()
}

// 添加数据
function addTable() {
  router.push({ name: 'foodRestaurantEdit', params: { id: 0 } });
}

// 编辑数据
function handleEdit(record: Recordable) {
  router.push({ name: 'foodRestaurantEdit', params: { id: record.id } });
}

// 套餐管理
function handleGoods(record: Recordable) {
  router.push({ name: 'foodGoodsIndex', params: { id: record.id } });
}

// // 查看详情
// function handleView(record: Recordable) {
//   viewRef.value.openModal(record);
// }
//
// // 单个删除
// function handleDelete(record: Recordable) {
//   dialog.warning({
//     title: '警告',
//     content: '你确定要删除？',
//     positiveText: '确定',
//     negativeText: '取消',
//     onPositiveClick: () => {
//       Delete(record).then((_res) => {
//         message.success('删除成功');
//         reloadTable();
//       });
//     },
//   });
// }

// 批量删除
// function handleBatchDelete() {
//   if (checkedIds.value.length < 1){
//     message.error('请至少选择一项要删除的数据');
//     return;
//   }
//
//   dialog.warning({
//     title: '警告',
//     content: '你确定要批量删除？',
//     positiveText: '确定',
//     negativeText: '取消',
//     onPositiveClick: () => {
//       Delete({ id: checkedIds.value }).then((_res) => {
//         checkedIds.value = [];
//         message.success('删除成功');
//         reloadTable();
//       });
//     },
//   });
// }

// 修改状态
function handleStatus(record: Recordable, status: string) {
  statusRef.value.openModal(record.id,status);


  // Status({ id: record.id, status: status }).then((_res) => {
  //   message.success('设为' + getOptionLabel(options.value.open_status, status) + '成功');
  //   setTimeout(() => {
  //     reloadTable();
  //   });
  // });
}

// 重置核销码
function handleResetVerifyCode(record: Recordable) {

  dialog.warning({
    title: '警告',
    content: '你确定要重置核销码？',
    positiveText: '确定',
    negativeText: '取消',
    onPositiveClick: () => {
      ResetVerifyCode(record).then((_res) => {
        message.success('重置成功');
        reloadTable();
      });
    },
  });
}

onMounted(() => {
  loadOptions();
});
</script>

<style lang="less" scoped></style>

