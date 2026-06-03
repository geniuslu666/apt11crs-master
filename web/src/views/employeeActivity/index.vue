<template>
  <div>
    <div class="n-layout-page-header">
      <n-card :bordered="false" :header-style="{
                    padding: '20px',
                  }">
        <template #header>
          <text style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">活动列表</text>
        </template>
        <!--  这是由系统生成的CURD表格，你可以将此行注释改为表格的描述 -->
      </n-card>
    </div>
    <n-card :bordered="false" class="proCard" :content-style="{
                    padding: '0 20px 20px',
                  }">
      <div class="statusTab">
        <div :class="tabValue==0 ? 'active' : ''"><span @click="handleUpdateValue(0)">全部</span></div>
        <div :class="tabValue==1 ? 'active' : ''"><span @click="handleUpdateValue(1)">未开始</span></div>
        <div :class="tabValue==2 ? 'active' : ''"><span @click="handleUpdateValue(2)">进行中</span></div>
        <div :class="tabValue==3 ? 'active' : ''"><span @click="handleUpdateValue(3)">已结束</span></div>
      </div>
      <BasicForm  ref="searchFormRef" @register="register" @submit="reloadTable" @reset="reloadTable" @keyup.enter="reloadTable">
        <template #statusSlot="{ model, field }">
          <n-input v-model:value="model[field]" />
        </template>
      </BasicForm>
      <BasicTable  ref="actionRef" openChecked :columns="columns" :request="loadDataTable" :row-key="(row) => row.id" :actionColumn="actionColumn" :scroll-x="scrollX" :resizeHeightOffset="-10000"  :checked-row-keys="checkedIds" @update:checked-row-keys="handleOnCheckedRow">
        <template #tableTitle>
          <n-button type="primary"  @click="addTable" class="min-left-space" v-if="hasPermission(['/employeeActivity/edit'])">
            <template #icon>
              <n-icon>
                <PlusOutlined />
              </n-icon>
            </template>
            添加
          </n-button>
          <n-button type="error" @click="handleBatchDelete" class="min-left-space" v-if="hasPermission(['/employeeActivity/delete'])">
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
   <Edit ref="editRef" @reloadTable="reloadTable" />
   <View ref="viewRef" />
   <CouponRecord ref="couponRecordRef" />
  </div>
</template>

<script lang="ts" setup>
import { h, reactive, ref, computed, onMounted } from 'vue';
import { useDialog, useMessage } from 'naive-ui';
import { BasicTable, TableAction } from '@/components/Table';
import { BasicForm, useForm } from '@/components/Form/index';
import { usePermission } from '@/hooks/web/usePermission';
import { List, Delete, Status, BatchIssueCoupons } from '@/api/employeeActivity';
import { PlusOutlined, DeleteOutlined } from '@vicons/antd';
import { columns, schemas, options, loadOptions } from './model';
import { adaTableScrollX, getOptionLabel } from '@/utils/hotgo';
import Edit from './edit.vue';
import View from './view.vue';
import CouponRecord from './couponRecord.vue';

const dialog = useDialog();
const message = useMessage();
const { hasPermission } = usePermission();
const actionRef = ref();
const searchFormRef = ref<any>({});
const editRef = ref();
const viewRef = ref();
const couponRecordRef = ref();
const checkedIds = ref([]);
const tabValue = ref(0)

const actionColumn = reactive({
  width: 440,
  title: '操作',
  key: 'action',
  fixed: 'right',
  render(record) {
    return h(TableAction as any, {
      style: 'button',
      actions: [
        {
          label: '详情',
          onClick: handleView.bind(null, record),
          // auth: ['/employeeActivity/view'],
        },
        {
          label: '编辑',
          onClick: handleEdit.bind(null, record),
          auth: ['/employeeActivity/edit'],
        },

        {
          label: '禁用',
          onClick: handleStatus.bind(null, record, 2),
          ifShow: () => {
            return record.isEnabled === 1;
          },
          auth: ['/employeeActivity/status'],
        },
        {
          label: '启用',
          onClick: handleStatus.bind(null, record, 1),
          ifShow: () => {
            return record.isEnabled === 2;
          },
          auth: ['/employeeActivity/status'],
        },
        {
          label: '删除',
          onClick: handleDelete.bind(null, record),
          auth: ['/employeeActivity/delete'],
        },
        {
          label: '批量发放',
          onClick: handleBatchIssueCoupons.bind(null, record),
          ifShow: () => {
            return record.status === 2 && record.isEnabled === 1;
          },
          auth: ['/employeeActivity/batchIssueCoupons'],
        },
        {
          label: '券领取记录',
          onClick: handleCouponRecord.bind(null, record),
          auth: ['/employeeActivity/couponRecordList'],
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

// 加载表格数据
const loadDataTable = async (res) => {
  res.status = tabValue.value;
  return await List({ ...searchFormRef.value?.formModel, ...res, isLanguage: true});
};

function handleUpdateValue(e){
  tabValue.value = e;
  reloadTable()
}

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
  editRef.value.openModal(null);
}

// 编辑数据
function handleEdit(record: Recordable) {
  editRef.value.openModal(record);
}

// 查看详情
function handleView(record: Recordable) {
  viewRef.value.openModal(record.id);
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
    message.success('设为' + getOptionLabel(options.value.sys_normal_disable, status) + '成功');
    setTimeout(() => {
      reloadTable();
    });
  });
}

// 查看券领取记录
function handleCouponRecord(record: Recordable) {
  couponRecordRef.value.openModal(record);
}

// 批量发放预约券
function handleBatchIssueCoupons(record: Recordable) {
  dialog.warning({
    title: '批量发放',
    content: `确定要为活动「${record.name}」批量发放预约券吗？将为所有符合条件的员工发放需要预约的券。`,
    positiveText: '确定发放',
    negativeText: '取消',
    onPositiveClick: () => {
      const loadingMsg = message.loading('正在批量发放中，请稍候...', { duration: 0 });
      BatchIssueCoupons({ activityId: record.id })
        .then((res) => {
          loadingMsg.destroy();
          const data = res;
          dialog.success({
            title: '发放完成',
            content: `成功发放 ${data.issuedCount} 张券，失败 ${data.failedCount} 张。`,
            positiveText: '确定',
          });
          reloadTable();
        })
        .catch((_err) => {
          loadingMsg.destroy();
        });
    },
  });
}

onMounted(() => {
  loadOptions();
});
</script>

<style lang="less" scoped>
.statusTab{
  display: flex;
  margin-bottom: 30px;
  div{
    padding: 0 16px;
    height: 28px;
    line-height: 28px;
    text-align: center;
    font-size: 14px;
    color: #4E5969;
    span{
      cursor: pointer;
    }
    &.active{
      background: #F2F3F8;
      border-radius: 28px;
      color: #1664FF;
      font-weight: 500;
    }
  }
}
</style>

