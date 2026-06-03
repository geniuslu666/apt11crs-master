<template>
  <div>
    <div class="n-layout-page-header">
      <n-card :bordered="false" :header-style="{
                    padding: '20px',
                  }">
        <template #header>
          <text style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">{{ restaurantName ? restaurantName + '-套餐管理' : '套餐管理' }}</text>
        </template>
        <!--  这是由系统生成的CURD表格，你可以将此行注释改为表格的描述 -->
      </n-card>
    </div>
    <n-card :bordered="false" class="proCard" :content-style="{
                    padding: '0 20px 20px',
                  }">
      <n-tabs type="line" animated v-model:value="tabValue" @update:value="handleUpdateValue">
        <n-tab-pane :name="0" tab="全部">
        </n-tab-pane>
        <n-tab-pane :name="1" tab="销售中">
        </n-tab-pane>
        <n-tab-pane :name="2" tab="仓库中">
        </n-tab-pane>
      </n-tabs>
      <BasicForm  ref="searchFormRef" @register="register" @submit="reloadTable" @reset="reloadTable" @keyup.enter="reloadTable">
        <template #statusSlot="{ model, field }">
          <n-input v-model:value="model[field]" />
        </template>
      </BasicForm>
      <BasicTable  ref="actionRef" openChecked :columns="columns" :request="loadDataTable" :row-key="(row) => row.id" :actionColumn="actionColumn" :scroll-x="scrollX" :resizeHeightOffset="-10000"  :checked-row-keys="checkedIds" @update:checked-row-keys="handleOnCheckedRow">
        <template #tableTitle>
          <n-button type="primary"  @click="addTable" class="min-left-space" v-if="hasPermission(['/foodGoods/edit'])">
            <template #icon>
              <n-icon>
                <PlusOutlined />
              </n-icon>
            </template>
            添加
          </n-button>
          <n-button type="error" @click="handleBatchDelete" class="min-left-space" v-if="hasPermission(['/foodGoods/delete'])">
            <template #icon>
              <n-icon>
                <DeleteOutlined />
              </n-icon>
            </template>
            批量删除
          </n-button>
        </template>
      </BasicTable>
    </n-card>
    <View ref="viewRef" />
    <Edit ref="editRef" @reloadTable="reloadTable" />
    <Copy ref="copyRef" @reloadTable="reloadTable" />
  </div>
</template>

<script lang="ts" setup>
import {h, reactive, ref, computed, onMounted} from 'vue';
import { useDialog, useMessage } from 'naive-ui';
import { BasicTable, TableAction } from '@/components/Table';
import { BasicForm, useForm } from '@/components/Form/index';
import { usePermission } from '@/hooks/web/usePermission';
import { List, Delete, Status } from '@/api/foodGoods';
import { PlusOutlined, DeleteOutlined } from '@vicons/antd';
import { columns, schemas } from './model';
import { adaTableScrollX } from '@/utils/hotgo';
import View from './view.vue';
import Edit from './edit.vue';
import Copy from './copy.vue';
import {useRouter} from "vue-router";
import { View as RestaurantView } from '@/api/foodRestaurant';

const restaurantName = ref('');
const tabValue = ref(0)
const dialog = useDialog();
const message = useMessage();
const { hasPermission } = usePermission();
const actionRef = ref();
const searchFormRef = ref<any>({});
const viewRef = ref();
const editRef = ref();
const copyRef = ref();
const checkedIds = ref([]);
const router = useRouter();
const params = router.currentRoute.value.params;

const actionColumn = reactive({
  width: 256,
  title: '操作',
  key: 'action',
  fixed: 'right',
  render(record) {
    return h(TableAction as any, {
      style: 'button',
      actions: [
        {
          label: '编辑',
          onClick: handleEdit.bind(null, record),
          auth: ['/foodGoods/edit'],
        },
        {
          label: '下架',
          onClick: handleStatus.bind(null, record, 2),
          type: 'warning',
          ifShow: () => {
            return record.goodsState === 1;
          },
          auth: ['/foodGoods/status'],
        },
        {
          label: '上架',
          onClick: handleStatus.bind(null, record, 1),
          type: 'success',
          ifShow: () => {
            return record.goodsState === 2;
          },
          auth: ['/foodGoods/status'],
        },
        {
          label: '复制',
          onClick: handleCopy.bind(null, record),
          auth: ['/foodGoods/edit'],
        },
        {
          label: '删除',
          onClick: handleDelete.bind(null, record),
          auth: ['/foodGoods/delete'],
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
  labelWidth: 80,
  schemas,
});

function handleUpdateValue(e){
  tabValue.value = e;
  reloadTable()
}

// 加载表格数据
const loadDataTable = async (res) => {
  res.restaurantId = params.id
  if(tabValue.value != 0){
    res.goodsState = tabValue.value
  }
  return await List({ ...searchFormRef.value?.formModel, ...res });
};

// 更新选中的行
function handleOnCheckedRow(rowKeys) {
  checkedIds.value = rowKeys;
}

// 重新加载表格数据
function reloadTable() {
  actionRef.value?.reload();
}

// 添加数据
function addTable() {
  // router.push({ name: 'foodGoodsEdit', params: { id: params.id, goodsId: 0 } });
  editRef.value.openModal(null, params.id);
}

// 编辑数据
function handleEdit(record: Recordable) {
  // router.push({ name: 'foodGoodsEdit', params: { id: record.restaurantId, goodsId: record.id } });
  editRef.value.openModal(record, record.restaurantId);
}

// 复制数据
function handleCopy(record: Recordable) {
  // router.push({ name: 'foodGoodsCopy', params: { id: record.restaurantId, goodsId: record.id } });
  copyRef.value.openModal(record, record.restaurantId);
}

// 查看详情
function handleView(record: Recordable) {
  viewRef.value.openModal(record);
}

// 单个删除
function handleDelete(record: Recordable) {
  dialog.warning({
    title: '警告',
    content: '你确定要删除？',
    positiveText: '确定',
    negativeText: '取消',
    onPositiveClick: () => {
      Delete(record).then((_res) => {
        message.success('删除成功');
        reloadTable();
      });
    },
  });
}

// 批量删除
function handleBatchDelete() {
  if (checkedIds.value.length < 1){
    message.error('请至少选择一项要删除的数据');
    return;
  }

  dialog.warning({
    title: '警告',
    content: '你确定要批量删除？',
    positiveText: '确定',
    negativeText: '取消',
    onPositiveClick: () => {
      Delete({ id: checkedIds.value }).then((_res) => {
        checkedIds.value = [];
        message.success('删除成功');
        reloadTable();
      });
    },
  });

}

// 修改状态
function handleStatus(record: Recordable, status: number) {
  Status({ id: record.id, status: status }).then((_res) => {
    message.success('设为' + (status == 1 ? '上架' : '下架') + '成功');
    setTimeout(() => {
      reloadTable();
    });
  });
}

onMounted(() => {
  RestaurantView({ id: params.id })
    .then((res) => {
      restaurantName.value = res.name;
    });
});

</script>

<style lang="less" scoped></style>

