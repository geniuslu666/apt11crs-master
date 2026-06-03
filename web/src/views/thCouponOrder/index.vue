<template>
  <div>
    <div class="n-layout-page-header" style="margin: 0">
      <n-card :bordered="false" :header-style="{
        padding: '20px',
      }">
        <template #header>
          <text style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">礼品券订单</text>
        </template>
        <!--  这是由系统生成的CURD表格，你可以将此行注释改为表格的描述 -->
      </n-card>
    </div>
    <n-card :bordered="false" class="proCard" :content-style="{
      padding: '0 20px 20px',
    }">
      <BasicForm ref="searchFormRef" @register="register" @submit="reloadTable" @reset="reloadTable"
        @keyup.enter="reloadTable">
        <template #statusSlot="{ model, field }">
          <n-input v-model:value="model[field]" />
        </template>
        <template #verifyTimeSlot="{ model, field }">
          <n-date-picker
            v-model:formatted-value="model[field]"
            type="datetimerange"
            value-format="yyyy-MM-dd HH:mm:ss"
            clearable
            :shortcuts="defRangeShortcuts()"
            style="width: 100%"
          />
        </template>
      </BasicForm>
      <BasicTable ref="actionRef" :openChecked="false" :columns="columns" :request="loadDataTable"
        :row-key="(row) => row.id" :actionColumn="actionColumn" :scroll-x="scrollX" :resizeHeightOffset="-10000">
        <template #tableTitle>
          <n-button type="primary" @click="handleExport" class="min-left-space"
            v-if="hasPermission(['/thMemberCoupon/export'])">
            <template #icon>
              <n-icon>
                <ExportOutlined />
              </n-icon>
            </template>
            导出
          </n-button>
        </template>
      </BasicTable>
    </n-card>
    <View ref="viewRef" />
    <Verify ref="verifyRef" @reloadTable="reloadTable" />
  </div>
</template>

<script lang="ts" setup>
import { ref, computed, onMounted, reactive, h } from 'vue';
import { BasicTable, TableAction } from '@/components/Table';
import { usePermission } from '@/hooks/web/usePermission';
import { List, Export, Recycle } from '@/api/thMemberCoupon';
import { columns, schemas, loadOptions } from './model';
import { adaTableScrollX } from '@/utils/hotgo';
import { defRangeShortcuts } from '@/utils/dateUtil';
import { BasicForm, useForm } from "@/components/Form";
import { useRouter } from "vue-router";
import { ExportOutlined } from "@vicons/antd";
import { useMessage, useDialog } from "naive-ui";
import View from "@/views/thCouponOrder/view.vue";
import Verify from "@/views/thCouponOrder/verify.vue";

const router = useRouter();
const dialog = useDialog();
const message = useMessage();
const { hasPermission } = usePermission();
const actionRef = ref();
const searchFormRef = ref<any>({});
const viewRef = ref();
const verifyRef = ref();

const actionColumn = reactive({
  width: 235,
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
          auth: ['/thMemberCoupon/view'],
        },
        {
          label: '回收',
          onClick: handleRecycle.bind(null, record),
          ifShow: () => {
            return record.state === 1 || record.state === 2;
          },
          auth: ['/thMemberCoupon/recycle'],
          type: 'warning'
        },
        {
          label: '手动核销',
          onClick: manualVerify.bind(null, record),
          ifShow: () => {
            return record.state === 1 || record.state === 2;
          },
        },
      ],
    });
  },
});

const scrollX = computed(() => {
  return adaTableScrollX(columns, actionColumn.width);
});

const [register, { }] = useForm({
  gridProps: { cols: '1 s:1 m:3 l:4 xl:5 2xl:4' },
  labelWidth: 90,
  schemas,
});


// 加载表格数据
const loadDataTable = async (res) => {
  return await List({ ...searchFormRef.value?.formModel, ...res });
};

// 重新加载表格数据
function reloadTable() {
  actionRef.value?.reload();
}

// 导出
function handleExport() {
  message.loading('正在导出列表...', { duration: 1200 });
  Export(searchFormRef.value?.formModel);
}

// 查看详情
function handleView(record: Recordable) {
  viewRef.value.openModal(record.id);
  // router.push({ name: 'thCoupon_view', params: { id: record.id } });
}

// 回收礼品券
function handleRecycle(record: Recordable) {
  dialog.warning({
    title: '警告',
    content: '回收后用户无法再使用，确认回收？',
    positiveText: '确定',
    negativeText: '取消',
    onPositiveClick: () => {
      Recycle({ id: record.id }).then((_res) => {
        message.success('操作成功');
        setTimeout(() => {
          reloadTable();
        });
      });
    },
  });
}

// 手动核销
function manualVerify(record: Recordable) {
  verifyRef.value.openModal(record);
}

onMounted(() => {
  loadOptions();
});
</script>

<style lang="less" scoped></style>
